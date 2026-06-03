// @noEmit: true

type Head<S> = S extends `${infer C}${string}` ? C : never;
type Rest<S> = S extends `${string}${infer R}` ? R : never;
type A = [Head<"😀abc">, Rest<"😀abc">];
declare let a: A;
const chk: ["😀", "abc"] = a;
