
module.exports = function (grunt) {
    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),

        sass: {
            options: {
                includePaths: [
                    'node_modules/foundation-sites/scss',
                    'node_modules/font-awesome/scss',
                    'node_modules/sass-flex-mixin/',
                    'node_modules/angularjs-datepicker/src/css/'
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
                        flatten: false,
                        expand: true,
                        cwd: 'node_modules/twemoji/svg/',
                        src: ['**/*.svg'],
                        dest: 'web/images/twemoji/svg/',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/twemoji/twemoji.npm.js'],
                        dest: 'web/assets/js/twemoji.min.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/markdown-it-emoji/dist/markdown-it-emoji.min.js'],
                        dest: 'web/assets/js/markdown-it-emoji.min.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/markdown-it/dist/markdown-it.js'],
                        dest: 'web/assets/js/markdown-it.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angular-storage/dist/angular-storage.js'],
                        dest: 'web/assets/js/angular-storage.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angular-lodash/angular-lodash.js'],
                        dest: 'web/assets/js/angular-lodash.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/lodash/lodash.js'],
                        dest: 'web/assets/js/lodash.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angular-mm-foundation/src/topbar/topbar.js'],
                        dest: 'web/assets/js/topbar.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angular-mm-foundation/src/dropdownToggle/dropdownToggle.js'],
                        dest: 'web/assets/js/dropdownToggle.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angular-file-upload/dist/angular-file-upload.js'],
                        dest: 'web/assets/js/angular-file-upload.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angular-mm-foundation/src/position/position.js'],
                        dest: 'web/assets/js/position.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angular-mm-foundation/src/typeahead/typeahead.js'],
                        dest: 'web/assets/js/typeahead.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angular-mm-foundation/src/bindHtml/bindHtml.js'],
                        dest: 'web/assets/js/bindHtml.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angular-mm-foundation/src/mediaQueries/mediaQueries.js'],
                        dest: 'web/assets/js/mediaQueries.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angular-mm-foundation/src/tabs/tabs.js'],
                        dest: 'web/assets/js/tabs.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angular-loading-bar/build/loading-bar.js'],
                        dest: 'web/assets/js/loading-bar.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angular-loading-bar/build/loading-bar.css'],
                        dest: 'web/assets/css/loading-bar.css',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angular-mm-foundation/src/transition/transition.js'],
                        dest: 'web/assets/js/transition.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/ng-sortable/dist/ng-sortable.js'],
                        dest: 'web/assets/js/ng-sortable.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angular/angular.js'],
                        dest: 'web/assets/js/angular.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angular-ui-router/release/angular-ui-router.min.js'],
                        dest: 'web/assets/js/angular-ui-router.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/ng-sortable/dist/ng-sortable.min.css'],
                        dest: 'web/assets/css/ng-sortable.min.css',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        src: ['node_modules/angularjs-datepicker/dist/angular-datepicker.min.js'],
                        dest: 'web/assets/js/angularjs-datepicker.min.js',
                        filter: 'isFile'
                    },
                    {
                        flatten: false,
                        expand: true,
                        cwd: 'node_modules/angular-mm-foundation/template',
                        src: '**',
                        dest: 'web/template/',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        expand: true,
                        cwd: 'node_modules/foundation-sites/js/foundation/',
                        src: '**',
                        dest: 'web/assets/js',
                        filter: 'isFile'
                    },
                    {
                        flatten: true,
                        expand: true,
                        cwd: 'node_modules/font-awesome/fonts/',
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
                        src: ['**/*.svg', '**/*.png'],
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
