(function(angular){
    'use strict';

    angular.module('gitlabKBApp.user').factory('AuthService',
        [
            '$http',
            '$q',
            'store',
            function ($http, $q, store) {
                return {
                    current: undefined,
                    roles: {
                        anon: 0,
                        user: 1
                    },
                    register: function (data) {
                        return $http.post('/api/register', {
                            _username: data.username,
                            _email: data.email,
                            _password: data.password,
                            _token: data.token
                        }).then(function (result) {
                            store.set('id_token', result.data.token);
                            return store.get('id_token');
                        });
                    },
                    authenticate: function (data) {
                        return $http.post('/api/login', {
                            _username: data.username,
                            _password: data.password
                        }).then(function (result) {
                            store.set('id_token', result.data.token);
                            return store.get('id_token');
                        });
                    },
                    getCurrent: function () {
                        return store.get('id_token');
                    },
                    isAuthenticated: function () {
                        return this.getCurrent() !== null;
                    },
                    authorized: function (state) {
                        var roles = this.roles;
                        return !!(this.isAuthenticated()
                        || state.data.access === undefined
                        || state.data.access == roles.anon);
                    },
                    logout: function() {
                        return store.remove('id_token');
                    }
                };
            }
        ]
    );
})(window.angular);


