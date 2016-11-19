(function(angular){
    'use strict';

    angular.module("gitlabKBApp").component("userAvatar", {
        bindings: {
            user: '=',
        },
        templateUrl: CLIENT_VERSION + "/assets/html/board/views/avatar.html"
    });
}(window.angular));
