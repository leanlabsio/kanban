(function(angular) {
    'use strict';
    angular.module('gitlabKBApp.board').factory('MilestoneService',
        [
            '$http',
            '$q', function($http, $q) {
                return {
                    milestones: {},
                    list: function(projectId) {
                        var _this = this;
                        if (_.isEmpty(_this.milestones[projectId])) {
                            _this.milestones[projectId] = $http.get('/api/milestones', {
                                params: {
                                    project_id: projectId
                                }
                            }).then(function(result) {
                                _this.milestones[projectId] = result.data.data;
                                return _this.milestones[projectId];
                            });
                        }

                        return $q.when(_this.milestones[projectId]);
                    },
                    create: function(data) {
                        return $http.post('/api/milestones', data).then(function(result){
                            if (_.isEmpty(this.milestones[data.project_id])) {
                                return this.list(data.project_id)
                            }
                            this.milestones[data.project_id].push(result.data.data)
                        }.bind(this));
                    },
                    findByName: function(projectId, name) {
                        return this.list(projectId).then(function(milestones) {
                            return _.find(milestones, function(milestone) {
                                return milestone.title == name;
                            });
                        }.bind(this));
                    }
                };
            }
        ]
    );
})(window.angular);


