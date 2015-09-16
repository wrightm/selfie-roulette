(function () {

    angular.module('photoGallery').controller('MainPhotoGalleryController', mainPhotoGalleryController);


    function mainPhotoGalleryController($scope, $interval, $modal, photos, photoGalleryService, settings, countdownTimerService) {

        var self = this;
        self.photos = photos;
        self.doShufflePhotos = true;

        countdownTimerService.setCount(settings.winnerTimeInterval);
        countdownTimerService.start();

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
            return self.photos[chosenPhotoPosition].filename;
        }

        function shufflePhotos() {
            randomisePhotoPositions();
            calculatePhotoGridPosition();
        }

        $interval(function onComplete() {
            if (self.doShufflePhotos) {
                shufflePhotos();
            }
        }, 1500);

        $interval(function onComplete() {
            photoGalleryService.getPhotos().then(function (photos) {
                self.photos = photos;
            });
        }, settings.refreshPhotoTimeInterval);

        $scope.$on('winning-photo:show', function onEvent() {
            self.doShufflePhotos = true;
            countdownTimerService.stop();
        });

        $scope.$on('winning-photo:hide', function onEvent() {
            self.doShufflePhotos = true;
            countdownTimerService.start();
        });

        $scope.$watch(function watch() {
            return countdownTimerService.getCount().count;
        }, function onChange() {
            /*if(countdownTimerService.getCount().count === 0) {
             $modal.open({
             templateUrl: 'photo-gallery/modal/winning-photo/winning-photo.html',
             controller: 'WinningPhotoController',
             controllerAs : 'winningPhotoController',
             backdrop : 'static',
             resolve : {
             winningPhotoFile : function(){
             return pickRandomPhoto();
             }

             }
             }).result.then(function (result) {
             //do something with the result
             });
             }*/
        });

    }

})();
