(function(angular){
    'use strict';

    angular.module("ll.select").component("llSelect",{
        transclude: true,
        bindings: {
            selected: '=',
            multiply: '='
        },
        templateUrl: CLIENT_VERSION + "/assets/html/select/views/select.html",
        controller: 'SelectController'
    }).controller('SelectController', [
        '$document',
        '$scope',
        function($document, $scope) {
            var ctrl = this;
            ctrl.is_open = false;
            ctrl.options = [];

            if (ctrl.multiply) {
                ctrl.selected = [];
            } else {
                ctrl.selected = null;
            }
            $document.on('click', function(e){
                ctrl.is_open = false;
                $scope.$apply();
            });

            ctrl.onSelect = function(option){
                if (ctrl.multiply) {
                    if (ctrl.selected.indexOf(option) !== -1) {
                        ctrl.selected.splice(ctrl.selected.indexOf(option), 1);
                    } else {
                        ctrl.selected.push(option);
                    }
                } else {
                    ctrl.selected = option;
                }
            };

            ctrl.toggle = function() {
                ctrl.is_open = !ctrl.is_open;
            };

            ctrl.addOption = function(option) {
                ctrl.options.push(option);
            };
        }
    ]);
}(window.angular));
