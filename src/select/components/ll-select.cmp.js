(function(angular){
    'use strict';

    angular.module("ll.select").component("llSelect",{
        transclude: true,
        bindings: {
            selected: '=',
            multiply: '=',
            onUpdate: '&',
            callback: '&'
        },
        templateUrl: CLIENT_VERSION + "/assets/html/select/views/select.html",
        controller: 'SelectController'
    }).controller('SelectController', [
        '$document',
        '$scope',
        '$parse',
        function($document, $scope, $parse) {
            var ctrl = this;

            ctrl.$onInit = function(){
                ctrl.is_open = false;
                ctrl.options = [];
            }

            $document.on('click', function(e){
                ctrl.is_open = false;
                $scope.$apply();
            });

            ctrl.onSelect = function(option){
                if (ctrl.multiply) {
                    if (ctrl.selected == undefined) {
                        ctrl.selected = [];
                    }

                    if (ctrl.selected.indexOf(option) !== -1) {
                        ctrl.selected.splice(ctrl.selected.indexOf(option), 1);
                    } else {
                        ctrl.selected.push(option);
                    }
                } else {
                    ctrl.selected = option;
                }

                ctrl.callback({value: option});
            };

            ctrl.toggle = function() {
                ctrl.is_open = !ctrl.is_open;
            };

            ctrl.addOption = function(option) {
                ctrl.options.push(option);
            };

            ctrl.isSelected = function(){
                return ! _.isEmpty(ctrl.selected);
            }
        }
    ]);
}(window.angular));
