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
                $scope.data = {
                    errors: []
                }
                $scope.isSaving = false;
                $scope.host_url = host_url;

                $scope.signup = function () {
                    $scope.isSaving = true;
                    AuthService.register($scope.user).then(function (result) {
                        $http.defaults.headers.common['X-KB-Access-Token'] = result;
                        $state.go('board.boards');
                    }, function (result) {
                          $scope.data.errors.push(result.data.message);
                          $scope.isSaving = false;
                    });
                };
            }
        ]
    );
})(window.angular);

