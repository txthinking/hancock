// args: [{name: xxx, value: xxx}...]
// return ["full command", null] or [null, "error message"]

function brook_server(args){
  
  var listen = args.find(v=>v.name == 'listen');
  if(!listen.value) return [null, "missing listen"];
  var password = args.find(v=>v.name == 'password');
  if(!password.value) return [null, "missing password"];
  var tcpTimeout = args.find(v=>v.name == 'tcpTimeout');
  if(isNaN(parseInt(tcpTimeout.value))) return [null, "tcpTimeout must be int"];
  var udpTimeout = args.find(v=>v.name == 'udpTimeout');
  if(isNaN(parseInt(udpTimeout.value))) return [null, "udpTimeout must be int"];
  
  var s = `brook server --listen ${listen.value} --password ${password.value} --tcpTimeout ${parseInt(tcpTimeout.value))} --udpTimeout ${parseInt(udpTimeout.value))}`;
  return [s, null];
}
