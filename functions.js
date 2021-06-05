// args: [{name: xxx, value: xxx}...]
// return ["full command", null] or [null, "error message"]

function brook_server(args){
  var s = "brook server";
  
  var listen = args.find(v=>v.name == 'listen');
  if(!listen.value) return [null, "missing listen"];
  s += " --listen '" + listen.value + "'";
  var password = args.find(v=>v.name == 'password');
  if(!password.value) return [null, "missing password"];
  s += " --password '" + password.value + "'";
  var tcpTimeout = args.find(v=>v.name == 'tcpTimeout');
  if(isNaN(parseInt(tcpTimeout.value))) return [null, "tcpTimeout must be int"];
  s += " --tcpTimeout " + parseInt(tcpTimeout.value);
  var udpTimeout = args.find(v=>v.name == 'udpTimeout');
  if(isNaN(parseInt(udpTimeout.value))) return [null, "udpTimeout must be int"];
  s += " --udpTimeout " + parseInt(udpTimeout.value);
  
  return [s, null];
}
