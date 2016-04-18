(function(angular) {
    'use strict';
    angular.module('gitlabKBApp.board').factory('LabelService', [
        '$q',
        '$http',
        'stage_regexp',
        'priority_regexp',
        'CardPriority',
        function($q, $http, stage_regexp, priority_regexp, CardPriority) {
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
                listPriorities: function(projectId) {
                    return _.chain(this.labels[projectId])
                        .filter(function(label) {
                            return priority_regexp.test(label.name);
                        })
                        .map(function(label){
                            return  new CardPriority(label);
                        })
                        .sortBy(function(label){
                            return label.index * -1;
                        }).value();
                },
                listViewLabels: function(projectId) {
                    return _.chain(this.labels[projectId])
                           .filter(function(label) {
                                return !(stage_regexp.test(label.name) || priority_regexp.test(label.name));
                           })
                           .indexBy('name')
                           .value();
                },
                getPriority: function(projectId, label){
                    var priority =_.findWhere(this.labels[projectId], {name: label});
                    return new CardPriority(priority);
                },
                create: function(projectId, label, color) {
                    return $http.post("/api/labels/" + projectId, {
                        name: label,
                        color: color
                    });
                },
                update: function(projectId, oldLabel, newLabel, color) {
                    return $http.put('/api/labels/' + projectId, {
                        name: oldLabel,
                        color: color,
                        new_name: newLabel
                    });
                },
                delete: function(projectId, label) {
                    return $http.delete("/api/labels/" + projectId + "/" + label);
                }
            };
        }
    ]);
})(window.angular);
