(function(angular) {
    'use strict';

    angular.module('gitlabKBApp.board').controller('TopBarController',
        [
            '$scope',
            '$state',
            '$stateParams',
            'BoardService',
            'AuthService',
            '$window',
            'UserService',
            'MilestoneService',
            'LabelService', function($scope, $state, $stateParams, BoardService, AuthService, $window, UserService, MilestoneService, LabelService) {
                var milestoneRegex = /\^/;
                var usersRegex = /@/;
                var labelRegex = /~/;

                if ($stateParams.project_path !== undefined) {
                    BoardService.get($stateParams.project_path).then(function(board) {
                        $scope.project = board.project;
                    });
                }

                $scope.logout = function() {
                    AuthService.logout();
                    $window.location.reload();
                };

                $scope.filterval = $stateParams.assignee
                || $stateParams.milestone
                || $stateParams.label
                || '';

                $scope.filter = function(filterval) {
                    var stateParams = {
                        project_id: $stateParams.project_id,
                        project_path: $stateParams.project_path,
                        milestone: '',
                        assignee: '',
                        label: ''
                    };

                    if (milestoneRegex.test(filterval)) {
                        stateParams.milestone = filterval;
                    }

                    if (usersRegex.test(filterval)) {
                        stateParams.assignee = filterval;
                    }

                    if (labelRegex.test(filterval)) {
                        stateParams.label = filterval;
                    }

                    $state.go('board.cards', stateParams);
                };

                $scope.getFilter = function(val) {
                    if (usersRegex.test(val)) {

                        return BoardService.get($stateParams.project_path).then(function(board) {
                            return UserService.list(board.project.id).then(function(users) {
                                var u = _.map(users, function(user) {
                                    return '@' + user.username;
                                });

                                return _.filter(u, function(name) {
                                    return name.toLowerCase().indexOf(val.toLowerCase()) !== -1;
                                });
                            });
                        });
                    }

                    if (milestoneRegex.test(val)) {
                        return BoardService.get($stateParams.project_path).then(function(board) {
                            return MilestoneService.list(board.project.id).then(function(milestones) {
                                var m = _.map(milestones, function(milestone) {
                                    return '^' + milestone.title;
                                });

                                return m;
                            });
                        });
                    }

                    if (labelRegex.test(val)) {
                        return BoardService.get($stateParams.project_path).then(function(board) {
                            var m = _.map(board.viewLabels, function(l) {
                                return '~' + l.name;
                            });

                            return m;
                        });
                    }

                    return [];
                };

                $scope.reset = function() {
                    var params = {
                        project_id: $stateParams.project_id,
                        project_path: $stateParams.project_path,
                        group: ''
                    };
                    $state.go('board.cards', params);
                };

                $scope.group = function() {
                    var params = {
                        project_id: $stateParams.project_id,
                        project_path: $stateParams.project_path,
                        group: 'user'
                    };
                    $state.go('board.cards', params);
                };

                $scope.state = $state;
            }
        ]
    );
})(window.angular);

