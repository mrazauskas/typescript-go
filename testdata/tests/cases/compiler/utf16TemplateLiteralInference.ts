// @noEmit: true

type Heads<S> = S extends `${infer C}${infer R}` ? [C, R] : never;
type A = Heads<"😀abc">;
declare let a: A;
const chk: ["😀", "abc"] = a;
