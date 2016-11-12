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
        'LabelService',
        function($scope, $http, $stateParams, BoardService, $state, $window, UserService, stage_regexp, $rootScope, WebsocketService, MilestoneService, LabelService) {
            $window.scrollTo(0, 0);

            var filter = function(item) {
                return true;
            };
            var group = function(item) {
                return 0;
            };

            $scope.groupLabels = [{
                id: 0,
                name: "none"
            }];

            $scope.showDetails = {};

            var grouped = $stateParams.group;

            var tags = [];
            if ($stateParams.tags) {
                tags = tags.concat($stateParams.tags);
            }

            if (!_.isEmpty(tags)) {
                var fByUser = false,
                    fByMilestone = false,
                    fByLabel = false,
                    fByPriority = false;

                fByUser = _.some(tags, function(v) {
                    return v.indexOf('@') === 0;
                });
                fByMilestone = _.some(tags, function(v) {
                    return v.indexOf('^') === 0;
                });
                fByLabel = _.some(tags, function(v) {
                    return v.indexOf('~') === 0;
                });
                fByPriority = _.some(tags, function(v){
                    return v.indexOf('$') === 0;
                });

                filter = function(item) {
                    var item_tags = [],
                        uMatch = true,
                        mMatch = true,
                        lMatch = true,
                        pMatch = true;

                    if (fByUser) {
                        var uName = _.isEmpty(item.assignee) ? '@' : '@' + item.assignee.id;
                        uMatch = _.includes(tags, uName);
                    }

                    if (fByMilestone) {
                        var ms = _.isEmpty(item.milestone) ? '^' : '^' + item.milestone.id;
                        mMatch = _.includes(tags, ms);
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

                    if (fByPriority) {
                        var priority = _.isEmpty(item.priority.name) ? '$' : '$' + item.priority.name;
                        pMatch = _.includes(tags, priority);
                    }

                    return uMatch && mMatch && lMatch && pMatch && item.stage != "";
                };
            }

            if (grouped) {
                group = function(item) {
                    if (grouped == 'milestone') {
                        if (_.isEmpty(item.milestone)) {
                            return 0;
                        }
                        return item.milestone.id;
                    } else if (grouped == 'project') {
                        return item.project_id;
                    } else if (grouped == 'user') {
                        if (!item.assignee) {
                            return 0;
                        }
                        return item.assignee.id;
                    } else if (grouped == 'priority') {
                        if (_.isEmpty(item.priority.name)) {
                            return 0;
                        }
                        return item.priority.id;
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

                if (grouped == 'milestone') {
                    MilestoneService.list(board.project.id).then(function(milestones) {
                        $scope.groupLabels = _.sortBy(milestones, function(milestone){ return milestone.id });
                        $scope.groupLabels.push({
                            id: 0,
                            title: "No Milestone"
                        });
                    });
                } else if (grouped == 'user') {
                    UserService.list(board.project.id).then(function(users) {
                        $scope.groupLabels = _.clone(users);
                        $scope.groupLabels.push({
                            id: 0,
                            name: "Unassigned"
                        });
                    });
                } else if (grouped == 'project') {
                    BoardService.listConnected(board.project.id).then(function(connected){
                        $scope.groupLabels = _.clone(connected);
                        $scope.groupLabels.push({
                            id: board.project.id,
                            path_with_namespace: board.project.path_with_namespace
                        });
                    });
                } else if (grouped == 'priority') {
                    $scope.groupLabels = _.clone(board.priorities);
                    $scope.groupLabels.push({
                        id: 0,
                        viewName: "No priority"
                    });
                }

                board.grouped = grouped;

                $scope.groups = board.reset(filter, group);
                $scope.board = board;


                $scope.showDetail = function(groupId) {
                    $scope.showDetails[groupId] = ! $scope.showDetails[groupId];
                };

                $scope.showDetailInit = function(groupId) {
                    return grouped ? $scope.groups[groupId] != undefined : true;
                };

                $scope.isShowDetail = function(groupId) {
                    if ($scope.showDetails[groupId] == undefined) {
                        $scope.showDetails[groupId] = $scope.showDetailInit(groupId);
                    }

                    return $scope.showDetails[groupId];
                };

                $rootScope.$on('board.change', function() {
                    $scope.groups = board.reset(filter, group);
                });

                $scope.dragControlListeners = BoardService.dragControlListeners(grouped, board);

                WebsocketService.emit('subscribe', {
                    routing_key: 'kanban.' + board.project.id.toString()
                });
                BoardService.listConnected(board.project.id.toString()).then(function(connected){
                    angular.forEach(connected, function(item){
                        WebsocketService.emit('subscribe', {
                            routing_key: 'kanban.' + item.id.toString()
                        });
                    });
                });
            });

            $scope.state = $state;

            $scope.remove = function(card) {
                BoardService.removeCard($scope.board, card).then(function(result) {});
            };
        }
    ]);
})(window.angular);
