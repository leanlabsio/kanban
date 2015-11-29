(function (angular) {
    'use strict';

    // api for modal
    angular.module('ll.modal').factory('$modal', function() {
        return {
            status: null,
            close: function() {
                this.status = 'close';
            },
            submit: function() {
                this.status = 'submit';
            },
            clear: function() {
                this.status = null;
            }
        }
    });

    // Directive: modal
    angular.module('ll.modal').directive('llModal', ['$modal', '$document','version', '$state', function($modal, $document, version, $state) {
        return {
            restrict: 'A',
            link: linkFn,
            transclude: true,
            scope: {
                close: '@onClose',
                submit: '@onSubmit'
            },
            templateUrl: version + '/assets/html/modal/views/modal.html'
        };

        function linkFn($scope, $elem, $attrs) {

            var $body = angular.element(document.body);

            $body.addClass('modal-open');
            $scope.close = hideModal;
            $scope.api = $modal;
            $scope.class = $attrs.windowClass  || '';

            $scope.$watch('api.status', toggledisplay);

            $document.on('keydown', function (event) {
                if (event.which === 27) {
                    hideModal();
                }
            });

            $elem.on('$destroy', function() {
                $body.removeClass('modal-open');
                $document.off('keydown');
            });

            function toggledisplay() {
                if ($scope.api.status == 'submit' || $scope.api.status == 'close') {
                    hideModal();
                }
            };

            function hideModal() {
                $modal.clear();
                $body.removeClass('modal-open');
                $state.go('^');
            };
        };
    }]);
})(window.angular);
