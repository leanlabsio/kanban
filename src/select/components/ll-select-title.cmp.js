(function(angular){
    'use strict';

    angular.module("ll.select").component("llSelectTitle",{
        transclude: true,
        require: {
            selectCtrl: '^llSelect'
        },
        templateUrl: CLIENT_VERSION + "/assets/html/select/views/title.html",
        controller: function() {
            var ctrl = this;


            ctrl.toggle = function() {
                ctrl.selectCtrl.toggle();
            };
        },
    });

}(window.angular));
