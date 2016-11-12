(function(angular){
    'use strict';

    angular.module('gitlabKBApp.board').factory('CardPriority', [
        'priority_regexp',
        function(priority_regexp) {
            function CardPriority(label) {
                if (_.isEmpty(label)) {
                    return {
                        index: 0,
                        color: '#fffff'
                    }
                }
                var priority = priority_regexp.exec(label.name);
                return {
                    id: label.name,
                    name: label.name,
                    index: parseInt(priority[1]),
                    color: label.color,
                    viewName: priority[2]
                };
            }

            return CardPriority;
        }
    ]);

}(window.angular));
