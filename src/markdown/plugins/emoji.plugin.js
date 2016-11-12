(function(twemoji) {
    window.md_twemoji_plugin = function(md) {
        md.renderer.rules.emoji = function(token, idx) {
            return twemoji.parse(token[idx].content, {folder: "svg", ext: ".svg"});
        };
    };
}(window.twemoji));
