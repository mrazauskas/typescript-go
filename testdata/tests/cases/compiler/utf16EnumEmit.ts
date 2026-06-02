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
