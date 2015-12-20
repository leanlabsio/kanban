(function(angular) {
    'use strict';

    angular.module('gitlabKBApp.board').directive('boardfilter', [
        function() {
            var ddo = {
                templateUrl: CLIENT_VERSION + "/assets/html/board/views/filter.html",
                restrict: "AE",
                controller: 'FilterController',
                controllerAs: "filter",
                bindToController: true
            };

            return ddo;
        }
    ]);
}(window.angular));
