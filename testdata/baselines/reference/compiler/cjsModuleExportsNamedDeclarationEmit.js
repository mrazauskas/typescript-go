//// [tests/cases/compiler/cjsModuleExportsNamedDeclarationEmit.ts] ////

//// [index.js]
class Foo {
    /** @param {string} x */
    constructor(x) {
        this.x = x;
    }
}

exports.foo = new Foo("hello");
exports.bar = function(/** @type {Foo} */ f) { return f; };


//// [index.js]
"use strict";
class Foo {
    /** @param {string} x */
    constructor(x) {
        this.x = x;
    }
}
exports.foo = new Foo("hello");
exports.bar = function (/** @type {Foo} */ f) { return f; };


//// [index.d.ts]
declare class Foo {
    x: string;
    /** @param {string} x */
    constructor(x: string);
}
export declare var foo: Foo;
export declare var bar: (f: Foo) => Foo;
