(function(angular) {
    'use strict';

    angular.module('gitlabKBApp.board').controller('BoardController',
        [
            '$scope',
            '$http',
            '$stateParams',
            'BoardService',
            '$state',
            '$window',
            'UserService',
            'stage_regexp',
            '$rootScope',
            'WebsocketService', function($scope, $http, $stateParams, BoardService, $state, $window, UserService, stage_regexp, $rootScope, WebsocketService) {
                $window.scrollTo(0, 0);

                var filter = function(item) {return true;};
                var group = function(item) {return 'none';};
                var grouped = !!$stateParams.group;

                if (!!$stateParams.assignee) {
                    var username = $stateParams.assignee.replace(/@/, '');
                    filter = function(item) {
                        if (!item.assignee) {
                            return false;
                        }

                        return item.assignee.username == username;
                    };
                }

                if (!!$stateParams.milestone) {
                    var milestone = $stateParams.milestone.replace(/\^/, '');
                    filter = function(item) {
                        if (_.isEmpty(item.milestone)) {
                            return false;
                        }
                        return item.milestone.title == milestone;
                    };
                }

                if (!!$stateParams.label) {
                    var label = $stateParams.label.replace(/~/, '');
                    filter = function(item) {
                        return item.labels.indexOf(label) != -1;
                    };
                }

                if (grouped) {
                    group = function (item) {
                        if (!item.assignee) {
                            return 'Unassigned';
                        }

                        return item.assignee.name;
                    };
                }

                BoardService.get($stateParams.project_path).then(function(board) {
                    if (_.isEmpty(board.labels)) {
                        $state.go('board.import', {
                            project_id: board.project.id,
                            project_path: board.project.path_with_namespace
                        });
                    }

                    board.grouped = grouped;

                    var issues = _.filter(board.issues, filter);
                    var groups = _.groupBy(issues, group);
                    groups = _.isEmpty(groups) ? {none:{}} : groups;
                    groups = _.each(groups, board.byStage, board);

                    $scope.groups = groups;
                    $scope.board = board;

                    $rootScope.$on('board.change', function() {
                        var issues = _.filter(board.issues, filter);
                        var groups = _.groupBy(issues, group);
                        groups = _.isEmpty(groups) ? {none:{}} : groups;
                        groups = _.each(groups, board.byStage, board);
                        $scope.groups = groups;
                    });

                    $scope.dragControlListeners = {
                        accept: function() {return true;},
                        dragEnd: function(event) {
                            var id = event.source.itemScope.card.id;
                            var card = board.getCardById(id);

                            var oldLabel = event.source.sortableScope.$parent.stageName;
                            var newLabel = event.dest.sortableScope.$parent.stageName;

                            var oldGroup = event.source.sortableScope.$parent.$parent.groupName;
                            var newGroup = event.dest.sortableScope.$parent.$parent.groupName;

                            card.labels = _.filter(card.labels, function (label) {
                                return !stage_regexp.test(label);
                            });

                            card.labels.push(newLabel);
                            card.stage = newLabel;
                            card.properties.andon = 'none';

                            var data = {
                                project_id: card.project_id,
                                issue_id: card.id,
                                title: card.title,
                                labels: card.labels.join(', '),
                                todo: card.todo,
                                description: card.description,
                                properties: card.properties,
                                stage: {
                                    source: oldLabel,
                                    dest: newLabel
                                }
                            };

                            if (oldGroup != newGroup) {
                                return UserService.findByName(card.project_id, newGroup).then(function (user) {
                                    data.assignee_id = user === undefined ? 0 : user.id;
                                    card.assignee = user;
                                    return $http.put('/api/card/move', data);
                                });
                            } else {
                                return $http.put('/api/card/move', data);
                            }
                        },
                        containment: '#board'
                    };


                    WebsocketService.emit('subscribe', {
                        routing_key: 'kanban.' + board.project.id.toString()
                    });
                });

                $scope.state = $state;

                $scope.remove = function(card) {
                    BoardService.removeCard(card).then(function(result) {});
                };
            }
        ]
    );
})(window.angular);
