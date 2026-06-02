//// [tests/cases/compiler/utf16DeclarationEmitEmoji.ts] ////

//// [utf16DeclarationEmitEmoji.ts]
export const u2019 = "I can’t do this 😂 " as const;
export const speaker = (msg: string) => [`🔈`, `🔈 ${msg}`] as const;




//// [utf16DeclarationEmitEmoji.d.ts]
export declare const u2019: "I can’t do this 😂 ";
export declare const speaker: (msg: string) => readonly [`🔈`, `\uD83D\uDD08 ${string}`];
