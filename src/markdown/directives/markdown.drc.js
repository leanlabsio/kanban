(function (angular) {
    'use strict';

    angular.module('ll.markdown').directive('llMarkdown',
        [
            '$sce',
            '$markdown',
            'BoardService',
            '$stateParams',
            'host_url',
            function($sce, $markdown, BoardService, $stateParams, host_url) {
                return {
                    restrict: 'A',
                    scope: {
                        markdown: "="
                    },
                    link: function(scope, element, attributes) {
                        scope.$watch('markdown', function(newData, oldData) {
                            if (newData !== undefined) {
                                BoardService.get($stateParams.project_path).then(function(board) {
                                    var repoUrl = '/' + board.project.path_with_namespace;
                                    element.html($markdown.render(newData, {host_url: host_url + repoUrl}));
                                });
                            }
                        }, true);
                    }
                }
            }
        ]
    );
})(window.angular);
