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
