import{IndexOptions}from"./type.js";import{create_object}from"./common.js";import handler from"./worker/handler.js";import apply_async from"./async.js";let pid=0;export default function WorkerIndex(a={}){function b(b){function f(a){a=a.data||a;const b=a.id,c=b&&e.resolver[b];c&&(c(a.msg),delete e.resolver[b])}if(this.worker=b,this.resolver=create_object(),!!this.worker)return(d?this.worker.on("message",f):this.worker.onmessage=f,a.config)?new Promise(function(b){e.resolver[++pid]=function(){b(e),1e9<pid&&(pid=0)},e.worker.postMessage({id:pid,task:"init",factory:c,options:a})}):(this.worker.postMessage({task:"init",factory:c,options:a}),this.priority=a.priority||4,this)}if(!this||this.constructor!==WorkerIndex)return new WorkerIndex(a);let c="undefined"==typeof self?"undefined"==typeof window?null:window._factory:self._factory;c&&(c=c.toString());const d="undefined"==typeof window,e=this,f=create(c,d,a.worker);return f.then?f.then(function(a){return b.call(e,a)}):b.call(this,f)}register("add"),register("append"),register("search"),register("update"),register("remove"),register("clear"),register("export"),register("import"),apply_async(WorkerIndex.prototype);function register(a){WorkerIndex.prototype[a]=function(){const b=this,c=[].slice.call(arguments),d=c[c.length-1];let e;"function"==typeof d&&(e=d,c.pop());const f=new Promise(function(d){"export"===a&&"function"==typeof c[0]&&(c[0]=null),b.resolver[++pid]=d,b.worker.postMessage({task:a,id:pid,args:c})});return e?(f.then(e),this):f}}function create(a,b,c){let d;return d=b?"undefined"==typeof module?import("worker_threads").then(function(worker){return new worker["Worker"](import.meta.dirname+"/../node/node.mjs")}):new(require("worker_threads")["Worker"])(__dirname+"/worker/node.js"):a?new window.Worker(URL.createObjectURL(new Blob(["onmessage="+handler.toString()],{type:"text/javascript"}))):new window.Worker("string"==typeof c?c:import.meta.url.replace("/worker.js","/worker/worker.js").replace("flexsearch.bundle.module.min.js","module/worker/worker.js"),{type:"module"}),d}