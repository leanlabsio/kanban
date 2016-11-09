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

        function getLabelString(stage) {
            var label = "KB[stage][" + stage.index +
                    "][" +
                    stage.viewName.trim() +
                    "]";

            if (stage.wip) {
                label = label + "[" + stage.wip +"]";
            }

            return label;
        }
        $scope.saving = false;
        BoardService.get($stateParams.project_path).then(function(board) {
            $scope.board = board;
            $scope.stages = board.stages;
            $scope.priorities = board.priorities;
            $scope.project_id = board.project.id;

            BoardService.listConnected(board.project.id).then(function(boards){
                $scope.boards = boards;
            });
        });

        $scope.update = function(stage) {
            $scope.saving = true;
            $scope.board.stale = true;
            LabelService.update($scope.project_id, stage.name, getLabelString(stage), "#fff")
                .then(function(res) {
                    $scope.saving = false;
                    stage.name = res.data.data.name;
                    stage.id = res.data.data.name;
                });
        };

        $scope.delete = function(index, stage) {
            $scope.saving = true;
            var label = getLabelString(stage);

            $scope.board.stale = true;
            if (_.isEmpty(stage.name)) {
                $scope.saving = false;
                $scope.stages.splice(index, 1);
            } else {
                return LabelService.delete($scope.project_id, stage.name).then(function(res) {
                    $scope.saving = false;
                    $scope.stages.splice(index, 1);
                });
            }
        };

        $scope.add = function(stage) {
            $scope.stages.push({
                index: 0,
                viewName: "",
                wip: ""
            });
        };

        $scope.create = function(index, stage) {
            var name = getLabelString(stage);
            $scope.saving = true;
            $scope.board.stale = true;
            return LabelService.create($scope.project_id, name, "#fff").then(function(res) {
                $scope.saving = false;
                stage.name = res.data.data.name;
                stage.id   = res.data.data.name;
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

        $scope.addBoard = function() {
            $scope.boards.push({
                id: 0,
                path_with_namespace: ""
            });
        };

        $scope.createConnectedBoard = function(board) {
            $scope.saving = true;
            $scope.board.stale = true;
            BoardService.connectBoard($scope.project_id, board.path_with_namespace).then(function(res){
                if (res) {
                    BoardService.boardConnected = {};
                }
                return BoardService.listConnected($scope.project_id);
            }).then(function(boards){
                $scope.boards = boards;
                $scope.saving = false;
            });
        };

        $scope.deleteConnectedBoard = function(board) {
            $scope.saving = true;
            $scope.board.stale = true;
            BoardService.deleteConnected($scope.project_id, board.id).then(function(res){
                if (res) {
                    BoardService.boardConnected = {};
                }
                return BoardService.listConnected($scope.project_id);
            }).then(function(boards){
                $scope.boards = boards;
                $scope.saving = false;
            });
        };
    }
}(window.angular));
