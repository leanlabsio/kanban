function merge_request_plugin(md) {
    function mrrule(state, silent) {
        var start = state.pos,
            regex = /!\d+/,
            max   = state.posMax;

        if (state.src.charCodeAt(start) !== 0x21 || silent) {
            return false;
        }

        var matches = state.src.match(regex);

        if (! matches) {
            return false;
        }

        var match = matches[0];
        var mr    = match.slice(1);

        token = state.push('link_open', 'a');
        token.attrPush(['title', match]);
        token.attrPush(['href', state.env.host_url + '/merge_requests/' + mr]);
        token.attrPush(['target', '_blank']);
        token.nesting = 1;

        token = state.push('text');
        token.content = match;
        token.nesting = 0;

        token = state.push('link_close', 'a');
        token.nesting = -1;

        state.pos = start + match.length;
        return true;
    }

    md.inline.ruler.push('mrrule', mrrule);
}

