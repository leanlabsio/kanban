(function(angular){
    'use strict';

    angular.module("ll.select").component("llSelectOptions",{
        transclude: true,
        bindings: {
            model: '='
        },
        require: {
            selectCtrl: '^llSelect'
        },
        templateUrl: CLIENT_VERSION + "/assets/html/select/views/options.html",
    });
}(window.angular));
