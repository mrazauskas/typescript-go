// @declaration: true
// @allowJs: true
// @checkJs: true
// @module: commonjs
// @target: es6
// @outDir: ./out

// @filename: /index.js
class Foo {
    /** @param {string} x */
    constructor(x) {
        this.x = x;
    }
}

exports.foo = new Foo("hello");
exports.bar = function(/** @type {Foo} */ f) { return f; };
