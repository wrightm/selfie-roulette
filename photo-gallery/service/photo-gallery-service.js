(function () {

    angular.module('photoGallery').factory('photoGalleryService', photoGalleryService);

    function photoGalleryService($resource, $q) {

        var self = this;

        self.Photos = $resource('data/photos.json');
        self.getPhotos = getPhotos;

        return {
            getPhotos: self.getPhotos
        };


        function getPhotos() {
            var deferred = $q.defer();
            self.Photos.get().$promise.then(function success(response) {
                var filenames = response.filenames;
                var numberOfFiles = filenames.length;
                var photos = [];
                for (var index = 0; index < numberOfFiles; ++index) {
                    var filename = filenames[index];
                    var photo = {
                        filename: filename,
                        id: index
                    };
                    photos.push(photo);
                }
                deferred.resolve(photos);
            }, function error() {
                deferred.reject([]);
            });
            return deferred.promise;
        }
    }

})();

