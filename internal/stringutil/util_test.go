package stringutil

import (
	"slices"
	"testing"
)

func TestEncodeURI(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "encodes spaces as percent20",
			input:    "a b",
			expected: "a%20b",
		},
		{
			name:     "preserves reserved uri characters",
			input:    ";/?:@&=+$,#",
			expected: ";/?:@&=+$,#",
		},
		{
			name:     "encodes brackets and unicode using utf8 bytes",
			input:    "①Ⅻㄨㄩ U1[abc]",
			expected: "%E2%91%A0%E2%85%AB%E3%84%A8%E3%84%A9%20U1%5Babc%5D",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := EncodeURI(tt.input); got != tt.expected {
				t.Fatalf("EncodeURI(%q) = %q, expected %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestNormalizeJSString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "ascii unchanged",
			input:    "abc",
			expected: "abc",
		},
		{
			name:     "raw supplementary code point unchanged",
			input:    "😀abc",
			expected: "😀abc",
		},
		{
			name:     "lone high surrogate unchanged",
			input:    EncodeJSStringRune(0xD800) + "abc",
			expected: EncodeJSStringRune(0xD800) + "abc",
		},
		{
			name:     "lone low surrogate unchanged",
			input:    EncodeJSStringRune(0xDC00) + "abc",
			expected: EncodeJSStringRune(0xDC00) + "abc",
		},
		{
			name:     "non-surrogate before low surrogate unchanged",
			input:    EncodeJSStringRune(0xD7FF) + EncodeJSStringRune(0xDC00) + "abc",
			expected: EncodeJSStringRune(0xD7FF) + EncodeJSStringRune(0xDC00) + "abc",
		},
		{
			name:     "high surrogate before non-low surrogate unchanged",
			input:    EncodeJSStringRune(0xD800) + EncodeJSStringRune(0xE000) + "abc",
			expected: EncodeJSStringRune(0xD800) + EncodeJSStringRune(0xE000) + "abc",
		},
		{
			name:     "encoded surrogate pair canonicalized",
			input:    EncodeJSStringRune(0xD83D) + EncodeJSStringRune(0xDE00) + "abc",
			expected: "😀abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NormalizeJSString(tt.input); got != tt.expected {
				t.Fatalf("NormalizeJSString(%q) = %q, expected %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestJSStringRuneEncoding(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		runeValue    rune
		encodedBytes []byte
	}{
		{name: "bmp scalar", runeValue: 0xD7FF, encodedBytes: []byte("퟿")},
		{name: "high surrogate", runeValue: 0xD800, encodedBytes: []byte{0xED, 0xA0, 0x80}},
		{name: "low surrogate", runeValue: 0xDC00, encodedBytes: []byte{0xED, 0xB0, 0x80}},
		{name: "supplementary scalar", runeValue: '😀', encodedBytes: []byte("😀")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			encoded := EncodeJSStringRune(tt.runeValue)
			if got := []byte(encoded); string(got) != string(tt.encodedBytes) {
				t.Fatalf("EncodeJSStringRune(%#x) = % X, expected % X", tt.runeValue, got, tt.encodedBytes)
			}
			decoded, size := DecodeJSStringRune(encoded)
			if decoded != tt.runeValue || size != len(tt.encodedBytes) {
				t.Fatalf("DecodeJSStringRune(% X) = (%#x, %d), expected (%#x, %d)", tt.encodedBytes, decoded, size, tt.runeValue, len(tt.encodedBytes))
			}
		})
	}
}

func TestJSStringCodeUnitOperations(t *testing.T) {
	t.Parallel()
	text := "a" + EncodeJSStringRune(0xD800) + "😀" + "b"
	units := []uint16{'a', 0xD800, 0xD83D, 0xDE00, 'b'}
	if got := JSStringCodeUnits(text); !slices.Equal(got, units) {
		t.Fatalf("JSStringCodeUnits(%q) = %#v, expected %#v", text, got, units)
	}
	if got := JSStringCodeUnitLen(text); got != len(units) {
		t.Fatalf("JSStringCodeUnitLen(%q) = %d, expected %d", text, got, len(units))
	}
	if got := JSStringFromCodeUnits(units); got != text {
		t.Fatalf("JSStringFromCodeUnits(%#v) = %q, expected %q", units, got, text)
	}
	if !JSStringHasCodeUnitPrefix(text, "a"+EncodeJSStringRune(0xD800)) {
		t.Fatalf("JSStringHasCodeUnitPrefix failed for lone surrogate prefix")
	}
	if !JSStringHasCodeUnitSuffix(text, "😀b") {
		t.Fatalf("JSStringHasCodeUnitSuffix failed for supplementary suffix")
	}
	if got := JSStringCodeUnitIndex(text, "😀", 0); got != 2 {
		t.Fatalf("JSStringCodeUnitIndex(%q, %q, 0) = %d, expected 2", text, "😀", got)
	}
	if got := JSStringCodeUnitSlice(text, 1, 4); got != EncodeJSStringRune(0xD800)+"😀" {
		t.Fatalf("JSStringCodeUnitSlice(%q, 1, 4) = %q", text, got)
	}
}

func TestSurrogatePairConversion(t *testing.T) {
	t.Parallel()
	high, low := CodePointToSurrogatePair('😀')
	if high != 0xD83D || low != 0xDE00 {
		t.Fatalf("CodePointToSurrogatePair('😀') = (%#x, %#x), expected (0xD83D, 0xDE00)", high, low)
	}
	if got := SurrogatePairToCodePoint(high, low); got != '😀' {
		t.Fatalf("SurrogatePairToCodePoint(%#x, %#x) = %#x, expected %#x", high, low, got, '😀')
	}
}
