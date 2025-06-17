import { EncoderOptions } from "../../type.js";

/** @type EncoderOptions */
const options = {
    normalize: function (str) {
        return str.toLowerCase();
    }
    //numeric: false,
    //dedupe: false
};
export default options;