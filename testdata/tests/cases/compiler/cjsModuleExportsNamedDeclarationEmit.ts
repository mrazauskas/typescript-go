// @declaration: true
// @allowJs: true
// @checkJs: true
// @module: commonjs
// @target: es6
// @outDir: ./out

// @filename: /helper.js
class InternalClass {
    /** @param {string} x */
    constructor(x) {
        this.x = x;
    }
}
module.exports = { InternalClass };

// @filename: /index.js
/** @type {typeof import("./helper").InternalClass} */
const Cls = require("./helper").InternalClass;
exports.instance = new Cls("hello");

// @filename: /reexport.js
const { InternalClass } = require("./helper");

/**
 * @param {InternalClass} cls
 * @returns {InternalClass}
 */
function wrap(cls) {
    return cls;
}

exports.wrapped = wrap(new InternalClass("test"));
