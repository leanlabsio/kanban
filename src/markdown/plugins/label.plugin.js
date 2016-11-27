/**
 * Milestone link plugin for generate link to other milestone
 */
(function(){
    window.md_label_plugin = function(md) {
        function label_rule(state, silent) {
            var start = state.pos,
                regex = /~(\d+)/,
                max   = state.posMax;

            if (state.src.charCodeAt(start) !== 0x7E || silent) {
                return false;
            }

            var matches = state.src.match(regex);

            if (! matches) {
                return false;
            }

            var id = matches[1];
            var title = matches[0];

            var label = _.find(state.env.labels, function(label) {
                return label.id == id;
            });

            if (label) {
                token = state.push('span_open', 'span');
                token.attrPush(['style', 'background-color:' + label.color]);
                token.attrPush(['class', 'label']);
                token.nesting = 1;

                token = state.push('text');
                token.content = label.name;
                token.nesting = 0;

                token = state.push('span_open', 'span');
                token.nesting = -1;
            }

            state.pos = start + id.length + 1;
            return true;
        }

        md.inline.ruler.push('label_rule', label_rule);
    };
}());
