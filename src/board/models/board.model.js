(function (angular) {
    'use strict';
    // taken from http://www.sitepoint.com/javascript-generate-lighter-darker-color/
    // modify to use `alpha`
    // @todo should use angular's `filter` instead
    // @todo this use `alpha` without fallback/shim
    function ColorLuminance(hex, lum) {

    // validate hex string
    hex = String(hex).replace(/[^0-9a-f]/gi, '');
    if (hex.length < 6) {
        hex = hex[0]+hex[0]+hex[1]+hex[1]+hex[2]+hex[2];
    }
    lum = lum || 0;

    // convert to decimal and change luminosity
    var rgb = "rgba(", c, i;
    for (i = 0; i < 3; i++) {
        c = parseInt(hex.substr(i*2,2), 16);
        rgb += c + ',';
    }
    rgb += lum + ')';
    return rgb;
}
    angular.module('gitlabKBApp.board').factory('Board',
        [
            'UserService',
            'Stage',
            'State',
            'stage_regexp',
            'priority_regexp',
            '$rootScope',
            function (UserService, Stage, State, stage_regexp, priority_regexp, $rootScope) {
                function Board(labels, issues, project) {
                    this.stages = [];

                    this.issues = [];
                    this.stale = false;
                    this.labels = [];
                    this.project = project;
                    this.grouped = false;
                    this.defaultStages = {};
                    this.state = new State();
                    this.counts = {};

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

                        issue.priority ={
                            index: Infinity,
                            name: "None",
                            color: null,
                        }

                        if (!_.isEmpty(issue.labels)) {
                            var labels = issue.labels;
                            for (var i = 0; i < labels.length; i++) {
                                var label = this.viewLabels[labels[i]];
                                var v = priority_regexp.exec(labels[i]);
                                if (v && v[2]) {
                                    issue.priority = {
                                        index: v[1],
                                        name: v[2],
                                        // @todo should use angular's `filter` instead
                                        // 20% opaque / 80% transparent
                                        color: ColorLuminance(this.viewLabels[labels[i]].color, .2)
                                    }
                                    // continue
                                }
                                else if (label !== undefined) {
                                    issue.viewLabels.push(label);
                                }
                            }
                        }

                        return issue;
                    };

                    this.issues = _.map(issues, this.initViewLabels, this);

                    this.byStage = function (element, index, items) {
                        element = _.sortBy(element, function(item) {return item.priority.index*1;});
                        var stages = {};
                        for (var k in this.defaultStages) {
                            stages[k] = [];
                            this.counts[k] = 0;
                        }
                        items[index] = _.extend(stages, _.groupBy(element, 'stage', this));
                        for (var idx in items[index]) {
                            this.counts[idx] += items[index][idx].length;
                        }
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
