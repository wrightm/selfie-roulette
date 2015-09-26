(function () {

    angular.module('photoGallery').controller('MainPhotoGalleryController', mainPhotoGalleryController);


    function mainPhotoGalleryController($scope, $interval, $modal, photos, photoGalleryService, settings, countdownTimerService) {

        var self = this;
        self.photos = photos;
        self.pickWinningPhoto = pickWinningPhoto;
        self.showPickWinnerButton = true;

        function randomisePhotoPositions() {
            var chosenPhotoPositions = [];
            for (var photoPosition = 0; photoPosition < self.photos.length; ++photoPosition) {
                var chosenPhotoPosition = Math.floor(Math.random() * self.photos.length);
                while (doesPhotoPositionAlreadyExist(chosenPhotoPositions, chosenPhotoPosition)) {
                    chosenPhotoPosition = Math.floor(Math.random() * self.photos.length);
                }
                chosenPhotoPositions.push(chosenPhotoPosition);
                self.photos[photoPosition].order = chosenPhotoPosition;
            }
        }


        function doesPhotoPositionAlreadyExist(chosenPhotoPositions, chosenPhotoPosition) {
            for (var i = 0; i < chosenPhotoPositions.length; i++) {
                if (chosenPhotoPositions[i] === chosenPhotoPosition) {
                    return true;
                }
            }
            return false;
        }

        function calculatePhotoGridPosition() {
            for (var i = 0; i < self.photos.length; i++) {
                var photo = self.photos[i];

                // columns, left-to-right, top-to-bottom
                var columns = 5;
                photo.column = photo.order % columns;
                photo.row = Math.floor(photo.order / columns);

                // rows, top-to-bottom, left-to-right
                // var rows = 3;
                // item.column = Math.floor(item.order/rows);
                // item.row = item.order%rows;
            }
        }

        function pickRandomPhoto() {
            var chosenPhotoPosition = Math.floor(Math.random() * self.photos.length);
            return self.photos[chosenPhotoPosition];
        }

        function shufflePhotos() {
            randomisePhotoPositions();
            calculatePhotoGridPosition();
        }

        $interval(function onComplete() {
            shufflePhotos();
        }, 1500);

        $interval(function onComplete() {
            updatePhotos();
        }, settings.refreshPhotoTimeInterval);

        $scope.$on('winning-photo:show', function onEvent() {
            self.showPickWinnerButton = false;
        });

        $scope.$on('winning-photo:hide', function onEvent() {
            self.showPickWinnerButton = true;
        });

        function pickWinningPhoto() {
            var winningPhoto = pickRandomPhoto();
            $modal.open({
                templateUrl: 'photo-gallery/modal/winning-photo/winning-photo.html',
                controller: 'WinningPhotoController',
                controllerAs: 'winningPhotoController',
                backdrop: 'static',
                resolve: {
                    winningPhoto: function () {
                        return winningPhoto;
                    }
                }
            }).result.then(function (result) {
                    winningPhoto.winner = true;
                    photoGalleryService.winningPhoto(winningPhoto).then(function () {
                        updatePhotos();
                    });

                });
        }

        function updatePhotos() {
            photoGalleryService.getPhotos().then(function (photos) {
                self.photos = photos;
            });
        }

    }

})();
