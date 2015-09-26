(function () {
    angular.module('photoGallery').factory('photoGallerySettings', photoGallerySettings);

    function photoGallerySettings($resource, $q) {

        var self = this;
        self.PhotoGallerySettings = $resource('data/photo-gallery-settings.json');
        self.getPhotoGallerySettings = getPhotoGallerySettings;

        return {
            getPhotoGallerySettings: self.getPhotoGallerySettings
        };

        function getPhotoGallerySettings() {
            var deferred = $q.defer();
            self.PhotoGallerySettings.get().$promise.then(function success(response) {
                deferred.resolve(response);
            }, function error() {
                deferred.reject({
                    refreshPhotoTimeInterval: 10000,
                    winnerTimeInterval: 10,
                    serverAddress: "localhost:8080"
                });
            });
            return deferred.promise;
        }

    }

})();
