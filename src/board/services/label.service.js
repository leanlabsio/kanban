(function(angular) {
    'use strict';
    angular.module('gitlabKBApp.board').factory('LabelService', [
        '$q',
        '$http',
        'stage_regexp',
        function($q, $http, stage_regexp) {
            return {
                labels: [],
                list: function(projectId, withCache) {
                    withCache = (typeof withCache === 'undefined') ? true : withCache;

                    return $http.get('/api/labels/' + projectId, {
                        cache: withCache
                    }).then(function(result) {
                        var labels = result.data.data;
                        if (_.isEmpty(labels)) {
                            return {};
                        }

                        var reserved = _.sortBy(_.filter(labels, function(label) {
                            return stage_regexp.test(label.name);
                        }), 'name');

                        if (_.isEmpty(reserved)) {
                            return {};
                        }

                        this.labels[projectId] = labels;
                        return this.labels[projectId];
                    }.bind(this));
                },
                create: function(projectId, label) {},
                update: function() {},
                delete: function() {}
            };
        }
    ]);
})(window.angular);
