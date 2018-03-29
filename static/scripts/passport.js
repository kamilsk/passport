(function (base, signer, sender){
    'use strict';

    var payload = {'fingerprint': undefined, 'metadata': undefined, 'counter': undefined}, synced = false;

    function notify() {
        var url = base;
        if (url.substr(-1) === '/') { url = url.substr(0, url.length - 1); }
        url += '{{ .Endpoint }}';
        sender({
            type: 'POST',
            url: url,
            data: JSON.stringify(payload),
            contentType: 'application/json; charset=utf-8',
            success: function () { synced = true; }
        });
    }

    var corrector = setInterval(function () {
        new signer().get(function(result, components) {
            if (result !== payload.fingerprint) {
                payload.fingerprint = result;
                payload.metadata = components;
                payload.counter = 0;
            }
            if (++payload.counter > {{ .Threshold }}) {
                clearInterval(corrector);
                synced || notify();
            }
        })
    }, {{ .Correct }});

    var watcher = setInterval(function () {
        synced && clearInterval(watcher);
        synced || notify();
    }, {{ .Watch }});
}('{{ .BaseURL }}', window.Fingerprint2, window.jQuery.ajax));
