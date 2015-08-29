(function(angular) {
    'use strict';
    angular.module('gitlabKBApp.board').factory('BoardService',
        [
            '$http',
            '$q',
            '$sce',
            'LabelService',
            'WebsocketService',
            'Board', function($http, $q, $sce, LabelService, WebsocketService, Board) {
                var service = {
                    boards: {},
                    boardIdIndex: {},
                    boardPathIndex: {},
                    boardsList: {},
                    get: function(path) {
                        if (_.isEmpty(this.boards[path])) {
                            this.boards[path] = $http.get('/api/board', {
                                params: {
                                    project_id: path
                                }
                            }).then(function(project) {
                                project = project.data.data;
                                this.boards[path] = $q.all([
                                    LabelService.list(project.id),
                                    $http.get('/api/cards', {
                                        params: {
                                            project_id: project.id
                                        }
                                    })
                                ]).then(function(results) {
                                    var board = new Board(results[0], results[1].data.data, project);
                                    this.boards[path] = board;
                                    this.boardIdIndex[board.project.id] = board.project.path_with_namespace;
                                    this.boardPathIndex[board.project.path_with_namespace] = board.project.id;

                                    return this.boards[path];
                                }.bind(this));

                                return this.boards[path];
                            }.bind(this));
                        }

                        return $q.when(this.boards[path]);
                    },
                    getBoardById: function(id) {
                        var path = this.boardIdIndex[id];
                        return this.get(path);
                    },
                    getCard: function(boardId, cardId) {
                        return this.get(boardId).then(function(result) {
                                return _.find(result.issues, function(card) {
                                    return card.iid == cardId;
                                });
                        });
                    },
                    createCard: function(data) {
                        return $http.post('/api/card', data).then(function(newCard) {});
                    },
                    updateCard: function(card) {
                        return $http.put('/api/card', {
                            issue_id: card.id,
                            project_id: card.project_id,
                            assignee_id: card.assignee_id,
                            milestone_id: card.milestone_id,
                            title: card.title,
                            labels: card.labels.join(', '),
                            todo: card.todo,
                            description: card.description,
                            properties: card.properties
                        }).then(function(result) {});
                    },
                    removeCard: function(card) {
                        return $http.delete('/api/card', {
                            data: {
                                project_id: card.project_id,
                                issue_id: card.id,
                                closed: 1
                            },
                            headers: {
                                'Content-Type': 'application/json'
                            }
                        });
                    },
                    getBoards: function() {
                        var _this = this;
                        if (_.isEmpty(_this.boardsList)) {
                            _this.boardsList = $http.get('/api/boards').then(function(result) {
                                _this.boardsList = _.indexBy(result.data.data, 'id');
                                return _this.boardsList;
                            }, function(result) {
                                    _this.boardsList = {};
                                    return $q.reject(result);
                                });
                        }

                        return $q.when(_this.boardsList);
                    },
                    addCardToBoard: function(card) {
                        var _this = this;
                        return _this.getBoardById(card.project_id).then(function(board) {
                            board.add(card);
                        });
                    },
                    removeCardFromBoard: function(card) {
                        var _this = this;
                        return _this.getBoardById(card.project_id).then(function(board) {
                           return board.remove(card);
                        });
                    },
                    updateCardOnBoard: function(card) {
                        var _this = this;
                        return _this.getBoardById(card.project_id).then(function(board) {
                            return board.update(card);
                        });
                    }
                };

                WebsocketService.on('card.create', function(data) {
                    return service.addCardToBoard(data);
                });

                WebsocketService.on('card.move', function(data) {
                    return service.updateCardOnBoard(data);
                });

                WebsocketService.on('card.delete', function(data) {
                    return service.removeCardFromBoard(data);
                });

                WebsocketService.on('card.update', function(data) {
                    return service.updateCardOnBoard(data);
                });

                return service;
            }
        ]
    );
})(window.angular);

