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

                    this.stages = _.sortBy(_.filter(labels, function (label) {
                        return stage_regexp.test(label.name);
                    }), 'name');

                    this.priorities = LabelService.listPriorities(project.id);
                    this.viewLabels = LabelService.listViewLabels(project.id);
                    this.priorityLabels = _.pluck(this.priorities, 'name');
                    this.stagelabels = _.pluck(this.stages, 'name');
                    _.each(this.stagelabels, function (label) {
                        this.defaultStages[label] = [];
                    }, this);

                    this.initViewLabels = function (issue) {
                        issue.viewLabels = [];
                        issue.stage = _.intersection(this.stagelabels, issue.labels)[0] || this.stagelabels[0];
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

                    this.issues = _.map(issues, this.initViewLabels, this);

                    this.byStage = function (element, index, items) {
                        element = _.chain(element)
                                  .sortBy(function(item) {return item.id * -1})
                                  .sortBy(function(item) {return item.priority.index * -1})
                                  .value();

                        var stages = {};
                        for (var k in this.defaultStages) {
                            stages[k] = [];
                        }
                        items[index] = _.extend(stages, _.groupBy(element, 'stage', this));
                        for (var idx in items[index]) {
                            this.counts[idx] += items[index][idx].length;
                        }
                    };


                    this.add = function (card) {
                        card.stage = _.intersection(this.stagelabels, card.labels)[0] || this.stagelabels[0];
                        this.initViewLabels(card);
                        this.issues.push(card);
                        $rootScope.$emit('board.change');
                    };

                    this.update = function (card) {
                        var old = this.getCardById(card.id);
                        _.extend(old, card);
                        old.stage = _.intersection(this.stagelabels, old.labels)[0] || this.stagelabels[0];
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
                        groups = _.each(groups, this.byStage, this);

                        return groups;
                    };
                }

                return Board;
            }
        ]
    );
})(window.angular);
