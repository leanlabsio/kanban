(function(angular) {
    'use strict';
    
    angular.module("gitlabKBApp.board").component('boardBacklog', {
        templateUrl: CLIENT_VERSION + "/assets/html/board/views/backlog.html",
        controller: 'BacklogController'
    }).controller("BacklogController", [
        'BoardService', 
        '$state',
        '$rootScope',
        function(BoardService, $state, $rootScope) {
            var ctrl = this;
            ctrl.project_path = $state.params.project_path;
            var grouped = $state.params.group;

            var filter = function(item) {
                return item.stage == '';
            };
            
            BoardService.get($state.params.project_path).then(function(board) {
                ctrl.cards = board.listCard(filter);

                $rootScope.$on('board.change', function() {
                    ctrl.cards = board.listCard(filter);
                });

                ctrl.dragControlListeners = BoardService.dragControlListeners(grouped, board);
            });
        }
    ]);
}(window.angular));
