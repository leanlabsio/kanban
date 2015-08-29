(function(angular) {
    'use strict';

    angular.module('gitlabKBApp.board').factory('Stage',
        [function() {
                function Stage(label, cards, title) {
                    this.label = label;
                    this.cards = cards;
                    this.title = title;
                }

                return Stage;
            }
        ]
    );
})(window.angular);
