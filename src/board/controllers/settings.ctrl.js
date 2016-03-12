(function(angular) {
    'use strict';
    angular.module('gitlabKBApp.board').controller('SettingsController', [
        'LabelService',
        'BoardService',
        '$stateParams',
        '$scope',
        'stage_regexp',
        '$http',
        SettingsController
    ]);

    function SettingsController(LabelService, BoardService, $stateParams, $scope, stage_regexp, $http) {
        $scope.saving = false;
        BoardService.get($stateParams.project_path).then(function(board) {
            $scope.board = board;
            var stages = _.map(board.stages, function(stage) {
                var st = stage_regexp.exec(stage.name);
                st[0] = stage.name;
                return st;
            });

            $scope.stages = stages;
            $scope.project_id = board.project.id;
        });

        $scope.update = function(stage) {
            $scope.saving = true;
            var oldLabel = stage[0];
            var newLabel = 'KB[stage][' + stage[1].trim() + '][' + stage[2].trim() + ']';
            $scope.board.stale = true;
            $http.put('/api/labels/' + $scope.project_id, {
                name: oldLabel,
                new_name: newLabel
            }).then(function(res) {
                $scope.saving = false;
                stage[0] = res.data.data.name;
            });
        };

        $scope.delete = function(index, stage) {
            $scope.saving = true;
            var label = "KB[stage][" + stage[1].trim() + "][" + stage[2].trim() + "]";

            $scope.board.stale = true;
            if (!_.isEmpty(stage[3])) {
                $scope.saving = false;
                $scope.stages.splice(index, 1);
            } else {
                return $http.delete("/api/labels/" + $scope.project_id + "/" + stage[0]).then(function(res) {
                    $scope.saving = false;
                    $scope.stages.splice(index, 1);
                });
            }
        };

        $scope.add = function(stage) {
            $scope.stages.push(["", "", "", "1"]);
        };

        $scope.create = function(index, stage) {
            $scope.saving = true;
            var data = {
                name: "KB[stage][" + stage[1].trim() + "][" + stage[2].trim() + "]",
                color: "#fff"
            };
            $scope.board.stale = true;
            return $http.post("/api/labels/" + $scope.project_id, data).then(function(res) {
                $scope.saving = false;
                $scope.stages[index] = stage_regexp.exec(res.data.data.name);
            });
        };
    }
}(window.angular));
