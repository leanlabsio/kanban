(function(window, angular) {
    'use strict';

    angular.module('gitlabKBApp.user').controller('SigninController', [
        '$scope',
        '$http',
        '$state',
        'AuthService',
        'store',
        'host_url',
        'enable_signup',
        function($scope, $http, $state, AuthService, store, host_url, enable_signup) {
            var goToKanban = function() {
                var state = store.get('state');
                if (state) {
                    $state.go(state.name, state.params);
                    store.remove('state');
                } else {
                    $state.go('board.boards');
                }
            };

            if (AuthService.isAuthenticated()) {
                goToKanban();
            }

            $scope.host_url = host_url;
            $scope.enable_signup = enable_signup;

            $scope.data = {
                signin: {},
                isSaving: false,
                errors: []
            };
            $scope.oauth = function() {
                var authWindow = window.open('/api/oauth?provider=gitlab', 'Auth', "menubar=no,location=0,resizable=yes,scrollbars=yes,status=0");
                var listener = function(event) {
                    authWindow.close();
                    window.removeEventListener('message', listener);
                    $http.post('/api/oauth', {
                        code: event.data,
                        provider: "gitlab"
                    }).then(function(result) {
                        store.set('id_token', result.data.token);
                        $http.defaults.headers.common['X-KB-Access-Token'] = result.data.token;
                        goToKanban();
                    });
                };

                window.addEventListener('message', listener);
            };

            $scope.authenticate = function() {
                $scope.data.errors = [];
                $scope.data.isSaving = true;

                AuthService.authenticate({
                    username: $scope.data.signin.username,
                    password: $scope.data.signin.password
                }).then(function(result) {
                    $http.defaults.headers.common['X-KB-Access-Token'] = result;
                    goToKanban();
                }, function(result) {
                    $scope.data.errors.push(result.data.message);
                    $scope.data.isSaving = false;
                });
            };
        }
    ]);
})(window, window.angular);
