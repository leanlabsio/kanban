(function(angular, CLIENT_VERSION, GITLAB_HOST, enableSignup) {
    'use strict';

    var app = angular.module('gitlabKBApp', [
            'ui.router',
            'gitlabKBApp.user',
            'gitlabKBApp.board',
            'angular-loading-bar',
            'angular-lodash',
            'mm.foundation.topbar',
            'angular-storage'
        ])
        .run([
            '$rootScope', '$state', '$http', 'AuthService', 'store',
            function($rootScope, $state, $http, AuthService, store) {
                if (AuthService.isAuthenticated()) {
                    $http.defaults.headers.common['X-KB-Access-Token'] = AuthService.getCurrent();
                }

                $rootScope.$on('$stateChangeStart', function(event, toState, toParams) {
                    if (!AuthService.authorized(toState)) {
                        event.preventDefault();
                        if (!store.get('state')) {
                            store.set('state', {
                                name: toState.name,
                                params: toParams
                            });
                        }
                        $state.go('login');
                    }
                });
            }
        ])
        .constant('host_url', GITLAB_HOST)
        .constant('version', CLIENT_VERSION)
        .constant('enable_signup', enableSignup)
        .config(
            [
                '$locationProvider',
                function($locationProvider) {
                    $locationProvider.html5Mode(true);
                }
            ]
        )
        .factory("KBStore", ["store", function(store) {
            return store.getNamespacedStore("kb", "localStorage", ":");
        }]);
})(window.angular, window.CLIENT_VERSION, window.GITLAB_HOST, window.ENABLE_SIGNUP);
