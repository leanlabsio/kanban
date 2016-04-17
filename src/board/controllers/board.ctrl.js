(function(angular) {
    'use strict';

    angular.module('gitlabKBApp.board').controller('BoardController', [
        '$scope',
        '$http',
        '$stateParams',
        'BoardService',
        '$state',
        '$window',
        'UserService',
        'stage_regexp',
        '$rootScope',
        'WebsocketService',
        'MilestoneService',
        function($scope, $http, $stateParams, BoardService, $state, $window, UserService, stage_regexp, $rootScope, WebsocketService, MilestoneService) {
            $window.scrollTo(0, 0);

            var filter = function(item) {
                return true;
            };
            var group = function(item) {
                return 'none';
            };
            var grouped = $stateParams.group;

            var tags = [];
            if ($stateParams.tags) {
                tags = tags.concat($stateParams.tags);
            }

            if (!_.isEmpty(tags)) {
                var fByUser = false,
                    fByMilestone = false,
                    fByLabel = false;

                fByUser = _.some(tags, function(v) {
                    return v.indexOf('@') === 0;
                });
                fByMilestone = _.some(tags, function(v) {
                    return v.indexOf('^') === 0;
                });
                fByLabel = _.some(tags, function(v) {
                    return v.indexOf('~') === 0;
                });

                filter = function(item) {
                    var item_tags = [],
                        uMatch = true,
                        mMatch = true,
                        lMatch = true;

                    if (fByUser) {
                        var uName = _.isEmpty(item.assignee) ? '@' : '@' + item.assignee.id;
                        uMatch = _.contains(tags, uName);
                    }

                    if (fByMilestone) {
                        var ms = _.isEmpty(item.milestone) ? '^' : '^' + item.milestone.id;
                        mMatch = _.contains(tags, ms);
                    }

                    if (fByLabel) {
                        var labels = ['~'];

                        if (!_.isEmpty(item.viewLabels)) {
                            labels = _.map(item.viewLabels, function(l) {
                                return '~' + l.name;
                            });
                        }

                        lMatch = _.intersection(tags, labels).length > 0;
                    }

                    return uMatch && mMatch && lMatch;
                };
            }

            if (grouped) {
                group = function(item) {
                    if (grouped == 'milestone') {
                        if (_.isEmpty(item.milestone)) {
                            return 'No Milestone';
                        }

                        return item.milestone.title;
                    } else if (grouped == 'user') {
                        if (!item.assignee) {
                            return 'Unassigned';
                        }

                        return item.assignee.name;
                    }
                };
            }

            BoardService.get($stateParams.project_path).then(function(board) {
                if (_.isEmpty(board.stages)) {
                    $state.go('board.import', {
                        project_id: board.project.id,
                        project_path: board.project.path_with_namespace
                    });
                }

                board.grouped = grouped;

                $scope.groups = board.reset(filter, group);
                $scope.board = board;

                $rootScope.$on('board.change', function() {
                    $scope.groups = board.reset(filter, group);
                });

                $scope.dragControlListeners = {
                    accept: function() {
                        return true;
                    },
                    dragEnd: function(event) {
                        var id = event.source.itemScope.card.id;
                        var card = board.getCardById(id);

                        var oldLabel = event.source.sortableScope.$parent.stageName;
                        var newLabel = event.dest.sortableScope.$parent.stageName;

                        var oldGroup = event.source.sortableScope.$parent.$parent.groupName;
                        var newGroup = event.dest.sortableScope.$parent.$parent.groupName;

                        card.labels = _.filter(card.labels, function(label) {
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
                            if (grouped == 'milestone') {
                                return MilestoneService.findByName(card.project_id, newGroup).then(function(milestone) {
                                    data.milestone_id = milestone === undefined ? 0 : milestone.id;
                                    card.milestone = milestone;
                                    return $http.put('/api/card/move', data);
                                });
                            } else if (grouped == 'user') {
                                return UserService.findByName(card.project_id, newGroup).then(function(user) {
                                    data.assignee_id = user === undefined ? 0 : user.id;
                                    card.assignee = user;
                                    return $http.put('/api/card/move', data);
                                });
                            }
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
    ]);
})(window.angular);
