// args: [{name: xxx, value: xxx}...]
// return ["full command", null] or [null, "error message"]

function brook_server(args){
  var s = 'brook server';
  args.forEach(function(v){
    s += " --"+v.name + " " + v.value;
  });
  return [s, null];
}
