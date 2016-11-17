(function (angular) {
    'use strict';

    angular.module('ll.markdown').directive('llMarkdown',
        [
            '$sce',
            '$markdown',
            'BoardService',
            '$stateParams',
            'host_url',
            'MilestoneService',
            'LabelService',
            function($sce, $markdown, BoardService, $stateParams, host_url, MilestoneService, LabelService) {
                return {
                    restrict: 'A',
                    scope: {
                        markdown: "="
                    },
                    link: function(scope, element, attributes) {
                        scope.$watch('markdown', function(newData, oldData) {
                            if (newData !== undefined) {
                                var board;
                                var milestones;
                                BoardService.get($stateParams.project_path).then(function(res) {
                                    board = res;
                                    return MilestoneService.list(board.project.id);
                                }).then(function(res) {
                                    milestones = res;
                                    return LabelService.list(board.project.id);
                                }).then(function(labels){
                                    element.html($markdown.render(newData, {
                                        host_url: host_url,
                                        board_url: board.project.path_with_namespace,
                                        milestones: milestones,
                                        labels: labels
                                    }));
                                });
                            }
                        }, true);
                    }
                }
            }
        ]
    );
})(window.angular);
