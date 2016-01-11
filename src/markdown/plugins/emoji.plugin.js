(function(twemoji) {
    window.md_twemoji_plugin = function(md) {
        md.renderer.rules.emoji = function(token, idx) {
            return twemoji.parse(token[idx].content, {base: "/", folder: "images/twemoji/svg", ext: ".svg"});
        };
    };
}(window.twemoji));
