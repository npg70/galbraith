import{EncoderOptions}from"../../type.js";const regex=/[\x00-\x7F]+/g,options={normalize:!1,dedupe:!0,split:"",prepare:function(a){return(""+a).replace(regex,"")}};export default options;