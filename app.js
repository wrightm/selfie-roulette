angular.module('selfieRoulette', ['ui.bootstrap', 'ui.utils', 'ui.router', 'ngResource', 'ngAnimate', 'photoGallery', 'photoBooth']);

angular.module('selfieRoulette').config(function ($stateProvider, $urlRouterProvider) {

    /* Add New States Above */
    $urlRouterProvider.otherwise('/photo-gallery');

});

angular.module('selfieRoulette').run(function ($rootScope) {

    $rootScope.safeApply = function (fn) {
        var phase = $rootScope.$$phase;
        if (phase === '$apply' || phase === '$digest') {
            if (fn && (typeof(fn) === 'function')) {
                fn();
            }
        } else {
            this.$apply(fn);
        }
    };

});
