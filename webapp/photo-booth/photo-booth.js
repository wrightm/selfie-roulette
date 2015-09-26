angular.module('photoBooth', ['ui.bootstrap', 'ui.utils', 'ui.router', 'ngAnimate']);

angular.module('photoBooth').config(function ($stateProvider) {

    $stateProvider.state('main', {
        url: 'photo-booth',
        templateUrl: 'photo-booth/partial/main/main.html'
    });
    /* Add New States Above */

});

