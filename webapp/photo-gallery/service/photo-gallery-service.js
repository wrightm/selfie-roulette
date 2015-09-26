(function () {

    angular.module('photoGallery').factory('photoGalleryService', photoGalleryService);

    function photoGalleryService($resource, $q) {

        var self = this;
        self.Photos = $resource('http://localhost:8080/photos/:id',
            {},
            {
                update: {
                    method: 'PUT'
                }
            }
        );

        return {
            getPhotos: getPhotos,
            winningPhoto: winningPhoto
        };


        function getPhotos() {
            var deferred = $q.defer();
            self.Photos.query().$promise.then(function success(response) {
                var numberOfFiles = response.length;
                var photos = [];
                for (var index = 0; index < numberOfFiles; ++index) {
                    var photo = response[index];
                    if (!photo.winner) {
                        photos.push(photo);
                    }
                }
                deferred.resolve(photos);
            }, function error() {
                deferred.reject([]);
            });
            return deferred.promise;
        }

        function winningPhoto(photo) {
            delete photo.column;
            delete photo.order;
            delete photo.row;
            return self.Photos.update({id: photo.id}, photo).$promise;
        }
    }

})();

