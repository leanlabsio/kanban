(function(angular){
    'use strict';

    angular.module("ll.select").component("llSelectSelected",{
        transclude: true,
        bindings: {
            model: '='
        },
        require: {
            selectCtrl: '^llSelect'
        },
        templateUrl: CLIENT_VERSION + "/assets/html/select/views/selected.html",
        controller: function(){
            var ctrl = this;
        }
    });

}(window.angular));
