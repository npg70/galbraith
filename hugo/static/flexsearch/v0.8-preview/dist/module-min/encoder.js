import{parse_option}from"./common.js";import normalize_polyfill from"./charset/normalize.js";import{EncoderOptions}from"./type.js";const whitespace=/[^\p{L}\p{N}]+/u,numeric_split_length=/(\d{3})/g,numeric_split_prev_char=/(\D)(\d{3})/g,numeric_split_next_char=/(\d{3})(\D)/g,normalize=/[\u0300-\u036f]/g;export default function Encoder(){if(!this)return new Encoder(...arguments);for(let a=0;a<arguments.length;a++)this.assign(arguments[a])}Encoder.prototype.assign=function(a){this.normalize=parse_option(a.normalize,!0,this.normalize);let b=a.include,c=b||a.exclude||a.split;if("object"==typeof c){let d=!b,e="";a.include||(e+="\\p{Z}"),c.letter&&(e+="\\p{L}"),c.number&&(e+="\\p{N}",d=!!b),c.symbol&&(e+="\\p{S}"),c.punctuation&&(e+="\\p{P}"),c.control&&(e+="\\p{C}"),(c=c.char)&&(e+="object"==typeof c?c.join(""):c);try{this.split=new RegExp("["+(b?"^":"")+e+"]+","u")}catch(a){this.split=/\s+/}this.numeric=d}else{try{this.split=parse_option(c,whitespace,this.split)}catch(a){this.split=/\s+/}this.numeric=parse_option(this.numeric,!0)}if(this.prepare=parse_option(a.prepare,null,this.prepare),this.finalize=parse_option(a.finalize,null,this.finalize),!normalize&&(this.mapper=new Map(normalize_polyfill)),this.rtl=a.rtl||!1,this.dedupe=parse_option(a.dedupe,!0,this.dedupe),this.filter=parse_option((c=a.filter)&&new Set(c),null,this.filter),this.matcher=parse_option((c=a.matcher)&&new Map(c),null,this.matcher),this.mapper=parse_option((c=a.mapper)&&new Map(c),null,this.mapper),this.stemmer=parse_option((c=a.stemmer)&&new Map(c),null,this.stemmer),this.replacer=parse_option(a.replacer,null,this.replacer),this.minlength=parse_option(a.minlength,1,this.minlength),this.maxlength=parse_option(a.maxlength,0,this.maxlength),this.cache=c=parse_option(a.cache,!0,this.cache),c&&(this.timer=null,this.cache_size="number"==typeof c?c:2e5,this.cache_enc=new Map,this.cache_prt=new Map,this.cache_enc_length=128,this.cache_prt_length=128),this.matcher_str="",this.matcher_test=null,this.stemmer_str="",this.stemmer_test=null,this.matcher)for(const a of this.matcher.keys())this.matcher_str+=(this.matcher_str?"|":"")+a;if(this.stemmer)for(const a of this.stemmer.keys())this.stemmer_str+=(this.stemmer_str?"|":"")+a;return this},Encoder.prototype.addMatcher=function(a,b){return"object"==typeof a?this.addReplacer(a,b):2>a.length?this.addMapper(a,b):(this.matcher||(this.matcher=new Map),this.matcher.set(a,b),this.matcher_str+=(this.matcher_str?"|":"")+a,this.matcher_test=null,this.cache&&this.invalidate(),this)},Encoder.prototype.addStemmer=function(a,b){return this.stemmer||(this.stemmer=new Map),this.stemmer.set(a,b),this.stemmer_str+=(this.stemmer_str?"|":"")+a,this.stemmer_test=null,this.cache&&this.invalidate(),this},Encoder.prototype.addFilter=function(a){return this.filter||(this.filter=new Set),this.filter.add(a),this.cache&&this.invalidate(),this},Encoder.prototype.addMapper=function(a,b){return"object"==typeof a?this.addReplacer(a,b):1<a.length?this.addMatcher(a,b):(this.mapper||(this.mapper=new Map),this.mapper.set(a,b),this.cache&&this.invalidate(),this)},Encoder.prototype.addReplacer=function(a,b){return"string"==typeof a&&(a=new RegExp(a,"g")),this.replacer||(this.replacer=[]),this.replacer.push(a,b||""),this.cache&&this.invalidate(),this},Encoder.prototype.invalidate=function(){this.cache_enc.clear(),this.cache_prt.clear()},Encoder.prototype.encode=function(a){if(this.cache&&a.length<=this.cache_enc_length)if(!this.timer)this.timer=setTimeout(clear,0,this);else if(this.cache_enc.has(a))return this.cache_enc.get(a);this.normalize&&("function"==typeof this.normalize?a=this.normalize(a):normalize?a=a.normalize("NFKD").replace(normalize,"").toLowerCase():a=a.toLowerCase()),this.prepare&&(a=this.prepare(a)),this.numeric&&3<a.length&&(a=a.replace(numeric_split_prev_char,"$1 $2").replace(numeric_split_next_char,"$1 $2").replace(numeric_split_length,"$1 "));const b=!(this.dedupe||this.mapper||this.filter||this.matcher||this.stemmer||this.replacer);let c=[],d=this.split||""===this.split?a.split(this.split):a;for(let e,f,g=0;g<d.length;g++){if(!(e=f=d[g]))continue;if(e.length<this.minlength)continue;if(b){c.push(e);continue}if(this.filter&&this.filter.has(e))continue;if(this.cache&&e.length<=this.cache_prt_length)if(this.timer){const a=this.cache_prt.get(e);if(a||""===a){a&&c.push(a);continue}}else this.timer=setTimeout(clear,0,this);let a;if(this.stemmer&&2<e.length&&(this.stemmer_test||(this.stemmer_test=new RegExp("(?!^)("+this.stemmer_str+")$")),e=e.replace(this.stemmer_test,a=>this.stemmer.get(a)),a=1),this.matcher&&1<e.length&&(this.matcher_test||(this.matcher_test=new RegExp("("+this.matcher_str+")","g")),e=e.replace(this.matcher_test,a=>this.matcher.get(a)),a=1),e&&a&&(e.length<this.minlength||this.filter&&this.filter.has(e))&&(e=""),e&&(this.mapper||this.dedupe&&1<e.length)){let a="";for(let b,c,d=0,f="";d<e.length;d++)b=e.charAt(d),b===f&&this.dedupe||(c=this.mapper&&this.mapper.get(b),c||""===c?(c!==f||!this.dedupe)&&(f=c)&&(a+=c):a+=f=b);e=a}if(e&&this.replacer)for(let a=0;e&&a<this.replacer.length;a+=2)e=e.replace(this.replacer[a],this.replacer[a+1]);this.cache&&f.length<=this.cache_prt_length&&(this.cache_prt.set(f,e),this.cache_prt.size>this.cache_size&&(this.cache_prt.clear(),this.cache_prt_length=0|this.cache_prt_length/1.1)),e&&c.push(e)}return this.finalize&&(c=this.finalize(c)||c),this.cache&&a.length<=this.cache_enc_length&&(this.cache_enc.set(a,c),this.cache_enc.size>this.cache_size&&(this.cache_enc.clear(),this.cache_enc_length=0|this.cache_enc_length/1.1)),c};function clear(a){a.timer=null,a.cache_enc.clear(),a.cache_prt.clear()}