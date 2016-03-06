(function(angular) {
    'use strict';

    angular.module('gitlabKBApp.board').controller(
        'ConfigurationController',
        [
            '$scope',
            '$http',
            '$state',
            '$stateParams',
            'BoardService', function($scope, $http, $state, $stateParams, BoardService) {
                $scope.isSaving = false;

                $scope.configure = function() {
                    $scope.isSaving = true;
                    $http.post(
                        '/api/boards/configure',
                        {
                            project_id: $stateParams.project_path
                        }
                    ).then(function(result) {
                        if (!_.isEmpty(BoardService.boards[$stateParams.project_path])) {
                            BoardService.boards[$stateParams.project_path].stale = true;
                        }
                        BoardService.get($stateParams.project_path).then(function(board) {
                            $scope.isSaving = false;
                            $state.go('board.cards', {
                                project_id: $stateParams.project_id,
                                project_path: $stateParams.project_path
                            });
                        });
                    });
                };
            }
        ]
    );
})(window.angular);
