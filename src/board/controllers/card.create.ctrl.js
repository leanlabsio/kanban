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
            '$modal',
            'KBStore',
            function($scope, $http, $stateParams, $location, BoardService, UserService, MilestoneService, LabelService, $modal, store) {
                $scope.isSaving = false;
                $scope.modal = $modal;
                $scope.card = {};

                var getHashKey = function() {
                    return $scope.card.project_id + ":card:new";
                };

                BoardService.get($stateParams.project_path).then(function(board) {
                    UserService.list(board.project.id).then(function(users) {
                        $scope.options = users;
                    });

                    MilestoneService.list(board.project.id).then(function(milestones) {
                        $scope.milestones = milestones;
                    });

                    $scope.labels = _.toArray(board.viewLabels);
                    $scope.priorities = board.priorities;
                    $scope.card = {
                        project_id: board.project.id,
                        labels: []
                    };

                    var card = store.get(getHashKey());
                    if (card !== null) {
                        $scope.card = card;

                        if (!_.isEmpty(card.labels)) {
                            var labels = card.labels.slice(0);
                            angular.forEach(labels, function(value){
                                var newLabel = _.findWhere($scope.labels, {name: value.name});
                                if (!_.isEmpty(newLabel)){
                                    $scope.updateLabels(value);
                                    $scope.updateLabels(newLabel);
                                }
                            });
                        }
                    }
                });

                $scope.update = function(user) {
                    $scope.card.assignee = user;
                };

                $scope.updateMilestone = function(milestone) {
                    $scope.card.milestone = milestone;
                };

                $scope.updatePriority = function(priority) {
                    $scope.card.priority = priority;
                };

                $scope.updateLabels = function(label) {
                    if ($scope.card.labels.indexOf(label) !== -1) {
                        $scope.card.labels.splice($scope.card.labels.indexOf(label), 1);
                    } else {
                        $scope.card.labels.push(label);
                    }
                };

                $scope.$watch('card', function(newV, oldV) {
                    if (oldV !== newV) {
                        store.set(getHashKey(), newV);
                    }
                }, true);

                $scope.cancelCreate = function() {
                        $modal.close();
                        store.remove(getHashKey());
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
                        var labels = [_.first(board.stagelabels)];

                        if (!_.isEmpty($scope.card.labels)) {
                            for (var i = 0; i < $scope.card.labels.length; i++) {
                                labels.push($scope.card.labels[i].name);
                            }

                        }

                        if (!_.isEmpty($scope.card.priority)) {
                            labels.push($scope.card.priority.name);
                        }

                        data.labels = labels.join(', ');

                        BoardService.createCard(data).then(function() {
                            $modal.close();
                            store.remove(getHashKey());
                        });
                    });
                };
            }
        ]
    );
})(window.angular);

