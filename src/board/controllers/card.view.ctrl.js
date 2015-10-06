(function(angular) {
    'use strict';

    angular.module('gitlabKBApp.board').controller('ViewController',
        [
            '$scope',
            '$http',
            '$stateParams',
            '$state',
            'BoardService',
            '$sce',
            'CommentService',
            'LabelService',
            'UserService',
            'MilestoneService',
            '$modal',
            'host_url',
            function($scope, $http, $stateParams, $state, BoardService, $sce, CommentService, LabelService, UserService, MilestoneService, $modal, host_url) {
                BoardService.get($stateParams.project_path).then(function(board) {
                    $scope.labels = _.toArray(board.viewLabels);
                });

                $scope.card_url = host_url + "/" + $stateParams.project_path;

                BoardService.getCard($stateParams.project_path, $stateParams.issue_id).then(function(card) {
                    $scope.card = card;

                    CommentService.list(card.project_id, card.id).then(function(data) {
                        $scope.comments = data;
                    });

                    $scope.submitComment = function() {
                        $scope.isSaving = true;

                        CommentService.create(card.project_id, card.id, $scope.commentFormData.comment).then(function(data) {
                            $scope.isSaving = false;
                            $scope.commentFormData = {};
                            $scope.comments.push(data);
                        });
                    };
                });

                BoardService.get($stateParams.project_path).then(function(board) {
                    MilestoneService.list(board.project.id).then(function(milestones) {
                        $scope.milestones = milestones;
                    });

                    UserService.list(board.project.id).then(function(users) {
                        $scope.options = users;
                    });
                });

                $scope.card_properties = {};
                $scope.commentFormData = {};
                $scope.blockedFormData = {};
                $scope.model = {};
                $scope.modal = $modal;
                $scope.todoFormData = {};

                $scope.submitTodo = function(card) {
                    $scope.isSavingTodo = true;

                    card.todo.push({
                        'checked': false,
                        'body': $scope.todoFormData.body
                    });
                    BoardService.updateCard(card).then(function() {
                        $scope.isSavingTodo = false;
                        $scope.todoFormData = {};
                        $scope.isTodoAdd = true;
                    });
                };

                $scope.remove = function(card) {
                    BoardService.removeCard(card).then(function(result) {
                        $modal.close();
                    });
                };

                $scope.updateTodo = function(card) {
                    $scope.isSavingTodo = true;
                    return BoardService.updateCard(card).then(function() {
                        $scope.isSavingTodo = false;
                    });
                };

                $scope.updateCard = function(card) {
                    $scope.newCard  = undefined;
                    $scope.isSaving = true;
                    return BoardService.updateCard(card).then(function() {
                        $scope.isSaving = false;
                    });
                };

                $scope.editCard = function(card){
                  $scope.newCard = _.clone(card);
                };

                $scope.removeTodo = function(index, card) {
                    $scope.isSavingTodo = true;
                    card.todo.splice(index, 1);
                    return BoardService.updateCard(card).then(function() {
                        $scope.isSavingTodo = false;
                    });
                };

                $scope.update = function(card, user) {
                    if (!card.assignee || (card.assignee.id != user.id)) {
                        card.assignee_id = user.id;
                        return BoardService.updateCard(card);
                    }
                };

                $scope.markAsBlocked = function(card, comment) {
                    CommentService.create(card.project_id, card.id, "Marked as **blocked**: " + comment).then(function(data) {
                        $scope.comments.push(data);
                    });

                    return BoardService.updateCard(card);
                };

                $scope.markAsUnBlocked = function(card) {
                    if (card.properties.andon !== 'none') {
                        return;
                    }

                    var comment = 'Marked as **unblocked**';
                    CommentService.create(card.project_id, card.id, comment).then(function(data) {
                        $scope.comments.push(data);
                    });

                    return BoardService.updateCard(card);
                };

                $scope.markAsReady = function (card) {
                    if (card.properties.andon === 'ready') {
                        CommentService.create(card.project_id, card.id, "Marked as **ready** for next stage").then(function(data) {
                            $scope.comments.push(data);
                        });
                    }

                    return BoardService.updateCard(card);
                };

                $scope.updateMilestone = function(card, milestone) {
                    if (!card.milestone || (card.milestone.id != milestone.id)) {
                        card.milestone_id = milestone.id;
                        return BoardService.updateCard(card);
                    }
                };

                $scope.updateLabels = function(card, label) {
                    BoardService.get($stateParams.project_path).then(function(board) {
                        if (card.labels.length === card.viewLabels.length) {
                            card.labels.push(_.first(board.labels));
                        }

                        if (card.labels.indexOf(label.name) !== -1) {
                            card.viewLabels.splice(card.viewLabels.indexOf(label), 1);
                            card.labels.splice(card.labels.indexOf(label.name), 1);
                        } else {
                            card.viewLabels.push(label);
                            card.labels.push(label.name);
                        }

                        return BoardService.updateCard(card);
                    });
                };
            }
        ]
    );
})(window.angular);
