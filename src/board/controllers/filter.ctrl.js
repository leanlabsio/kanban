(function(angular, _) {
    'use strict';

    angular.module('gitlabKBApp.board').controller('FilterController', [
        '$scope',
        '$state',
        '$stateParams',
        'BoardService',
        'UserService',
        'MilestoneService',
        function($scope, $state, $stateParams, BoardService, UserService, MilestoneService) {
            var labels = [],
                milestones = [],
                users = [],
                priority = [];

            this.tags = _.isArray($stateParams.tags) ? $stateParams.tags : [$stateParams.tags];

            BoardService.get($stateParams.project_path).then(function(board) {
                this.labels = _.toArray(board.viewLabels);
                this.priorities = board.priorities;

                MilestoneService.list(board.project.id).then(function(milestones) {
                    this.milestones = milestones;
                }.bind(this));

                UserService.list(board.project.id).then(function(users) {
                    this.users = users;
                }.bind(this));

            }.bind(this));

            /**
             * Apply selected filtering criteria
             */
            this.apply = function(tag) {
                var params = {
                    project_id: $stateParams.project_id,
                    project_path: $stateParams.project_path,
                    tags: $stateParams.tags
                };

                if (_.isArray(params.tags)) {
                    var idx = params.tags.indexOf(tag);

                    if (idx == -1) {
                        params.tags = params.tags.concat([tag]);
                    } else {
                        params.tags.splice(idx, 1);
                    }
                } else if (_.isString(params.tags)) {
                    if (params.tags == tag) {
                        params.tags = [];
                    } else {
                        params.tags = [params.tags, tag];
                    }
                } else {
                    params.tags = [tag];
                }

                $state.go('board.cards', params);
            };

            this.applyAll = function(tagPrefix, tags, identifyBy, enable) {
                var params = {
                    project_id: $stateParams.project_id,
                    project_path: $stateParams.project_path,
                    tags: Array.isArray($stateParams.tags) ? $stateParams.tags : $stateParams.tags ? [$stateParams.tags] : []
                };

                tags = _(tags).values().map(function(tag) { return tagPrefix + tag[identifyBy]; }).value().concat(tagPrefix);

                if (enable) {
                    params.tags = _.uniq((params.tags || []).concat(tags));
                } else {
                    params.tags && _.pullAll(params.tags, tags);
                }

                $state.go('board.cards', params);
            };

            /**
             * Clear all filters
             */
            this.clear = function() {
                $state.go('board.cards', {
                    tags: []
                });
            };

            this.checked = function(obj) {
                return _.includes(this.tags, obj);
            }
        }
    ]);
}(window.angular, window._));
