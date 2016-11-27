/**
 * Milestone link plugin for generate link to other milestone
 */
(function(){
    window.md_milestone_link_plugin = function(md) {
        function milestone_rule(state, silent) {
            var start = state.pos,
                regex = /([a-z0-9\.-_\/]+)?%(\d+)/,
                max   = state.posMax;

            if (state.src.charCodeAt(start) !== 0x25 || silent) {
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
                board_name =  matches[1];
                state.pending = state.pending.slice(0, -matches[1].length);
            }
            var milestone = _.find(state.env.milestones, function(milestone) {
                return milestone.iid == id;
            });

            token = state.push('link_open', 'a');
            token.attrPush(['title', title]);
            token.attrPush(['href', '/boards/' + board_name + '?tags=^' + milestone.id]);
            token.attrPush(['data-link-local', '']);
            token.nesting = 1;

            token = state.push('text');
            if (milestone) {
                token.content = milestone.title;
            } else {
                token.content = title;
            }
            token.nesting = 0;

            token = state.push('link_close', 'a');
            token.nesting = -1;
            state.pos = start + id.length + 1;
            return true;
        }

        md.inline.ruler.push('card_rule', milestone_rule);
    };
}());
