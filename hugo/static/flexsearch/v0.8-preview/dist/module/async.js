import Document from "./document.js";
import Index from "./index.js";

export default function (prototype) {
    register.call(prototype, "add");
    register.call(prototype, "append");
    register.call(prototype, "search");
    register.call(prototype, "update");
    register.call(prototype, "remove");
}

// let cycle;
// let budget = 0;
//
// function tick(resolve){
//     cycle = null;
//     budget = 0;
//     resolve();
// }

/**
 * @param {!string} key
 * @this {Index|Document}
 */

function register(key) {
    this[key + "Async"] = function () {

        // // prevent stack overflow of adding too much tasks to the same event loop
        // // actually limit stack to 1,000,000 tasks every ~4ms
        // cycle || (
        //     cycle = new Promise(resolve => setTimeout(tick, 0, resolve))
        // );
        //
        // // apply different performance budgets
        // if(key === "update" || key === "remove" && this.fastupdate === false){
        //     budget += 1000 * this.resolution;
        //     if(this.depth)
        //         budget += 1000 * this.resolution_ctx;
        // }
        // else if(key === "search"){
        //     budget++;
        // }
        // else{
        //     budget += 20 * this.resolution;
        //     if(this.depth)
        //         budget += 20 * this.resolution_ctx;
        // }
        //
        // // wait for the event loop cycle
        // if(budget >= 1e6){
        //     await cycle;
        // }

        const args = /*[].slice.call*/arguments,
              arg = args[args.length - 1];

        let callback;

        if ("function" == typeof arg) {
            callback = arg;
            delete args[args.length - 1];
        }

        this.async = !0;
        const res = this[key].apply(this, args);
        this.async = !1;
        if (callback) {
            res.then ? res.then(callback) : callback(res);
        }
        return res;
    };
}