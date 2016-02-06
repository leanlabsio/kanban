(function (angular) {
    'use strict';

    angular.module('gitlabKBApp.board').filter('stagename', [
        'stage_regexp',
        function (stage_regexp) {
            return function (input) {
                return input.match(stage_regexp)[2];
            };
        }]);
}(window.angular));
