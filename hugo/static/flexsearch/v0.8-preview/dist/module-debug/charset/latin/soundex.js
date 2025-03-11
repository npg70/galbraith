import { EncoderOptions } from "../../type.js";

/** @type EncoderOptions */
const options = {
    normalize: !0,
    dedupe: !1,
    include: {
        letter: !0
    },
    finalize: function (arr) {
        for (let i = 0; i < arr.length; i++) {
            arr[i] = soundex(arr[i]);
        }
    }
};
export default options;

const codes = {
    a: "", e: "", i: "", o: "", u: "", y: "",
    b: 1, f: 1, p: 1, v: 1,
    c: 2, g: 2, j: 2, k: 2, q: 2, s: 2, x: 2, z: 2, ß: 2,
    d: 3, t: 3,
    l: 4,
    m: 5, n: 5,
    r: 6
};

function soundex(stringToEncode) {
    let encodedString = stringToEncode.charAt(0),
        last = codes[encodedString];

    for (let i = 1, char; i < stringToEncode.length; i++) {
        char = stringToEncode.charAt(i);
        // Remove all occurrences of "h" and "w"
        if ("h" !== char && "w" !== char) {
            // Replace all consonants with digits
            char = codes[char];
            // Remove all occurrences of a,e,i,o,u,y except first letter
            if (char) {
                // Replace all adjacent same digits with one digit
                if (char !== last) {
                    encodedString += char;
                    last = char;
                    if (4 === encodedString.length) {
                        break;
                    }
                }
            }
        }
    }
    // while(encodedString.length < 4){
    //     encodedString += "0";
    // }
    return encodedString;
}