(function(angular, markdownit) {
    'use strict';

    angular.module('ll.markdown').provider('$markdown', [function() {
            return {
                opts: {},
                plugins: [],
                config: function(options) {
                    this.opts = options;
                },
                registerPlugin: function(plugin) {
                    this.plugins.push(plugin);
                },
                $get: function() {
                    var md = markdownit(this.opts);

                    if (this.plugins.length !== 0) {
                        for (var i = 0; i < this.plugins.length; i++) {
                            md.use(this.plugins[i]);
                        }
                    }

                    return md;
                }
            };

    }]);
})(window.angular, window.markdownit);
