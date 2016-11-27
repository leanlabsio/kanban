/**
 * Card link plugin for generate link to other card
 */
(function(){
    window.md_user_link_plugin = function(md) {
        user_rule = function(state, silent){
            var start = state.pos,
                regex = /@([a-z0-9\._-]+)/,
                max   = state.posMax;

            if (state.src.charCodeAt(start) !== 0x40 || silent) {
                return false;
            }

            var matches = state.src.match(regex);

            if (! matches) {
                return false;
            }

            var login = matches[1];
            var title = matches[0];
            var board_name = state.env.host_url + '/' + state.env.board_url;

            token = state.push('link_open', 'a');
            token.attrPush(['title', title]);
            token.attrPush(['href', state.env.host_url + '/' + login]);
            token.nesting = 1;

            token = state.push('text');
            token.content = title;
            token.nesting = 0;

            token = state.push('link_close', 'a');
            token.nesting = -1;
            state.pos = start + title.length;
            return true;
        };

        md.inline.ruler.push('user_rule', user_rule);
    };
}());
