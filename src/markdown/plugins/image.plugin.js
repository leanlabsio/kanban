/**
 * Markdown-it extension
 *
 * Wraps default image renderer to provide additional
 * src attribute checks.
 *
 * GitLab returns relative images paths, but to display images
 * correctly we need absolute paths including path to repo.
 *
 * This plugin checks if path is relative and prefixes src
 * with GITLAB_HOST / GITLAB_REPO_URL if so.
 */
(function() {
    window.md_image_plugin = function(md) {
        var defaultRenderer = md.renderer.rules.image;

        function imageRenderer(tokens, idx, options, env, self) {
            var token = tokens[idx],
                srcIndex = token.attrIndex('src'),
                srcAttr  = token.attrs[srcIndex][1];

            if (srcAttr.indexOf("http://") !== 0 && srcAttr.indexOf("https://") !== 0) {
                token.attrs[srcIndex][1] = env.host_url + '/' + env.board_url + srcAttr;
            }

            return defaultRenderer(tokens, idx, options, env, self);
        }

        md.renderer.rules.image = imageRenderer;
    };
}());
