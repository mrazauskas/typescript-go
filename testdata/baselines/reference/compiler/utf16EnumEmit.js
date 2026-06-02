//// [tests/cases/compiler/utf16EnumEmit.ts] ////

//// [utf16EnumEmit.ts]
enum E {
    "\uD800lone" = 1,
}
const enum S {
    C = "\uD800",
}
const c = S.C;

enum EscapedName {
    \u0041 = 1,
}
const a = EscapedName.A;


//// [utf16EnumEmit.js]
"use strict";
var E;
(function (E) {
    E[E["\uD800lone"] = 1] = "\uD800lone";
})(E || (E = {}));
const c = "\uD800" /* S.C */;
var EscapedName;
(function (EscapedName) {
    EscapedName[EscapedName["A"] = 1] = "A";
})(EscapedName || (EscapedName = {}));
const a = EscapedName.A;
