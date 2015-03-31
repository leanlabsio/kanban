(function() {
    'use strict';



    var onReady = function() {
       var box = document.createElement('div');
        css(box, {
            with: "200px",
            height: "200px",
            border: "5px solid red",
            "background-color": "#000"
        });
       prepend(box, document.body);
    };

    var on = function(elem, event, handler) {
        elem.addEventListener(event, handler, false);
    };

    if (document.readyState === 'complete') {
        onReady();
    } else {
        on(document, 'DOMContentLoaded', onReady);
    }

    var css = function(elem, style) {
        for (var prop in style) {
            elem.style[prop] = style[prop];
        }
    };

    var prepend = function(elem, parent) {
        if (parent.children && parent.children.length) {
            parent.insertBefore(elem, parent.children[0]);
        } else {
            parent.appendChild(elem);
        }
    };

    try {

    } catch (ex) {}

})();
