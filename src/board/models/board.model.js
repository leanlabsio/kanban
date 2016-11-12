(function (angular) {
    'use strict';

    angular.module('gitlabKBApp.board').factory('Board',
        [
            'UserService',
            'Stage',
            'State',
            'stage_regexp',
            'priority_regexp',
            'LabelService',
            '$rootScope',
            function (UserService, Stage, State, stage_regexp, priority_regexp, LabelService, $rootScope) {
                function Board(labels, issues, project) {
                    this.stages = [];

                    this.issues = [];
                    this.stale = false;
                    this.project = project;
                    this.grouped = false;
                    this.defaultStages = {};
                    this.state = new State();
                    this.counts = {};
                    this.stages = LabelService.listStages(project.id);
                    this.priorities = LabelService.listPriorities(project.id);
                    this.viewLabels = LabelService.listViewLabels(project.id);
                    this.priorityLabels = _.map(this.priorities, 'name');
                    this.stagelabels = _.map(this.stages, 'name');
                    _.each(this.stages, _.bind(function (stage) {
                        this.defaultStages[stage.viewName] = [];
                    }, this));

                    this.initViewLabels = function (issue) {
                        issue.viewLabels = [];
                        issue.stage = LabelService.getStage(project.id,
                            _.find(issue.labels, function(l){return stage_regexp.test(l)}
                        ));
                        if (! issue.stage) {
                            issue.stage = this.stages[0];
                        }
                        issue.priority = LabelService.getPriority(project.id, _.intersection(this.priorityLabels, issue.labels)[0]);

                        if (!_.isEmpty(issue.labels)) {
                            var labels = issue.labels;
                            for (var i = 0; i < labels.length; i++) {
                                var label = this.viewLabels[labels[i]];
                                if (label !== undefined) {
                                    issue.viewLabels.push(label);
                                }
                            }
                        }

                        return issue;
                    };

                    this.issues = _.map(issues, _.bind(this.initViewLabels, this));

                    this.byStage = function (element, index, items) {
                        element = _.chain(element)
                                  .sortBy(function(item) {return item.id * -1})
                                  .sortBy(function(item) {return item.priority.index * -1})
                                  .value();

                        var stages = {};
                        for (var k in this.defaultStages) {
                            stages[k] = [];
                        }

                        items[index] = _.extend(stages, _.groupBy(element, function(el){
                            return el.stage.viewName;
                        }, this));
                        for (var idx in items[index]) {
                            this.counts[idx] += items[index][idx].length;
                        }
                    };

                    this.add = function (card) {
                        var old = this.getCardById(card.id);

                        if (_.isEmpty(old)) {
                            card.stage = LabelService.getStage(project.id,
                                _.find(card.labels, function(l){return stage_regexp.test(l)}
                            ));
                            this.initViewLabels(card);
                            this.issues.push(card);
                            $rootScope.$emit('board.change');
                        }
                    };

                    this.update = function (card) {
                        var old = this.getCardById(card.id);
                        _.extend(old, card);
                        old.stage = LabelService.getStage(project.id,
                            _.find(old.labels, function(l){return stage_regexp.test(l)}
                        ));
                        this.initViewLabels(old);
                        $rootScope.$emit('board.change');
                    };

                    this.remove = function (card) {
                        var old = this.getCardById(card.id);
                        this.issues.splice(this.issues.indexOf(old), 1);
                        $rootScope.$emit('board.change');
                    };

                    this.getCardById = function (id) {
                        return _.find(this.issues, function (item) {
                            return item.id == id;
                        });
                    };

                    this.reset = function(filter, group) {
                        for (var k in this.defaultStages) {
                            this.counts[k] = 0;
                        }

                        var issues = _.filter(this.issues, filter);
                        var groups = _.groupBy(issues, group);
                        groups = _.isEmpty(groups) ? {
                            none: {}
                        } : groups;
                        groups = _.each(groups, _.bind(this.byStage, this));

                        return groups;
                    };

                    this.listCard = function(filter) {
                        return _.chain(this.issues)
                                .filter(filter)
                                .sortBy(function(item) {return item.id * -1})
                                .sortBy(function(item) {return item.priority.index * -1})
                                .value()
                    }
                }

                return Board;
            }
        ]
    );
})(window.angular);
