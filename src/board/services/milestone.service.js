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
                    }
                };
            }
        ]
    );
})(window.angular);


