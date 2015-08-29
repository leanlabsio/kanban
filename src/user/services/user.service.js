(function(angular) {
    'use strict';

    angular.module('gitlabKBApp.user').factory('UserService',
        [
            '$http',
            '$q', function($http, $q) {
                return {
                    users: {},
                    list: function(projectId) {
                        var _this = this;
                        if (_.isEmpty(_this.users[projectId])) {
                            _this.users[projectId] = $http.get('/api/users', {
                                params: {
                                    project_id: projectId
                                }
                            }).then(function(result) {
                                _this.users[projectId] = result.data;
                                return _this.users[projectId];
                            });
                        }
                        return $q.when(_this.users[projectId]);
                    },
                    findByName: function(projectId, name) {
                        var _this = this;

                        return _this.list(projectId).then(function(users) {
                            return _.find(users, function(user) {
                                return user.name == name;
                            });
                        });
                    }
                };
            }
        ]
    );
})(window.angular);
