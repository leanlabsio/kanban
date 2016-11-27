(function(angular){
    'use strict';

    angular.module("ll.select").component("llSelectOption",{
        transclude: true,
        bindings: {
            model: '=',
            onUpdate : '&'
        },
        require: {
            selectCtrlOptions: '^llSelectOptions'
        },
        templateUrl: CLIENT_VERSION + "/assets/html/select/views/option.html",
        controller: function() {
            var ctrl = this;
            this.$onInit = function() {
                ctrl.selectCtrlOptions.selectCtrl.addOption(this.model);
            };

            ctrl.onSelect = function() {
                if (!ctrl.selectCtrlOptions.selectCtrl.multiply) {
                    ctrl.selectCtrlOptions.selectCtrl.toggle();
                }
                ctrl.selectCtrlOptions.selectCtrl.onSelect(this.model);
            };

            ctrl.isChecked = function(){
                return ctrl.selectCtrlOptions.selectCtrl.selected.indexOf(ctrl.model) !== -1
            };
        },
    });
}(window.angular));
