(function(){
    'use strict';

    angular.module('gitlabKBApp.board').filter('colorLuminance', [
        function(){
            // taken from http://www.sitepoint.com/javascript-generate-lighter-darker-color/
            // modify to use `alpha`
            return function(hex, lum) {

                // validate hex string
                hex = String(hex).replace(/[^0-9a-f]/gi, '');
                if (hex.length < 6) {
                    hex = hex[0]+hex[0]+hex[1]+hex[1]+hex[2]+hex[2];
                }
                lum = lum || 0;

                // convert to decimal and change luminosity
                var rgb = "rgba(", c, i;
                for (i = 0; i < 3; i++) {
                    c = parseInt(hex.substr(i*2,2), 16);
                    rgb += c + ',';
                }
                rgb += lum + ')';
                return rgb;
            }
        }
    ]);
}(window.angular));
