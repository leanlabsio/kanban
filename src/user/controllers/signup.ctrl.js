(function() {
    'use strict';

    angular.module('gitlabKBApp.user').controller('SignupController', 
        [
            '$scope', 
            '$http', 
            '$state', 
            'AuthService',
            'host_url',
            function ($scope, $http, $state, AuthService, host_url) {
                $scope.user = {};
                $scope.isSaving = false;
                $scope.host_url = host_url;

                $scope.signup = function () {
                    $scope.isSaving = true;
                    AuthService.register($scope.user).then(function (result) {
                        AuthService.authenticate($scope.user).then(function(auth) {
                            $http.defaults.headers.common['X-KB-Access-Token'] = auth;
                            $state.go('board.boards');
                        });
                    });
                };
            }
        ]
    );
})(window.angular);

