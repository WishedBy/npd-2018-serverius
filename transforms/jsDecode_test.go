//Test cases for JSDecode
//
//Based on https://github.com/SpiderLabs/ModSecurity/blob/v2/master/tests/tfn/jsDecode.t.
package transforms_test

import (
	"fmt"
	"serverius/waf/transforms"
	"testing"
)

func equalJSDecode(t *testing.T, underTest string, expected string) {
	got := transforms.JSDecode(underTest)
	if got != expected {
		fmt.Printf("Input:    '%s'\n", underTest)
		fmt.Printf("Expected: '% x'\n", expected)
		fmt.Printf("Got:      '% x'\n", got)
		t.Fail()
	}
}

func Test_JSDecode_Empty(t *testing.T) {
	equalJSDecode(
		t,
		"",
		"",
	)
}

func Test_JSDecode_Nothing1(t *testing.T) {
	equalJSDecode(
		t,
		"Test\x00Case",
		"Test\x00Case",
	)
}

func Test_JSDecode_Nothing2(t *testing.T) {
	equalJSDecode(
		t,
		"TestCase",
		"TestCase",
	)
}

func Test_JSDecode_Valid1(t *testing.T) {
	equalJSDecode(
		t,
		"\\a\\b\\f\\n\\r\\t\\v\\?\\'\\\"\\0\\12\\123\\x00\\xff\\u0021\\uff01",
		"\a\b\f\x0a\x0d\t\x0b?'\"\x00\x0a\x53\x00\xff\x21\x21",
	)
}

func Test_JSDecode_Valid2(t *testing.T) {
	equalJSDecode(
		t,
		"\\a\\b\\f\\n\\r\\t\\v\x00\\?\\'\\\"\\0\\12\\123\\x00\\xff\\u0021\\uff01",
		"\a\b\f\x0a\x0d\t\x0b\x00?'\"\x00\x0a\x53\x00\xff\x21\x21",
	)
}

/* Invalid Sequences
\8 and \9 are not octal
\666 is \66 + '6' (JS does not allow the overflow as C does)
\u00ag, \u00ga, \u0zaa, \uz0aa are not hex
\xag and \xga are not hex,
\0123 is \012 + '3' */
func Test_JSDecode_Invalid1(t *testing.T) {
	equalJSDecode(
		t,
		"\\8\\9\\666\\u00ag\\u00ga\\u0zaa\\uz0aa\\xag\\xga\\0123\\u00a",
		"89\x366u00agu00gau0zaauz0aaxagxga\x0a3u00a",
	)
}

func Test_JSDecode_Invalid2(t *testing.T) {
	equalJSDecode(
		t,
		"\\x",
		"x",
	)
}

func Test_JSDecode_Invalid3(t *testing.T) {
	equalJSDecode(
		t,
		"\\x\\x0",
		"xx0",
	)
}

func Test_JSDecode_Invalid4(t *testing.T) {
	equalJSDecode(
		t,
		"\\x\\x0\x00",
		"xx0\x00",
	)
}

func Test_JSDecode_Invalid5(t *testing.T) {
	equalJSDecode(
		t,
		"\\u",
		"u",
	)
}

func Test_JSDecode_Invalid6(t *testing.T) {
	equalJSDecode(
		t,
		"\\u\\u0",
		"uu0",
	)
}

func Test_JSDecode_Invalid7(t *testing.T) {
	equalJSDecode(
		t,
		"\\u\\u0\\u01",
		"uu0u01",
	)
}

func Test_JSDecode_Invalid8(t *testing.T) {
	equalJSDecode(
		t,
		"\\u\\u0\\u01\\u012",
		"uu0u01u012",
	)
}

func Test_JSDecode_Invalid9(t *testing.T) {
	equalJSDecode(
		t,
		"\\u\\u0\\u01\\u012\x00",
		"uu0u01u012\x00",
	)
}

func Test_JSDecode_Invalid10(t *testing.T) {
	equalJSDecode(
		t,
		"\\",
		"\\",
	)
}

func Test_JSDecode_Invalid11(t *testing.T) {
	equalJSDecode(
		t,
		"\\\x00",
		"\x00",
	)
}
