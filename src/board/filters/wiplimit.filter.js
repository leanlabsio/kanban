(function(angular) {
    'use strict';

    angular.module('gitlabKBApp.board').filter('wiplimit', [
        'stage_regexp',
        function(stage_regexp) {
            return function(input) {
                return input.match(stage_regexp)[3];
            };
        }
    ]);
}(window.angular));
