const{parentPort}=require("worker_threads"),{join}=require("path"),{Index}=require("../flexsearch.bundle.min.js");let index;parentPort.on("message",function(a){const b=a.args,c=a.task,d=a.id;switch(c){case"init":let e=a.options||{},f=e.config;if(f&&"/"!==f[0]&&"\\"!==f[0]){const a=process.cwd();f=join(a,f)}f&&(e=require(f)),index=new Index(e);break;default:const g=index[c].apply(index,b);parentPort.postMessage("search"===c?{id:d,msg:g}:{id:d});}});