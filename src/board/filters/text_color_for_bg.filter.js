(function(){
    'use strict';

    angular.module('gitlabKBApp.board').filter('textColorForBg', [
        function(){
            return function(hex) {

                var rgb = [], i, brightness = 0;

                hex = String(hex).replace(/[^0-9a-f]/gi, '');
                if (hex.length < 6) {
                    hex = hex[0]+hex[0]+hex[1]+hex[1]+hex[2]+hex[2];
                }

                for (i = 0; i < 3; i++) {
                    brightness += parseInt(hex.substr(i*2,2), 16);
                }

                return (brightness > 500) ? '#333' : '#fff';
            }
        }
    ]);
}(window.angular));
