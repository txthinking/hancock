// args: [{"name": "xxx", "value": "xxx"}...]
// return ["full command", null] or [null, "error message"]

function brook_server(args){
    var s = "brook server";
    for(var i=0; i<args.length; i++){
        if(args[i].name == '--listen'){
            if(!args[i].value){
                    return [null, "missing listen"];
            }
            s += " --listen '" + args[i].value + "'";
        }
        if(args[i].name == '--password'){
            if(!args[i].value){
                    return [null, "missing password"];
            }
            s += " --password '" + args[i].value + "'";
        }
        if(args[i].name == '--tcpTimeout'){
            if(isNaN(parseInt(args[i].value))){
                    return [null, "tcpTimeout must be int"];
            }
            s += " --tcpTimeout " + parseInt(args[i].value);
        }
        if(args[i].name == '--udpTimeout'){
            if(isNaN(parseInt(args[i].value))){
                    return [null, "udpTimeout must be int"];
            }
            s += " --udpTimeout " + parseInt(args[i].value);
        }
    }
    return [s, null];
}

function brook_wsserver(args){
    var s = "brook wsserver";
    for(var i=0; i<args.length; i++){
        if(args[i].name == '--listen'){
            if(!args[i].value){
                    return [null, "missing listen"];
            }
            s += " --listen '" + args[i].value + "'";
        }
        if(args[i].name == '--password'){
            if(!args[i].value){
                    return [null, "missing password"];
            }
            s += " --password '" + args[i].value + "'";
        }
        if(args[i].name == '--path'){
            if(!args[i].value){
                    return [null, "missing path"];
            }
            s += " --path '" + args[i].value + "'";
        }
        if(args[i].name == '--tcpTimeout'){
            if(isNaN(parseInt(args[i].value))){
                    return [null, "tcpTimeout must be int"];
            }
            s += " --tcpTimeout " + parseInt(args[i].value);
        }
        if(args[i].name == '--udpTimeout'){
            if(isNaN(parseInt(args[i].value))){
                    return [null, "udpTimeout must be int"];
            }
            s += " --udpTimeout " + parseInt(args[i].value);
        }
    }
    return [s, null];
}

function brook_wssserver(args){
    var s = "brook wssserver";
    for(var i=0; i<args.length; i++){
        if(args[i].name == '--domain'){
            if(!args[i].value){
                    return [null, "missing domain"];
            }
            s += " --domain '" + args[i].value + "'";
        }
        if(args[i].name == '--password'){
            if(!args[i].value){
                    return [null, "missing password"];
            }
            s += " --password '" + args[i].value + "'";
        }
        if(args[i].name == '--path'){
            if(!args[i].value){
                    return [null, "missing path"];
            }
            s += " --path '" + args[i].value + "'";
        }
        if(args[i].name == '--tcpTimeout'){
            if(isNaN(parseInt(args[i].value))){
                    return [null, "tcpTimeout must be int"];
            }
            s += " --tcpTimeout " + parseInt(args[i].value);
        }
        if(args[i].name == '--udpTimeout'){
            if(isNaN(parseInt(args[i].value))){
                    return [null, "udpTimeout must be int"];
            }
            s += " --udpTimeout " + parseInt(args[i].value);
        }
    }
    return [s, null];
}

function brook_socks5(args){
    var s = "brook socks5";
    for(var i=0; i<args.length; i++){
        if(args[i].name == '--socks5'){
            if(!args[i].value){
                    return [null, "missing socks5"];
            }
            s += " --socks5 '" + args[i].value + "'";
        }
        if(args[i].name == '--username'){
            if(args[i].value){
                s += " --username '" + args[i].value + "'";
            }
        }
        if(args[i].name == '--password'){
            if(args[i].value){
                s += " --password '" + args[i].value + "'";
            }
        }
        if(args[i].name == '--tcpTimeout'){
            if(isNaN(parseInt(args[i].value))){
                    return [null, "tcpTimeout must be int"];
            }
            s += " --tcpTimeout " + parseInt(args[i].value);
        }
        if(args[i].name == '--udpTimeout'){
            if(isNaN(parseInt(args[i].value))){
                    return [null, "udpTimeout must be int"];
            }
            s += " --udpTimeout " + parseInt(args[i].value);
        }
    }
    return [s, null];
}

function brook_relay(args){
    var s = "brook relay";
    for(var i=0; i<args.length; i++){
        if(args[i].name == '--from'){
            if(!args[i].value){
                    return [null, "missing from"];
            }
            s += " --from '" + args[i].value + "'";
        }
        if(args[i].name == '--to'){
            if(!args[i].value){
                    return [null, "missing to"];
            }
            s += " --to '" + args[i].value + "'";
        }
        if(args[i].name == '--tcpTimeout'){
            if(isNaN(parseInt(args[i].value))){
                    return [null, "tcpTimeout must be int"];
            }
            s += " --tcpTimeout " + parseInt(args[i].value);
        }
        if(args[i].name == '--udpTimeout'){
            if(isNaN(parseInt(args[i].value))){
                    return [null, "udpTimeout must be int"];
            }
            s += " --udpTimeout " + parseInt(args[i].value);
        }
    }
    return [s, null];
}
