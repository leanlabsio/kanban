(function (angular) {
    'use strict';

    angular.module('gitlabKBApp.board').filter('priorityName', [
        'priority_regexp',
        function (priority_regexp) {
            return function (input) {
                var priority = input.match(priority_regexp);
                if (_.isEmpty(priority)) {
                    return input;
                }
                return priority[2];
            };
        }]);
}(window.angular));
