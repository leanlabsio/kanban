(function(angular) {
    'use strict';

    angular.module('gitlabKBApp.board').factory('State',[
        function() {
            function State() {
                this.showActionBar = false;
            }

            return State;
        }
    ]);
}(window.angular));
