/**!
 * FlexSearch.js v0.8.0 (Light/Debug)
 * Author and Copyright: Thomas Wilkerling
 * Licence: Apache-2.0
 * Hosted by Nextapps GmbH
 * https://github.com/nextapps-de/flexsearch
 */
(function(self){'use strict';
function t(a, c, b) {
  const d = typeof b, e = typeof a;
  if ("undefined" !== d) {
    if ("undefined" !== e) {
      if (b) {
        if ("function" === e && d === e) {
          return function(h) {
            return a(b(h));
          };
        }
        c = a.constructor;
        if (c === b.constructor) {
          if (c === Array) {
            return b.concat(a);
          }
          if (c === Map) {
            var g = new Map(b);
            for (var f of a) {
              g.set(f[0], f[1]);
            }
            return g;
          }
          if (c === Set) {
            f = new Set(b);
            for (g of a.values()) {
              f.add(g);
            }
            return f;
          }
        }
      }
      return a;
    }
    return b;
  }
  return "undefined" === e ? c : a;
}
function u() {
  return Object.create(null);
}
function v(a, c) {
  return c.length - a.length;
}
;const w = /[^\p{L}\p{N}]+/u, x = /(\d{3})/g, y = /(\D)(\d{3})/g, z = /(\d{3})(\D)/g, B = "".normalize && /[\u0300-\u036f]/g;
function C(a) {
  if (!this) {
    return new C(...arguments);
  }
  for (let c = 0; c < arguments.length; c++) {
    this.assign(arguments[c]);
  }
}
C.prototype.assign = function(a) {
  this.normalize = t(a.normalize, !0, this.normalize);
  let c = a.include, b = c || a.exclude || a.split;
  if ("object" === typeof b) {
    let d = !c, e = "";
    a.include || (e += "\\p{Z}");
    b.letter && (e += "\\p{L}");
    b.number && (e += "\\p{N}", d = !!c);
    b.symbol && (e += "\\p{S}");
    b.punctuation && (e += "\\p{P}");
    b.control && (e += "\\p{C}");
    if (b = b.char) {
      e += "object" === typeof b ? b.join("") : b;
    }
    try {
      this.split = new RegExp("[" + (c ? "^" : "") + e + "]+", "u");
    } catch (g) {
      this.split = /\s+/;
    }
    this.numeric = d;
  } else {
    try {
      this.split = t(b, w, this.split);
    } catch (d) {
      this.split = /\s+/;
    }
    this.numeric = t(this.numeric, !0);
  }
  this.prepare = t(a.prepare, null, this.prepare);
  this.finalize = t(a.finalize, null, this.finalize);
  this.rtl = a.rtl || !1;
  this.dedupe = t(a.dedupe, !0, this.dedupe);
  this.filter = t((b = a.filter) && new Set(b), null, this.filter);
  this.matcher = t((b = a.matcher) && new Map(b), null, this.matcher);
  this.mapper = t((b = a.mapper) && new Map(b), null, this.mapper);
  this.stemmer = t((b = a.stemmer) && new Map(b), null, this.stemmer);
  this.replacer = t(a.replacer, null, this.replacer);
  this.minlength = t(a.minlength, 1, this.minlength);
  this.maxlength = t(a.maxlength, 0, this.maxlength);
  if (this.cache = b = t(a.cache, !0, this.cache)) {
    this.j = null, this.v = "number" === typeof b ? b : 2e5, this.h = new Map(), this.i = new Map(), this.l = this.g = 128;
  }
  this.m = "";
  this.s = null;
  this.o = "";
  this.u = null;
  if (this.matcher) {
    for (const d of this.matcher.keys()) {
      this.m += (this.m ? "|" : "") + d;
    }
  }
  if (this.stemmer) {
    for (const d of this.stemmer.keys()) {
      this.o += (this.o ? "|" : "") + d;
    }
  }
  return this;
};
C.prototype.encode = function(a) {
  if (this.cache && a.length <= this.g) {
    if (this.j) {
      if (this.h.has(a)) {
        return this.h.get(a);
      }
    } else {
      this.j = setTimeout(D, 0, this);
    }
  }
  this.normalize && (a = "function" === typeof this.normalize ? this.normalize(a) : B ? a.normalize("NFKD").replace(B, "").toLowerCase() : a.toLowerCase());
  this.prepare && (a = this.prepare(a));
  this.numeric && 3 < a.length && (a = a.replace(y, "$1 $2").replace(z, "$1 $2").replace(x, "$1 "));
  const c = !(this.dedupe || this.mapper || this.filter || this.matcher || this.stemmer || this.replacer);
  let b = [], d = this.split || "" === this.split ? a.split(this.split) : a;
  for (let g = 0, f, h; g < d.length; g++) {
    if (!(f = h = d[g])) {
      continue;
    }
    if (f.length < this.minlength) {
      continue;
    }
    if (c) {
      b.push(f);
      continue;
    }
    if (this.filter && this.filter.has(f)) {
      continue;
    }
    if (this.cache && f.length <= this.l) {
      if (this.j) {
        var e = this.i.get(f);
        if (e || "" === e) {
          e && b.push(e);
          continue;
        }
      } else {
        this.j = setTimeout(D, 0, this);
      }
    }
    let k;
    this.stemmer && 2 < f.length && (this.u || (this.u = new RegExp("(?!^)(" + this.o + ")$")), f = f.replace(this.u, m => this.stemmer.get(m)), k = 1);
    this.matcher && 1 < f.length && (this.s || (this.s = new RegExp("(" + this.m + ")", "g")), f = f.replace(this.s, m => this.matcher.get(m)), k = 1);
    f && k && (f.length < this.minlength || this.filter && this.filter.has(f)) && (f = "");
    if (f && (this.mapper || this.dedupe && 1 < f.length)) {
      e = "";
      for (let m = 0, l = "", p, r; m < f.length; m++) {
        p = f.charAt(m), p === l && this.dedupe || ((r = this.mapper && this.mapper.get(p)) || "" === r ? r === l && this.dedupe || !(l = r) || (e += r) : e += l = p);
      }
      f = e;
    }
    if (f && this.replacer) {
      for (e = 0; f && e < this.replacer.length; e += 2) {
        f = f.replace(this.replacer[e], this.replacer[e + 1]);
      }
    }
    this.cache && h.length <= this.l && (this.i.set(h, f), this.i.size > this.v && (this.i.clear(), this.l = this.l / 1.1 | 0));
    f && b.push(f);
  }
  this.finalize && (b = this.finalize(b) || b);
  this.cache && a.length <= this.g && (this.h.set(a, b), this.h.size > this.v && (this.h.clear(), this.g = this.g / 1.1 | 0));
  return b;
};
function D(a) {
  a.j = null;
  a.h.clear();
  a.i.clear();
}
;function E(a) {
  this.limit = a && !0 !== a ? a : 1000;
  this.cache = new Map();
  this.g = "";
}
E.prototype.set = function(a, c) {
  this.cache.set(this.g = a, c);
  this.cache.size > this.limit && this.cache.delete(this.cache.keys().next().value);
};
E.prototype.get = function(a) {
  const c = this.cache.get(a);
  c && this.g !== a && (this.cache.delete(a), this.cache.set(this.g = a, c));
  return c;
};
E.prototype.remove = function(a) {
  for (const c of this.cache) {
    const b = c[0];
    c[1].includes(a) && this.cache.delete(b);
  }
};
E.prototype.clear = function() {
  this.cache.clear();
  this.g = "";
};
const F = {normalize:function(a) {
  return a.toLowerCase();
}, dedupe:!1};
const G = {memory:{resolution:1}, performance:{resolution:6, fastupdate:!0, context:{depth:1, resolution:3}}, match:{tokenize:"forward"}, score:{resolution:9, context:{depth:2, resolution:9}}};
u();
H.prototype.add = function(a, c, b, d) {
  if (c && (a || 0 === a)) {
    if (!d && !b && this.reg.has(a)) {
      return this.update(a, c);
    }
    c = this.encoder.encode(c);
    if (d = c.length) {
      const m = u(), l = u(), p = this.depth, r = this.resolution;
      for (let n = 0; n < d; n++) {
        let q = c[this.rtl ? d - 1 - n : n];
        var e = q.length;
        if (e && (p || !l[q])) {
          var g = this.score ? this.score(c, q, n, null, 0) : I(r, d, n), f = "";
          switch(this.tokenize) {
            case "full":
              if (2 < e) {
                for (g = 0; g < e; g++) {
                  for (var h = e; h > g; h--) {
                    f = q.substring(g, h);
                    var k = this.score ? this.score(c, q, n, f, g) : I(r, d, n, e, g);
                    J(this, l, f, k, a, b);
                  }
                }
                break;
              }
            case "reverse":
              if (1 < e) {
                for (h = e - 1; 0 < h; h--) {
                  f = q[h] + f, k = this.score ? this.score(c, q, n, f, h) : I(r, d, n, e, h), J(this, l, f, k, a, b);
                }
                f = "";
              }
            case "forward":
              if (1 < e) {
                for (h = 0; h < e; h++) {
                  f += q[h], J(this, l, f, g, a, b);
                }
                break;
              }
            default:
              if (J(this, l, q, g, a, b), p && 1 < d && n < d - 1) {
                for (e = u(), f = this.A, g = q, h = Math.min(p + 1, d - n), e[g] = 1, k = 1; k < h; k++) {
                  if ((q = c[this.rtl ? d - 1 - n - k : n + k]) && !e[q]) {
                    e[q] = 1;
                    const A = this.score ? this.score(c, g, n, q, k) : I(f + (d / 2 > f ? 0 : 1), d, n, h - 1, k - 1), M = this.bidirectional && q > g;
                    J(this, m, M ? g : q, A, a, b, M ? q : g);
                  }
                }
              }
          }
        }
      }
      this.fastupdate || this.reg.add(a);
    }
  }
  return this;
};
function J(a, c, b, d, e, g, f) {
  let h = f ? a.ctx : a.map, k;
  c[b] && f && (k = c[b])[f] || (f ? (c = k || (c[b] = u()), c[f] = 1, (k = h.get(f)) ? h = k : h.set(f, h = new Map())) : c[b] = 1, (k = h.get(b)) ? h = k : h.set(b, h = []), h = h[d] || (h[d] = []), g && h.includes(e) || (h.push(e), a.fastupdate && ((c = a.reg.get(e)) ? c.push(h) : a.reg.set(e, [h]))));
}
function I(a, c, b, d, e) {
  return b && 1 < a ? c + (d || 0) <= a ? b + (e || 0) : (a - 1) / (c + (d || 0)) * (b + (e || 0)) + 1 | 0 : 0;
}
;function K(a, c, b) {
  if (1 === a.length) {
    return a = a[0], a = b || a.length > c ? c ? a.slice(b, b + c) : a.slice(b) : a;
  }
  let d = [];
  for (let e = 0, g, f; e < a.length; e++) {
    if ((g = a[e]) && (f = g.length)) {
      if (b) {
        if (b >= f) {
          b -= f;
          continue;
        }
        b < f && (g = c ? g.slice(b, b + c) : g.slice(b), f = g.length, b = 0);
      }
      if (d.length) {
        f > c && (g = g.slice(0, c), f = g.length), d.push(g);
      } else {
        if (f >= c) {
          return f > c && (g = g.slice(0, c)), g;
        }
        d = [g];
      }
      c -= f;
      if (!c) {
        break;
      }
    }
  }
  return d.length ? d = 1 < d.length ? [].concat.apply([], d) : d[0] : d;
}
;H.prototype.search = function(a, c, b) {
  b || (c || "object" !== typeof a ? "object" === typeof c && (b = c, c = 0) : (b = a, a = ""));
  var d = [], e = 0;
  if (b) {
    a = b.query || a;
    c = b.limit || c;
    e = b.offset || 0;
    var g = b.context;
    var f = b.suggest;
  }
  a = this.encoder.encode(a);
  b = a.length;
  c || (c = 100);
  if (1 === b) {
    return L.call(this, a[0], "", c, e);
  }
  g = this.depth && !1 !== g;
  if (2 === b && g && !f) {
    return L.call(this, a[0], a[1], c, e);
  }
  var h = 0, k = 0;
  if (1 < b) {
    var m = u();
    const p = [];
    for (let r = 0, n; r < b; r++) {
      if ((n = a[r]) && !m[n]) {
        if (f || N(this, n)) {
          p.push(n), m[n] = 1;
        } else {
          return d;
        }
        const q = n.length;
        h = Math.max(h, q);
        k = k ? Math.min(k, q) : q;
      }
    }
    a = p;
    b = a.length;
  }
  if (!b) {
    return d;
  }
  m = 0;
  if (1 === b) {
    return L.call(this, a[0], "", c, e);
  }
  if (2 === b && g && !f) {
    return L.call(this, a[0], a[1], c, e);
  }
  if (1 < b) {
    if (g) {
      var l = a[0];
      m = 1;
    } else {
      9 < h && 3 < h / k && a.sort(v);
    }
  }
  for (let p, r; m < b; m++) {
    r = a[m];
    l ? (p = N(this, r, l), p = O(p, d, f, this.A, c, e, 2 === b), f && !1 === p && d.length || (l = r)) : (p = N(this, r), p = O(p, d, f, this.resolution, c, e, 1 === b));
    if (p) {
      return p;
    }
    if (f && m === b - 1) {
      g = d.length;
      if (!g) {
        if (l) {
          l = "";
          m = -1;
          continue;
        }
        return d;
      }
      if (1 === g) {
        return K(d[0], c, e);
      }
    }
  }
  a: {
    a = d;
    d = this.resolution;
    l = a.length;
    b = [];
    g = u();
    for (let p = 0, r, n, q, A; p < d; p++) {
      for (k = 0; k < l; k++) {
        if (q = a[k], p < q.length && (r = q[p])) {
          for (m = 0; m < r.length; m++) {
            n = r[m], (h = g[n]) ? g[n]++ : (h = 0, g[n] = 1), A = b[h] || (b[h] = []), A.push(n);
          }
        }
      }
    }
    if (a = b.length) {
      if (f) {
        if (1 < b.length) {
          f = b;
          a = [];
          d = u();
          g = f.length;
          for (k = 0; k < g; k++) {
            for (l = f[k], h = l.length, m = 0; m < h; m++) {
              if (b = l[m], !d[b]) {
                if (d[b] = 1, e) {
                  e--;
                } else {
                  if (a.push(b), a.length === c) {
                    break;
                  }
                }
              }
            }
          }
          c = a;
        } else {
          c = b[0];
        }
        b = c;
      } else {
        if (a < l) {
          d = [];
          break a;
        }
        b = b[a - 1];
        if (b.length > c || e) {
          b = b.slice(e, c + e);
        }
      }
    }
    d = b;
  }
  return d;
};
function L(a, c, b, d) {
  return (a = N(this, a, c)) && a.length ? K(a, b, d) : [];
}
function O(a, c, b, d, e, g, f) {
  let h = [];
  if (a) {
    d = Math.min(a.length, d);
    for (let k = 0, m = 0, l; k < d; k++) {
      if (l = a[k]) {
        if (g && l && f && (l.length <= g ? (g -= l.length, l = null) : (l = l.slice(g), g = 0)), l && (h[k] = l, f && (m += l.length, m >= e))) {
          break;
        }
      }
    }
    if (h.length) {
      if (f) {
        return K(h, e, 0);
      }
      c.push(h);
      return;
    }
  }
  return !b && h;
}
function N(a, c, b) {
  let d;
  b && (d = a.bidirectional && c > b);
  a = b ? (a = a.ctx.get(d ? c : b)) && a.get(d ? b : c) : a.map.get(c);
  return a;
}
;H.prototype.remove = function(a, c) {
  const b = this.reg.size && (this.fastupdate ? this.reg.get(a) : this.reg.has(a));
  if (b) {
    if (this.fastupdate) {
      for (let d = 0, e; d < b.length; d++) {
        if (e = b[d]) {
          if (2 > e.length) {
            e.pop();
          } else {
            const g = e.indexOf(a);
            g === b.length - 1 ? e.pop() : e.splice(g, 1);
          }
        }
      }
    } else {
      P(this.map, a), this.depth && P(this.ctx, a);
    }
    c || this.reg.delete(a);
  }
  this.cache && this.cache.remove(a);
  return this;
};
function P(a, c) {
  let b = 0;
  if (a.constructor === Array) {
    for (let d = 0, e, g; d < a.length; d++) {
      if ((e = a[d]) && e.length) {
        if (g = e.indexOf(c), 0 <= g) {
          1 < e.length ? (e.splice(g, 1), b++) : delete a[d];
          break;
        } else {
          b++;
        }
      }
    }
  } else {
    for (let d of a) {
      const e = d[0], g = P(d[1], c);
      g ? b += g : a.delete(e);
    }
  }
  return b;
}
;function H(a, c) {
  if (!this) {
    return new H(a);
  }
  if (a) {
    var b = "string" === typeof a ? a : a.preset;
    b && (G[b] || console.warn("Preset not found: " + b), a = Object.assign({}, G[b], a));
  } else {
    a = {};
  }
  b = a.context || {};
  const d = a.encode || a.encoder || F;
  this.encoder = d.encode ? d : "object" === typeof d ? new C(d) : {encode:d};
  let e;
  this.resolution = a.resolution || 9;
  this.tokenize = e = a.tokenize || "strict";
  this.depth = "strict" === e && b.depth || 0;
  this.bidirectional = !1 !== b.bidirectional;
  this.fastupdate = !!a.fastupdate;
  this.score = a.score || null;
  e = !1;
  this.map = new Map();
  this.ctx = new Map();
  this.reg = c || (this.fastupdate ? new Map() : new Set());
  this.A = b.resolution || 1;
  this.rtl = d.rtl || a.rtl || !1;
  this.cache = (e = a.cache || null) && new E(e);
}
H.prototype.clear = function() {
  this.map.clear();
  this.ctx.clear();
  this.reg.clear();
  this.cache && this.cache.clear();
  return this;
};
H.prototype.append = function(a, c) {
  return this.add(a, c, !0);
};
H.prototype.contain = function(a) {
  return this.reg.has(a);
};
H.prototype.update = function(a, c) {
  if (this.async) {
    const b = this, d = this.remove(a);
    return d.then ? d.then(() => b.add(a, c)) : this.add(a, c);
  }
  return this.remove(a).add(a, c);
};
function Q(a) {
  let c = 0;
  if (a.constructor === Array) {
    for (let b = 0, d; b < a.length; b++) {
      (d = a[b]) && (c += d.length);
    }
  } else {
    for (const b of a) {
      const d = b[0], e = Q(b[1]);
      e ? c += e : a.delete(d);
    }
  }
  return c;
}
H.prototype.cleanup = function() {
  if (!this.fastupdate) {
    return console.info('Cleanup the index isn\'t required when not using "fastupdate".'), this;
  }
  Q(this.map);
  this.depth && Q(this.ctx);
  return this;
};
H.prototype.searchCache = function(a, c, b) {
  a = ("object" === typeof a ? "" + a.query : a).toLowerCase();
  let d = this.cache.get(a);
  if (!d) {
    d = this.search(a, c, b);
    if (d.then) {
      const e = this;
      d.then(function(g) {
        e.cache.set(a, g);
        return g;
      });
    }
    this.cache.set(a, d);
  }
  return d;
};
const R = {Index:H, Charset:null, Encoder:C, Document:null, Worker:null, Resolver:null, IndexedDB:null, Language:{}}, S = self;
let T;
(T = S.define) && T.amd ? T([], function() {
  return R;
}) : "object" === typeof S.exports ? S.exports = R : S.FlexSearch = R;
}(this));
