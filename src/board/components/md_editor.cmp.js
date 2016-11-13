(function(angular){
    'use strict';

    angular.module("gitlabKBApp").component("mdEditor",{
        templateUrl: CLIENT_VERSION + "/assets/html/board/views/md-editor.html",
        controller: 'MdEditorController',
        bindings: {
            description: '=',
            rows: '='
        }
    }).controller('MdEditorController', [
        'FileUploader',
        '$state',
        'AuthService',
        'BoardService',
        function(FileUploader,$state, AuthService, BoardService){
            var ctrl = this;
            ctrl.project_path = $state.params.project_path;
            ctrl.progress = false;
            ctrl.uploader = new FileUploader({
                autoUpload: true
            });

            BoardService.get($state.params.project_path).then(function(board) {
                ctrl.uploader.url = '/api/boards/' + board.project.id + '/upload';
            });

            ctrl.uploader.headers = {
                'X-KB-Access-Token': AuthService.getCurrent()
            };

            ctrl.uploader.onProgressItem = function(item) {
                ctrl.progress = true;
            };

            ctrl.uploader.onCompleteItem = function(item, response, status, headers) {
                if (ctrl.description == undefined) {
                    ctrl.description = '';
                }
                ctrl.progress = false;
                ctrl.description += '\n' + '\n' + response.data.markdown;
            };
        }
    ]);
}(window.angular));
