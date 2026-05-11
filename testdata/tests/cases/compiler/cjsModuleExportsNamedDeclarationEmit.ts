// @declaration: true
// @allowJs: true
// @checkJs: true
// @module: commonjs
// @target: es6
// @outDir: ./out

// @filename: /index.js
class InternalClass {
    /** @param {string} x */
    constructor(x) {
        this.x = x;
    }
    /**
     * @param {InternalClass} other
     * @returns {boolean}
     */
    equals(other) {
        return this.x === other.x;
    }
}
exports.createInstance = function(/** @type {string} */ name) { return new InternalClass(name); };
