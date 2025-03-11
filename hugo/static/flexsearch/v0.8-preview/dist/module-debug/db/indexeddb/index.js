
import Document from "../../document.js";
import { PersistentOptions } from "../../type.js";

const VERSION = 1,
      IndexedDB = "undefined" != typeof window && (window.indexedDB || window.mozIndexedDB || window.webkitIndexedDB || window.msIndexedDB),
      IDBTransaction = "undefined" != typeof window && (window.IDBTransaction || window.webkitIDBTransaction || window.msIDBTransaction),
      IDBKeyRange = "undefined" != typeof window && (window.IDBKeyRange || window.webkitIDBKeyRange || window.msIDBKeyRange),
      fields = ["map", "ctx", "tag", "reg", "cfg"];

import StorageInterface from "../interface.js";
import { toArray } from "../../common.js";

function sanitize(str) {
    return str.toLowerCase().replace(/[^a-z0-9_\-]/g, "");
}

/**
 * @param {string|PersistentOptions=} name
 * @param {PersistentOptions=} config
 * @constructor
 * @implements StorageInterface
 */

export default function IdxDB(name, config = {}) {
    if (!this) {
        return new IdxDB(name, config);
    }
    if ("object" == typeof name) {
        name = name.name;
        config = name;
    }
    if (!name) {
        console.info("Default storage space was used, because a name was not passed.");
    }
    this.id = "flexsearch" + (name ? ":" + sanitize(name) : "");
    this.field = config.field ? sanitize(config.field) : "";
    this.support_tag_search = !1;
    this.db = null;
    this.trx = {};
}

IdxDB.prototype.mount = function (flexsearch) {
    if (flexsearch.constructor === Document) {
        return flexsearch.mount(this);
    }
    flexsearch.db = this;
    return this.open();
};

IdxDB.prototype.open = function () {

    let self = this;

    navigator.storage && navigator.storage.persist();

    return this.db || new Promise(function (resolve, reject) {

        const req = IndexedDB.open(self.id + (self.field ? ":" + self.field : ""), VERSION);

        req.onupgradeneeded = function () {

            const db = self.db = this.result;

            // Using Indexes + IDBKeyRange on schema map => [key, res, id] performs
            // too bad and blows up amazingly in size
            // The schema map:key => [res][id] is currently used instead
            // In fact that bypass the idea of a storage solution,
            // IndexedDB is such a poor contribution :(

            fields.forEach(ref => {
                db.objectStoreNames.contains(ref) || db.createObjectStore(ref); //{ autoIncrement: true /*keyPath: "id"*/ }
                //.createIndex("idx", "ids", { multiEntry: true, unique: false });
            });

            // switch(event.oldVersion){ // existing db version
            //     case 0:
            //     // version 0 means that the client had no database
            //     // perform initialization
            //     case 1:
            //     // client had version 1
            //     // update
            // }
        };

        req.onblocked = function (event) {
            // this event shouldn't trigger if we handle onversionchange correctly
            // it means that there's another open connection to the same database
            // and it wasn't closed after db.onversionchange triggered for it
            console.error("blocked", event);
            reject();
        };

        req.onerror = function (event) {
            console.error(this.error, event);
            reject();
        };

        req.onsuccess = function () {
            self.db = this.result; //event.target.result;
            self.db.onversionchange = function () {
                //database is outdated
                self.close();
            };
            resolve(self);
        };
    });
};

IdxDB.prototype.close = function () {
    this.db.close();
    this.db = null;
};

IdxDB.prototype.destroy = function () {
    this.db && this.close();
    return IndexedDB.deleteDatabase(this.id + (this.field ? ":" + this.field : ""));
};

// IdxDB.prototype.set = function(ref, key, ctx, data){
//     const transaction = this.db.transaction(ref, "readwrite");
//     const map = transaction.objectStore(ref);
//     const req = map.put(data, ctx ? ctx + ":" + key : key);
//     return transaction;//promisfy(req, callback);
// };

// IdxDB.prototype.delete = function(ref, key, ctx){
//     const transaction = this.db.transaction(ref, "readwrite");
//     const map = transaction.objectStore(ref);
//     const req = map.delete(ctx ? ctx + ":" + key : key);
//     return transaction;//promisfy(req, callback);
// };

IdxDB.prototype.clear = function () {
    const transaction = this.db.transaction(fields, "readwrite");
    for (let i = 0; i < fields.length; i++) {
        transaction.objectStore(fields[i]).clear();
    }
    return promisfy(transaction);
};

IdxDB.prototype.get = function (key, ctx, limit = 0, offset = 0, resolve = /* tag? */!0, enrich = !1) {
    const transaction = this.db.transaction(ctx ? "ctx" : "map", "readonly"),
          map = transaction.objectStore(ctx ? "ctx" : "map"),
          req = map.get(ctx ? ctx + ":" + key : key),
          self = this;

    return promisfy(req).then(function (res) {
        let result = [];
        if (!res || !res.length) return result;
        if (resolve) {
            if (!limit && !offset && 1 === res.length) {
                return res[0];
            }
            for (let i = 0, arr; i < res.length; i++) {
                if ((arr = res[i]) && arr.length) {
                    if (offset >= arr.length) {
                        offset -= arr.length;
                        continue;
                    }
                    const end = limit ? offset + Math.min(arr.length - offset, limit) : arr.length;
                    for (let j = offset; j < end; j++) {
                        result.push(arr[j]);
                    }
                    offset = 0;
                    if (result.length === limit) {
                        break;
                    }
                }
            }
            return enrich ? self.enrich(result) : result;
        } else {
            return res;
        }
    });
};

IdxDB.prototype.tag = function (tag, limit = 0, offset = 0, enrich = !1) {
    const transaction = this.db.transaction("tag", "readonly"),
          map = transaction.objectStore("tag"),
          req = map.get(tag),
          self = this;

    return promisfy(req).then(function (ids) {
        if (!ids || !ids.length || offset >= ids.length) return [];
        if (!limit && !offset) return ids;
        const result = ids.slice(offset, offset + limit);
        return enrich ? self.enrich(result) : result;
    });
};


IdxDB.prototype.enrich = function (ids) {
    if ("object" != typeof ids) {
        ids = [ids];
    }
    const transaction = this.db.transaction("reg", "readonly"),
          map = transaction.objectStore("reg"),
          promises = [];

    for (let i = 0; i < ids.length; i++) {
        promises[i] = promisfy(map.get(ids[i]));
    }
    return Promise.all(promises).then(function (docs) {
        for (let i = 0; i < docs.length; i++) {
            docs[i] = {
                id: ids[i],
                doc: docs[i] ? JSON.parse(docs[i]) : null
            };
        }
        return docs;
    });
};


IdxDB.prototype.has = function (id) {
    const transaction = this.db.transaction("reg", "readonly"),
          map = transaction.objectStore("reg"),
          req = map.getKey(id);

    return promisfy(req);
};

IdxDB.prototype.search = null;

// IdxDB.prototype.has = function(ref, key, ctx){
//     const transaction = this.db.transaction(ref, "readonly");
//     const map = transaction.objectStore(ref);
//     const req = map.getKey(ctx ? ctx + ":" + key : key);
//     return promisfy(req);
// };

IdxDB.prototype.info = function () {
    // todo
};

/**
 * @param {!string} ref
 * @param {!string} modifier
 * @param {!Function} task
 */

IdxDB.prototype.transaction = function (ref, modifier, task) {

    let store = this.trx[ref + ":" + modifier];
    if (store) return task.call(this, store);

    let transaction = this.db.transaction(ref, modifier);
    this.trx[ref + ":" + modifier] = store = transaction.objectStore(ref);

    return new Promise((resolve, reject) => {
        transaction.onerror = err => {
            this.trx[ref + ":" + modifier] = null;
            transaction.abort();
            transaction = store = null;
            reject(err);
            //db.close;
        };
        transaction.oncomplete = res => {
            this.trx[ref + ":" + modifier] = null;
            transaction = store = null;
            resolve(res || !0);
            //db.close;
        };
        return task.call(this, store);
    });
};

IdxDB.prototype.commit = async function (flexsearch, _replace, _append) {

    // process cleanup tasks
    if (_replace) {
        await this.clear();
        // there are just removals in the task queue
        flexsearch.commit_task = [];
    } else {
        let tasks = flexsearch.commit_task;
        flexsearch.commit_task = [];
        for (let i = 0, task; i < tasks.length; i++) {
            task = tasks[i];
            // there are just removals in the task queue
            if (task.clear) {
                await this.clear();
                _replace = !0;
                break;
            } else {
                tasks[i] = task.del;
            }
        }
        if (!_replace) {
            if (!_append) {
                tasks = tasks.concat(toArray(flexsearch.reg));
            }
            tasks.length && (await this.remove(tasks));
        }
    }

    if (!flexsearch.reg.size) {
        return;
    }

    await this.transaction("map", "readwrite", function (store) {

        for (const item of flexsearch.map) {
            const key = item[0],
                  value = item[1];

            if (!value.length) continue;

            if (_replace) {
                store.put(value, key);
                continue;
            }

            store.get(key).onsuccess = function () {
                let result = this.result,
                    changed;

                if (result && result.length) {
                    const maxlen = Math.max(result.length, value.length);
                    for (let i = 0, res, val; i < maxlen; i++) {
                        val = value[i];
                        if (val && val.length) {
                            res = result[i];
                            if (res && res.length) {
                                for (let j = 0; j < val.length; j++) {
                                    res.push(val[j]);
                                }
                                changed = 1;
                                //result[i] = res.concat(val);
                                //result[i] = new Set([...result[i], ...value[i]]);
                                //result[i] = result[i].union(new Set(value[i]));
                            } else {
                                result[i] = val;
                                changed = 1;
                                //result[i] = new Set(value[i])
                            }
                        }
                    }
                } else {
                    result = value;
                    changed = 1;
                    //result = [];
                    //for(let i = 0; i < value.length; i++){
                    //    if(value[i]) result[i] = new Set(value[i]);
                    //}
                }

                changed && store.put(result, key);
            };
        }
    });

    await this.transaction("ctx", "readwrite", function (store) {

        for (const ctx of flexsearch.ctx) {
            const ctx_key = ctx[0],
                  ctx_value = ctx[1];


            for (const item of ctx_value) {
                const key = item[0],
                      value = item[1];

                if (!value.length) continue;

                if (_replace) {
                    store.put(value, ctx_key + ":" + key);
                    continue;
                }

                store.get(ctx_key + ":" + key).onsuccess = function () {
                    let result = this.result,
                        changed;

                    if (result && result.length) {
                        const maxlen = Math.max(result.length, value.length);
                        for (let i = 0, res, val; i < maxlen; i++) {
                            val = value[i];
                            if (val && val.length) {
                                res = result[i];
                                if (res && res.length) {
                                    for (let j = 0; j < val.length; j++) {
                                        res.push(val[j]);
                                    }
                                    //result[i] = res.concat(val);
                                    changed = 1;
                                } else {
                                    result[i] = val;
                                    changed = 1;
                                }
                            }
                        }
                    } else {
                        result = value;
                        changed = 1;
                    }

                    changed && store.put(result, ctx_key + ":" + key);
                };
            }
        }
    });

    if (flexsearch.store) {
        await this.transaction("reg", "readwrite", function (store) {
            for (const item of flexsearch.store) {
                const id = item[0],
                      doc = item[1];

                store.put("object" == typeof doc ? JSON.stringify(doc) : 1, id);
            }
        });
    } else if (!flexsearch.bypass) {
        await this.transaction("reg", "readwrite", function (store) {
            for (const id of flexsearch.reg.keys()) {
                store.put(1, id);
            }
        });
    }

    if (flexsearch.tag) {
        await this.transaction("tag", "readwrite", function (store) {
            for (const item of flexsearch.tag) {
                const tag = item[0],
                      ids = item[1];

                if (!ids.length) continue;

                store.get(tag).onsuccess = function () {
                    let result = this.result;
                    result = result && result.length ? result.concat(ids) : ids;
                    store.put(result, tag);
                };
            }
        });
    }

    // TODO
    // await this.transaction("cfg", "readwrite", function(store){
    //     store.put({
    //         "charset": flexsearch.charset,
    //         "tokenize": flexsearch.tokenize,
    //         "resolution": flexsearch.resolution,
    //         "fastupdate": flexsearch.fastupdate,
    //         "compress": flexsearch.compress,
    //         "encoder": {
    //             "minlength": flexsearch.encoder.minlength
    //         },
    //         "context": {
    //             "depth": flexsearch.depth,
    //             "bidirectional": flexsearch.bidirectional,
    //             "resolution": flexsearch.resolution_ctx
    //         }
    //     }, "current");
    // });

    flexsearch.map.clear();
    flexsearch.ctx.clear();

    flexsearch.tag && flexsearch.tag.clear();

    flexsearch.store && flexsearch.store.clear();

    flexsearch.document || flexsearch.reg.clear();
};

/**
 * @param {IDBCursorWithValue} cursor
 * @param {Array} ids
 * @param {boolean=} _tag
 */

function handle(cursor, ids, _tag) {

    const arr = cursor.value;
    let changed,
        parse,
        count = 0;


    for (let x = 0, result; x < arr.length; x++) {
        // tags has no resolution layer
        if (result = _tag ? arr : arr[x]) {
            for (let i = 0, pos, id; i < ids.length; i++) {
                id = ids[i];
                pos = result.indexOf(parse ? parseInt(id, 10) : id);
                if (0 > pos && !parse && "string" == typeof id && !isNaN(id)) {
                    pos = result.indexOf(parseInt(id, 10));
                    pos && (parse = 1);
                }
                if (0 <= pos) {
                    changed = 1;
                    if (1 < result.length) {
                        result.splice(pos, 1);
                    } else {
                        arr[x] = [];
                        break;
                    }
                }
            }

            count += result.length;
        }
        if (_tag) break;
    }

    if (!count) {

        cursor.delete();
        //store.delete(cursor.key);
    } else if (changed) {

        //await new Promise(resolve => {
        cursor.update(arr); //.onsuccess = resolve;
        //});
    }

    cursor.continue();
}

IdxDB.prototype.remove = function (ids) {

    if ("object" != typeof ids) {
        ids = [ids];
    }

    return (/** @type {!Promise<undefined>} */Promise.all([this.transaction("map", "readwrite", function (store) {
            store.openCursor().onsuccess = function () {
                const cursor = this.result;
                cursor && handle(cursor, ids);
            };
        }), this.transaction("ctx", "readwrite", function (store) {
            store.openCursor().onsuccess = function () {
                const cursor = this.result;
                cursor && handle(cursor, ids);
            };
        }), this.transaction("tag", "readwrite", function (store) {
            store.openCursor().onsuccess = function () {
                const cursor = this.result;
                cursor && handle(cursor, ids, !0);
            };
        }),
        // let filtered = [];
        this.transaction("reg", "readwrite", function (store) {
            for (let i = 0; i < ids.length; i++) {
                store.delete(ids[i]);
            }
            // return new Promise(resolve => {
            //     store.openCursor().onsuccess = function(){
            //         const cursor = this.result;
            //         if(cursor){
            //             const id = cursor.value;
            //             if(ids.includes(id)){
            //                 filtered.push(id);
            //                 cursor.delete();
            //             }
            //             cursor.continue();
            //         }
            //         else{
            //             resolve();
            //         }
            //     };
            // });
        })
        // ids = filtered;
        ])
    );
};

/**
 * @param {IDBRequest} req
 * @param {Function=} callback
 * @return {!Promise<undefined>}
 */

function promisfy(req, callback) {
    return new Promise((resolve, reject) => {
        /** @this IDBRequest */
        req.onsuccess = function () {
            callback && callback(this.result);
            resolve(this.result);
        };
        /** @this IDBRequest */
        req.oncomplete = function () {
            callback && callback(this.result);
            resolve(this.result);
        };
        req.onerror = reject;
        req = null;
    });
}