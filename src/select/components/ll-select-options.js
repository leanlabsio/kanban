(function(angular){
    'use strict';

    angular.module("ll.select").component("llSelectOptions",{
        transclude: true,
        bindings: {
            model: '=',
            onUpdate : '&'
        },
        require: {
            selectCtrl: '^llSelect'
        },
        templateUrl: CLIENT_VERSION + "/assets/html/select/views/options.html",
    });
}(window.angular));
