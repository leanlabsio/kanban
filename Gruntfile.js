
module.exports = function (grunt) {
    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),

        sass: {
            options: {
                includePaths: [
                    'bower_components/foundation/scss',
                    'bower_components/font-awesome/scss/',
                    'bower_components/sass-flex-mixin/'
                ]
            },
            dist: {
                options: {
                    outputStyle: 'compressed'
                },
                files: {
                    'web/assets/css/app.css': 'src/scss/app.scss'
                }
            }
        },

        concat: {
            dist: {
                src: [
                    "src/**/*.module.js", 
                    "src/**/**!(.module).js"
                ],
                dest: "web/assets/js/app.js"
            }
        },

        uglify: {
            dist: {
                files: {
                    "web/assets/js/app.min.js": ["web/assets/js/app.js"]
                }
            }
        },

        copy: {
            main: {
                files: [
                    {
                        flatten: true,
                        src: ['bower_components/angular-mocks/angular-mocks.js'],
                        dest: 'web/assets/js/angular-mocks.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['bower_components/reconnectingWebsocket/reconnecting-websocket.js'],
                        dest: 'web/assets/js/reconnecting-websocket.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['bower_components/markdown-it/dist/markdown-it.js'],
                        dest: 'web/assets/js/markdown-it.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['bower_components/a0-angular-storage/dist/angular-storage.js'],
                        dest: 'web/assets/js/angular-storage.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['bower_components/angular-underscore/angular-underscore.js'],
                        dest: 'web/assets/js/angular-underscore.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['bower_components/underscore/underscore.js'],
                        dest: 'web/assets/js/underscore.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true, 
                        src: ['bower_components/angular-foundation/src/topbar/topbar.js'],
                        dest: 'web/assets/js/topbar.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['bower_components/angular-foundation/src/dropdownToggle/dropdownToggle.js'],
                        dest: 'web/assets/js/dropdownToggle.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true, 
                        src: ['bower_components/angular-foundation/src/position/position.js'], 
                        dest: 'web/assets/js/position.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['bower_components/angular-foundation/src/typeahead/typeahead.js'],
                        dest: 'web/assets/js/typeahead.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['bower_components/angular-foundation/src/bindHtml/bindHtml.js'],
                        dest: 'web/assets/js/bindHtml.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true, 
                        src: ['bower_components/angular-foundation/src/mediaQueries/mediaQueries.js'], 
                        dest: 'web/assets/js/mediaQueries.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true, 
                        src: ['bower_components/marked/lib/marked.js'], 
                        dest: 'web/assets/js/marked.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true, 
                        src: ['bower_components/angular-loading-bar/build/loading-bar.js'], 
                        dest: 'web/assets/js/loading-bar.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true, 
                        src: ['bower_components/angular-loading-bar/build/loading-bar.css'], 
                        dest: 'web/assets/css/loading-bar.css',
                        filter: 'isFile'
                    },
                    {
                        flatten: true, 
                        src: ['bower_components/angular-foundation/src/transition/transition.js'], 
                        dest: 'web/assets/js/transition.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true, 
                        src: ['bower_components/ng-sortable/dist/ng-sortable.js'], 
                        dest: 'web/assets/js/ng-sortable.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true, 
                        src: ['bower_components/angular/angular.js'], 
                        dest: 'web/assets/js/angular.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true, 
                        src: ['bower_components/angular-ui-router/release/angular-ui-router.min.js'], 
                        dest: 'web/assets/js/angular-ui-router.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true, 
                        src: ['bower_components/ng-sortable/dist/ng-sortable.min.css'], 
                        dest: 'web/assets/css/ng-sortable.min.css',
                        filter: 'isFile'
                    },
                    {
                        flatten: true, 
                        src: ['bower_components/jquery/dist/jquery.min.js'], 
                        dest: 'web/assets/js/jquery.min.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: false, 
                        expand: true, 
                        cwd: 'bower_components/angular-foundation/template', 
                        src: '**', 
                        dest: 'web/template/', 
                        filter: 'isFile'
                    },
                    {
                        flatten: true, 
                        expand: true, 
                        cwd: 'bower_components/foundation/js/foundation/', 
                        src: '**', 
                        dest: 'web/assets/js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true, 
                        expand: true, 
                        cwd: 'bower_components/font-awesome/fonts/', 
                        src: '**', 
                        dest: 'web/assets/fonts/',
                        filter: 'isFile'
                    },
                    {
                        flatten: false, 
                        expand: true, 
                        cwd: 'src/', 
                        src: ['**/*.js'], 
                        dest: 'web/assets/js/',
                        filter: 'isFile'
                    },
                    {
                        flatten: false,
                        expand: true, 
                        cwd: 'src/', 
                        src: ['**/*.html'], 
                        dest: 'web/assets/html/',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['src/user/views/oauth.html'],
                        dest: 'web/assets/html/user/views/oauth.html',
                        filter: 'isFile'
                    },
                    {
                        flatten: false,
                        expand: true, 
                        cwd: 'src/static/images/', 
                        src: ['**/*.svg'], 
                        dest: 'web/images/', 
                        filter: 'isFile'
                    }
        ]
    }
},

watch: {
    grunt: {
        files: ['Gruntfile.js'],
            tasks: ['sass', 'copy']
        },

        sass: {
            files: 'src/scss/*.scss',
            tasks: ['sass']
        },

        copy: {
            files: ['src/**/*.js', 'src/**/*.html'],
            tasks: ['copy']
        },

        concat: {
            files: ['src/**/*.js'],
            tasks: ['concat']
        },

        uglify: {
            files: ['web/assets/js/app.js'],
            tasks: ['uglify']
        }
    }
});

grunt.loadNpmTasks('grunt-sass');
grunt.loadNpmTasks('grunt-contrib-watch');
grunt.loadNpmTasks('grunt-contrib-copy');
grunt.loadNpmTasks('grunt-contrib-concat');
grunt.loadNpmTasks('grunt-contrib-uglify');

grunt.registerTask('build', ['sass', 'copy', 'concat', 'uglify']);
grunt.registerTask('default', ['build', 'watch']);
};
