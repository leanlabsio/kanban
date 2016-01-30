(function(angular, CLIENT_VERSION) {
    'use strict';

    angular.module('gitlabKBApp.user', ['ui.router', 'angular-storage']).config(['$stateProvider', '$urlRouterProvider', function($stateProvider, $urlRouterProvider) {
        $stateProvider.decorator('views', function(state, parent) {
            var result = {},
                views = parent(state);

            angular.forEach(views, function(config, name) {
                config.templateUrl = CLIENT_VERSION + "/" + config.templateUrl;
                result[name] = config;
            });

            return result;
        });

        $stateProvider
            .state('login', {
                url: '/',
                templateUrl: 'assets/html/user/views/signin.html',
                controller: 'SigninController',
                data: {
                    access: 0
                }
            })
            .state('signup', {
                url: '/signup',
                templateUrl: 'assets/html/user/views/signup.html',
                controller: 'SignupController',
                data: {
                    access: 0
                }
            });
        $urlRouterProvider.otherwise('/');
    }]).config(['$httpProvider', function($httpProvider) {
        $httpProvider.interceptors.push(['$q', '$injector', 'store', function($q, $injector, store) {
            return {
                response: function(response) {
                    if (response.status === 401) {}
                    return response || $q.when(response);
                },
                responseError: function(rejection) {
                    var $state;
                    if (rejection.status === 401 || rejection.status === undefined) {
                        $state = $injector.get('$state');
                        store.remove('id_token');
                        $httpProvider.defaults.headers.common['X-KB-Access-Token'] = '';
                        if (!store.get('state')) {
                            store.set('state', {
                                name: $state.current.name,
                                params: $state.params
                            });
                        }
                        window.location.pathname = '/';
                    }
                    if (rejection.status === 403) {
                        alert('Access denied');
                    }

                    return $q.reject(rejection);
                }
            };
        }]);
    }]);
})(window.angular, window.CLIENT_VERSION);
