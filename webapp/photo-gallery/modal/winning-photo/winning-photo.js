(function () {

    angular.module('photoGallery').controller('WinningPhotoController', WinningPhotoController);

    function WinningPhotoController($scope, $rootScope, winningPhoto) {

        var self = this;
        self.closeModal = closeModal;
        self.winningPhotoFile = winningPhoto.filename;

        $rootScope.$broadcast('winning-photo:show');

        function closeModal(close) {
            $rootScope.$broadcast('winning-photo:hide');
            close();
        }

    }

})();
