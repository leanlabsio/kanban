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
    window.md_link_target_blank_plugin = function(md) {
        var defaultRenderer = md.renderer.rules.link_open || function(token, idx, options, env, self) {
            return self.renderToken(token, idx, options);
        };

        function imageRenderer(tokens, idx, options, env, self) {
            // If you are sure other plugins can't add `target` - drop check below
            var aIndex = tokens[idx].attrIndex('target');

            if (aIndex < 0) {
                tokens[idx].attrPush(['target', '_blank']); // add new attribute
            } else {
                tokens[idx].attrs[aIndex][1] = '_blank';    // replace value of existing attr
            }

            // pass token to default renderer.
            return defaultRenderer(tokens, idx, options, env, self);
        }

        md.renderer.rules.link_open = imageRenderer;
    };
}());
