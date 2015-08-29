(function(angular) {
    'use strict';
    angular.module('gitlabKBApp.board').controller('NewIssueController',
        [
            '$scope',
            '$http',
            '$stateParams',
            '$location',
            'BoardService',
            'UserService',
            'MilestoneService',
            'LabelService',
            '$modal', function($scope, $http, $stateParams, $location, BoardService, UserService, MilestoneService, LabelService, $modal) {
                $scope.isSaving = false;
                $scope.modal = $modal;

                BoardService.get($stateParams.project_path).then(function(board) {
                    UserService.list(board.project.id).then(function(users) {
                        $scope.options = users;
                    });

                    MilestoneService.list(board.project.id).then(function(milestones) {
                        $scope.milestones = milestones;
                    });
                });

                BoardService.get($stateParams.project_path).then(function(board) {
                    $scope.labels = _.toArray(board.viewLabels);
                    $scope.card = {
                        project_id: board.project.id,
                        labels: []
                    };
                });

                $scope.update = function(user) {
                    $scope.card.assignee = user;
                };

                $scope.updateMilestone = function(milestone) {
                    $scope.card.milestone = milestone;
                };

                $scope.updateLabels = function(label) {
                    if ($scope.card.labels.indexOf(label) !== -1) {
                        $scope.card.labels.splice($scope.card.labels.indexOf(label), 1);
                    } else {
                        $scope.card.labels.push(label);
                    }
                };

                $scope.createIssue = function() {
                    $scope.isSaving = true;

                    var data = {
                        project_id: $scope.card.project_id,
                        title: $scope.card.title,
                        description: $scope.card.description,
                    };

                    if (!_.isEmpty($scope.card.assignee)) {
                        data.assignee_id = $scope.card.assignee.id;
                    }
                    if (!_.isEmpty($scope.card.milestone)) {
                        data.milestone_id = $scope.card.milestone.id;
                    }

                    BoardService.getBoardById(data.project_id).then(function(board) {
                        var labels = [_.first(board.labels)];

                        if (!_.isEmpty($scope.card.labels)) {
                            for (var i = 0; i < $scope.card.labels.length; i++) {
                                labels.push($scope.card.labels[i].name);
                            }

                        }
                        data.labels = labels.join(', ');

                        BoardService.createCard(data).then(function() {
                            $modal.close();
                        });
                    });

                };
            }
        ]
    );
})(window.angular);

