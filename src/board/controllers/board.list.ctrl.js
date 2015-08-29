(function(angular) {
    'use strict';

    angular.module('gitlabKBApp.board').controller('BoardListController',
        [
            '$scope',
            'host_url',
            'BoardService', function($scope, host_url, BoardService) {

                $scope.host_url = host_url;

                BoardService.getBoards().then(function(result) {
                    $scope.boards = _.groupBy(result, function(val) {
                        return val.namespace.name;
                    });
                });

                $scope.setVisible = function(key) {
                    $scope.boards[key].visible = ($scope.boards[key].visible == undefined || $scope.boards[key].visible != true);
                };
            }
        ]
    );
})(window.angular);

