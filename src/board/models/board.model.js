(function (angular) {
    'use strict';

    angular.module('gitlabKBApp.board').factory('Board',
        [
            'UserService',
            'Stage',
            'stage_regexp',
            '$rootScope',
            function (UserService, Stage, stage_regexp, $rootScope) {
                function Board(labels, issues, project) {
                    this.stages = [];

                    this.issues = [];
                    this.labels = [];
                    this.project = project;
                    this.grouped = false;
                    this.defaultStages = {};

                    this.stages = _.sortBy(_.filter(labels, function (label) {
                        return stage_regexp.test(label.name);
                    }), 'name');

                    this.viewLabels = _.indexBy(_.difference(labels, this.stages), 'name');
                    this.labels = _.pluck(this.stages, 'name');
                    _.each(this.labels, function (label) {
                        this.defaultStages[label] = [];
                    }, this);

                    this.initViewLabels = function (issue) {
                        issue.viewLabels = [];
                        issue.stage = _.intersection(this.labels, issue.labels)[0] || this.labels[0];

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
                        element = _.sortBy(element, function(item) {return item.id * -1;});
                        var stages = {};
                        for (var k in this.defaultStages) {
                            stages[k] = [];
                        }
                        items[index] = _.extend(stages, _.groupBy(element, 'stage', this));
                    };


                    this.add = function (card) {
                        card.stage = _.intersection(this.labels, card.labels)[0] || this.labels[0];
                        this.initViewLabels(card);
                        this.issues.push(card);
                        $rootScope.$emit('board.change');
                    };

                    this.update = function (card) {
                        var old = this.getCardById(card.id);
                        _.extend(old, card);
                        old.stage = _.intersection(this.labels, old.labels)[0] || this.labels[0];
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
                    }
                }

                return Board;
            }
        ]
    );
})(window.angular);
