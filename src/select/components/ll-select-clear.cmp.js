(function(angular){
    'use strict';

    angular.module("ll.select").component("llSelectClear",{
        transclude: true,
        require: {
            selectCtrlOptions: '^llSelectOptions'
        },
        templateUrl: CLIENT_VERSION + "/assets/html/select/views/clear.html",
        controller: function(){
            var ctrl = this;
            ctrl.onClear = function(){
                ctrl.selectCtrlOptions.selectCtrl.toggle();
                ctrl.selectCtrlOptions.selectCtrl.onSelect(null);
            }
        }
    });
}(window.angular));
