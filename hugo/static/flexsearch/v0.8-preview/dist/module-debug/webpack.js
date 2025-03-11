
import { SearchOptions, ContextOptions, DocumentDescriptor, DocumentSearchOptions, DocumentIndexOptions, IndexOptions, DocumentOptions, TagOptions, StoreOptions, EncoderOptions, EncoderSplitOptions, PersistentOptions } from "./type.js";
import StorageInterface from "./db/interface.js";
import Document from "./document.js";
import Index from "./index.js";
import WorkerIndex from "./worker.js";
import Resolver from "./resolver.js";
import Encoder from "./encoder.js";
import IdxDB from "./db/indexeddb/index.js";
import Charset from "./charset.js";

/** @export */Index.prototype.add;
/** @export */Index.prototype.append;
/** @export */Index.prototype.search;
/** @export */Index.prototype.update;
/** @export */Index.prototype.remove;
/** @export */Index.prototype.contain;
/** @export */Index.prototype.clear;
/** @export */Index.prototype.cleanup;
/** @export */Index.prototype.searchCache;
/** @export */Index.prototype.addAsync;
/** @export */Index.prototype.appendAsync;
/** @export */Index.prototype.searchAsync;
/** @export */Index.prototype.updateAsync;
/** @export */Index.prototype.removeAsync;
/** @export */Index.prototype.export;
/** @export */Index.prototype.import;
/** @export */Index.prototype.serialize;
/** @export */Index.prototype.mount;
/** @export */Index.prototype.commit;

/** @export */Index.prototype.reg;
/** @export */Index.prototype.map;
/** @export */Index.prototype.ctx;

/** @export */Index.prototype.db;
/** @export */Index.prototype.tag;
/** @export */Index.prototype.store;
/** @export */Index.prototype.depth;
/** @export */Index.prototype.bidirectional;
/** @export */Index.prototype.commit_task;
/** @export */Index.prototype.commit_timer;
/** @export */Index.prototype.cache;
/** @export */Index.prototype.bypass;
/** @export */Index.prototype.document;

/** @export */Document.prototype.add;
/** @export */Document.prototype.append;
/** @export */Document.prototype.search;
/** @export */Document.prototype.update;
/** @export */Document.prototype.remove;
/** @export */Document.prototype.contain;
/** @export */Document.prototype.clear;
/** @export */Document.prototype.cleanup;
/** @export */Document.prototype.addAsync;
/** @export */Document.prototype.appendAsync;
/** @export */Document.prototype.searchAsync;
/** @export */Document.prototype.updateAsync;
/** @export */Document.prototype.removeAsync;
/** @export */Document.prototype.mount;
/** @export */Document.prototype.commit;
/** @export */Document.prototype.export;
/** @export */Document.prototype.import;
/** @export */Document.prototype.searchCache;
/** @export */Document.prototype.get;
/** @export */Document.prototype.set;

/** @export */Resolver.prototype.limit;
/** @export */Resolver.prototype.offset;
/** @export */Resolver.prototype.boost;
/** @export */Resolver.prototype.resolve;
/** @export */Resolver.prototype.or;
/** @export */Resolver.prototype.and;
/** @export */Resolver.prototype.xor;
/** @export */Resolver.prototype.not;

/** @export */StorageInterface.db;
/** @export */StorageInterface.id;
/** @export */StorageInterface.support_tag_search;
/** @export */StorageInterface.prototype.mount;
/** @export */StorageInterface.prototype.open;
/** @export */StorageInterface.prototype.close;
/** @export */StorageInterface.prototype.destroy;
/** @export */StorageInterface.prototype.clear;
/** @export */StorageInterface.prototype.get;
/** @export */StorageInterface.prototype.tag;
/** @export */StorageInterface.prototype.enrich;
/** @export */StorageInterface.prototype.has;
/** @export */StorageInterface.prototype.search;
/** @export */StorageInterface.prototype.info;
/** @export */StorageInterface.prototype.commit;
/** @export */StorageInterface.prototype.remove;

/** @export */Charset.LatinExact;
/** @export */Charset.LatinDefault;
/** @export */Charset.LatinSimple;
/** @export */Charset.LatinBalance;
/** @export */Charset.LatinAdvanced;
/** @export */Charset.LatinExtra;
/** @export */Charset.LatinSoundex;
/** @export */Charset.ArabicDefault;
/** @export */Charset.CjkDefault;
/** @export */Charset.CyrillicDefault;

/** @export */IndexOptions.preset;
/** @export */IndexOptions.context;
/** @export */IndexOptions.encoder;
/** @export */IndexOptions.encode;
/** @export */IndexOptions.resolution;
/** @export */IndexOptions.tokenize;
/** @export */IndexOptions.fastupdate;
/** @export */IndexOptions.score;
/** @export */IndexOptions.keystore;
/** @export */IndexOptions.rtl;
/** @export */IndexOptions.cache;
/** @export */IndexOptions.resolve;
/** @export */IndexOptions.db;
/** @export */IndexOptions.config;

/** @export */DocumentIndexOptions.preset;
/** @export */DocumentIndexOptions.context;
/** @export */DocumentIndexOptions.encoder;
/** @export */DocumentIndexOptions.encode;
/** @export */DocumentIndexOptions.resolution;
/** @export */DocumentIndexOptions.tokenize;
/** @export */DocumentIndexOptions.fastupdate;
/** @export */DocumentIndexOptions.score;
/** @export */DocumentIndexOptions.keystore;
/** @export */DocumentIndexOptions.rtl;
/** @export */DocumentIndexOptions.cache;
/** @export */DocumentIndexOptions.db;
/** @export */DocumentIndexOptions.config;
// /** @export */ DocumentIndexOptions.resolve;
/** @export */DocumentIndexOptions.field;
/** @export */DocumentIndexOptions.filter;
/** @export */DocumentIndexOptions.custom;

/** @export */DocumentOptions.context;
/** @export */DocumentOptions.encoder;
/** @export */DocumentOptions.encode;
/** @export */DocumentOptions.resolution;
/** @export */DocumentOptions.tokenize;
/** @export */DocumentOptions.fastupdate;
/** @export */DocumentOptions.score;
/** @export */DocumentOptions.keystore;
/** @export */DocumentOptions.rtl;
/** @export */DocumentOptions.cache;
/** @export */DocumentOptions.db;
/** @export */DocumentOptions.doc;
/** @export */DocumentOptions.document;
/** @export */DocumentOptions.worker;

/** @export */ContextOptions.depth;
/** @export */ContextOptions.bidirectional;
/** @export */ContextOptions.resolution;

/** @export */DocumentDescriptor.field;
/** @export */DocumentDescriptor.index;
/** @export */DocumentDescriptor.tag;
/** @export */DocumentDescriptor.store;

/** @export */TagOptions.field;
/** @export */TagOptions.tag;
/** @export */TagOptions.filter;
/** @export */TagOptions.custom;
/** @export */TagOptions.keystore;
/** @export */TagOptions.db;
/** @export */TagOptions.config;

/** @export */StoreOptions.field;
/** @export */StoreOptions.filter;
/** @export */StoreOptions.custom;
/** @export */StoreOptions.config;

/** @export */SearchOptions.query;
/** @export */SearchOptions.limit;
/** @export */SearchOptions.offset;
/** @export */SearchOptions.context;
/** @export */SearchOptions.suggest;
/** @export */SearchOptions.resolve;
/** @export */SearchOptions.enrich;

/** @export */DocumentSearchOptions.query;
/** @export */DocumentSearchOptions.limit;
/** @export */DocumentSearchOptions.offset;
/** @export */DocumentSearchOptions.context;
/** @export */DocumentSearchOptions.suggest;
/** @export */DocumentSearchOptions.enrich;
/** @export */DocumentSearchOptions.tag;
/** @export */DocumentSearchOptions.field;
/** @export */DocumentSearchOptions.index;
/** @export */DocumentSearchOptions.pluck;
/** @export */DocumentSearchOptions.merge;

/** @export */EncoderOptions.rtl;
/** @export */EncoderOptions.dedupe;
/** @export */EncoderOptions.split;
/** @export */EncoderOptions.include;
/** @export */EncoderOptions.exclude;
/** @export */EncoderOptions.prepare;
/** @export */EncoderOptions.finalize;
/** @export */EncoderOptions.filter;
/** @export */EncoderOptions.matcher;
/** @export */EncoderOptions.mapper;
/** @export */EncoderOptions.stemmer;
/** @export */EncoderOptions.replacer;
/** @export */EncoderOptions.minlength;
/** @export */EncoderOptions.maxlength;
/** @export */EncoderOptions.cache;

/** @export */EncoderSplitOptions.letter;
/** @export */EncoderSplitOptions.number;
/** @export */EncoderSplitOptions.symbol;
/** @export */EncoderSplitOptions.punctuation;
/** @export */EncoderSplitOptions.control;
/** @export */EncoderSplitOptions.char;

/** @export */PersistentOptions.name;
/** @export */PersistentOptions.field;
/** @export */PersistentOptions.type;
/** @export */PersistentOptions.db;

const FlexSearch = {
    Index: Index,
    Charset: Charset,
    Encoder: Encoder,
    Document: Document,
    Worker: WorkerIndex,
    Resolver: Resolver,
    IndexedDB: IdxDB,
    Language: {}
};

// Export as library (Bundle)
// --------------------------------

{

    FlexSearch.Language = {};

    const root = self;
    let prop;

    // AMD (RequireJS)
    if ((prop = root.define) && prop.amd) {
        prop([], function () {
            return FlexSearch;
        });
    }
    // CommonJS
    else if ("object" == typeof root.exports) {
            root.exports = FlexSearch;
        }
        // Global (window)
        else {
                /** @export */
                root.FlexSearch = FlexSearch;
            }
}