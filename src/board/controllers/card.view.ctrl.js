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
            'KBStore',
            'CardPriority',
            function($scope, $http, $stateParams, $state, BoardService, $sce, CommentService, LabelService, UserService, MilestoneService, $modal, host_url, store, CardPriority) {
                BoardService.get($stateParams.project_path).then(function(board) {
                    $scope.labels = _.toArray(board.viewLabels);
                    $scope.priorities = board.priorities;
                    $scope.board = board;

                    BoardService.listConnected(board.project.id).then(function(projects){
                        $scope.connected_projects = projects;
                    });

                    return BoardService.getCard($stateParams.project_path, $stateParams.path_with_namespace, $stateParams.issue_id);
                }).then(function(card) {
                    $scope.card = card;

                    CommentService.list(card.project_id, card.id).then(function(data) {
                        $scope.comments = data;
                    });

                    MilestoneService.list(card.project_id).then(function(milestones) {
                        $scope.milestones = milestones;
                    });

                    UserService.list(card.project_id).then(function(users) {
                        $scope.users = users;
                    });

                    $scope.commentFormData = store.get(getCommentHashKey()) || {};

                    $scope.submitComment = function() {
                        $scope.isSaving = true;
                        card.user_comments_count += 1;

                        CommentService.create(card.project_id, card.id, $scope.commentFormData.comment).then(function(data) {
                            $scope.isSaving = false;
                            $scope.commentFormData = {};
                            $scope.comments.push(data);
                            $scope.discardCommentDraft();
                        });
                    };
                });

                $scope.card_url = host_url + "/" + $stateParams.project_path;
                $scope.card_properties = {};
                $scope.commentFormData = {};
                $scope.blockedFormData = {};
                $scope.model = {};
                $scope.modal = $modal;
                $scope.todoFormData = {};
                $scope.newCard = {};
                $scope.default_priority = new CardPriority();

                var getCommentHashKey = function() {
                    return $scope.card.project_id + ":card:" + $scope.card.iid + ":comment";
                };

                var getCardHashKey = function() {
                    return $scope.card.project_id + ":card:" + $scope.card.iid;
                };

                $scope.changeProject = function(project){
                    return BoardService.changeProject($scope.board, $scope.card, project).then(function(card){
                        $state.go('board.cards.view', {
                            project_id: $stateParams.project_id,
                            project_path: $stateParams.project_path,
                            path_with_namespace: card.path_with_namespace,
                            issue_id: card.iid
                        });
                    });
                };

                $scope.discardCommentDraft = function() {
                    store.remove(getCommentHashKey());
                    $scope.commentFormData = {};
                };

                $scope.discardCardDraft = function() {
                    store.remove(getCardHashKey());
                };

                $scope.$watch('commentFormData', function(newV, oldV) {
                    if (oldV !== newV) {
                        store.set(getCommentHashKey(), newV);
                    }
                }, true);

                $scope.$watch('newCard', function(newV, oldV) {
                    if (oldV !== newV) {
                        store.set(getCardHashKey(), {
                            title: newV.title,
                            description: newV.description
                        });
                    }
                }, true);

                $scope.submitTodo = function(card) {
                    $scope.isSavingTodo = true;

                    card.todo.push({
                        'checked': false,
                        'body': $scope.todoFormData.body
                    });
                    BoardService.updateCard($scope.board, card).then(function() {
                        $scope.isSavingTodo = false;
                        $scope.todoFormData = {};
                        $scope.isTodoAdd = true;
                    });
                };

                $scope.remove = function(card) {
                    BoardService.removeCard($scope.board, card).then(function(result) {
                        $modal.close();
                    });
                };

                $scope.updateTodo = function(card) {
                    $scope.isSavingTodo = true;
                    return BoardService.updateCard($scope.board, card).then(function() {
                        $scope.isSavingTodo = false;
                    });
                };

                $scope.updateCard = function(card) {
                    $scope.newCard  = {};
                    $scope.isSaving = true;
                    return BoardService.updateCard($scope.board, card).then(function() {
                        $scope.isSaving = false;
                        $scope.discardCardDraft();
                    });
                };

                $scope.editCard = function(card){
                    var draft = store.get(getCardHashKey());
                    $scope.newCard = _.clone(card);
                    if (draft !== null) {
                        $scope.newCard.title = draft.title;
                        $scope.newCard.description = draft.description;
                    }
                };

                $scope.removeTodo = function(index, card) {
                    $scope.isSavingTodo = true;
                    card.todo.splice(index, 1);
                    return BoardService.updateCard($scope.board, card).then(function() {
                        $scope.isSavingTodo = false;
                    });
                };

                /**
                 * Update card assignee
                 */
                $scope.update = function(card, user) {
                    if (_.isEmpty(card.assignee) || card.assignee.id != user.id) {
                        card.assignee_id = user.id;
                        return BoardService.updateCard($scope.board, card);
                    } else {
                        card.assignee_id = 0;
                        return BoardService.updateCard($scope.board, card);
                    }
                };

                $scope.markAsBlocked = function(card, comment) {
                    CommentService.create(card.project_id, card.id, "Marked as **blocked**: " + comment).then(function(data) {
                        $scope.comments.push(data);
                    });

                    return BoardService.updateCard($scope.board, card);
                };

                $scope.markAsUnBlocked = function(card) {
                    if (card.properties.andon !== 'none') {
                        return;
                    }

                    var comment = 'Marked as **unblocked**';
                    CommentService.create(card.project_id, card.id, comment).then(function(data) {
                        $scope.comments.push(data);
                    });

                    return BoardService.updateCard($scope.board, card);
                };

                $scope.markAsReady = function (card) {
                    if (card.properties.andon === 'ready') {
                        CommentService.create(card.project_id, card.id, "Marked as **ready** for next stage").then(function(data) {
                            $scope.comments.push(data);
                        });
                    }

                    return BoardService.updateCard($scope.board, card);
                };

                $scope.updateMilestone = function(milestone) {
                    $scope.card.milestone = milestone;
                    return BoardService.updateCard($scope.board, $scope.card);
                };

                $scope.updateAssignee = function(user){
                    $scope.card.assignee = user;
                    return BoardService.updateCard($scope.board, $scope.card);
                };

                $scope.updateLabel = function(label) {
                    var card = $scope.card;
                    if (card.labels.indexOf(label.name) !== -1) {
                        card.labels.splice(card.labels.indexOf(label.name), 1);
                    } else {
                        card.labels.push(label.name);
                        card.viewLabels.push(label);
                    }

                    return BoardService.updateCard($scope.board, card);
                };

                $scope.updatePriority = function (priority) {
                    var card = $scope.card;
                    var index = card.labels.indexOf(card.priority.name);
                    if (index !== -1) {
                        card.labels.splice(index, 1);
                    }

                    if (_.isEmpty(card.priority.name) || card.priority.name != priority.name) {
                        if (! _.isEmpty(priority.name)) {
                          card.labels.push(priority.name);
                        }
                        card.priority = priority;
                    } else {
                        card.priority = LabelService.getPriority(card.project_id, "");
                    }

                    return BoardService.updateCard($scope.board, card);
                }
            }
        ]
    );
})(window.angular);
