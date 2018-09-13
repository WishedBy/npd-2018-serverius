//Test cases for urlDecode
//
//Based on https://github.com/SpiderLabs/ModSecurity/blob/v2/master/tests/tfn/urlDecode.t.
package transforms_test

import (
	"fmt"
	"serverius/waf/transforms"
	"testing"
)

func equalHTMLEntitiesDecode(t *testing.T, underTest string, expected string) {
	got := transforms.HTMLEntitiesDecode(underTest)
	if got != expected {
		fmt.Printf("Input:         '% x'\n", underTest)
		fmt.Printf("Input str:     '%s'\n", underTest)
		fmt.Printf("Expected:      '% x'\n", expected)
		fmt.Printf("Expected str:  '%s'\n", expected)
		fmt.Printf("Got:           '% x'\n", got)
		fmt.Printf("Got str:       '%s'\n", got)
		t.Fail()
	}
}

func Test_HTMLEntitiesDecode_Empty(t *testing.T) {
	equalHTMLEntitiesDecode(
		t,
		"",
		"",
	)
}

func Test_HTMLEntitiesDecode_Nothing1(t *testing.T) {
	equalHTMLEntitiesDecode(
		t,
		"TestCase",
		"TestCase",
	)
}

func Test_HTMLEntitiesDecode_Nothing2(t *testing.T) {
	equalHTMLEntitiesDecode(
		t,
		"Test\x00Case",
		"Test\x00Case",
	)
}

func Test_HTMLEntitiesDecode_Valid1(t *testing.T) {
	equalHTMLEntitiesDecode(
		t,
		"&#x0;&#X0;&#x20;&#X20;&#0;&#32;\x00&#100;&quot;&amp;&lt;&gt;&nbsp;&apos",
		"\x00\x00\x20\x20\x00\x20\x00\x64\"&<>\xa0'",
	)
}

func Test_HTMLEntitiesDecode_Valid2(t *testing.T) {
	equalHTMLEntitiesDecode(
		t,
		"&#x0&#X0&#x20&#X20&#0&#32\x00&#100&quot&amp&lt&gt&nbsp",
		"\x00\x00\x20\x20\x00\x20\x00\x64\"&<>\xa0",
	)
}

func Test_HTMLEntitiesDecode_Invalid1(t *testing.T) {
	equalHTMLEntitiesDecode(
		t,
		"&#xg;&#Xg;&#xg0;&#X2g;&#a;\x00&#a2;&#3a&#a00;&#1a0;&#10a;&foo;",
		"&#xg;&#Xg;&#xg0;\x02g;&#a;\x00&#a2;\x03a&#a00;\x01a0;\x0aa;&foo;",
	)
}

func Test_HTMLEntitiesDecode_Invalid2(t *testing.T) {
	equalHTMLEntitiesDecode(
		t,
		"&#xg&#Xg&#xg0&#X2g&#a\x00&#a2&#3a&#a00&#1a0&#10a&foo",
		"&#xg&#Xg&#xg0\x02g&#a\x00&#a2\x03a&#a00\x01a0\x0aa&foo",
	)
}

func Test_HTMLEntitiesDecode_To100percent(t *testing.T) {
	equalHTMLEntitiesDecode(
		t,
		"&#x0&#X0&#x20&#XAf&#0&#32\x00&#100",
		"\x00\x00\x20¯\x00\x20\x00\x64",
	)
}
