(function(angular, ReconnectingWebsocket) {
    'use strict';

    angular.module('gitlabKBApp.websocket').factory('WebsocketService',
        [
            '$location', function($location) {
                var protocol = $location.protocol();
                var wsProtocol = protocol === 'http' ? 'ws://' : 'wss://';
                this.ws = new WebSocket(wsProtocol + window.location.host + '/ws/');
                this.missedPings = 0;
                this.queue = [];
                this.handlers = {};

                this.ws.onopen = function(event) {
                    if (this.queue.length > 0) {
                        for (var index in this.queue) {
                            this.ws.send(this.queue[index]);
                        }
                    }
                }.bind(this);

                this.ws.onclose = function(event) {
                    window.location.reload();
                };

                this.ws.onmessage = function(event) {
                    this.handle(event);
                }.bind(this);

                this.on = function(eventId, callback) {
                    this.handlers[eventId] = callback;
                };

                this.handle = function(event) {
                    try {
                        var data = angular.fromJson(event.data);
                        var handler = this.handlers[data.event];
                        handler(data.data);
                    } catch ( e ) {
                        //do nothing
                    }
                };

                this.emit = function(eventId, payload) {
                    var data = {
                        event: eventId,
                        data: payload
                    };
                    if (this.ws.readyState === WebSocket.OPEN) {
                        this.ws.send(angular.toJson(data));
                    } else {
                        this.queue.push(data);
                    }
                };

                this.on('system.ping', function(data) {
                    this.missedPings--;
                }.bind(this));

                return this;
            }
        ]
    );

})(window.angular, window.ReconnectingWebSocket);
