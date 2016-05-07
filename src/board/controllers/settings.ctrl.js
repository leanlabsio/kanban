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
                st[1] = parseInt(st[1]);
                return st;
            });

            $scope.stages = stages;
            $scope.priorities = board.priorities;
            $scope.project_id = board.project.id;
        });

        $scope.update = function(stage) {
            $scope.saving = true;
            var oldLabel = stage[0];
            var newLabel = 'KB[stage][' + stage[1] + '][' + stage[2].trim() + ']';
            $scope.board.stale = true;
            LabelService.update($scope.project_id, oldLabel, newLabel, "#fff")
                .then(function(res) {
                    $scope.saving = false;
                    stage[0] = res.data.data.name;
                });
        };

        $scope.delete = function(index, stage) {
            $scope.saving = true;
            var label = "KB[stage][" + stage[1] + "][" + stage[2].trim() + "]";

            $scope.board.stale = true;
            if (!_.isEmpty(stage[3])) {
                $scope.saving = false;
                $scope.stages.splice(index, 1);
            } else {
                return LabelService.delete($scope.project_id, stage[0]).then(function(res) {
                    $scope.saving = false;
                    $scope.stages.splice(index, 1);
                });
            }
        };

        $scope.add = function(stage) {
            $scope.stages.push(["", "", "", "1"]);
        };

        $scope.create = function(index, stage) {
            var name = "KB[stage][" + stage[1] + "][" + stage[2].trim() + "]";
            $scope.saving = true;
            $scope.board.stale = true;
            return LabelService.create($scope.project_id, name, "#fff").then(function(res) {
                $scope.saving = false;
                $scope.stages[index] = stage_regexp.exec(res.data.data.name);
                $scope.stages[index][1] = parseInt($scope.stages[index][1]);
            });
        };

        $scope.addPriority = function() {
          $scope.priorities.push({
              name: "",
              index: 0,
              color: "#ffffff",
              viewName: ""
          });
        };

        $scope.createPriority = function(priority) {
            $scope.saving = true;
            var name = "KB[priority][" + priority.index + "][" + priority.viewName.trim() + "]";
            LabelService.create($scope.project_id, name, priority.color)
                .then(function(res) {
                    $scope.saving = false;
                    priority.name = res.data.data.name;
                });
        };

        $scope.updatePriority = function(priority) {
            $scope.saving = true;
            var name = "KB[priority][" + priority.index + "][" + priority.viewName.trim() + "]";
            LabelService.update($scope.project_id, priority.name, name, priority.color)
                .then(function(res) {
                    $scope.saving = false;
                    priority.name = res.data.data.name;
                });
        };

        $scope.deletePriority = function(priority) {
            if (_.isEmpty(priority.name)) {
                $scope.priorities.splice($scope.priorities.indexOf(priority), 1);
            } else {
                $scope.saving = true;
                LabelService.delete($scope.project_id, priority.name)
                    .then(function(res) {
                        $scope.saving = false;
                        $scope.priorities.splice($scope.priorities.indexOf(priority), 1);
                    });
            }
        };
    }
}(window.angular));
