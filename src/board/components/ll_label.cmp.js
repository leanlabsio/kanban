(function(angular){
    'use strict';

    angular.module("gitlabKBApp").component("llLabel", {
        bindings: {
            label: '=',
        },
        templateUrl: CLIENT_VERSION + "/assets/html/board/views/label.html"
    });
}(window.angular));
