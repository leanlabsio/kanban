(function(angular) {
    'use strict';
    angular.module('gitlabKBApp.board').factory('BoardService', [
        '$http',
        '$q',
        '$sce',
        'LabelService',
        'WebsocketService',
        'Board',
        'stage_regexp',
        'UserService',
        'MilestoneService',
        function($http, $q, $sce, LabelService, WebsocketService, Board, stage_regexp, UserService, MilestoneService) {
            var service = {
                boards: {},
                boardIdIndex: {},
                boardPathIndex: {},
                boardsList: {},
                get: function(path) {
                    if (_.isEmpty(this.boards[path]) || this.boards[path].stale) {
                        var withCache = true;
                        if (!_.isEmpty(this.boards[path])) {
                            withCache = !this.boards[path].stale;
                        }
                        this.boards[path] = $http.get('/api/board', {
                            params: {
                                project_id: path
                            }
                        }).then(function(project) {
                            project = project.data.data;
                            this.boards[path] = $q.all([
                                LabelService.list(project.id, withCache),
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
                            issue_id: card.id,
                            project_id: card.project_id,
                            assignee_id: card.assignee_id,
                            milestone_id: card.milestone_id,
                            title: card.title,
                            labels: card.labels.join(', '),
                            todo: card.todo,
                            description: card.description,
                            properties: card.properties,
                            closed: 1
                        },
                        headers: {
                            'Content-Type': 'application/json'
                        }
                    });
                },
                moveCard: function(card, oldStage, newStage) {
                    return $http.put('/api/card/move', {
                        project_id: card.project_id,
                        issue_id: card.id,
                        assignee_id: card.assignee_id,
                        milestone_id: card.milestone_id,
                        title: card.title,
                        labels: card.labels.join(', '),
                        todo: card.todo,
                        description: card.description,
                        properties: card.properties,
                        stage: {
                            source: oldStage,
                            dest: newStage
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
                },
                dragControlListeners: function(grouped, board) {
                    var _this = this;
                    return {
                        accept: function() {
                            return true;
                        },
                        dragEnd: function(event) {
                            var id = event.source.itemScope.card.id;
                            var card = board.getCardById(id);

                            var oldLabel = event.source.sortableScope.$parent.stageName;
                            var newLabel = event.dest.sortableScope.$parent.stageName;

                            var oldGroup = event.source.sortableScope.$parent.$parent.group;
                            var newGroup = event.dest.sortableScope.$parent.$parent.group;

                            card.labels = _.filter(card.labels, function(label) {
                                return !stage_regexp.test(label);
                            });

                            if (newLabel == undefined) {
                                card.stage = "";
                            } else {
                                card.labels.push(newLabel);
                                card.stage = newLabel;
                            }
                            card.properties.andon = 'none';

                            if (oldGroup != newGroup && card.stage != "") {
                                if (grouped == 'milestone') {
                                    card.milestone_id = newGroup.id === undefined ? 0 : newGroup.id;
                                    card.milestone = newGroup;
                                } else if (grouped == 'user') {
                                    card.assignee_id = newGroup.id === undefined ? 0 : newGroup.id;
                                    card.assignee = newGroup;
                                } else if (grouped == 'priority') {
                                    var index = card.labels.indexOf(oldGroup.name);
                                    if (index !== -1) {
                                        card.labels.splice(index, 1);
                                    }
                                    if (!_.isEmpty(newGroup.name)) {
                                        card.priority = newGroup;
                                        card.labels.push(newGroup.name);
                                    }
                                }

                                return _this.moveCard(card, oldLabel, newLabel);
                            } else {
                                return _this.moveCard(card, oldLabel, newLabel);
                            }
                        },
                        containment: '#board'
                    }
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
    ]);
})(window.angular);
