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
                current: {},
                boards: {},
                boardIdIndex: {},
                boardPathIndex: {},
                boardsList: {},
                boardConnectedIndex: {},
                boardConnected: {},
                get: function(path) {
                    if (_.isEmpty(this.boards[path]) || this.boards[path].stale) {
                        var withCache = true;
                        if (!_.isEmpty(this.boards[path])) {
                            withCache = !this.boards[path].stale;
                        }
                        this.boards[path] = $http.get('/api/board', {
                            params: {
                                project_id: path,
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
                listConnected: function(id) {
                    if (_.isEmpty(this.boardConnected[id])) {
                        this.boardConnected[id] = $http.get('/api/boards/' + id + '/connect', {
                        }).then(function(result){
                            angular.forEach(result.data.data, function(item){
                                if (! this.boardConnectedIndex[item.id]) {
                                    this.boardConnectedIndex[item.id] = [];
                                }
                                this.boardConnectedIndex[item.id].push(id);
                                this.boardConnectedIndex[item.id] = _.uniq(this.boardConnectedIndex[item.id]);
                            }, this);

                            return result.data.data;
                        }.bind(this));
                    }
                    return $q.when(this.boardConnected[id]);
                },
                connectBoard: function(boardId, connectBoardId) {
                    return $http.post('/api/boards/' + boardId + '/connect', {
                        project_id: connectBoardId
                    });
                },
                deleteConnected: function(boardId, connectBoardId) {
                    return $http.delete('/api/boards/' + boardId + '/connect', {
                        params: {
                            board_id: connectBoardId
                        }
                    });
                },
                getBoardById: function(id) {
                    var path = this.boardIdIndex[id];

                    return this.get(path);
                },
                getCard: function(boardId, path_with_namespace, cardId) {
                    var project_path = path_with_namespace || boardId;
                    return this.get(boardId).then(function(result) {
                        return _.find(result.issues, function(card) {
                            return card.iid == cardId && card.path_with_namespace == project_path;
                        });
                    });
                },
                createCard: function(board, data) {
                    return $http.post('/api/card/' + board.project.id, data).then(function(newCard) {});
                },
                sanitize: function(json) {
                    for (var key in json) {
                            if (json.hasOwnProperty(key) && json[key] === null ) {
                                delete json[key];
                            }
                    }
                    return json;
                },
                updateCard: function(board, card) {
                    return $http.put('/api/card/' + board.project.id, this.sanitize({
                        issue_id: card.id,
                        project_id: card.project_id,
                        assignee_id: card.assignee ? card.assignee.id : null,
                        milestone_id: card.milestone ? card.milestone.id : null,
                        title: card.title,
                        labels: card.labels.join(', '),
                        todo: card.todo,
                        description: card.description,
                        properties: card.properties,
                        due_date: card.due_date
                    })).then(function(result) {});
                },
                removeCard: function(board, card) {
                    return $http.delete('/api/card/' + board.project.id, {
                        data: {
                            issue_id: card.id,
                            project_id: card.project_id,
                            assignee_id: card.assignee ? card.assignee.id : null,
                            milestone_id: card.milestone ? card.milestone.id : null,
                            title: card.title,
                            labels: card.labels.join(', '),
                            todo: card.todo,
                            description: card.description,
                            properties: card.properties,
                            closed: 1,
                            due_date: card.due_date
                        },
                        headers: {
                            'Content-Type': 'application/json'
                        }
                    });
                },
                moveCard: function(board, card, oldStage, newStage) {
                    return $http.put('/api/card/' + board.project.id + '/move', this.sanitize({
                        project_id: card.project_id,
                        issue_id: card.id,
                        assignee_id: card.assignee ? card.assignee.id : null,
                        milestone_id: card.milestone ? card.milestone.id : null,
                        title: card.title,
                        labels: card.labels.join(', '),
                        todo: card.todo,
                        description: card.description,
                        properties: card.properties,
                        stage: {
                            source: oldStage,
                            dest: newStage
                        }
                    }));
                },
                changeProject: function(board, card, project) {
                    return LabelService.getStageByName(project.id, card.stage.viewName).then(function(stage){
                        card.labels = _.filter(card.labels, function(label) {
                            return !stage_regexp.test(label);
                        });
                        if (stage) {
                            card.labels.push(stage.name);
                        }

                        return $http.post('/api/card/' + board.project.id + '/move/' + project.id, this.sanitize({
                            issue_id: card.id,
                            project_id: card.project_id,
                            assignee_id: card.assignee ? card.assignee.id : null,
                            milestone_id: card.milestone ? card.milestone.id : null,
                            title: card.title,
                            labels: card.labels.join(', '),
                            todo: card.todo,
                            description: card.description,
                            properties: card.properties
                        })).then(function(result) {
                            this.addCardToBoard(result.data.data);
                            return result.data.data;
                        }.bind(this));
                    }.bind(this));
                },
                getBoards: function() {
                    var _this = this;
                    if (_.isEmpty(_this.boardsList)) {
                        _this.boardsList = $http.get('/api/boards').then(function(result) {
                            _this.boardsList = _.keyBy(result.data.data, 'id');
                            return _this.boardsList;
                        }, function(result) {
                            _this.boardsList = {};
                            return $q.reject(result);
                        });
                    }

                    return $q.when(_this.boardsList);
                },
                getStarredBoards: function() {
                    return $http.get('/api/boards/starred').then(function(result) {
                            return result.data.data;
                        }, function(result) {
                            return $q.reject(result);
                        });
                },
                addCardToBoard: function(card) {
                    var _this = this;
                    var ids = this.boardConnectedIndex[card.project_id];

                    if (!_.isEmpty(ids)) {
                        angular.forEach(ids, function(id){
                            _this.getBoardById(id).then(function(board) {
                                board.add(card);
                            });
                        });
                    } else {
                        _this.getBoardById(card.project_id).then(function(board) {
                            board.add(card);
                        });
                    }
                },
                removeCardFromBoard: function(card) {
                    var _this = this;
                    var ids = this.boardConnectedIndex[card.project_id];

                    if (!_.isEmpty(ids)) {
                        angular.forEach(ids, function(id){
                            _this.getBoardById(id).then(function(board) {
                                board.remove(card);
                            });
                        });
                    } else {
                        _this.getBoardById(card.project_id).then(function(board) {
                            board.remove(card);
                        });
                    }
                },
                updateCardOnBoard: function(card) {
                    var _this = this;
                    var ids = this.boardConnectedIndex[card.project_id];

                    if (!_.isEmpty(ids)) {
                        angular.forEach(ids, function(id){
                            _this.getBoardById(id).then(function(board) {
                                board.update(card);
                            });
                        });
                    } else {
                        _this.getBoardById(card.project_id).then(function(board) {
                            board.update(card);
                        });
                    }
                },
                dragControlListeners: function(grouped, board) {
                    var _this = this;
                    return {
                        longTouch: true,
                        accept: function() {
                            return true;
                        },
                        dragEnd: function(event) {
                            var id = event.source.itemScope.card.id,
                                card = board.getCardById(id),
                                oldLabel = "",
                                newLabel = "",
                                oldGroup = event.source.sortableScope.$parent.$parent.group,
                                newGroup = event.dest.sortableScope.$parent.$parent.group,
                                oldStage = null,
                                newStage = null;

                            LabelService.getStageByName(
                                card.project_id,
                                event.source.sortableScope.$parent.stageName
                            ).then(function(res) {
                                oldStage = res;
                                if (oldStage) {
                                    oldLabel = oldStage.name;
                                }
                                return LabelService.getStageByName(
                                    card.project_id,
                                    event.dest.sortableScope.$parent.stageName
                                );
                            }).then(function(res) {
                                newStage = res;
                                if (newStage) {
                                    newLabel = newStage.name;
                                }

                                card.labels = _.filter(card.labels, function(label) {
                                    return !stage_regexp.test(label);
                                });

                                if (newLabel == "") {
                                    card.stage = null;
                                } else {
                                    card.labels.push(newLabel);
                                    card.stage = newStage;
                                }
                                card.properties.andon = 'none';

                                if (oldGroup != newGroup && card.stage != null) {
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
                                    } else if (grouped == 'project') {
                                        return _this.changeProject(board, card, newGroup);
                                    }

                                    return _this.moveCard(board, card, oldLabel, newLabel);
                                } else {
                                    return _this.moveCard(board, card, oldLabel, newLabel);
                                }
                            });
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
