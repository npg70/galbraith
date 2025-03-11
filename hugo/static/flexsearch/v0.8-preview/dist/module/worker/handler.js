import Index from "../index.js";
import { IndexOptions } from "../type.js";

export default (async function (data) {

            data = data.data;

            /** @type Index */
            const index = self._index,
                  args = data.args,
                  task = data.task;


            switch (task) {

                        case "init":

                                    /** @type IndexOptions */
                                    let options = data.options || {},
                                        filepath = options.config;

                                    if (filepath) {
                                                options = filepath;
                                                // will be replaced after build with the line below because
                                                // there is an issue with closure compiler dynamic import
                                                //options = await import(filepath);
                                    }

                                    // deprecated:
                                    // const encode = options.encode;
                                    // if(encode && (encode.indexOf("function") === 0)){
                                    //     options.encode = Function("return " + encode)();
                                    // }

                                    const factory = data.factory;

                                    if (factory) {

                                                // export the FlexSearch global payload to "self"
                                                Function("return " + factory)()(self);

                                                /** @type Index */
                                                self._index = new self.FlexSearch.Index(options);

                                                // destroy the exported payload
                                                delete self.FlexSearch;
                                    } else {

                                                self._index = new Index(options);
                                    }

                                    postMessage({ id: data.id });
                                    break;

                        default:
                                    const id = data.id,
                                          message = index[task].apply(index, args);

                                    postMessage("search" === task ? { id: id, msg: message } : { id: id });
            }
});