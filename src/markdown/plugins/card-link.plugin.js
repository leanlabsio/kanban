/**
 * Card link plugin for generate link to other card
 */
(function(){
    window.md_card_link_plugin = function(md) {
        function card_rule(state, silent) {
            var start = state.pos,
                regex = /([a-z0-9\.-_\/]+)?#(\d+)/,
                max   = state.posMax;

            if (state.src.charCodeAt(start) !== 0x23 || silent) {
                return false;
            }

            var matches = state.src.match(regex);

            if (! matches) {
                return false;
            }

            var id = matches[2];
            var title = matches[0];
            var board_name = state.env.board_url;
            if (matches[1]) {
                board_name = matches[1];
                state.pending = state.pending.slice(0, -board_name.length);
            }

            token = state.push('link_open', 'a');
            token.attrPush(['title', title]);
            token.attrPush(['href', '/boards/' + board_name + '/issues/' + id]);
            token.attrPush(['data-link-local', '']);
            token.nesting = 1;

            token = state.push('text');
            token.content = title;
            token.nesting = 0;

            token = state.push('link_close', 'a');
            token.nesting = -1;
            state.pos = start + id.length + 1;
            return true;
        }

        md.inline.ruler.push('card_rule', card_rule);
    };
}());
