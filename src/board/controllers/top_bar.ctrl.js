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
                
                $scope.tags = [];
                if ($stateParams.tags) {
                    $scope.tags = JSON.parse(decodeURIComponent($stateParams.tags));
                }
                
                $scope.applyFilter = function(tags) {
                    var stateParams = {
                        project_id: $stateParams.project_id,
                        project_path: $stateParams.project_path,
                        tags: ''
                    };
                    
                    // Put the object into storage
                    stateParams.tags = JSON.stringify(tags);
                    $state.go('board.cards', stateParams);
                };
                
                $scope.filterTypeahead = function(val) {
                    if (usersRegex.test(val)) {
                        return BoardService.get($stateParams.project_path).then(function(board) {
                            return UserService.list(board.project.id).then(function(users) {
                                var u = _.map(users, function(user) {
                                    return {
                                        idname : "@" + user.id,
                                        name : user.name,
                                        username : user.username,
                                        avatar_url : user.avatar_url
                                        };
                                });
                                
                                var val2 = val.replace(/@/, '').toLowerCase();
                                return _.filter(u, function(obj) {
                                    return obj.username.toLowerCase().indexOf(val2)!==-1 || 
                                        obj.name.toLowerCase().indexOf(val2)!==-1;
                                });
                            });
                        });
                    };
                    
                    if (milestoneRegex.test(val)) {
                        return BoardService.get($stateParams.project_path).then(function(board) {
                            return MilestoneService.list(board.project.id).then(function(milestones) {
                                var m = _.map(milestones, function(milestone) {
                                    return {
                                        idname : '^' + milestone.id,
                                        name : milestone.title
                                    };
                                });

                                var val2 = val.replace(/\^/, '').toLowerCase();
                                return _.filter(m, function(obj) {
                                    return obj.name.toLowerCase().indexOf(val2)!==-1;
                                });
                            });
                        });
                    };
 

                    if (labelRegex.test(val)) {
                        return BoardService.get($stateParams.project_path).then(function(board) {
                            var m = _.map(board.viewLabels, function(l) {
                                return {
                                    idname : "~" + l.name,
                                    name : l.name,
                                    color: l.color
                                };
                            });

                            var val2 = val.replace(/~/, '').toLowerCase();
                            return _.filter(m, function(obj) {
                                return obj.name.toLowerCase().indexOf(val2)!==-1;
                            });
                        });
                    };

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

                $scope.group = function(field) {
                    var params = {
                        project_id: $stateParams.project_id,
                        project_path: $stateParams.project_path,
                        group: field
                    };
                    $state.go('board.cards', params);
                };

                $scope.state = $state;
            }
        ]
    );
})(window.angular);

