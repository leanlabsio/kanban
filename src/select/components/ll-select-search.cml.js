(function(angular){
    'use strict';

    angular.module("ll.select").component("llSelectSearch",{
        transclude: true,
        bindings: {
            model: '=',
            placeholder: '@'
        },
        require: {
            selectCtrlOptions: '^llSelectOptions'
        },
        templateUrl: CLIENT_VERSION + "/assets/html/select/views/search.html",
        controller: function(){
            var ctrl = this;
            ctrl.$onInit = function(){
                ctrl.model = '';
            }
        }
    });
}(window.angular));
