//Test cases for CSSDecode
//
//Based on https://github.com/SpiderLabs/ModSecurity/blob/v2/master/tests/tfn/CSSDecode.t.
package transforms_test

import (
	"fmt"
	"serverius/waf/transforms"
	"testing"
)

func equalCSSDecode(t *testing.T, underTest string, expected string) {
	got := transforms.CSSDecode(underTest)
	if got != expected {
		fmt.Printf("Input:    '%s'\n", underTest)
		fmt.Printf("Expected: '% x'\n", expected)
		fmt.Printf("Got:      '% x'\n", got)
		t.Fail()
	}
}

func Test_CSSDecode_Empty(t *testing.T) {
	equalCSSDecode(
		t,
		"",
		"",
	)
}

func Test_CSSDecode_Nothing1(t *testing.T) {
	equalCSSDecode(
		t,
		"Test\x00Case",
		"Test\x00Case",
	)
}

func Test_CSSDecode_Nothing2(t *testing.T) {
	equalCSSDecode(
		t,
		"TestCase",
		"TestCase",
	)
}

func Test_CSSDecode_Valid(t *testing.T) {
	equalCSSDecode(
		t,
		"test\\a\\b\\f\\n\\r\\t\\v\\?\\'\\\"\\0\\12\\123\\1234\\12345\\123456\\ff01\\ff5e\\\n\\0  string",
		"test\x0a\x0b\x0fnrtv?'\"\x00\x12\x23\x34\x45\x56\x21\x7e\x00 string",
	)
}

//Trailing escape == line continuation with no line following (ie nothing)
func Test_CSSDecode_Invalid1(t *testing.T) {
	equalCSSDecode(
		t,
		"test\\",
		"test",
	)
}

/*
 Edge cases
  "\1A" == "\x1A"
  "\1 A" == "\x01A"
  "\1234567" == "\x567"
  "\123456 7" == "\x567"
  "\1x" == "\x01x"
  "\1 x" == "\x01 x"
*/
func Test_CSSDecode_Invalid2(t *testing.T) {
	equalCSSDecode(
		t,
		"\\1A\\1 A\\1234567\\123456 7\\1x\\1 x",
		"\x1A\x01A\x567\x567\x01x\x01x",
	)
}

//Test the 5 and 6 length fullwidth checks
func Test_CSSDecode_to100percent1(t *testing.T) {
	equalCSSDecode(
		t,
		"\\0000a\\0000Af",
		"\x0a\xaf",
	)
}
