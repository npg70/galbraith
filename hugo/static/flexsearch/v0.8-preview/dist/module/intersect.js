import { create_object, concat, sort_by_length_up, get_max_len } from "./common.js";

/*

 from -> result[
    res[score][id],
    res[score][id],
 ]

 to -> [id]

*/

/**
 * @param arrays
 * @param {number} resolution
 * @param limit
 * @param offset
 * @param suggest
 * @returns {Array}
 */

export function intersect(arrays, resolution, limit, offset, suggest) {

    const length = arrays.length;

    // todo remove
    // if(length < 2){
    //     throw new Error("Not optimized intersect");
    // }

    let result = [],
        check,
        count;


    // alternatively the results could be sorted by length up
    //arrays.sort(sort_by_length_up);

    check = create_object();

    for (let y = 0, ids, id, res_arr, tmp; y < resolution; y++) {

        for (let x = 0; x < length; x++) {

            res_arr = arrays[x];

            if (y < res_arr.length && (ids = res_arr[y])) {

                for (let z = 0; z < ids.length; z++) {

                    id = ids[z];

                    if (count = check[id]) {
                        check[id]++;
                        // tmp.count++;
                        // tmp.sum += y;
                    } else {
                        count = 0;
                        check[id] = 1;
                        // check[id] = {
                        //     count: 1,
                        //     sum: y
                        // };
                    }

                    tmp = result[count] || (result[count] = []);
                    tmp.push(id);
                }
            }
        }
    }

    // result.sort(function(a, b){
    //     return check[a] - check[b];
    // });

    const result_len = result.length;

    if (result_len) {

        if (!suggest) {

            if (result_len < length) {
                return [];
            }

            result = result[result_len - 1];

            if (result.length > limit || offset) {
                result = result.slice(offset, limit + offset);
            }
        } else {

            result = 1 < result.length ? union(result, offset, limit) : result[0];
        }
    }

    return result;
}

function union(arrays, offset, limit) {
    const result = [],
          check = create_object();

    let ids,
        id,
        arr_len = arrays.length,
        ids_len;

    for (let i = 0; i < arr_len; i++) {

        ids = arrays[i];
        ids_len = ids.length;

        for (let j = 0; j < ids_len; j++) {
            id = ids[j];
            if (!check[id]) {
                check[id] = 1;
                if (offset) {
                    offset--;
                } else {
                    result.push(id);
                    if (result.length === limit) {
                        break;
                    }
                }
            }
        }
    }

    return result;
}

/**
 * @param {Array<number|string>} mandatory
 * @param {Array<Array<number|string>>} arrays
 * @returns {Array<number|string>}
 */

export function intersect_union(mandatory, arrays) {
    const check = create_object(),
          result = [];


    for (let x = 0, ids; x < arrays.length; x++) {
        ids = arrays[x];
        for (let i = 0; i < ids.length; i++) {
            check[ids[i]] = 1;
        }
    }

    for (let i = 0, id; i < mandatory.length; i++) {
        id = mandatory[i];
        if (1 === check[id]) {
            result.push(id);
            check[id] = 2;
        }
    }

    return result;
}

// export function intersect_union(mandatory, arrays, resolution) {
//
//     const check = create_object();
//     const union = create_object();
//     const result = [];
//
//     for(let x = 0; x < mandatory.length; x++){
//         check[mandatory[x]] = 1;
//     }
//
//
//     for(let y = 0, ids, id; y < resolution; y++){
//         for(let x = 0; x < arrays.length; x++){
//
//             ids = arrays[x];
//
//             if(y < ids.length){
//
//                 id = ids[y];
//
//                 if(check[id]){
//
//                     if(!union[id]){
//
//                         union[id] = 1;
//                         result.push(id);
//                     }
//                 }
//             }
//         }
//     }
//
//     return result;
// }

//
// /**
//  * Implementation based on Object[key] provides better suggestions
//  * capabilities and has less performance scaling issues on large indexes.
//  *
//  * @param arrays
//  * @param limit
//  * @param offset
//  * @param {boolean|Array=} suggest
//  * @param {boolean=} resolve
//  * @returns {Array}
//  */
//
// export function intersect(arrays, limit, offset, suggest, resolve) {
//
//     const length = arrays.length;
//
//     // todo remove
//     // if(length < 2){
//     //     throw new Error("Not optimized intersect");
//     // }
//
//     let result = [];
//     let size = 0;
//     let check;
//     let check_suggest;
//     let check_new;
//     let res_arr;
//
//     if(suggest){
//         suggest = [];
//     }
//
//     // 1. a reversed order prioritize the order of words from a query
//     //    because it ends with the first term.
//     // 2. process terms in reversed order often has advantage for
//     //    the fast path "end reached".
//
//     // alternatively the results could be sorted by length up
//     //arrays.sort(sort_by_length_up);
//
//     // todo the outer loop should be the res array instead of term array,
//     // this isn't just simple because the intersection calculation needs to reflect this
//     //const maxlen = get_max_len(arrays);
//
//     for(let x = length - 1, found; x >= 0; x--){
//     //for(let x = 0, found; x < length; x++){
//
//         res_arr = arrays[x];
//         check_new = create_object();
//         found = !check;
//
//         // process relevance in forward order (direction is
//         // important for adding IDs during the last round)
//
//         for(let y = 0, ids; y < res_arr.length; y++){
//
//             ids = res_arr[y];
//             if(!ids || !ids.length) continue;
//
//             for(let z = 0, id; z < ids.length; z++){
//
//                 id = ids[z];
//
//                 // check exists starting from the 2nd slot
//                 if(check){
//                     if(check[id]){
//
//                         // check if in last round
//                         if(!x){
//                         //if(x === length - 1){
//
//                             if(offset){
//                                 offset--;
//                             }
//                             else{
//
//                                 result[size++] = id;
//
//                                 if(size === limit){
//                                     // fast path "end reached"
//                                     return result /*resolve === false
//                                         ? { result, suggest }
//                                         :*/
//                                 }
//                             }
//                         }
//
//                         if(x || suggest){
//                         //if((x < length - 1) || suggest){
//                             check_new[id] = 1;
//                         }
//
//                         found = true;
//                     }
//
//                     if(suggest){
//
//                         if(!check_suggest[id]){
//                             check_suggest[id] = 1;
//                             const arr = suggest[y] || (suggest[y] = []);
//                             arr.push(id);
//                         }
//
//                         // OLD:
//                         //
//                         // check_idx = (check_suggest[id] || 0) + 1;
//                         // check_suggest[id] = check_idx;
//                         //
//                         // // do not adding IDs which are already included in the result (saves one loop)
//                         // // the first intersection match has the check index 2, so shift by -2
//                         //
//                         // if(check_idx < length){
//                         //
//                         //     const tmp = suggest[check_idx - 2] || (suggest[check_idx - 2] = []);
//                         //     tmp[tmp.length] = id;
//                         // }
//                     }
//                 }
//                 else{
//
//                     // pre-fill in first round
//                     check_new[id] = 1;
//                 }
//             }
//         }
//
//         if(suggest){
//
//             // re-use the first pre-filled check for suggestions
//             check || (check_suggest = check_new);
//         }
//         else if(!found){
//
//             return [];
//         }
//
//         check = check_new;
//     }
//
//     // return intermediate result
//     // if(resolve === false){
//     //     return { result, suggest };
//     // }
//
//     if(suggest){
//
//         // needs to iterate in reverse direction
//         for(let x = suggest.length - 1, ids, len; x >= 0; x--){
//
//             ids = suggest[x];
//             len = ids.length;
//
//             for(let y = 0, id; y < len; y++){
//
//                 id = ids[y];
//
//                 if(!check[id]){
//
//                     if(offset){
//                         offset--;
//                     }
//                     else{
//
//                         result[size++] = id;
//
//                         if(size === limit){
//                             // fast path "end reached"
//                             return result;
//                         }
//                     }
//
//                     check[id] = 1;
//                 }
//             }
//         }
//     }
//
//     return result;
// }

/**
 * Implementation based on Array.includes() provides better performance,
 * but it needs at least one word in the query which is less frequent.
 * Also on large indexes it does not scale well performance-wise.
 * This strategy also lacks of suggestion capabilities (matching & sorting).
 *
 * @param arrays
 * @param limit
 * @param offset
 * @param {boolean|Array=} suggest
 * @returns {Array}
 */

// export function intersect(arrays, limit, offset, suggest) {
//
//     const length = arrays.length;
//     let result = [];
//     let check;
//
//     // determine shortest array and collect results
//     // from the sparse relevance arrays
//
//     let smallest_size;
//     let smallest_arr;
//     let smallest_index;
//
//     for(let x = 0; x < length; x++){
//
//         const arr = arrays[x];
//         const len = arr.length;
//
//         let size = 0;
//
//         for(let y = 0, tmp; y < len; y++){
//
//             tmp = arr[y];
//
//             if(tmp){
//
//                 size += tmp.length;
//             }
//         }
//
//         if(!smallest_size || (size < smallest_size)){
//
//             smallest_size = size;
//             smallest_arr = arr;
//             smallest_index = x;
//         }
//     }
//
//     smallest_arr = smallest_arr.length === 1 ?
//
//         smallest_arr[0]
//     :
//         concat(smallest_arr);
//
//     if(suggest){
//
//         suggest = [smallest_arr];
//         check = create_object();
//     }
//
//     let size = 0;
//     let steps = 0;
//
//     // process terms in reversed order often results in better performance.
//     // the outer loop must be the words array, using the
//     // smallest array here disables the "fast fail" optimization.
//
//     for(let x = length - 1; x >= 0; x--){
//
//         if(x !== smallest_index){
//
//             steps++;
//
//             const word_arr = arrays[x];
//             const word_arr_len = word_arr.length;
//             const new_arr = [];
//
//             let count = 0;
//
//             for(let z = 0, id; z < smallest_arr.length; z++){
//
//                 id = smallest_arr[z];
//
//                 let found;
//
//                 // process relevance in forward order (direction is
//                 // important for adding IDs during the last round)
//
//                 for(let y = 0; y < word_arr_len; y++){
//
//                     const arr = word_arr[y];
//
//                     if(arr.length){
//
//                         found = arr.includes(id);
//
//                         if(found){
//
//                             // check if in last round
//
//                             if(steps === length - 1){
//
//                                 if(offset){
//
//                                     offset--;
//                                 }
//                                 else{
//
//                                     result[size++] = id;
//
//                                     if(size === limit){
//
//                                         // fast path "end reached"
//
//                                         return result;
//                                     }
//                                 }
//
//                                 if(suggest){
//
//                                     check[id] = 1;
//                                 }
//                             }
//
//                             break;
//                         }
//                     }
//                 }
//
//                 if(found){
//
//                     new_arr[count++] = id;
//                 }
//             }
//
//             if(suggest){
//
//                 suggest[steps] = new_arr;
//             }
//             else if(!count){
//
//                 return [];
//             }
//
//             smallest_arr = new_arr;
//         }
//     }
//
//     if(suggest){
//
//         // needs to iterate in reverse direction
//
//         for(let x = suggest.length - 1, arr, len; x >= 0; x--){
//
//             arr = suggest[x];
//             len = arr && arr.length;
//
//             if(len){
//
//                 for(let y = 0, id; y < len; y++){
//
//                     id = arr[y];
//
//                     if(!check[id]){
//
//                         check[id] = 1;
//
//                         if(offset){
//
//                             offset--;
//                         }
//                         else{
//
//                             result[size++] = id;
//
//                             if(size === limit){
//
//                                 // fast path "end reached"
//
//                                 return result;
//                             }
//                         }
//                     }
//                 }
//             }
//         }
//     }
//
//     return result;
// }