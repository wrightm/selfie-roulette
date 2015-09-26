package persistence

import (
    "model"
    "config"
    "io/ioutil"
    "log"
    "path/filepath"
    "strings"
    "os"
    "encoding/json"
    "errors"
)

func RepoFindPhoto(id string) (model.Photo, error) {
    return GetPhoto(id)
}

func RepoCreatePhoto(photo model.Photo) model.Photo {
    // Todo
    return photo
}

func RepoUpdatePhoto(newPhoto model.Photo) (model.Photo, error) {
    if DoesPhotoExist(newPhoto.Id) {
        newPhoto, err := UpdatePhotoFile(newPhoto)
        return newPhoto, err
    } else {
        return newPhoto, errors.New("Photo with id " + newPhoto.Id + " does not exist")
    }
}

func RepoDeletePhoto(id string) bool {
    // Todo
    return false
}

func RepoPhotos() model.Photos {
    return GetPhotos()
}

func FindFiles(dirname string, fileType string) []string {
    fileList, err := ioutil.ReadDir(dirname)
    if err != nil {
        log.Println(err)
    }

    var fileNames []string
    for _, info := range fileList {
        name := info.Name()
        if info.IsDir() {
            FindFiles(filepath.Join(dirname, name), fileType)
        } else {
            fileName := filepath.Join(dirname, name)
            ext := strings.ToLower(filepath.Ext(name))
            if ext[1:] != fileType {
                continue
            }
            fileNames = append(fileNames, fileName)
        }
    }
    return fileNames
}

func FindFile(fileDescriptor string, dirname string) (string, error) {
    fileList, err := ioutil.ReadDir(dirname)
    if err != nil {
        log.Println(err)
    }

    for _, info := range fileList {
        name := info.Name()
        if info.IsDir() {
            FindFile(fileDescriptor, filepath.Join(dirname, name))
        } else {
            fileName := filepath.Join(dirname, name)
            if strings.Contains(name, fileDescriptor) {
                return fileName, nil
            }
        }
    }
    return "", errors.New("File not found for descriptor " + fileDescriptor)
}

func GetPhotos() model.Photos {
    CreateJsonFilesFromImages(config.Config.ImagesDir, config.Config.ImageFileType, config.Config.RepoDir)
    files := FindFiles(config.Config.RepoDir, "json")
    return MakePhotosFromFiles(files)
}

func MakePhotosFromFiles(files []string) model.Photos {
    var photos model.Photos = make([]model.Photo, 0)
    for _, file := range files {
        openedFile, _ := os.Open(file)
        decoder := json.NewDecoder(openedFile)
        photo := model.Photo{}
        err := decoder.Decode(&photo)
        if err != nil {
            log.Fatal("Cannot Decode photo with file: ", files)
        }
        photos = append(photos, photo)
        defer openedFile.Close()
    }
    return photos
}

func GetPhoto(photoId string) (model.Photo, error) {
    photos := GetPhotos()
    for _, photo := range photos {
        if photo.Id == photoId {
            return photo, nil
        }
    }
    return model.Photo{}, errors.New("Photo with id " + photoId + " does not exist")
}

func DoesPhotoExist(photoId string) bool {
    photos := GetPhotos()
    for _, photo := range photos {
        if photo.Id == photoId {
            return true
        }
    }
    return false
}

func UpdatePhotoFile(newPhoto model.Photo) (model.Photo, error) {
    fileName, err := FindFile(newPhoto.Id, config.Config.RepoDir)

    if err != nil {
        return model.Photo{}, err
    } else {
        return SavePhotoToFile(fileName, newPhoto)
    }

    return newPhoto, nil
}

func SavePhotoToFile(fileName string, newPhoto model.Photo) (model.Photo, error) {
    newPhotoJson, _ := json.Marshal(newPhoto)
    err := ioutil.WriteFile(fileName, newPhotoJson, 0644)
    return newPhoto, err
}



func CreateJsonFilesFromImages(imagesDir string, imageFileType string, dirToSaveJSONFilesTo string) {
    imageFileNames := FindFiles(imagesDir, imageFileType)
    imageFileNames = FilterFileNames(imageFileNames, DoesImageMetaFileNotExist)
    CreateImageMetaFilesFromImages(imageFileNames, dirToSaveJSONFilesTo)
}


type filterFunc func(string) bool

func FilterFileNames(fileNames []string, predicate filterFunc) []string {
    var filteredFileNames []string = make([]string, 0)
    for _, fileName := range fileNames {
        if predicate(fileName) {
            filteredFileNames = append(filteredFileNames, fileName)
        }
    }
    return filteredFileNames
}

func DoesImageMetaFileNotExist(imageFileName string) bool {
    imageFileMetaName := GetImageFileMetaName(GetFileNameWithoutTypeAndDir(imageFileName))
    if _, err := os.Stat(imageFileMetaName); err == nil {
        return false
    }
    return true
}

func CreateImageMetaFilesFromImages(imageFileNames []string, dirToSaveJSONFilesTo string) {
    webAppImageDir := GetWebAppImageDir()
    for _, imageFileName := range imageFileNames {
        imageFileNameWithoutDir := ParseFileDir(imageFileName)
        imageFileMetaName := GetImageFileMetaName(GetFileNameWithoutTypeAndDir(imageFileNameWithoutDir))
        photo := CreatePhotoFromImageFile(webAppImageDir, imageFileNameWithoutDir)
        SavePhotoToFile(imageFileMetaName, photo)
    }
}

func GetWebAppImageDir() string {
    webAppImageDirSplit := strings.Split(config.Config.ImagesDir, "/")
    return webAppImageDirSplit[len(webAppImageDirSplit) - 1]
}

func ParseFileDir(filename string) string {
    filenameDirSplit := strings.Split(filename, "/")
    return filenameDirSplit[len(filenameDirSplit) - 1]
}

func CreatePhotoFromImageFile(webAppImageDir string, imageFileName string) model.Photo {
    imageName := strings.Split(imageFileName, ".")[0]
    return model.Photo{
        Id : imageName,
        Name : imageName,
        Filename : webAppImageDir + "/" + imageFileName,
        Winner : false,
    }
}

func GetImageFileMetaName(imageName string) string {
    return config.Config.RepoDir + "/image-meta-" + imageName + ".json"
}

func GetFileNameWithoutTypeAndDir(fileName string) string {
    fileNameWithoutDir := ParseFileDir(fileName)
    return strings.Split(fileNameWithoutDir, ".")[0]
}
