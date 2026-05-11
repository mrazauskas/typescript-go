//// [tests/cases/compiler/cjsModuleExportsNamedDeclarationEmit.ts] ////

//// [helper.js]
class InternalClass {
    /** @param {string} x */
    constructor(x) {
        this.x = x;
    }
}
module.exports = { InternalClass };

//// [index.js]
/** @type {typeof import("./helper").InternalClass} */
const Cls = require("./helper").InternalClass;
exports.instance = new Cls("hello");

//// [reexport.js]
const { InternalClass } = require("./helper");

/**
 * @param {InternalClass} cls
 * @returns {InternalClass}
 */
function wrap(cls) {
    return cls;
}

exports.wrapped = wrap(new InternalClass("test"));


//// [helper.js]
"use strict";
class InternalClass {
    /** @param {string} x */
    constructor(x) {
        this.x = x;
    }
}
module.exports = { InternalClass };
//// [reexport.js]
"use strict";
const { InternalClass } = require("./helper");
/**
 * @param {InternalClass} cls
 * @returns {InternalClass}
 */
function wrap(cls) {
    return cls;
}
exports.wrapped = wrap(new InternalClass("test"));


//// [helper.d.ts]
declare class InternalClass {
    /** @param {string} x */
    constructor(x: string);
}
declare const _default: {
    InternalClass: typeof InternalClass;
};
export = _default;
//// [reexport.d.ts]
export declare var wrapped: InternalClass;
