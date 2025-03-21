/**!
 * FlexSearch.js v0.8.0 (Bundle/Module/Debug)
 * Author and Copyright: Thomas Wilkerling
 * Licence: Apache-2.0
 * Hosted by Nextapps GmbH
 * https://github.com/nextapps-de/flexsearch
 */
var u;
function y(a, b, c) {
  const d = typeof c, e = typeof a;
  if ("undefined" !== d) {
    if ("undefined" !== e) {
      if (c) {
        if ("function" === e && d === e) {
          return function(h) {
            return a(c(h));
          };
        }
        b = a.constructor;
        if (b === c.constructor) {
          if (b === Array) {
            return c.concat(a);
          }
          if (b === Map) {
            var f = new Map(c);
            for (var g of a) {
              f.set(g[0], g[1]);
            }
            return f;
          }
          if (b === Set) {
            g = new Set(c);
            for (f of a.values()) {
              g.add(f);
            }
            return g;
          }
        }
      }
      return a;
    }
    return c;
  }
  return "undefined" === e ? b : a;
}
function z() {
  return Object.create(null);
}
function aa(a, b) {
  return b.length - a.length;
}
function B(a) {
  return "string" === typeof a;
}
function C(a) {
  return "object" === typeof a;
}
function ba(a) {
  const b = [];
  for (const c of a.keys()) {
    b.push(c);
  }
  return b;
}
function ca(a, b) {
  if (B(b)) {
    a = a[b];
  } else {
    for (let c = 0; a && c < b.length; c++) {
      a = a[b[c]];
    }
  }
  return a;
}
function da(a) {
  let b = 0;
  for (let c = 0, d; c < a.length; c++) {
    (d = a[c]) && b < d.length && (b = d.length);
  }
  return b;
}
;var ea = [["\u00aa", "a"], ["\u00b2", "2"], ["\u00b3", "3"], ["\u00b9", "1"], ["\u00ba", "o"], ["\u00bc", "1\u20444"], ["\u00bd", "1\u20442"], ["\u00be", "3\u20444"], ["\u00e0", "a"], ["\u00e1", "a"], ["\u00e2", "a"], ["\u00e3", "a"], ["\u00e4", "a"], ["\u00e5", "a"], ["\u00e7", "c"], ["\u00e8", "e"], ["\u00e9", "e"], ["\u00ea", "e"], ["\u00eb", "e"], ["\u00ec", "i"], ["\u00ed", "i"], ["\u00ee", "i"], ["\u00ef", "i"], ["\u00f1", "n"], ["\u00f2", "o"], ["\u00f3", "o"], ["\u00f4", "o"], ["\u00f5", 
"o"], ["\u00f6", "o"], ["\u00f9", "u"], ["\u00fa", "u"], ["\u00fb", "u"], ["\u00fc", "u"], ["\u00fd", "y"], ["\u00ff", "y"], ["\u0101", "a"], ["\u0103", "a"], ["\u0105", "a"], ["\u0107", "c"], ["\u0109", "c"], ["\u010b", "c"], ["\u010d", "c"], ["\u010f", "d"], ["\u0113", "e"], ["\u0115", "e"], ["\u0117", "e"], ["\u0119", "e"], ["\u011b", "e"], ["\u011d", "g"], ["\u011f", "g"], ["\u0121", "g"], ["\u0123", "g"], ["\u0125", "h"], ["\u0129", "i"], ["\u012b", "i"], ["\u012d", "i"], ["\u012f", "i"], ["\u0133", 
"ij"], ["\u0135", "j"], ["\u0137", "k"], ["\u013a", "l"], ["\u013c", "l"], ["\u013e", "l"], ["\u0140", "l"], ["\u0144", "n"], ["\u0146", "n"], ["\u0148", "n"], ["\u0149", "n"], ["\u014d", "o"], ["\u014f", "o"], ["\u0151", "o"], ["\u0155", "r"], ["\u0157", "r"], ["\u0159", "r"], ["\u015b", "s"], ["\u015d", "s"], ["\u015f", "s"], ["\u0161", "s"], ["\u0163", "t"], ["\u0165", "t"], ["\u0169", "u"], ["\u016b", "u"], ["\u016d", "u"], ["\u016f", "u"], ["\u0171", "u"], ["\u0173", "u"], ["\u0175", "w"], ["\u0177", 
"y"], ["\u017a", "z"], ["\u017c", "z"], ["\u017e", "z"], ["\u017f", "s"], ["\u01a1", "o"], ["\u01b0", "u"], ["\u01c6", "dz"], ["\u01c9", "lj"], ["\u01cc", "nj"], ["\u01ce", "a"], ["\u01d0", "i"], ["\u01d2", "o"], ["\u01d4", "u"], ["\u01d6", "u"], ["\u01d8", "u"], ["\u01da", "u"], ["\u01dc", "u"], ["\u01df", "a"], ["\u01e1", "a"], ["\u01e3", "ae"], ["\u00e6", "ae"], ["\u01fd", "ae"], ["\u01e7", "g"], ["\u01e9", "k"], ["\u01eb", "o"], ["\u01ed", "o"], ["\u01ef", "\u0292"], ["\u01f0", "j"], ["\u01f3", 
"dz"], ["\u01f5", "g"], ["\u01f9", "n"], ["\u01fb", "a"], ["\u01ff", "\u00f8"], ["\u0201", "a"], ["\u0203", "a"], ["\u0205", "e"], ["\u0207", "e"], ["\u0209", "i"], ["\u020b", "i"], ["\u020d", "o"], ["\u020f", "o"], ["\u0211", "r"], ["\u0213", "r"], ["\u0215", "u"], ["\u0217", "u"], ["\u0219", "s"], ["\u021b", "t"], ["\u021f", "h"], ["\u0227", "a"], ["\u0229", "e"], ["\u022b", "o"], ["\u022d", "o"], ["\u022f", "o"], ["\u0231", "o"], ["\u0233", "y"], ["\u02b0", "h"], ["\u02b1", "h"], ["\u0266", "h"], 
["\u02b2", "j"], ["\u02b3", "r"], ["\u02b4", "\u0279"], ["\u02b5", "\u027b"], ["\u02b6", "\u0281"], ["\u02b7", "w"], ["\u02b8", "y"], ["\u02e0", "\u0263"], ["\u02e1", "l"], ["\u02e2", "s"], ["\u02e3", "x"], ["\u02e4", "\u0295"], ["\u0390", "\u03b9"], ["\u03ac", "\u03b1"], ["\u03ad", "\u03b5"], ["\u03ae", "\u03b7"], ["\u03af", "\u03b9"], ["\u03b0", "\u03c5"], ["\u03ca", "\u03b9"], ["\u03cb", "\u03c5"], ["\u03cc", "\u03bf"], ["\u03cd", "\u03c5"], ["\u03ce", "\u03c9"], ["\u03d0", "\u03b2"], ["\u03d1", 
"\u03b8"], ["\u03d2", "\u03a5"], ["\u03d3", "\u03a5"], ["\u03d4", "\u03a5"], ["\u03d5", "\u03c6"], ["\u03d6", "\u03c0"], ["\u03f0", "\u03ba"], ["\u03f1", "\u03c1"], ["\u03f2", "\u03c2"], ["\u03f5", "\u03b5"], ["\u0439", "\u0438"], ["\u0450", "\u0435"], ["\u0451", "\u0435"], ["\u0453", "\u0433"], ["\u0457", "\u0456"], ["\u045c", "\u043a"], ["\u045d", "\u0438"], ["\u045e", "\u0443"], ["\u0477", "\u0475"], ["\u04c2", "\u0436"], ["\u04d1", "\u0430"], ["\u04d3", "\u0430"], ["\u04d7", "\u0435"], ["\u04db", 
"\u04d9"], ["\u04dd", "\u0436"], ["\u04df", "\u0437"], ["\u04e3", "\u0438"], ["\u04e5", "\u0438"], ["\u04e7", "\u043e"], ["\u04eb", "\u04e9"], ["\u04ed", "\u044d"], ["\u04ef", "\u0443"], ["\u04f1", "\u0443"], ["\u04f3", "\u0443"], ["\u04f5", "\u0447"]];
const fa = /[^\p{L}\p{N}]+/u, ha = /(\d{3})/g, ia = /(\D)(\d{3})/g, ja = /(\d{3})(\D)/g, ka = "".normalize && /[\u0300-\u036f]/g;
function G(a) {
  if (!this) {
    return new G(...arguments);
  }
  for (let b = 0; b < arguments.length; b++) {
    this.assign(arguments[b]);
  }
}
G.prototype.assign = function(a) {
  this.normalize = y(a.normalize, !0, this.normalize);
  let b = a.include, c = b || a.exclude || a.split;
  if ("object" === typeof c) {
    let d = !b, e = "";
    a.include || (e += "\\p{Z}");
    c.letter && (e += "\\p{L}");
    c.number && (e += "\\p{N}", d = !!b);
    c.symbol && (e += "\\p{S}");
    c.punctuation && (e += "\\p{P}");
    c.control && (e += "\\p{C}");
    if (c = c.char) {
      e += "object" === typeof c ? c.join("") : c;
    }
    try {
      this.split = new RegExp("[" + (b ? "^" : "") + e + "]+", "u");
    } catch (f) {
      this.split = /\s+/;
    }
    this.numeric = d;
  } else {
    try {
      this.split = y(c, fa, this.split);
    } catch (d) {
      this.split = /\s+/;
    }
    this.numeric = y(this.numeric, !0);
  }
  this.prepare = y(a.prepare, null, this.prepare);
  this.finalize = y(a.finalize, null, this.finalize);
  ka || (this.mapper = new Map(ea));
  this.rtl = a.rtl || !1;
  this.dedupe = y(a.dedupe, !0, this.dedupe);
  this.filter = y((c = a.filter) && new Set(c), null, this.filter);
  this.matcher = y((c = a.matcher) && new Map(c), null, this.matcher);
  this.mapper = y((c = a.mapper) && new Map(c), null, this.mapper);
  this.stemmer = y((c = a.stemmer) && new Map(c), null, this.stemmer);
  this.replacer = y(a.replacer, null, this.replacer);
  this.minlength = y(a.minlength, 1, this.minlength);
  this.maxlength = y(a.maxlength, 0, this.maxlength);
  if (this.cache = c = y(a.cache, !0, this.cache)) {
    this.L = null, this.S = "number" === typeof c ? c : 2e5, this.H = new Map(), this.J = new Map(), this.A = this.h = 128;
  }
  this.B = "";
  this.O = null;
  this.M = "";
  this.P = null;
  if (this.matcher) {
    for (const d of this.matcher.keys()) {
      this.B += (this.B ? "|" : "") + d;
    }
  }
  if (this.stemmer) {
    for (const d of this.stemmer.keys()) {
      this.M += (this.M ? "|" : "") + d;
    }
  }
  return this;
};
G.prototype.encode = function(a) {
  if (this.cache && a.length <= this.h) {
    if (this.L) {
      if (this.H.has(a)) {
        return this.H.get(a);
      }
    } else {
      this.L = setTimeout(la, 0, this);
    }
  }
  this.normalize && (a = "function" === typeof this.normalize ? this.normalize(a) : ka ? a.normalize("NFKD").replace(ka, "").toLowerCase() : a.toLowerCase());
  this.prepare && (a = this.prepare(a));
  this.numeric && 3 < a.length && (a = a.replace(ia, "$1 $2").replace(ja, "$1 $2").replace(ha, "$1 "));
  const b = !(this.dedupe || this.mapper || this.filter || this.matcher || this.stemmer || this.replacer);
  let c = [], d = this.split || "" === this.split ? a.split(this.split) : a;
  for (let f = 0, g, h; f < d.length; f++) {
    if (!(g = h = d[f])) {
      continue;
    }
    if (g.length < this.minlength) {
      continue;
    }
    if (b) {
      c.push(g);
      continue;
    }
    if (this.filter && this.filter.has(g)) {
      continue;
    }
    if (this.cache && g.length <= this.A) {
      if (this.L) {
        var e = this.J.get(g);
        if (e || "" === e) {
          e && c.push(e);
          continue;
        }
      } else {
        this.L = setTimeout(la, 0, this);
      }
    }
    let k;
    this.stemmer && 2 < g.length && (this.P || (this.P = new RegExp("(?!^)(" + this.M + ")$")), g = g.replace(this.P, l => this.stemmer.get(l)), k = 1);
    this.matcher && 1 < g.length && (this.O || (this.O = new RegExp("(" + this.B + ")", "g")), g = g.replace(this.O, l => this.matcher.get(l)), k = 1);
    g && k && (g.length < this.minlength || this.filter && this.filter.has(g)) && (g = "");
    if (g && (this.mapper || this.dedupe && 1 < g.length)) {
      e = "";
      for (let l = 0, m = "", n, q; l < g.length; l++) {
        n = g.charAt(l), n === m && this.dedupe || ((q = this.mapper && this.mapper.get(n)) || "" === q ? q === m && this.dedupe || !(m = q) || (e += q) : e += m = n);
      }
      g = e;
    }
    if (g && this.replacer) {
      for (e = 0; g && e < this.replacer.length; e += 2) {
        g = g.replace(this.replacer[e], this.replacer[e + 1]);
      }
    }
    this.cache && h.length <= this.A && (this.J.set(h, g), this.J.size > this.S && (this.J.clear(), this.A = this.A / 1.1 | 0));
    g && c.push(g);
  }
  this.finalize && (c = this.finalize(c) || c);
  this.cache && a.length <= this.h && (this.H.set(a, c), this.H.size > this.S && (this.H.clear(), this.h = this.h / 1.1 | 0));
  return c;
};
function la(a) {
  a.L = null;
  a.H.clear();
  a.J.clear();
}
;function ma(a, b, c) {
  a = ("object" === typeof a ? "" + a.query : a).toLowerCase();
  let d = this.cache.get(a);
  if (!d) {
    d = this.search(a, b, c);
    if (d.then) {
      const e = this;
      d.then(function(f) {
        e.cache.set(a, f);
        return f;
      });
    }
    this.cache.set(a, d);
  }
  return d;
}
function H(a) {
  this.limit = a && !0 !== a ? a : 1000;
  this.cache = new Map();
  this.h = "";
}
H.prototype.set = function(a, b) {
  this.cache.set(this.h = a, b);
  this.cache.size > this.limit && this.cache.delete(this.cache.keys().next().value);
};
H.prototype.get = function(a) {
  const b = this.cache.get(a);
  b && this.h !== a && (this.cache.delete(a), this.cache.set(this.h = a, b));
  return b;
};
H.prototype.remove = function(a) {
  for (const b of this.cache) {
    const c = b[0];
    b[1].includes(a) && this.cache.delete(c);
  }
};
H.prototype.clear = function() {
  this.cache.clear();
  this.h = "";
};
const na = {normalize:function(a) {
  return a.toLowerCase();
}, dedupe:!1};
const oa = new Map([["b", "p"], ["v", "f"], ["w", "f"], ["z", "s"], ["x", "s"], ["d", "t"], ["n", "m"], ["c", "k"], ["g", "k"], ["j", "k"], ["q", "k"], ["i", "e"], ["y", "e"], ["u", "o"]]);
const pa = new Map([["ai", "ei"], ["ae", "a"], ["oe", "o"], ["ue", "u"], ["sh", "s"], ["ch", "c"], ["th", "t"], ["ph", "f"], ["pf", "f"]]), qa = [/([^aeo])h([aeo$])/g, "$1$2", /([aeo])h([^aeo]|$)/g, "$1$2"];
const ra = {a:"", e:"", i:"", o:"", u:"", y:"", b:1, f:1, p:1, v:1, c:2, g:2, j:2, k:2, q:2, s:2, x:2, z:2, "\u00df":2, d:3, t:3, l:4, m:5, n:5, r:6};
const sa = /[\x00-\x7F]+/g;
const ta = /[\x00-\x7F]+/g;
const ua = /[\x00-\x7F]+/g;
var va = {LatinExact:{normalize:!1, dedupe:!1}, LatinDefault:na, LatinSimple:{normalize:!0, dedupe:!0}, LatinBalance:{normalize:!0, dedupe:!0, mapper:oa}, LatinAdvanced:{normalize:!0, dedupe:!0, mapper:oa, replacer:qa, matcher:pa}, LatinExtra:{normalize:!0, dedupe:!0, mapper:oa, replacer:qa.concat([/(?!^)[aeoy]/g, ""]), matcher:pa}, LatinSoundex:{normalize:!0, dedupe:!1, include:{letter:!0}, finalize:function(a) {
  for (let c = 0; c < a.length; c++) {
    var b = a[c];
    let d = b.charAt(0), e = ra[d];
    for (let f = 1, g; f < b.length && (g = b.charAt(f), "h" === g || "w" === g || !(g = ra[g]) || g === e || (d += g, e = g, 4 !== d.length)); f++) {
    }
    a[c] = d;
  }
}}, ArabicDefault:{rtl:!0, normalize:!1, dedupe:!0, prepare:function(a) {
  return ("" + a).replace(sa, " ");
}}, CjkDefault:{normalize:!1, dedupe:!0, split:"", prepare:function(a) {
  return ("" + a).replace(ta, "");
}}, CyrillicDefault:{normalize:!1, dedupe:!0, prepare:function(a) {
  return ("" + a).replace(ua, " ");
}}};
function wa(a, b, c, d) {
  let e = [];
  for (let f = 0, g; f < a.index.length; f++) {
    if (g = a.index[f], b >= g.length) {
      b -= g.length;
    } else {
      b = g[d ? "splice" : "slice"](b, c);
      const h = b.length;
      if (h && (e = e.length ? e.concat(b) : b, c -= h, d && (a.length -= h), !c)) {
        break;
      }
      b = 0;
    }
  }
  return e;
}
function K(a) {
  if (!this) {
    return new K(a);
  }
  this.index = a ? [a] : [];
  this.length = a ? a.length : 0;
  const b = this;
  return new Proxy([], {get(c, d) {
    if ("length" === d) {
      return b.length;
    }
    if ("push" === d) {
      return function(e) {
        b.index[b.index.length - 1].push(e);
        b.length++;
      };
    }
    if ("pop" === d) {
      return function() {
        if (b.length) {
          return b.length--, b.index[b.index.length - 1].pop();
        }
      };
    }
    if ("indexOf" === d) {
      return function(e) {
        let f = 0;
        for (let g = 0, h, k; g < b.index.length; g++) {
          h = b.index[g];
          k = h.indexOf(e);
          if (0 <= k) {
            return f + k;
          }
          f += h.length;
        }
        return -1;
      };
    }
    if ("includes" === d) {
      return function(e) {
        for (let f = 0; f < b.index.length; f++) {
          if (b.index[f].includes(e)) {
            return !0;
          }
        }
        return !1;
      };
    }
    if ("slice" === d) {
      return function(e, f) {
        return wa(b, e || 0, f || b.length, !1);
      };
    }
    if ("splice" === d) {
      return function(e, f) {
        return wa(b, e || 0, f || b.length, !0);
      };
    }
    if ("constructor" === d) {
      return Array;
    }
    if ("symbol" !== typeof d) {
      return (c = b.index[d / 2 ** 31 | 0]) && c[d];
    }
  }, set(c, d, e) {
    c = d / 2 ** 31 | 0;
    (b.index[c] || (b.index[c] = []))[d] = e;
    b.length++;
    return !0;
  }});
}
K.prototype.clear = function() {
  this.index.length = 0;
};
K.prototype.destroy = function() {
  this.proxy = this.index = null;
};
K.prototype.push = function() {
};
function L(a = 8) {
  if (!this) {
    return new L(a);
  }
  this.index = z();
  this.B = [];
  this.size = 0;
  32 < a ? (this.h = xa, this.A = BigInt(a)) : (this.h = ya, this.A = a);
}
L.prototype.get = function(a) {
  const b = this.index[this.h(a)];
  return b && b.get(a);
};
L.prototype.set = function(a, b) {
  var c = this.h(a);
  let d = this.index[c];
  d ? (c = d.size, d.set(a, b), (c -= d.size) && this.size++) : (this.index[c] = d = new Map([[a, b]]), this.B.push(d));
};
function M(a = 8) {
  if (!this) {
    return new M(a);
  }
  this.index = z();
  this.h = [];
  32 < a ? (this.B = xa, this.A = BigInt(a)) : (this.B = ya, this.A = a);
}
M.prototype.add = function(a) {
  var b = this.B(a);
  let c = this.index[b];
  c ? (b = c.size, c.add(a), (b -= c.size) && this.size++) : (this.index[b] = c = new Set([a]), this.h.push(c));
};
u = L.prototype;
u.has = M.prototype.has = function(a) {
  const b = this.index[this.B(a)];
  return b && b.has(a);
};
u.delete = M.prototype.delete = function(a) {
  const b = this.index[this.B(a)];
  b && b.delete(a) && this.size--;
};
u.clear = M.prototype.clear = function() {
  this.index = z();
  this.h = [];
  this.size = 0;
};
u.values = M.prototype.values = function*() {
  for (let a = 0; a < this.h.length; a++) {
    for (let b of this.h[a].values()) {
      yield b;
    }
  }
};
u.keys = M.prototype.keys = function*() {
  for (let a = 0; a < this.h.length; a++) {
    for (let b of this.h[a].keys()) {
      yield b;
    }
  }
};
u.entries = M.prototype.entries = function*() {
  for (let a = 0; a < this.h.length; a++) {
    for (let b of this.h[a].entries()) {
      yield b;
    }
  }
};
function ya(a) {
  let b = 2 ** this.A - 1;
  if ("number" == typeof a) {
    return a & b;
  }
  let c = 0, d = this.A + 1;
  for (let e = 0; e < a.length; e++) {
    c = (c * d ^ a.charCodeAt(e)) & b;
  }
  return 32 === this.A ? c + 2 ** 31 : c;
}
function xa(a) {
  let b = BigInt(2) ** this.A - BigInt(1);
  var c = typeof a;
  if ("bigint" === c) {
    return a & b;
  }
  if ("number" === c) {
    return BigInt(a) & b;
  }
  c = BigInt(0);
  let d = this.A + BigInt(1);
  for (let e = 0; e < a.length; e++) {
    c = (c * d ^ BigInt(a.charCodeAt(e))) & b;
  }
  return c;
}
;function za(a, b, c, d, e, f, g, h) {
  (d = a(c ? c + "." + d : d, JSON.stringify(g))) && d.then ? d.then(function() {
    b.export(a, b, c, e, f + 1, h);
  }) : b.export(a, b, c, e, f + 1, h);
}
;const Aa = {memory:{resolution:1}, performance:{resolution:6, fastupdate:!0, context:{depth:1, resolution:3}}, match:{tokenize:"forward"}, score:{resolution:9, context:{depth:2, resolution:9}}};
function Ba(a) {
  N.call(a, "add");
  N.call(a, "append");
  N.call(a, "search");
  N.call(a, "update");
  N.call(a, "remove");
}
function N(a) {
  this[a + "Async"] = function() {
    var b = arguments;
    const c = b[b.length - 1];
    let d;
    "function" === typeof c && (d = c, delete b[b.length - 1]);
    this.async = !0;
    b = this[a].apply(this, b);
    this.async = !1;
    d && (b.then ? b.then(d) : d(b));
    return b;
  };
}
;z();
O.prototype.add = function(a, b, c, d) {
  if (b && (a || 0 === a)) {
    if (!d && !c && this.reg.has(a)) {
      return this.update(a, b);
    }
    b = this.encoder.encode(b);
    if (d = b.length) {
      const l = z(), m = z(), n = this.depth, q = this.resolution;
      for (let r = 0; r < d; r++) {
        let p = b[this.rtl ? d - 1 - r : r];
        var e = p.length;
        if (e && (n || !m[p])) {
          var f = this.score ? this.score(b, p, r, null, 0) : P(q, d, r), g = "";
          switch(this.tokenize) {
            case "full":
              if (2 < e) {
                for (f = 0; f < e; f++) {
                  for (var h = e; h > f; h--) {
                    g = p.substring(f, h);
                    var k = this.score ? this.score(b, p, r, g, f) : P(q, d, r, e, f);
                    Q(this, m, g, k, a, c);
                  }
                }
                break;
              }
            case "reverse":
              if (1 < e) {
                for (h = e - 1; 0 < h; h--) {
                  g = p[h] + g, k = this.score ? this.score(b, p, r, g, h) : P(q, d, r, e, h), Q(this, m, g, k, a, c);
                }
                g = "";
              }
            case "forward":
              if (1 < e) {
                for (h = 0; h < e; h++) {
                  g += p[h], Q(this, m, g, f, a, c);
                }
                break;
              }
            default:
              if (Q(this, m, p, f, a, c), n && 1 < d && r < d - 1) {
                for (e = z(), g = this.R, f = p, h = Math.min(n + 1, d - r), e[f] = 1, k = 1; k < h; k++) {
                  if ((p = b[this.rtl ? d - 1 - r - k : r + k]) && !e[p]) {
                    e[p] = 1;
                    const t = this.score ? this.score(b, f, r, p, k) : P(g + (d / 2 > g ? 0 : 1), d, r, h - 1, k - 1), x = this.bidirectional && p > f;
                    Q(this, l, x ? f : p, t, a, c, x ? p : f);
                  }
                }
              }
          }
        }
      }
      this.fastupdate || this.reg.add(a);
    } else {
      b = "";
    }
  }
  this.db && (b || this.commit_task.push({del:a}), this.T && Ca(this));
  return this;
};
function Q(a, b, c, d, e, f, g) {
  let h = g ? a.ctx : a.map, k;
  if (!b[c] || !g || !(k = b[c])[g]) {
    if (g ? (b = k || (b[c] = z()), b[g] = 1, (k = h.get(g)) ? h = k : h.set(g, h = new Map())) : b[c] = 1, (k = h.get(c)) ? h = k : h.set(c, h = k = []), h = h[d] || (h[d] = []), !f || !h.includes(e)) {
      if (h.length === 2 ** 31 - 1) {
        b = new K(h);
        if (a.fastupdate) {
          for (let l of a.reg.values()) {
            l.includes(h) && (l[l.indexOf(h)] = b);
          }
        }
        k[d] = h = b;
      }
      h.push(e);
      a.fastupdate && ((d = a.reg.get(e)) ? d.push(h) : a.reg.set(e, [h]));
    }
  }
}
function P(a, b, c, d, e) {
  return c && 1 < a ? b + (d || 0) <= a ? c + (e || 0) : (a - 1) / (b + (d || 0)) * (c + (e || 0)) + 1 | 0 : 0;
}
;function R(a, b, c, d) {
  if (1 === a.length) {
    return a = a[0], a = c || a.length > b ? b ? a.slice(c, c + b) : a.slice(c) : a, d ? Da(a) : a;
  }
  let e = [];
  for (let f = 0, g, h; f < a.length; f++) {
    if ((g = a[f]) && (h = g.length)) {
      if (c) {
        if (c >= h) {
          c -= h;
          continue;
        }
        c < h && (g = b ? g.slice(c, c + b) : g.slice(c), h = g.length, c = 0);
      }
      if (e.length) {
        h > b && (g = g.slice(0, b), h = g.length), e.push(g);
      } else {
        if (h >= b) {
          return h > b && (g = g.slice(0, b)), d ? Da(g) : g;
        }
        e = [g];
      }
      b -= h;
      if (!b) {
        break;
      }
    }
  }
  if (!e.length) {
    return e;
  }
  e = 1 < e.length ? [].concat.apply([], e) : e[0];
  return d ? Da(e) : e;
}
function Da(a) {
  for (let b = 0; b < a.length; b++) {
    a[b] = {score:b, id:a[b]};
  }
  return a;
}
;S.prototype.or = function() {
  const a = this;
  let b = arguments;
  var c = b[0];
  if (c.then) {
    return c.then(function() {
      return a.or.apply(a, b);
    });
  }
  if (c[0] && c[0].index) {
    return this.or.apply(this, c);
  }
  let d = [];
  c = [];
  let e = 0, f = 0, g, h;
  for (let k = 0, l; k < b.length; k++) {
    if (l = b[k]) {
      let m;
      if (l.constructor === S) {
        m = l.result;
      } else if (l.constructor === Array) {
        m = l;
      } else if (l.index) {
        l.resolve = !1, m = l.index.search(l).result;
      } else if (l.and) {
        m = this.and(l.and);
      } else if (l.xor) {
        m = this.xor(l.xor);
      } else if (l.not) {
        m = this.not(l.not);
      } else {
        e = l.limit || 0;
        f = l.offset || 0;
        g = l.enrich;
        h = l.resolve;
        continue;
      }
      d[k] = m;
      m.then && c.push(m);
    }
  }
  if (c.length) {
    return Promise.all(c).then(function() {
      a.result.length && (d = [a.result].concat(d));
      a.result = Ea(d, e, f, g, h, a.D);
      return h ? a.result : a;
    });
  }
  this.result.length && (d = [this.result].concat(d));
  this.result = Ea(d, e, f, g, h, a.D);
  return h ? this.result : this;
};
function Ea(a, b, c, d, e, f) {
  if (!a.length) {
    return a;
  }
  "object" === typeof b && (c = b.offset || 0, d = b.enrich || !1, b = b.limit || 0);
  if (2 > a.length) {
    return e ? R(a[0], b, c, d) : a[0];
  }
  d = [];
  let g = 0, h = z(), k = da(a);
  for (let l = 0, m; l < k; l++) {
    for (let n = 0, q; n < a.length; n++) {
      if (q = a[n]) {
        if (m = q[l]) {
          for (let r = 0, p; r < m.length; r++) {
            if (p = m[r], !h[p]) {
              if (h[p] = 1, c) {
                c--;
              } else {
                if (e) {
                  d.push(p);
                } else {
                  const t = l + (n ? f : 0);
                  d[t] || (d[t] = []);
                  d[t].push(p);
                }
                if (b && ++g === b) {
                  return d;
                }
              }
            }
          }
        }
      }
    }
  }
  return d;
}
;S.prototype.and = function() {
  if (this.result.length) {
    const b = this;
    let c = arguments;
    var a = c[0];
    if (a.then) {
      return a.then(function() {
        return b.and.apply(b, c);
      });
    }
    if (a[0] && a[0].index) {
      return this.and.apply(this, a);
    }
    let d = [];
    a = [];
    let e = 0, f = 0, g;
    for (let h = 0, k; h < c.length; h++) {
      if (k = c[h]) {
        let l;
        if (k.constructor === S) {
          l = k.result;
        } else if (k.constructor === Array) {
          l = k;
        } else if (k.index) {
          k.resolve = !1, l = k.index.search(k).result;
        } else if (k.or) {
          l = this.or(k.or);
        } else if (k.xor) {
          l = this.xor(k.xor);
        } else if (k.not) {
          l = this.not(k.not);
        } else {
          e = k.limit || 0;
          f = k.offset || 0;
          g = k.resolve;
          continue;
        }
        d[h] = l;
        l.then && a.push(l);
      }
    }
    if (a.length) {
      return Promise.all(a).then(function() {
        d = [b.result].concat(d);
        b.result = Fa(d, e, f, g, b.D);
        return g ? b.result : b;
      });
    }
    d = [this.result].concat(d);
    this.result = Fa(d, e, f, g, b.D);
    return g ? this.result : this;
  }
  return this;
};
function Fa(a, b, c, d, e) {
  if (2 > a.length) {
    return [];
  }
  let f = [], g = 0, h = z(), k = da(a);
  if (!k) {
    return f;
  }
  for (let l = 0, m; l < a.length; l++) {
    m = a[l];
    if (!m || !m.length) {
      return [];
    }
    let n = z(), q = 0, r = l === a.length - 1;
    for (let p = 0, t; p < k; p++) {
      if (t = m[p]) {
        for (let x = 0, A, w; x < t.length; x++) {
          if (A = t[x], !l) {
            n[A] = p + 1 + (l ? e : 0), q = 1;
          } else if (r) {
            if (w = h[A]) {
              if (q = 1, c) {
                c--;
              } else {
                if (d ? f.push(A) : (w--, p < w && (w = p), f[w] || (f[w] = []), f[w].push(A)), b && ++g === b) {
                  return f;
                }
              }
            }
          } else if (w = h[A]) {
            p + 1 < w && (w = p + 1), n[A] = w, q = 1;
          }
        }
      }
    }
    if (!q) {
      return [];
    }
    h = n;
  }
  return f;
}
;S.prototype.xor = function() {
  const a = this;
  let b = arguments;
  var c = b[0];
  if (c.then) {
    return c.then(function() {
      return a.xor.apply(a, b);
    });
  }
  if (c[0] && c[0].index) {
    return this.xor.apply(this, c);
  }
  let d = [];
  c = [];
  let e = 0, f = 0, g, h;
  for (let k = 0, l; k < b.length; k++) {
    if (l = b[k]) {
      let m;
      if (l.constructor === S) {
        m = l.result;
      } else if (l.constructor === Array) {
        m = l;
      } else if (l.index) {
        l.resolve = !1, m = l.index.search(l).result;
      } else if (l.or) {
        m = this.or(l.or);
      } else if (l.and) {
        m = this.and(l.and);
      } else if (l.not) {
        m = this.not(l.not);
      } else {
        e = l.limit || 0;
        f = l.offset || 0;
        g = l.enrich;
        h = l.resolve;
        continue;
      }
      d[k] = m;
      m.then && c.push(m);
    }
  }
  if (c.length) {
    return Promise.all(c).then(function() {
      a.result.length && (d = [a.result].concat(d));
      a.result = Ga(d, e, f, g, !h, a.D);
      return h ? a.result : a;
    });
  }
  this.result.length && (d = [this.result].concat(d));
  this.result = Ga(d, e, f, g, !h, a.D);
  return h ? this.result : this;
};
function Ga(a, b, c, d, e, f) {
  if (!a.length) {
    return a;
  }
  if (2 > a.length) {
    return e ? R(a[0], b, c, d) : a[0];
  }
  b = [];
  c = z();
  for (let g = 0, h; g < a.length; g++) {
    if (h = a[g]) {
      for (let k = 0, l; k < h.length; k++) {
        if (l = h[k]) {
          for (let m = 0, n; m < l.length; m++) {
            n = l[m], c[n] ? c[n]++ : c[n] = 1;
          }
        }
      }
    }
  }
  for (let g = 0, h; g < a.length; g++) {
    if (h = a[g]) {
      for (let k = 0, l; k < h.length; k++) {
        if (l = h[k]) {
          for (let m = 0, n; m < l.length; m++) {
            n = l[m], 1 === c[n] && (e ? b.push(n) : (d = k + (g ? f : 0), b[d] || (b[d] = []), b[d].push(n)));
          }
        }
      }
    }
  }
  return b;
}
;S.prototype.not = function() {
  const a = this;
  let b = arguments;
  var c = b[0];
  if (c.then) {
    return c.then(function() {
      return a.not.apply(a, b);
    });
  }
  if (c[0] && c[0].index) {
    return this.not.apply(this, c);
  }
  let d = [];
  c = [];
  let e;
  for (let f = 0, g; f < b.length; f++) {
    if (g = b[f]) {
      let h;
      if (g.constructor === S) {
        h = g.result;
      } else if (g.constructor === Array) {
        h = g;
      } else if (g.index) {
        g.resolve = !1, h = g.index.search(g).result;
      } else if (g.or) {
        h = this.or(g.or);
      } else if (g.and) {
        h = this.and(g.and);
      } else if (g.xor) {
        h = this.xor(g.xor);
      } else {
        e = g.resolve;
        continue;
      }
      d[f] = h;
      h.then && c.push(h);
    }
  }
  if (c.length) {
    return Promise.all(c).then(function() {
      a.result = Ha.call(a, d, e);
      return e ? a.result : a;
    });
  }
  this.result = Ha.call(this, d, e);
  return e ? this.result : this;
};
function Ha(a, b) {
  if (!a.length) {
    return this.result;
  }
  const c = [];
  a = new Set(a.flat().flat());
  for (let d = 0, e; d < this.result.length; d++) {
    if (e = this.result[d]) {
      for (let f = 0, g; f < e.length; f++) {
        g = e[f], a.has(g) || (b ? c.push(g) : (c[d] || (c[d] = []), c[d].push(g)));
      }
    }
  }
  return c;
}
;function S(a) {
  if (!this) {
    return new S(a);
  }
  if (a && a.index) {
    return a.resolve = !1, this.index = a.index, a.index.search(a);
  }
  if (a.constructor === S) {
    return a;
  }
  this.index = null;
  this.result = a || [];
  this.D = 0;
}
S.prototype.limit = function(a) {
  if (this.result.length) {
    const b = [];
    let c = 0;
    for (let d = 0, e; d < this.result.length; d++) {
      if (e = this.result[d], e.length + c < a) {
        b[d] = e, c += e.length;
      } else {
        b[d] = e.slice(0, a - c);
        this.result = b;
        break;
      }
    }
  }
  return this;
};
S.prototype.offset = function(a) {
  if (this.result.length) {
    const b = [];
    let c = 0;
    for (let d = 0, e; d < this.result.length; d++) {
      e = this.result[d], e.length + c < a ? c += e.length : (b[d] = e.slice(a - c), c = a);
    }
    this.result = b;
  }
  return this;
};
S.prototype.boost = function(a) {
  this.D += a;
  return this;
};
S.prototype.resolve = function(a, b, c) {
  T = 1;
  const d = this.result;
  this.result = this.index = null;
  return d.length ? ("object" === typeof a && (c = a.enrich, b = a.offset, a = a.limit), R(d, a || 100, b, c)) : d;
};
function Ia(a, b, c, d, e) {
  var f = a.length, g = [], h;
  var k = z();
  for (let n = 0, q, r, p, t; n < b; n++) {
    for (var l = 0; l < f; l++) {
      if (p = a[l], n < p.length && (q = p[n])) {
        for (var m = 0; m < q.length; m++) {
          r = q[m], (h = k[r]) ? k[r]++ : (h = 0, k[r] = 1), t = g[h] || (g[h] = []), t.push(r);
        }
      }
    }
  }
  if (a = g.length) {
    if (e) {
      if (1 < g.length) {
        e = g;
        f = [];
        g = z();
        k = e.length;
        for (l = 0; l < k; l++) {
          for (a = e[l], h = a.length, m = 0; m < h; m++) {
            if (b = a[m], !g[b]) {
              if (g[b] = 1, d) {
                d--;
              } else {
                if (f.push(b), f.length === c) {
                  break;
                }
              }
            }
          }
        }
        c = f;
      } else {
        c = g[0];
      }
      g = c;
    } else {
      if (a < f) {
        return [];
      }
      g = g[a - 1];
      if (g.length > c || d) {
        g = g.slice(d, c + d);
      }
    }
  }
  return g;
}
function Ja(a, b) {
  const c = z(), d = [];
  for (let e = 0, f; e < b.length; e++) {
    f = b[e];
    for (let g = 0; g < f.length; g++) {
      c[f[g]] = 1;
    }
  }
  for (let e = 0, f; e < a.length; e++) {
    f = a[e], 1 === c[f] && (d.push(f), c[f] = 2);
  }
  return d;
}
;let T = 1;
O.prototype.search = function(a, b, c) {
  c || (!b && C(a) ? (c = a, a = "") : C(b) && (c = b, b = 0));
  let d = [], e;
  let f, g = 0, h, k, l;
  if (c) {
    a = c.query || a;
    b = c.limit || b;
    g = c.offset || 0;
    var m = c.context;
    f = c.suggest;
    (h = T && !1 !== c.resolve) || (T = 0);
    k = h && c.enrich;
    l = this.db && c.tag;
  } else {
    h = this.resolve || T;
  }
  a = this.encoder.encode(a);
  e = a.length;
  b || !h || (b = 100);
  if (1 === e) {
    return U.call(this, a[0], "", b, g, h, k, l);
  }
  m = this.depth && !1 !== m;
  if (2 === e && m && !f) {
    return U.call(this, a[0], a[1], b, g, h, k, l);
  }
  let n = c = 0;
  if (1 < e) {
    const p = z(), t = [];
    for (let x = 0, A; x < e; x++) {
      if ((A = a[x]) && !p[A]) {
        if (f || this.db || V(this, A)) {
          t.push(A), p[A] = 1;
        } else {
          return h ? d : new S(d);
        }
        const w = A.length;
        c = Math.max(c, w);
        n = n ? Math.min(n, w) : w;
      }
    }
    a = t;
    e = a.length;
  }
  if (!e) {
    return h ? d : new S(d);
  }
  let q = 0, r;
  if (1 === e) {
    return U.call(this, a[0], "", b, g, h, k, l);
  }
  if (2 === e && m && !f) {
    return U.call(this, a[0], a[1], b, g, h, k, l);
  }
  1 < e && (m ? (r = a[0], q = 1) : 9 < c && 3 < c / n && a.sort(aa));
  if (this.db) {
    if (this.db.search && (m = this.db.search(this, a, b, g, f, h, k, l), !1 !== m)) {
      return m;
    }
    const p = this;
    return async function() {
      for (let t, x; q < e; q++) {
        x = a[q];
        r ? (t = await V(p, x, r), t = Ka(t, d, f, p.R, b, g, 2 === e), f && !1 === t && d.length || (r = x)) : (t = await V(p, x), t = Ka(t, d, f, p.resolution, b, g, 1 === e));
        if (t) {
          return t;
        }
        if (f && q === e - 1) {
          let A = d.length;
          if (!A) {
            if (r) {
              r = "";
              q = -1;
              continue;
            }
            return d;
          }
          if (1 === A) {
            return h ? R(d[0], b, g) : new S(d[0]);
          }
        }
      }
      return h ? Ia(d, p.resolution, b, g, f) : new S(d[0]);
    }();
  }
  for (let p, t; q < e; q++) {
    t = a[q];
    r ? (p = V(this, t, r), p = Ka(p, d, f, this.R, b, g, 2 === e), f && !1 === p && d.length || (r = t)) : (p = V(this, t), p = Ka(p, d, f, this.resolution, b, g, 1 === e));
    if (p) {
      return p;
    }
    if (f && q === e - 1) {
      m = d.length;
      if (!m) {
        if (r) {
          r = "";
          q = -1;
          continue;
        }
        return d;
      }
      if (1 === m) {
        return h ? R(d[0], b, g) : new S(d[0]);
      }
    }
  }
  d = Ia(d, this.resolution, b, g, f);
  return h ? d : new S(d);
};
function U(a, b, c, d, e, f, g) {
  a = V(this, a, b, c, d, e, f, g);
  return this.db ? a.then(function(h) {
    return e ? h : h && h.length ? e ? R(h, c, d) : new S(h) : e ? [] : new S([]);
  }) : a && a.length ? e ? R(a, c, d) : new S(a) : e ? [] : new S([]);
}
function Ka(a, b, c, d, e, f, g) {
  let h = [];
  if (a) {
    d = Math.min(a.length, d);
    for (let k = 0, l = 0, m; k < d; k++) {
      if (m = a[k]) {
        if (f && m && g && (m.length <= f ? (f -= m.length, m = null) : (m = m.slice(f), f = 0)), m && (h[k] = m, g && (l += m.length, l >= e))) {
          break;
        }
      }
    }
    if (h.length) {
      if (g) {
        return R(h, e, 0);
      }
      b.push(h);
      return;
    }
  }
  return !c && h;
}
function V(a, b, c, d, e, f, g, h) {
  let k;
  c && (k = a.bidirectional && b > c);
  if (a.db) {
    return c ? a.db.get(k ? c : b, k ? b : c, d, e, f, g, h) : a.db.get(b, "", d, e, f, g, h);
  }
  a = c ? (a = a.ctx.get(k ? b : c)) && a.get(k ? c : b) : a.map.get(b);
  return a;
}
;O.prototype.remove = function(a, b) {
  const c = this.reg.size && (this.fastupdate ? this.reg.get(a) : this.reg.has(a));
  if (c) {
    if (this.fastupdate) {
      for (let d = 0, e; d < c.length; d++) {
        if (e = c[d]) {
          if (2 > e.length) {
            e.pop();
          } else {
            const f = e.indexOf(a);
            f === c.length - 1 ? e.pop() : e.splice(f, 1);
          }
        }
      }
    } else {
      La(this.map, a), this.depth && La(this.ctx, a);
    }
    b || this.reg.delete(a);
  }
  this.db && (this.commit_task.push({del:a}), this.T && Ca(this));
  this.cache && this.cache.remove(a);
  return this;
};
function La(a, b) {
  let c = 0;
  if (a.constructor === Array) {
    for (let d = 0, e, f; d < a.length; d++) {
      if ((e = a[d]) && e.length) {
        if (f = e.indexOf(b), 0 <= f) {
          1 < e.length ? (e.splice(f, 1), c++) : delete a[d];
          break;
        } else {
          c++;
        }
      }
    }
  } else {
    for (let d of a) {
      const e = d[0], f = La(d[1], b);
      f ? c += f : a.delete(e);
    }
  }
  return c;
}
;function O(a, b) {
  if (!this) {
    return new O(a);
  }
  if (a) {
    var c = B(a) ? a : a.preset;
    c && (Aa[c] || console.warn("Preset not found: " + c), a = Object.assign({}, Aa[c], a));
  } else {
    a = {};
  }
  c = a.context || {};
  const d = a.encode || a.encoder || na;
  this.encoder = d.encode ? d : "object" === typeof d ? new G(d) : {encode:d};
  let e;
  this.resolution = a.resolution || 9;
  this.tokenize = e = a.tokenize || "strict";
  this.depth = "strict" === e && c.depth || 0;
  this.bidirectional = !1 !== c.bidirectional;
  this.fastupdate = !!a.fastupdate;
  this.score = a.score || null;
  (e = a.keystore || 0) && (this.keystore = e);
  this.map = e ? new L(e) : new Map();
  this.ctx = e ? new L(e) : new Map();
  this.reg = b || (this.fastupdate ? e ? new L(e) : new Map() : e ? new M(e) : new Set());
  this.R = c.resolution || 1;
  this.rtl = d.rtl || a.rtl || !1;
  this.cache = (e = a.cache || null) && new H(e);
  this.resolve = !1 !== a.resolve;
  if (e = a.db) {
    this.db = e.mount(this);
  }
  this.T = !1 !== a.commit;
  this.commit_task = [];
  this.commit_timer = null;
}
u = O.prototype;
u.mount = function(a) {
  this.commit_timer && (clearTimeout(this.commit_timer), this.commit_timer = null);
  return a.mount(this);
};
u.commit = function(a, b) {
  this.commit_timer && (clearTimeout(this.commit_timer), this.commit_timer = null);
  return this.db.commit(this, a, b);
};
function Ca(a) {
  a.commit_timer || (a.commit_timer = setTimeout(function() {
    a.commit_timer = null;
    a.db.commit(a, void 0, void 0);
  }, 0));
}
u.clear = function() {
  this.map.clear();
  this.ctx.clear();
  this.reg.clear();
  this.cache && this.cache.clear();
  this.db && (this.commit_timer && clearTimeout(this.commit_timer), this.commit_timer = null, this.commit_task = [{clear:!0}]);
  return this;
};
u.append = function(a, b) {
  return this.add(a, b, !0);
};
u.contain = function(a) {
  return this.db ? this.db.has(a) : this.reg.has(a);
};
u.update = function(a, b) {
  if (this.async) {
    const c = this, d = this.remove(a);
    return d.then ? d.then(() => c.add(a, b)) : this.add(a, b);
  }
  return this.remove(a).add(a, b);
};
function Ma(a) {
  let b = 0;
  if (a.constructor === Array) {
    for (let c = 0, d; c < a.length; c++) {
      (d = a[c]) && (b += d.length);
    }
  } else {
    for (const c of a) {
      const d = c[0], e = Ma(c[1]);
      e ? b += e : a.delete(d);
    }
  }
  return b;
}
u.cleanup = function() {
  if (!this.fastupdate) {
    return console.info('Cleanup the index isn\'t required when not using "fastupdate".'), this;
  }
  Ma(this.map);
  this.depth && Ma(this.ctx);
  return this;
};
u.searchCache = ma;
u.export = function(a, b, c, d, e, f) {
  let g = !0;
  "undefined" === typeof f && (g = new Promise(l => {
    f = l;
  }));
  let h, k;
  switch(e || (e = 0)) {
    case 0:
      h = "reg";
      if (this.fastupdate) {
        k = z();
        for (let l of this.reg.keys()) {
          k[l] = 1;
        }
      } else {
        k = this.reg;
      }
      break;
    case 1:
      h = "cfg";
      k = {doc:0, opt:this.h ? 1 : 0};
      break;
    case 2:
      h = "map";
      k = this.map;
      break;
    case 3:
      h = "ctx";
      k = this.ctx;
      break;
    default:
      "undefined" === typeof c && f && f();
      return;
  }
  za(a, b || this, c, h, d, e, k, f);
  return g;
};
u.import = function(a, b) {
  if (b) {
    switch(B(b) && (b = JSON.parse(b)), a) {
      case "cfg":
        this.h = !!b.opt;
        break;
      case "reg":
        this.fastupdate = !1;
        this.reg = b;
        break;
      case "map":
        this.map = b;
        break;
      case "ctx":
        this.ctx = b;
    }
  }
};
u.serialize = function(a = !0) {
  if (!this.reg.size) {
    return "";
  }
  let b = "", c = "";
  for (var d of this.reg.keys()) {
    c || (c = typeof d), b += (b ? "," : "") + ("string" === c ? '"' + d + '"' : d);
  }
  b = "index.reg=new Set([" + b + "]);";
  d = "";
  for (var e of this.map.entries()) {
    var f = e[0], g = e[1], h = "";
    for (let m = 0, n; m < g.length; m++) {
      n = g[m] || [""];
      var k = "";
      for (var l = 0; l < n.length; l++) {
        k += (k ? "," : "") + ("string" === c ? '"' + n[l] + '"' : n[l]);
      }
      k = "[" + k + "]";
      h += (h ? "," : "") + k;
    }
    h = '["' + f + '",[' + h + "]]";
    d += (d ? "," : "") + h;
  }
  d = "index.map=new Map([" + d + "]);";
  e = "";
  for (const m of this.ctx.entries()) {
    f = m[0];
    g = m[1];
    for (const n of g.entries()) {
      g = n[0];
      h = n[1];
      k = "";
      for (let q = 0, r; q < h.length; q++) {
        r = h[q] || [""];
        l = "";
        for (let p = 0; p < r.length; p++) {
          l += (l ? "," : "") + ("string" === c ? '"' + r[p] + '"' : r[p]);
        }
        l = "[" + l + "]";
        k += (k ? "," : "") + l;
      }
      k = 'new Map([["' + g + '",[' + k + "]]])";
      k = '["' + f + '",' + k + "]";
      e += (e ? "," : "") + k;
    }
  }
  e = "index.ctx=new Map([" + e + "]);";
  return a ? "function inject(index){" + b + d + e + "}" : b + d + e;
};
Ba(O.prototype);
async function Na(a) {
  a = a.data;
  var b = self._index;
  const c = a.args;
  var d = a.task;
  switch(d) {
    case "init":
      d = a.options || {};
      (b = d.config) && (d = b);
      (b = a.factory) ? (Function("return " + b)()(self), self._index = new self.FlexSearch.Index(d), delete self.FlexSearch) : self._index = new O(d);
      postMessage({id:a.id});
      break;
    default:
      a = a.id, b = b[d].apply(b, c), postMessage("search" === d ? {id:a, msg:b} : {id:a});
  }
}
;let Oa = 0;
function W(a) {
  function b(f) {
    f = f.data || f;
    const g = f.id, h = g && e.h[g];
    h && (h(f.msg), delete e.h[g]);
  }
  if (!this) {
    return new W(a);
  }
  a || (a = {});
  let c = (self || window)._factory;
  c && (c = c.toString());
  const d = "undefined" === typeof window && self.exports, e = this;
  this.worker = Pa(c, d, a.worker);
  this.h = z();
  if (this.worker) {
    d ? this.worker.on("message", b) : this.worker.onmessage = b;
    if (a.config) {
      return new Promise(function(f) {
        e.h[++Oa] = function() {
          f(e);
        };
        e.worker.postMessage({id:Oa, task:"init", factory:c, options:a});
      });
    }
    this.worker.postMessage({task:"init", factory:c, options:a});
  }
}
X("add");
X("append");
X("search");
X("update");
X("remove");
function X(a) {
  W.prototype[a] = W.prototype[a + "Async"] = function() {
    const b = this, c = [].slice.call(arguments);
    var d = c[c.length - 1];
    let e;
    "function" === typeof d && (e = d, c.splice(c.length - 1, 1));
    d = new Promise(function(f) {
      b.h[++Oa] = f;
      b.worker.postMessage({task:a, id:Oa, args:c});
    });
    return e ? (d.then(e), this) : d;
  };
}
function Pa(a, b, c) {
  return b ? new (require("worker_threads")["Worker"])(__dirname + "/node/node.js") : a ? new window.Worker(URL.createObjectURL(new Blob(["onmessage=" + Na.toString()], {type:"text/javascript"}))) : new window.Worker(B(c) ? c : "worker/worker.js", {type:"module"});
}
;Y.prototype.add = function(a, b, c) {
  C(a) && (b = a, a = ca(b, this.key));
  if (b && (a || 0 === a)) {
    if (!c && this.reg.has(a)) {
      return this.update(a, b);
    }
    for (let h = 0, k; h < this.field.length; h++) {
      k = this.G[h];
      var d = this.index.get(this.field[h]);
      if ("function" === typeof k) {
        var e = k(b);
        e && d.add(a, e, !1, !0);
      } else {
        if (e = k.I, !e || e(b)) {
          k.constructor === String ? k = ["" + k] : B(k) && (k = [k]), Qa(b, k, this.K, 0, d, a, k[0], c);
        }
      }
    }
    if (this.tag) {
      for (d = 0; d < this.F.length; d++) {
        var f = this.F[d], g = this.N[d];
        e = this.tag.get(g);
        let h = z();
        if ("function" === typeof f) {
          if (f = f(b), !f) {
            continue;
          }
        } else {
          const k = f.I;
          if (k && !k(b)) {
            continue;
          }
          f.constructor === String && (f = "" + f);
          f = ca(b, f);
        }
        if (e && f) {
          B(f) && (f = [f]);
          for (let k = 0, l, m; k < f.length; k++) {
            if (l = f[k], !h[l] && (h[l] = 1, (g = e.get(l)) ? m = g : e.set(l, m = []), !c || !m.includes(a))) {
              if (m.length === 2 ** 31 - 1) {
                g = new K(m);
                if (this.fastupdate) {
                  for (let n of this.reg.values()) {
                    n.includes(m) && (n[n.indexOf(m)] = g);
                  }
                }
                e.set(l, m = g);
              }
              m.push(a);
              this.fastupdate && ((g = this.reg.get(a)) ? g.push(m) : this.reg.set(a, [m]));
            }
          }
        } else {
          e || console.warn("Tag '" + g + "' was not found");
        }
      }
    }
    if (this.store && (!c || !this.store.has(a))) {
      let h;
      if (this.C) {
        h = z();
        for (let k = 0, l; k < this.C.length; k++) {
          l = this.C[k];
          if ((c = l.I) && !c(b)) {
            continue;
          }
          let m;
          if ("function" === typeof l) {
            m = l(b);
            if (!m) {
              continue;
            }
            l = [l.U];
          } else if (B(l) || l.constructor === String) {
            h[l] = b[l];
            continue;
          }
          Ra(b, h, l, 0, l[0], m);
        }
      }
      this.store.set(a, h || b);
    }
  }
  return this;
};
function Ra(a, b, c, d, e, f) {
  a = a[e];
  if (d === c.length - 1) {
    b[e] = f || a;
  } else if (a) {
    if (a.constructor === Array) {
      for (b = b[e] = Array(a.length), e = 0; e < a.length; e++) {
        Ra(a, b, c, d, e);
      }
    } else {
      b = b[e] || (b[e] = z()), e = c[++d], Ra(a, b, c, d, e);
    }
  }
}
function Qa(a, b, c, d, e, f, g, h) {
  if (a = a[g]) {
    if (d === b.length - 1) {
      if (a.constructor === Array) {
        if (c[d]) {
          for (b = 0; b < a.length; b++) {
            e.add(f, a[b], !0, !0);
          }
          return;
        }
        a = a.join(" ");
      }
      e.add(f, a, h, !0);
    } else {
      if (a.constructor === Array) {
        for (g = 0; g < a.length; g++) {
          Qa(a, b, c, d, e, f, g, h);
        }
      } else {
        g = b[++d], Qa(a, b, c, d, e, f, g, h);
      }
    }
  } else {
    e.db && e.remove(f);
  }
}
;Y.prototype.search = function(a, b, c, d) {
  c || (!b && C(a) ? (c = a, a = "") : C(b) && (c = b, b = 0));
  let e = [], f = [], g;
  let h;
  let k;
  let l, m = 0;
  if (c) {
    c.constructor === Array && (c = {index:c});
    a = c.query || a;
    g = c.pluck;
    h = c.merge;
    k = g || c.field || c.index;
    var n = this.tag && c.tag;
    var q = this.store && c.enrich;
    var r = c.suggest;
    b = c.limit || b;
    l = c.offset || 0;
    b || (b = 100);
    if (n && (!this.db || !d)) {
      n.constructor !== Array && (n = [n]);
      var p = [];
      for (let w = 0, v; w < n.length; w++) {
        v = n[w];
        if (B(v)) {
          throw Error("A tag option can't be a string, instead it needs a { field: tag } format.");
        }
        if (v.field && v.tag) {
          var t = v.tag;
          if (t.constructor === Array) {
            for (var x = 0; x < t.length; x++) {
              p.push(v.field, t[x]);
            }
          } else {
            p.push(v.field, t);
          }
        } else {
          t = Object.keys(v);
          for (let I = 0, J, D; I < t.length; I++) {
            if (J = t[I], D = v[J], D.constructor === Array) {
              for (x = 0; x < D.length; x++) {
                p.push(J, D[x]);
              }
            } else {
              p.push(J, D);
            }
          }
        }
      }
      if (!p.length) {
        throw Error("Your tag definition within the search options is probably wrong. No valid tags found.");
      }
      n = p;
      if (!a) {
        r = [];
        if (p.length) {
          for (n = 0; n < p.length; n += 2) {
            if (this.db) {
              d = this.index.get(p[n]);
              if (!d) {
                console.warn("Tag '" + p[n] + ":" + p[n + 1] + "' will be skipped because there is no field '" + p[n] + "'.");
                continue;
              }
              r.push(d = d.db.tag(p[n + 1], b, l, q));
            } else {
              d = Sa.call(this, p[n], p[n + 1], b, l, q);
            }
            e.push({field:p[n], tag:p[n + 1], result:d});
          }
        }
        return r.length ? Promise.all(r).then(function(w) {
          for (let v = 0; v < w.length; v++) {
            e[v].result = w[v];
          }
          return e;
        }) : e;
      }
    }
    B(k) && (k = [k]);
  }
  k || (k = this.field);
  p = !d && (this.worker || this.async) && [];
  let A;
  for (let w = 0, v, I, J; w < k.length; w++) {
    I = k[w];
    if (this.db && this.tag && !this.G[w]) {
      continue;
    }
    let D;
    B(I) || (D = I, I = D.field, a = D.query || a, b = D.limit || b, r = D.suggest || r);
    if (d) {
      v = d[w];
    } else {
      if (t = D || c, x = this.index.get(I), n && (this.db && (t.tag = n, A = x.db.support_tag_search, t.field = k), A || (t.enrich = !1)), p) {
        p[w] = x.searchAsync(a, b, t);
        t && q && (t.enrich = q);
        continue;
      } else {
        v = x.search(a, b, t), t && q && (t.enrich = q);
      }
    }
    J = v && v.length;
    if (n && J) {
      t = [];
      x = 0;
      if (this.db && d) {
        if (!A) {
          for (let E = k.length; E < d.length; E++) {
            let F = d[E];
            if (F && F.length) {
              x++, t.push(F);
            } else if (!r) {
              return e;
            }
          }
        }
      } else {
        for (let E = 0, F, $a; E < n.length; E += 2) {
          F = this.tag.get(n[E]);
          if (!F) {
            if (console.warn("Tag '" + n[E] + ":" + n[E + 1] + "' will be skipped because there is no field '" + n[E] + "'."), r) {
              continue;
            } else {
              return e;
            }
          }
          if ($a = (F = F && F.get(n[E + 1])) && F.length) {
            x++, t.push(F);
          } else if (!r) {
            return e;
          }
        }
      }
      if (x) {
        v = Ja(v, t);
        J = v.length;
        if (!J && !r) {
          return e;
        }
        x--;
      }
    }
    if (J) {
      f[m] = I, e.push(v), m++;
    } else if (1 === k.length) {
      return e;
    }
  }
  if (p) {
    if (this.db && n && n.length && !A) {
      for (q = 0; q < n.length; q += 2) {
        d = this.index.get(n[q]);
        if (!d) {
          if (console.warn("Tag '" + n[q] + ":" + n[q + 1] + "' was not found because there is no field '" + n[q] + "'."), r) {
            continue;
          } else {
            return e;
          }
        }
        p.push(d.db.tag(n[q + 1], b, l, !1));
      }
    }
    const w = this;
    return Promise.all(p).then(function(v) {
      return v.length ? w.search(a, b, c, v) : v;
    });
  }
  if (!m) {
    return e;
  }
  if (g && (!q || !this.store)) {
    return e[0];
  }
  p = [];
  for (let w = 0, v; w < f.length; w++) {
    v = e[w];
    q && v.length && !v[0].doc && (this.db ? p.push(v = this.index.get(this.field[0]).db.enrich(v)) : v.length && (v = Ta.call(this, v)));
    if (g) {
      return v;
    }
    e[w] = {field:f[w], result:v};
  }
  return q && this.db && p.length ? Promise.all(p).then(function(w) {
    for (let v = 0; v < w.length; v++) {
      e[v].result = w[v];
    }
    return h ? Ua(e, b) : e;
  }) : h ? Ua(e, b) : e;
};
function Ua(a, b) {
  const c = [], d = z();
  for (let e = 0, f, g; e < a.length; e++) {
    f = a[e];
    g = f.result;
    for (let h = 0, k, l, m; h < g.length; h++) {
      if (l = g[h], k = l.id, m = d[k]) {
        m.push(f.field);
      } else {
        if (c.length === b) {
          return c;
        }
        l.field = d[k] = [f.field];
        c.push(l);
      }
    }
  }
  return c;
}
function Sa(a, b, c, d, e) {
  let f = this.tag.get(a);
  if (!f) {
    return console.warn("Tag '" + a + "' was not found"), [];
  }
  if ((a = (f = f && f.get(b)) && f.length - d) && 0 < a) {
    if (a > c || d) {
      f = f.slice(d, d + c);
    }
    e && (f = Ta.call(this, f));
    return f;
  }
}
function Ta(a) {
  const b = Array(a.length);
  for (let c = 0, d; c < a.length; c++) {
    d = a[c], b[c] = {id:d, doc:this.store.get(d)};
  }
  return b;
}
;function Y(a) {
  if (!this) {
    return new Y(a);
  }
  const b = a.document || a.doc || a;
  var c, d;
  this.G = [];
  this.field = [];
  this.K = [];
  this.key = (c = b.key || b.id) && Va(c, this.K) || "id";
  (d = a.keystore || 0) && (this.keystore = d);
  this.reg = (this.fastupdate = !!a.fastupdate) ? d ? new L(d) : new Map() : d ? new M(d) : new Set();
  this.C = (c = b.store || null) && !0 !== c && [];
  this.store = c && (d ? new L(d) : new Map());
  this.cache = (c = a.cache || null) && new H(c);
  a.cache = !1;
  this.worker = a.worker;
  this.async = !1;
  c = new Map();
  d = b.index || b.field || b;
  B(d) && (d = [d]);
  for (let e = 0, f, g; e < d.length; e++) {
    f = d[e];
    B(f) || (g = f, f = f.field);
    g = C(g) ? Object.assign({}, a, g) : a;
    if (this.worker) {
      const h = new W(g);
      c.set(f, h);
      h.worker || (this.worker = !1);
    }
    this.worker || c.set(f, new O(g, this.reg));
    g.custom ? this.G[e] = g.custom : (this.G[e] = Va(f, this.K), g.filter && ("string" === typeof this.G[e] && (this.G[e] = new String(this.G[e])), this.G[e].I = g.filter));
    this.field[e] = f;
  }
  if (this.C) {
    d = b.store;
    B(d) && (d = [d]);
    for (let e = 0, f, g; e < d.length; e++) {
      f = d[e], g = f.field || f, f.custom ? (this.C[e] = f.custom, f.custom.U = g) : (this.C[e] = Va(g, this.K), f.filter && ("string" === typeof this.C[e] && (this.C[e] = new String(this.C[e])), this.C[e].I = f.filter));
    }
  }
  this.index = c;
  this.tag = null;
  if (c = b.tag) {
    if ("string" === typeof c && (c = [c]), c.length) {
      this.tag = new Map();
      this.F = [];
      this.N = [];
      for (let e = 0, f, g; e < c.length; e++) {
        f = c[e];
        g = f.field || f;
        if (!g) {
          throw Error("The tag field from the document descriptor is undefined.");
        }
        f.custom ? this.F[e] = f.custom : (this.F[e] = Va(g, this.K), f.filter && ("string" === typeof this.F[e] && (this.F[e] = new String(this.F[e])), this.F[e].I = f.filter));
        this.N[e] = g;
        this.tag.set(g, new Map());
      }
    }
  }
  a.db && this.mount(a.db);
}
u = Y.prototype;
u.mount = function(a) {
  let b = this.field;
  if (this.tag) {
    for (let e = 0, f; e < this.N.length; e++) {
      f = this.N[e];
      var c = this.index.get(f);
      c || (this.index.set(f, c = new O({}, this.reg)), b === this.field && (b = b.slice(0)), b.push(f));
      c.tag = this.tag.get(f);
    }
  }
  c = [];
  const d = {db:a.db, type:a.type, fastupdate:a.fastupdate};
  for (let e = 0, f, g; e < b.length; e++) {
    d.field = g = b[e];
    f = this.index.get(g);
    const h = new a.constructor(a.id, d);
    h.id = a.id;
    c[e] = h.mount(f);
    f.document = !0;
    e ? f.bypass = !0 : f.store = this.store;
  }
  this.db = this.async = !0;
  return Promise.all(c);
};
u.commit = async function(a, b) {
  const c = [];
  for (const d of this.index.values()) {
    c.push(d.db.commit(d, a, b));
  }
  await Promise.all(c);
  this.reg.clear();
};
function Va(a, b) {
  const c = a.split(":");
  let d = 0;
  for (let e = 0; e < c.length; e++) {
    a = c[e], "]" === a[a.length - 1] && (a = a.substring(0, a.length - 2)) && (b[d] = !0), a && (c[d++] = a);
  }
  d < c.length && (c.length = d);
  return 1 < d ? c : c[0];
}
u.append = function(a, b) {
  return this.add(a, b, !0);
};
u.update = function(a, b) {
  return this.remove(a).add(a, b);
};
u.remove = function(a) {
  C(a) && (a = ca(a, this.key));
  for (var b of this.index.values()) {
    b.remove(a, !0);
  }
  if (this.reg.has(a)) {
    if (this.tag && !this.fastupdate) {
      for (let c of this.tag.values()) {
        for (let d of c) {
          b = d[0];
          const e = d[1], f = e.indexOf(a);
          -1 < f && (1 < e.length ? e.splice(f, 1) : c.delete(b));
        }
      }
    }
    this.store && this.store.delete(a);
    this.reg.delete(a);
  }
  this.cache && this.cache.remove(a);
  return this;
};
u.clear = function() {
  for (const a of this.index.values()) {
    a.clear();
  }
  if (this.tag) {
    for (const a of this.tag.values()) {
      a.clear();
    }
  }
  this.store && this.store.clear();
  return this;
};
u.contain = function(a) {
  return this.db ? this.index.get(this.field[0]).db.has(a) : this.reg.has(a);
};
u.cleanup = function() {
  for (const a of this.index.values()) {
    a.cleanup();
  }
  return this;
};
u.get = function(a) {
  return this.db ? this.index.get(this.field[0]).db.enrich(a).then(function(b) {
    return b[0] && b[0].doc;
  }) : this.store.get(a);
};
u.set = function(a, b) {
  this.store.set(a, b);
  return this;
};
u.searchCache = ma;
u.export = function(a, b, c, d, e, f) {
  let g;
  "undefined" === typeof f && (g = new Promise(k => {
    f = k;
  }));
  e || (e = 0);
  d || (d = 0);
  if (d < this.field.length) {
    c = this.field[d];
    var h = this.index[c];
    b = this;
    h.export(a, b, e ? c : "", d, e++, f) || (d++, b.export(a, b, c, d, 1, f));
  } else {
    switch(e) {
      case 1:
        b = "tag";
        h = this.A;
        c = null;
        break;
      case 2:
        b = "store";
        h = this.store;
        c = null;
        break;
      default:
        f();
        return;
    }
    za(a, this, c, b, d, e, h, f);
  }
  return g;
};
u.import = function(a, b) {
  if (b) {
    switch(B(b) && (b = JSON.parse(b)), a) {
      case "tag":
        this.A = b;
        break;
      case "reg":
        this.fastupdate = !1;
        this.reg = b;
        for (let d = 0, e; d < this.field.length; d++) {
          e = this.index[this.field[d]], e.reg = b, e.fastupdate = !1;
        }
        break;
      case "store":
        this.store = b;
        break;
      default:
        a = a.split(".");
        const c = a[0];
        a = a[1];
        c && a && this.index[c].import(a, b);
    }
  }
};
Ba(Y.prototype);
const Wa = "undefined" !== typeof window && (window.indexedDB || window.mozIndexedDB || window.webkitIndexedDB || window.msIndexedDB), Xa = ["map", "ctx", "tag", "reg", "cfg"];
function Ya(a, b = {}) {
  if (!this) {
    return new Ya(a, b);
  }
  "object" === typeof a && (b = a = a.name);
  a || console.info("Default storage space was used, because a name was not passed.");
  this.id = "flexsearch" + (a ? ":" + a.toLowerCase().replace(/[^a-z0-9_\-]/g, "") : "");
  this.field = b.field ? b.field.toLowerCase().replace(/[^a-z0-9_\-]/g, "") : "";
  this.support_tag_search = !1;
  this.db = null;
  this.h = {};
}
u = Ya.prototype;
u.mount = function(a) {
  if (a.constructor === Y) {
    return a.mount(this);
  }
  a.db = this;
  return this.open();
};
u.open = function() {
  let a = this;
  navigator.storage && navigator.storage.persist();
  return this.db || new Promise(function(b, c) {
    const d = Wa.open(a.id + (a.field ? ":" + a.field : ""), 1);
    d.onupgradeneeded = function() {
      const e = a.db = this.result;
      Xa.forEach(f => {
        e.objectStoreNames.contains(f) || e.createObjectStore(f);
      });
    };
    d.onblocked = function(e) {
      console.error("blocked", e);
      c();
    };
    d.onerror = function(e) {
      console.error(this.error, e);
      c();
    };
    d.onsuccess = function() {
      a.db = this.result;
      a.db.onversionchange = function() {
        a.close();
      };
      b(a);
    };
  });
};
u.close = function() {
  this.db.close();
  this.db = null;
};
u.destroy = function() {
  this.db && this.close();
  return Wa.deleteDatabase(this.id + (this.field ? ":" + this.field : ""));
};
u.clear = function() {
  const a = this.db.transaction(Xa, "readwrite");
  for (let b = 0; b < Xa.length; b++) {
    a.objectStore(Xa[b]).clear();
  }
  return Z(a);
};
u.get = function(a, b, c = 0, d = 0, e = !0, f = !1) {
  a = this.db.transaction(b ? "ctx" : "map", "readonly").objectStore(b ? "ctx" : "map").get(b ? b + ":" + a : a);
  const g = this;
  return Z(a).then(function(h) {
    let k = [];
    if (!h || !h.length) {
      return k;
    }
    if (e) {
      if (!c && !d && 1 === h.length) {
        return h[0];
      }
      for (let l = 0, m; l < h.length; l++) {
        if ((m = h[l]) && m.length) {
          if (d >= m.length) {
            d -= m.length;
            continue;
          }
          const n = c ? d + Math.min(m.length - d, c) : m.length;
          for (let q = d; q < n; q++) {
            k.push(m[q]);
          }
          d = 0;
          if (k.length === c) {
            break;
          }
        }
      }
      return f ? g.enrich(k) : k;
    }
    return h;
  });
};
u.tag = function(a, b = 0, c = 0, d = !1) {
  a = this.db.transaction("tag", "readonly").objectStore("tag").get(a);
  const e = this;
  return Z(a).then(function(f) {
    if (!f || !f.length || c >= f.length) {
      return [];
    }
    if (!b && !c) {
      return f;
    }
    f = f.slice(c, c + b);
    return d ? e.enrich(f) : f;
  });
};
u.enrich = function(a) {
  "object" !== typeof a && (a = [a]);
  const b = this.db.transaction("reg", "readonly").objectStore("reg"), c = [];
  for (let d = 0; d < a.length; d++) {
    c[d] = Z(b.get(a[d]));
  }
  return Promise.all(c).then(function(d) {
    for (let e = 0; e < d.length; e++) {
      d[e] = {id:a[e], doc:d[e] ? JSON.parse(d[e]) : null};
    }
    return d;
  });
};
u.has = function(a) {
  a = this.db.transaction("reg", "readonly").objectStore("reg").getKey(a);
  return Z(a);
};
u.search = null;
u.info = function() {
};
u.transaction = function(a, b, c) {
  let d = this.h[a + ":" + b];
  if (d) {
    return c.call(this, d);
  }
  let e = this.db.transaction(a, b);
  this.h[a + ":" + b] = d = e.objectStore(a);
  return new Promise((f, g) => {
    e.onerror = h => {
      this.h[a + ":" + b] = null;
      e.abort();
      e = d = null;
      g(h);
    };
    e.oncomplete = h => {
      e = d = this.h[a + ":" + b] = null;
      f(h || !0);
    };
    return c.call(this, d);
  });
};
u.commit = async function(a, b, c) {
  if (b) {
    await this.clear(), a.commit_task = [];
  } else {
    let d = a.commit_task;
    a.commit_task = [];
    for (let e = 0, f; e < d.length; e++) {
      if (f = d[e], f.clear) {
        await this.clear();
        b = !0;
        break;
      } else {
        d[e] = f.V;
      }
    }
    b || (c || (d = d.concat(ba(a.reg))), d.length && await this.remove(d));
  }
  a.reg.size && (await this.transaction("map", "readwrite", function(d) {
    for (const e of a.map) {
      const f = e[0], g = e[1];
      g.length && (b ? d.put(g, f) : d.get(f).onsuccess = function() {
        let h = this.result;
        var k;
        if (h && h.length) {
          const l = Math.max(h.length, g.length);
          for (let m = 0, n, q; m < l; m++) {
            if ((q = g[m]) && q.length) {
              if ((n = h[m]) && n.length) {
                for (k = 0; k < q.length; k++) {
                  n.push(q[k]);
                }
              } else {
                h[m] = q;
              }
              k = 1;
            }
          }
        } else {
          h = g, k = 1;
        }
        k && d.put(h, f);
      });
    }
  }), await this.transaction("ctx", "readwrite", function(d) {
    for (const e of a.ctx) {
      const f = e[0], g = e[1];
      for (const h of g) {
        const k = h[0], l = h[1];
        l.length && (b ? d.put(l, f + ":" + k) : d.get(f + ":" + k).onsuccess = function() {
          let m = this.result;
          var n;
          if (m && m.length) {
            const q = Math.max(m.length, l.length);
            for (let r = 0, p, t; r < q; r++) {
              if ((t = l[r]) && t.length) {
                if ((p = m[r]) && p.length) {
                  for (n = 0; n < t.length; n++) {
                    p.push(t[n]);
                  }
                } else {
                  m[r] = t;
                }
                n = 1;
              }
            }
          } else {
            m = l, n = 1;
          }
          n && d.put(m, f + ":" + k);
        });
      }
    }
  }), a.store ? await this.transaction("reg", "readwrite", function(d) {
    for (const e of a.store) {
      const f = e[0], g = e[1];
      d.put("object" === typeof g ? JSON.stringify(g) : 1, f);
    }
  }) : a.bypass || await this.transaction("reg", "readwrite", function(d) {
    for (const e of a.reg.keys()) {
      d.put(1, e);
    }
  }), a.tag && await this.transaction("tag", "readwrite", function(d) {
    for (const e of a.tag) {
      const f = e[0], g = e[1];
      g.length && (d.get(f).onsuccess = function() {
        let h = this.result;
        h = h && h.length ? h.concat(g) : g;
        d.put(h, f);
      });
    }
  }), a.map.clear(), a.ctx.clear(), a.tag && a.tag.clear(), a.store && a.store.clear(), a.document || a.reg.clear());
};
function Za(a, b, c) {
  const d = a.value;
  let e, f, g = 0;
  for (let h = 0, k; h < d.length; h++) {
    if (k = c ? d : d[h]) {
      for (let l = 0, m, n; l < b.length; l++) {
        if (n = b[l], m = k.indexOf(f ? parseInt(n, 10) : n), 0 > m && !f && "string" === typeof n && !isNaN(n) && (m = k.indexOf(parseInt(n, 10))) && (f = 1), 0 <= m) {
          if (e = 1, 1 < k.length) {
            k.splice(m, 1);
          } else {
            d[h] = [];
            break;
          }
        }
      }
      g += k.length;
    }
    if (c) {
      break;
    }
  }
  g ? e && a.update(d) : a.delete();
  a.continue();
}
u.remove = function(a) {
  "object" !== typeof a && (a = [a]);
  return Promise.all([this.transaction("map", "readwrite", function(b) {
    b.openCursor().onsuccess = function() {
      const c = this.result;
      c && Za(c, a);
    };
  }), this.transaction("ctx", "readwrite", function(b) {
    b.openCursor().onsuccess = function() {
      const c = this.result;
      c && Za(c, a);
    };
  }), this.transaction("tag", "readwrite", function(b) {
    b.openCursor().onsuccess = function() {
      const c = this.result;
      c && Za(c, a, !0);
    };
  }), this.transaction("reg", "readwrite", function(b) {
    for (let c = 0; c < a.length; c++) {
      b.delete(a[c]);
    }
  })]);
};
function Z(a) {
  return new Promise((b, c) => {
    a.onsuccess = function() {
      b(this.result);
    };
    a.oncomplete = function() {
      b(this.result);
    };
    a.onerror = c;
    a = null;
  });
}
;export default {Index:O, Charset:va, Encoder:G, Document:Y, Worker:W, Resolver:S, IndexedDB:Ya, Language:{}};

export const Index=O;export const  Charset=va;export const  Encoder=G;export const  Document=Y;export const  Worker=W;export const  Resolver=S;export const  IndexedDB=Ya;export const  Language={};