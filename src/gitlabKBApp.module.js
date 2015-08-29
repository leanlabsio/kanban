(function(angular) {
    'use strict';

    var app = angular.module('gitlabKBApp', 
        [
            'ui.router',
            'gitlabKBApp.user',
            'gitlabKBApp.board',
            'angular-loading-bar',
            'angular-underscore',
            'mm.foundation.topbar'
        ])
        .run(['$rootScope', '$state', '$http', 'AuthService', 'store', function ($rootScope, $state, $http, AuthService) {
            if (AuthService.isAuthenticated()) {
                $http.defaults.headers.common['X-KB-Access-Token'] = AuthService.getCurrent();
            }

            $rootScope.$on('$stateChangeStart', function(event, toState, toParams) {
                if (!AuthService.authorized(toState)) {
                    event.preventDefault();
                    $state.go('login');
                }
            });
        }])
        .constant('host_url', 'GITLAB_HOST')
        .config(
            [
                '$locationProvider',
                function($locationProvider) {
                    $locationProvider.html5Mode(true);
                }
            ]
        );
})(window.angular);

