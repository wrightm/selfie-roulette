(function () {

    angular.module('photoGallery').controller('WinningPhotoController', WinningPhotoController);

    function WinningPhotoController($scope, $rootScope, winningPhotoFile) {

        var self = this;
        self.closeModal = closeModal;
        self.winningPhotoFile = winningPhotoFile;

        $rootScope.$broadcast('winning-photo:show');

        function closeModal(close) {
            console.log('dsdsd');
            $rootScope.$broadcast('winning-photo:hide');
            close();
        }

    }

})();
