(function(angular) {
    'use strict';

    angular.module('gitlabKBApp.board').controller('TopBarController', [
        '$scope',
        '$state',
        '$stateParams',
        'BoardService',
        'AuthService',
        '$window',
        function($scope, $state, $stateParams, BoardService, AuthService, $window) {
            if ($stateParams.project_path !== undefined) {
                BoardService.get($stateParams.project_path).then(function(board) {
                    $scope.project = board.project;
                });
            }

            $scope.stateParams = $stateParams;

            $scope.logout = function() {
                AuthService.logout();
                $window.location.pathname = '/';
            };

            $scope.showActionBar = function() {
                BoardService.get($stateParams.project_path).then(function(board) {
                    board.state.showActionBar = !board.state.showActionBar;
                });
            };

            $scope.reset = function() {
                var params = {
                    project_id: $stateParams.project_id,
                    project_path: $stateParams.project_path,
                    group: ''
                };
                $state.go('board.cards', params);
            };

            $scope.group = function(field) {
                var params = {
                    project_id: $stateParams.project_id,
                    project_path: $stateParams.project_path,
                    group: field
                };
                $state.go('board.cards', params);
            };
        }
    ]);
}(window.angular));
