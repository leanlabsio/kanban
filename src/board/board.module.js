(function(angular, CLIENT_VERSION) {
    'use strict';
    /**
    * @todo избавиться от блоков resolve v будущем
    **/
    angular.module('gitlabKBApp.board', ['ui.router', 'as.sortable', 'mm.foundation.dropdownToggle', 'gitlabKBApp.websocket', 'll.markdown', 'll.modal', 'mm.foundation.typeahead'])
        .config(['$stateProvider', '$urlMatcherFactoryProvider', function($stateProvider, $urlMatcherFactoryProvider) {
            function valToString(val) {
                return val != null ? val.toString() : val;
            }
            function valFromString(val) {
                return val != null ? val.toString() : val;
            }
            function regexpMatches(val) {
                /*jshint validthis:true */ return this.pattern.test(val);
            }
            $urlMatcherFactoryProvider.type('MyType', {
                encode: valToString,
                decode: valFromString,
                is: regexpMatches,
                pattern: /[^/]+\/[^/]+/
            });

            $stateProvider.decorator('views', function(state, parent) {
                var result = {},
                    views = parent(state);

                angular.forEach(views, function(config, name) {
                    config.templateUrl =  CLIENT_VERSION + "/" + config.templateUrl;
                    result[name] = config;
                });

                return result;
            });

            $stateProvider
                .state('board', {
                    url: '/boards',
                    views: {
                        '': {
                            templateUrl: 'assets/html/board/views/index.html'
                        },
                        'top-bar@board': {
                            templateUrl: 'assets/html/board/views/top_bar.html',
                            controller: 'TopBarController'
                        }
                    },
                    data: {
                        access: 1
                    }
                })
                .state('board.boards', {
                    url: '/',
                    views: {
                        'content@board': {
                            templateUrl: 'assets/html/board/views/board/boards.html',
                            controller: 'BoardListController'
                        }
                    },
                    data: {
                        access: 1
                    }

                })
                .state('board.cards', {
                    url: '/{project_path:MyType}?assignee&milestone&label&group',
                    views: {
                        'content@board': {
                            templateUrl: 'assets/html/board/views/board/cards.html',
                            controller: 'BoardController'
                        },
                        'top-bar@board': {
                            templateUrl: 'assets/html/board/views/top_bar.html',
                            controller: 'TopBarController'
                        },
                        'title@': {
                            templateUrl: 'assets/html/board/views/title.html',
                            controller: 'TopBarController'
                        }
                    },
                    data: {
                        access: 1
                    }
                })
                .state('board.import', {
                    url: '/{project_path:MyType}/import',
                    views: {
                        'content@board': {
                            templateUrl: 'assets/html/board/views/board/configuration.html',
                            controller: 'ConfigurationController'
                        }
                    },
                    data: {
                        access: 1
                    }
                })
                .state('board.cards.create', {
                    url: '/issues/new',
                    views: {
                        'modal@board': {
                            templateUrl: 'assets/html/board/views/card/create.html',
                            controller: 'NewIssueController'
                        }
                    },
                    data: {
                        access: 1
                    }
                })
                .state('board.cards.view', {
                    url: '/issues/:issue_id',
                    views: {
                        'modal@board': {
                            templateUrl: 'assets/html/board/views/card/view.html',
                            controller: 'ViewController'
                        }
                    },
                    data: {
                        access: 1
                    }
                });
        }])
        .config(['$markdownProvider', function($markdownProvider) {
                $markdownProvider.config({
                    linkify: true,
                    html: true
                });
                $markdownProvider.registerPlugin(window.merge_request_plugin);
        }])
        .constant('stage_regexp', /KB\[stage\]\[\d\]\[(.*)\]/);
})(window.angular, window.CLIENT_VERSION);
