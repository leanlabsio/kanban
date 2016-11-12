(function(angular) {
    'use strict';

    angular.module('gitlabKBApp.board').factory('Stage',[
        'stage_regexp', function(stage_regexp) {
            function Stage(label) {
                if (_.isEmpty(label)) {
                    return {
                        index: 0,
                        color: '#fffff'
                    }
                }
                var stage = stage_regexp.exec(label.name);
                return {
                    id: label.name,
                    name: label.name,
                    index: parseInt(stage[1]),
                    color: label.color,
                    viewName: stage[2],
                    wip: stage[3]
                };
            }

            return Stage;
        }]
    );
})(window.angular);
