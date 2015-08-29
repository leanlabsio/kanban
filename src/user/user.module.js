(function(angular) {
    'use strict';

    angular.module('gitlabKBApp.user', ['ui.router', 'angular-storage']).config(['$stateProvider', '$urlRouterProvider', function ($stateProvider, $urlRouterProvider) {
        $stateProvider
            .state('login', {
                url: '/',
                templateUrl: 'assets/v1.2.0/html/user/views/signin.html',
                controller: 'SigninController',
                data: {
                    access: 0
                }
            })
            .state('signup', {
                url: '/signup',
                templateUrl: 'assets/v1.2.0/html/user/views/signup.html',
                controller: 'SignupController',
                data: {
                    access: 0
                }
            });
        $urlRouterProvider.otherwise('/');
    }]).config(['$httpProvider',function($httpProvider) {
        $httpProvider.interceptors.push(['$q', '$location', 'store', function($q, $location, store) {
            return {
                response: function(response){
                    if (response.status === 401) {
                    }
                    return response || $q.when(response);
                },
                responseError: function(rejection) {
                    if (rejection.status === 401) {
                        store.remove('id_token');
                        $httpProvider.defaults.headers.common['X-KB-Access-Token'] = '';
                        $location.path('/');
                    }
                    if (rejection.status === 403) {
                        alert('Access denied');
                    }

                    return $q.reject(rejection);
                }
            }
        }]);
    }]);
})(window.angular);
