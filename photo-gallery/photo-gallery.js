angular.module('photoGallery', ['ui.bootstrap', 'ui.utils', 'ui.router', 'ngAnimate']);

angular.module('photoGallery').config(function ($stateProvider) {

    $stateProvider.state('photoGallery', {
        url: '/photo-gallery',
        templateUrl: 'photo-gallery/partial/main/main.html',
        controller: 'MainPhotoGalleryController',
        controllerAs: 'mainPhotoGallery',
        resolve: {
            photos: function (photoGalleryService) {
                return photoGalleryService.getPhotos();
            },
            settings: function (photoGallerySettings) {
                return photoGallerySettings.getPhotoGallerySettings();
            }
        }
    });
});

