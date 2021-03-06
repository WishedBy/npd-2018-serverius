//Test cases for urlDecode
//
//Based on https://github.com/SpiderLabs/ModSecurity/blob/v2/master/tests/tfn/urlDecode.t.
package transforms_test

import (
	"fmt"
	"serverius/waf/transforms"
	"testing"
)

func equalURLDecode(t *testing.T, underTest string, expected string) {
	got := transforms.URLDecode(underTest)
	if got != expected {
		fmt.Printf("Input:    '%s'\n", underTest)
		fmt.Printf("Expected: '% x'\n", expected)
		fmt.Printf("Got:      '% x'\n", got)
		t.Fail()
	}
}

//func Test_URLDecode_
func Test_URLDecode_Empty(t *testing.T) {
	equalURLDecode(
		t,
		"",
		"",
	)
}

func Test_URLDecode_Nothing1(t *testing.T) {
	equalURLDecode(
		t,
		"TestCase",
		"TestCase",
	)
}

func Test_URLDecode_Nothing2(t *testing.T) {
	equalURLDecode(
		t,
		"Test\x0Case",
		"Test\x0Case",
	)
}

func Test_URLDecode_Valid1(t *testing.T) {
	equalURLDecode(
		t,
		"+%00%01%02%03%04%05%06%07%08%09%0a%0b%0c%0d%0e%0f%10%11%12%13%14%15%16%17%18%19%1a%1b%1c%1d%1e%1f%20%21%22%23%24%25%26%27%28%29%2a%2b%2c%2d%2e%2f%30%31%32%33%34%35%36%37%38%39%3a%3b%3c%3d%3e%3f%40%41%42%43%44%45%46%47%48%49%4a%4b%4c%4d%4e%4f%50%51%52%53%54%55%56%57%58%59%5a%5b%5c%5d%5e%5f%60%61%62%63%64%65%66%67%68%69%6a%6b%6c%6d%6e%6f%70%71%72%73%74%75%76%77%78%79%7a%7b%7c%7d%7e%7f%80%81%82%83%84%85%86%87%88%89%8a%8b%8c%8d%8e%8f%90%91%92%93%94%95%96%97%98%99%9a%9b%9c%9d%9e%9f%a0%a1%a2%a3%a4%a5%a6%a7%a8%a9%aa%ab%ac%ad%ae%af%b0%b1%b2%b3%b4%b5%b6%b7%b8%b9%ba%bb%bc%bd%be%bf%c0%c1%c2%c3%c4%c5%c6%c7%c8%c9%ca%cb%cc%cd%ce%cf%d0%d1%d2%d3%d4%d5%d6%d7%d8%d9%da%db%dc%dd%de%df%e0%e1%e2%e3%e4%e5%e6%e7%e8%e9%ea%eb%ec%ed%ee%ef%f0%f1%f2%f3%f4%f5%f6%f7%f8%f9%fa%fb%fc%fd%fe%ff",
		" \x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f \x21\x22\x23\x24\x25\x26\x27\x28\x29\x2a\x2b\x2c\x2d\x2e\x2f\x30\x31\x32\x33\x34\x35\x36\x37\x38\x39\x3a\x3b\x3c\x3d\x3e\x3f\x40\x41\x42\x43\x44\x45\x46\x47\x48\x49\x4a\x4b\x4c\x4d\x4e\x4f\x50\x51\x52\x53\x54\x55\x56\x57\x58\x59\x5a\x5b\x5c\x5d\x5e\x5f\x60\x61\x62\x63\x64\x65\x66\x67\x68\x69\x6a\x6b\x6c\x6d\x6e\x6f\x70\x71\x72\x73\x74\x75\x76\x77\x78\x79\x7a\x7b\x7c\x7d\x7e\x7f\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff",
	)
}

func Test_URLDecode_Valid2(t *testing.T) {
	equalURLDecode(
		t,
		"Test+Case",
		"Test Case",
	)
}

func Test_URLDecode_PartialInvalid1(t *testing.T) {
	equalURLDecode(
		t,
		"%+",
		"% ",
	)
}

func Test_URLDecode_PartialInvalid2(t *testing.T) {
	equalURLDecode(
		t,
		"%%20",
		"% ",
	)
}

func Test_URLDecode_PartialInvalid3(t *testing.T) {
	equalURLDecode(
		t,
		"%0g%20",
		"%0g ",
	)
}

func Test_URLDecode_PartialInvalid4(t *testing.T) {
	equalURLDecode(
		t,
		"%0%20",
		"%0 ",
	)
}

func Test_URLDecode_PartialInvalid5(t *testing.T) {
	equalURLDecode(
		t,
		"%g0%20",
		"%g0 ",
	)
}

func Test_URLDecode_PartialInvalid6(t *testing.T) {
	equalURLDecode(
		t,
		"%g%20",
		"%g ",
	)
}

func Test_URLDecode_Invalid1(t *testing.T) {

	equalURLDecode(
		t,
		"%0%1%2%3%4%5%6%7%8%9%0%a%b%c%d%e%f",
		"%0%1%2%3%4%5%6%7%8%9%0%a%b%c%d%e%f",
	)
}
func Test_URLDecode_Invalid2(t *testing.T) {
	equalURLDecode(
		t,
		"%g0%g1%g2%g3%g4%g5%g6%g7%g8%g9%g0%ga%gb%gc%gd%ge%gf",
		"%g0%g1%g2%g3%g4%g5%g6%g7%g8%g9%g0%ga%gb%gc%gd%ge%gf",
	)
}
func Test_URLDecode_Invalid3(t *testing.T) {
	equalURLDecode(
		t,
		"%0g%1g%2g%3g%4g%5g%6g%7g%8g%9g%0g%ag%bg%cg%dg%eg%fg",
		"%0g%1g%2g%3g%4g%5g%6g%7g%8g%9g%0g%ag%bg%cg%dg%eg%fg",
	)
}
func Test_URLDecode_Invalid4(t *testing.T) {
	equalURLDecode(
		t,
		"%",
		"%",
	)
}
func Test_URLDecode_Invalid5(t *testing.T) {
	equalURLDecode(
		t,
		"%0",
		"%0",
	)
}
func Test_URLDecode_Invalid6(t *testing.T) {
	equalURLDecode(
		t,
		"%%",
		"%%",
	)
}
func Test_URLDecode_Invalid7(t *testing.T) {
	equalURLDecode(
		t,
		"%0g",
		"%0g",
	)
}
func Test_URLDecode_Invalid8(t *testing.T) {
	equalURLDecode(
		t,
		"%gg",
		"%gg",
	)
}
