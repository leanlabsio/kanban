(function(angular){
    'use strict';

    angular.module("ll.select").component("llUnselected",{
        transclude: true,
        require: {
            selectCtrl: '^llSelect'
        },
        templateUrl: CLIENT_VERSION + "/assets/html/select/views/unselected.html"
    });
}(window.angular));
