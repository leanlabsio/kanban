(function (angular) {
    'use strict';
    angular.module('gitlabKBApp.board').controller('NewMilestoneController',
        [
            '$scope',
            '$http',
            '$stateParams',
            '$location',
            'BoardService',
            'MilestoneService',
            '$modal', function ($scope, $http, $stateParams, $location, BoardService, MilestoneService, $modal) {
            $scope.isSaving = false;
            $scope.modal = $modal;

            BoardService.get($stateParams.project_path).then(function (board) {
                $scope.milestone = {
                    project_id: board.project.id
                };
            });

            $scope.createMilestone = function () {
                $scope.isSaving = true;

                var data = {
                    project_id: $scope.milestone.project_id,
                    title: $scope.milestone.title,
                    description: $scope.milestone.description
                };

                if (!_.isEmpty($scope.milestone.due_date)) {
                    data.due_date = $scope.milestone.due_date;
                }

                MilestoneService.create(data).then(function () {
                    $modal.close();
                });

            };
        }
        ]
    );
}(window.angular));

