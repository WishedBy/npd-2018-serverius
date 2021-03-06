//Test cases for urlDecodeUni
//
//based on https://github.com/SpiderLabs/ModSecurity/blob/v2/master/tests/tfn/urlDecodeUni.t.

package transforms_test

import (
	"fmt"
	"serverius/waf/transforms"
	"testing"
)

func equalURLDecodeUniWithCodePage(t *testing.T, underTest string, expected string, codePage int) {
	got := transforms.URLDecodeUni(underTest, codePage)
	if got != expected {
		fmt.Printf("Input:    '%s'\n", underTest)
		fmt.Printf("Expected: '% x'\n", expected)
		fmt.Printf("Got:      '% x'\n", got)
		t.Fail()
	}
}

func equalURLDecodeUni(t *testing.T, underTest string, expected string) {
	got := transforms.URLDecodeUni(underTest, 20127)
	if got != expected {
		fmt.Printf("Input:    '%s'\n", underTest)
		fmt.Printf("Expected: '% x'\n", expected)
		fmt.Printf("Got:      '% x'\n", got)
		t.Fail()
	}
}

func Test_URLDecodeUni_Empty(t *testing.T) {
	equalURLDecodeUni(t, "", "")
}

func Test_URLDecodeUni_Nothing1(t *testing.T) {
	equalURLDecodeUni(t, "TestCase", "TestCase")
}

func Test_URLDecodeUni_Nothing2(t *testing.T) {
	equalURLDecodeUni(t, "Test\\0Case", "Test\\0Case")
}

func Test_URLDecodeUni_Valid(t *testing.T) {
	equalURLDecodeUni(t, "", "")
}

func Test_URLDecodeUni_Valid1(t *testing.T) {

	equalURLDecodeUni(t,
		"+%00%01%02%03%04%05%06%07%08%09%0a%0b%0c%0d%0e%0f%10%11%12%13%14%15%16%17%18%19%1a%1b%1c%1d%1e%1f%20%21%22%23%24%25%26%27%28%29%2a%2b%2c%2d%2e%2f%30%31%32%33%34%35%36%37%38%39%3a%3b%3c%3d%3e%3f%40%41%42%43%44%45%46%47%48%49%4a%4b%4c%4d%4e%4f%50%51%52%53%54%55%56%57%58%59%5a%5b%5c%5d%5e%5f%60%61%62%63%64%65%66%67%68%69%6a%6b%6c%6d%6e%6f%70%71%72%73%74%75%76%77%78%79%7a%7b%7c%7d%7e%7f%80%81%82%83%84%85%86%87%88%89%8a%8b%8c%8d%8e%8f%90%91%92%93%94%95%96%97%98%99%9a%9b%9c%9d%9e%9f%a0%a1%a2%a3%a4%a5%a6%a7%a8%a9%aa%ab%ac%ad%ae%af%b0%b1%b2%b3%b4%b5%b6%b7%b8%b9%ba%bb%bc%bd%be%bf%c0%c1%c2%c3%c4%c5%c6%c7%c8%c9%ca%cb%cc%cd%ce%cf%d0%d1%d2%d3%d4%d5%d6%d7%d8%d9%da%db%dc%dd%de%df%e0%e1%e2%e3%e4%e5%e6%e7%e8%e9%ea%eb%ec%ed%ee%ef%f0%f1%f2%f3%f4%f5%f6%f7%f8%f9%fa%fb%fc%fd%fe%ff",
		" \x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f \x21\x22\x23\x24\x25\x26\x27\x28\x29\x2a\x2b\x2c\x2d\x2e\x2f\x30\x31\x32\x33\x34\x35\x36\x37\x38\x39\x3a\x3b\x3c\x3d\x3e\x3f\x40\x41\x42\x43\x44\x45\x46\x47\x48\x49\x4a\x4b\x4c\x4d\x4e\x4f\x50\x51\x52\x53\x54\x55\x56\x57\x58\x59\x5a\x5b\x5c\x5d\x5e\x5f\x60\x61\x62\x63\x64\x65\x66\x67\x68\x69\x6a\x6b\x6c\x6d\x6e\x6f\x70\x71\x72\x73\x74\x75\x76\x77\x78\x79\x7a\x7b\x7c\x7d\x7e\x7f\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff",
	)
}

func Test_URLDecodeUni_Valid2(t *testing.T) {
	equalURLDecodeUni(t, "Test+Case", "Test Case")
}

func Test_URLDecodeUni_Valid3(t *testing.T) {
	//Use a codepage of 0 here since the test is based around a non existant codepage
	equalURLDecodeUniWithCodePage(t,
		"+%u0000%u0001%u0002%u0003%u0004%u0005%u0006%u0007%u0008%u0009%u000a%u000b%u000c%u000d%u000e%u000f%u0010%u0011%u0012%u0013%u0014%u0015%u0016%u0017%u0018%u0019%u001a%u001b%u001c%u001d%u001e%u001f%u0020%u0021%u0022%u0023%u0024%u0025%u0026%u0027%u0028%u0029%u002a%u002b%u002c%u002d%u002e%u002f%u0030%u0031%u0032%u0033%u0034%u0035%u0036%u0037%u0038%u0039%u003a%u003b%u003c%u003d%u003e%u003f%u0040%u0041%u0042%u0043%u0044%u0045%u0046%u0047%u0048%u0049%u004a%u004b%u004c%u004d%u004e%u004f%u0050%u0051%u0052%u0053%u0054%u0055%u0056%u0057%u0058%u0059%u005a%u005b%u005c%u005d%u005e%u005f%u0060%u0061%u0062%u0063%u0064%u0065%u0066%u0067%u0068%u0069%u006a%u006b%u006c%u006d%u006e%u006f%u0070%u0071%u0072%u0073%u0074%u0075%u0076%u0077%u0078%u0079%u007a%u007b%u007c%u007d%u007e%u007f%u0080%u0081%u0082%u0083%u0084%u0085%u0086%u0087%u0088%u0089%u008a%u008b%u008c%u008d%u008e%u008f%u0090%u0091%u0092%u0093%u0094%u0095%u0096%u0097%u0098%u0099%u009a%u009b%u009c%u009d%u009e%u009f%u00a0%u00a1%u00a2%u00a3%u00a4%u00a5%u00a6%u00a7%u00a8%u00a9%u00aa%u00ab%u00ac%u00ad%u00ae%u00af%u00b0%u00b1%u00b2%u00b3%u00b4%u00b5%u00b6%u00b7%u00b8%u00b9%u00ba%u00bb%u00bc%u00bd%u00be%u00bf%u00c0%u00c1%u00c2%u00c3%u00c4%u00c5%u00c6%u00c7%u00c8%u00c9%u00ca%u00cb%u00cc%u00cd%u00ce%u00cf%u00d0%u00d1%u00d2%u00d3%u00d4%u00d5%u00d6%u00d7%u00d8%u00d9%u00da%u00db%u00dc%u00dd%u00de%u00df%u00e0%u00e1%u00e2%u00e3%u00e4%u00e5%u00e6%u00e7%u00e8%u00e9%u00ea%u00eb%u00ec%u00ed%u00ee%u00ef%u00f0%u00f1%u00f2%u00f3%u00f4%u00f5%u00f6%u00f7%u00f8%u00f9%u00fa%u00fb%u00fc%u00fd%u00fe%u00ff",
		" \x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f \x21\x22\x23\x24\x25\x26\x27\x28\x29\x2a\x2b\x2c\x2d\x2e\x2f\x30\x31\x32\x33\x34\x35\x36\x37\x38\x39\x3a\x3b\x3c\x3d\x3e\x3f\x40\x41\x42\x43\x44\x45\x46\x47\x48\x49\x4a\x4b\x4c\x4d\x4e\x4f\x50\x51\x52\x53\x54\x55\x56\x57\x58\x59\x5a\x5b\x5c\x5d\x5e\x5f\x60\x61\x62\x63\x64\x65\x66\x67\x68\x69\x6a\x6b\x6c\x6d\x6e\x6f\x70\x71\x72\x73\x74\x75\x76\x77\x78\x79\x7a\x7b\x7c\x7d\x7e\x7f\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff",
		0,
	)
}

func Test_URLDecodeUni_Valid4(t *testing.T) {
	equalURLDecodeUni(t,
		"+%u1100%u1101%u1102%u1103%u1104%u1105%u1106%u1107%u1108%u1109%u110a%u110b%u110c%u110d%u110e%u110f%u1110%u1111%u1112%u1113%u1114%u1115%u1116%u1117%u1118%u1119%u111a%u111b%u111c%u111d%u111e%u111f%u1120%u1121%u1122%u1123%u1124%u1125%u1126%u1127%u1128%u1129%u112a%u112b%u112c%u112d%u112e%u112f%u1130%u1131%u1132%u1133%u1134%u1135%u1136%u1137%u1138%u1139%u113a%u113b%u113c%u113d%u113e%u113f%u1140%u1141%u1142%u1143%u1144%u1145%u1146%u1147%u1148%u1149%u114a%u114b%u114c%u114d%u114e%u114f%u1150%u1151%u1152%u1153%u1154%u1155%u1156%u1157%u1158%u1159%u115a%u115b%u115c%u115d%u115e%u115f%u1160%u1161%u1162%u1163%u1164%u1165%u1166%u1167%u1168%u1169%u116a%u116b%u116c%u116d%u116e%u116f%u1170%u1171%u1172%u1173%u1174%u1175%u1176%u1177%u1178%u1179%u117a%u117b%u117c%u117d%u117e%u117f%u1180%u1181%u1182%u1183%u1184%u1185%u1186%u1187%u1188%u1189%u118a%u118b%u118c%u118d%u118e%u118f%u1190%u1191%u1192%u1193%u1194%u1195%u1196%u1197%u1198%u1199%u119a%u119b%u119c%u119d%u119e%u119f%u11a0%u11a1%u11a2%u11a3%u11a4%u11a5%u11a6%u11a7%u11a8%u11a9%u11aa%u11ab%u11ac%u11ad%u11ae%u11af%u11b0%u11b1%u11b2%u11b3%u11b4%u11b5%u11b6%u11b7%u11b8%u11b9%u11ba%u11bb%u11bc%u11bd%u11be%u11bf%u11c0%u11c1%u11c2%u11c3%u11c4%u11c5%u11c6%u11c7%u11c8%u11c9%u11ca%u11cb%u11cc%u11cd%u11ce%u11cf%u11d0%u11d1%u11d2%u11d3%u11d4%u11d5%u11d6%u11d7%u11d8%u11d9%u11da%u11db%u11dc%u11dd%u11de%u11df%u11e0%u11e1%u11e2%u11e3%u11e4%u11e5%u11e6%u11e7%u11e8%u11e9%u11ea%u11eb%u11ec%u11ed%u11ee%u11ef%u11f0%u11f1%u11f2%u11f3%u11f4%u11f5%u11f6%u11f7%u11f8%u11f9%u11fa%u11fb%u11fc%u11fd%u11fe%u11ff",
		" \x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f \x21\x22\x23\x24\x25\x26\x27\x28\x29\x2a\x2b\x2c\x2d\x2e\x2f\x30\x31\x32\x33\x34\x35\x36\x37\x38\x39\x3a\x3b\x3c\x3d\x3e\x3f\x40\x41\x42\x43\x44\x45\x46\x47\x48\x49\x4a\x4b\x4c\x4d\x4e\x4f\x50\x51\x52\x53\x54\x55\x56\x57\x58\x59\x5a\x5b\x5c\x5d\x5e\x5f\x60\x61\x62\x63\x64\x65\x66\x67\x68\x69\x6a\x6b\x6c\x6d\x6e\x6f\x70\x71\x72\x73\x74\x75\x76\x77\x78\x79\x7a\x7b\x7c\x7d\x7e\x7f\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff",
	)
}

func Test_URLDecodeUni_FullWidthASCII(t *testing.T) {
	equalURLDecodeUni(t,
		"%uff01%uff02%uff03%uff04%uff05%uff06%uff07%uff08%uff09%uff0a%uff0b%uff0c%uff0d%uff0e%uff0f%uff10%uff11%uff12%uff13%uff14%uff15%uff16%uff17%uff18%uff19%uff1a%uff1b%uff1c%uff1d%uff1e%uff1f%uff20%uff21%uff22%uff23%uff24%uff25%uff26%uff27%uff28%uff29%uff2a%uff2b%uff2c%uff2d%uff2e%uff2f%uff30%uff31%uff32%uff33%uff34%uff35%uff36%uff37%uff38%uff39%uff3a%uff3b%uff3c%uff3d%uff3e%uff3f%uff40%uff41%uff42%uff43%uff44%uff45%uff46%uff47%uff48%uff49%uff4a%uff4b%uff4c%uff4d%uff4e%uff4f%uff50%uff51%uff52%uff53%uff54%uff55%uff56%uff57%uff58%uff59%uff5a%uff5b%uff5c%uff5d%uff5e",
		"\x21\x22\x23\x24\x25\x26\x27\x28\x29\x2a\x2b\x2c\x2d\x2e\x2f\x30\x31\x32\x33\x34\x35\x36\x37\x38\x39\x3a\x3b\x3c\x3d\x3e\x3f\x40\x41\x42\x43\x44\x45\x46\x47\x48\x49\x4a\x4b\x4c\x4d\x4e\x4f\x50\x51\x52\x53\x54\x55\x56\x57\x58\x59\x5a\x5b\x5c\x5d\x5e\x5f\x60\x61\x62\x63\x64\x65\x66\x67\x68\x69\x6a\x6b\x6c\x6d\x6e\x6f\x70\x71\x72\x73\x74\x75\x76\x77\x78\x79\x7a\x7b\x7c\x7d\x7e",
	)
}

func Test_URLDecodeUni_InvalidFullWidthASCII(t *testing.T) {
	equalURLDecodeUni(t,
		"%uff00%uff7f%uff80%uff81%uff82%uff83%uff84%uff85%uff86%uff87%uff88%uff89%uff8a%uff8b%uff8c%uff8d%uff8e%uff8f%uff90%uff91%uff92%uff93%uff94%uff95%uff96%uff97%uff98%uff99%uff9a%uff9b%uff9c%uff9d%uff9e%uff9f%uffa0%uffa1%uffa2%uffa3%uffa4%uffa5%uffa6%uffa7%uffa8%uffa9%uffaa%uffab%uffac%uffad%uffae%uffaf%uffb0%uffb1%uffb2%uffb3%uffb4%uffb5%uffb6%uffb7%uffb8%uffb9%uffba%uffbb%uffbc%uffbd%uffbe%uffbf%uffc0%uffc1%uffc2%uffc3%uffc4%uffc5%uffc6%uffc7%uffc8%uffc9%uffca%uffcb%uffcc%uffcd%uffce%uffcf%uffd0%uffd1%uffd2%uffd3%uffd4%uffd5%uffd6%uffd7%uffd8%uffd9%uffda%uffdb%uffdc%uffdd%uffde%uffdf%uffe0%uffe1%uffe2%uffe3%uffe4%uffe5%uffe6%uffe7%uffe8%uffe9%uffea%uffeb%uffec%uffed%uffee%uffef%ufff0%ufff1%ufff2%ufff3%ufff4%ufff5%ufff6%ufff7%ufff8%ufff9%ufffa%ufffb%ufffc%ufffd%ufffe%uffff",
		"\x00\x7f\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff",
	)
}

func Test_URLDecodeUni_PartialInvalid1(t *testing.T) {
	equalURLDecodeUni(t,
		"%+",
		"% ",
	)
}

func Test_URLDecodeUni_PartialInvalid2(t *testing.T) {
	equalURLDecodeUni(t,
		"%%20",
		"% ",
	)
}

func Test_URLDecodeUni_PartialInvalid3(t *testing.T) {
	equalURLDecodeUni(t,
		"%0g%20",
		"%0g ",
	)
}

func Test_URLDecodeUni_PartialInvalid4(t *testing.T) {
	equalURLDecodeUni(t,
		"%0%20",
		"%0 ",
	)
}

func Test_URLDecodeUni_PartialInvalid5(t *testing.T) {
	equalURLDecodeUni(t,
		"%g0%20",
		"%g0 ",
	)
}

func Test_URLDecodeUni_PartialInvalid6(t *testing.T) {
	equalURLDecodeUni(t,
		"%g%20",
		"%g ",
	)
}

func Test_URLDecodeUni_PartialInvalid7(t *testing.T) {
	equalURLDecodeUni(t,
		"%%u0020",
		"% ",
	)
}

func Test_URLDecodeUni_PartialInvalid8(t *testing.T) {
	equalURLDecodeUni(t,
		"%0g%u0020",
		"%0g ",
	)
}

func Test_URLDecodeUni_PartialInvalid9(t *testing.T) {
	equalURLDecodeUni(t,
		"%0%u0020",
		"%0 ",
	)
}

func Test_URLDecodeUni_PartialInvalid10(t *testing.T) {
	equalURLDecodeUni(t,
		"%g0%u0020",
		"%g0 ",
	)
}

func Test_URLDecodeUni_PartialInvalid11(t *testing.T) {
	equalURLDecodeUni(t,
		"%u%u0020",
		"%u ",
	)
}

func Test_URLDecodeUni_PartialInvalid12(t *testing.T) {
	equalURLDecodeUni(t,
		"%u0%u0020",
		"%u0 ",
	)
}

func Test_URLDecodeUni_PartialInvalid13(t *testing.T) {
	equalURLDecodeUni(t,
		"%u00%u0020",
		"%u00 ",
	)
}

func Test_URLDecodeUni_PartialInvalid14(t *testing.T) {
	equalURLDecodeUni(t,
		"%u000%u0020",
		"%u000 ",
	)
}

func Test_URLDecodeUni_PartialInvalid15(t *testing.T) {
	equalURLDecodeUni(t,
		"%u000g%u0020",
		"%u000g ",
	)
}

func TestInvalid1(t *testing.T) {
	equalURLDecodeUni(t,
		"%0%1%2%3%4%5%6%7%8%9%0%a%b%c%d%e%f",
		"%0%1%2%3%4%5%6%7%8%9%0%a%b%c%d%e%f",
	)
}

func Test_URLDecodeUni_Invalid2(t *testing.T) {
	equalURLDecodeUni(t,
		"%g0%g1%g2%g3%g4%g5%g6%g7%g8%g9%g0%ga%gb%gc%gd%ge%gf",
		"%g0%g1%g2%g3%g4%g5%g6%g7%g8%g9%g0%ga%gb%gc%gd%ge%gf",
	)
}

func Test_URLDecodeUni_Invalid3(t *testing.T) {
	equalURLDecodeUni(t,
		"%0g%1g%2g%3g%4g%5g%6g%7g%8g%9g%0g%ag%bg%cg%dg%eg%fg",
		"%0g%1g%2g%3g%4g%5g%6g%7g%8g%9g%0g%ag%bg%cg%dg%eg%fg",
	)
}

func Test_URLDecodeUni_Invalid4(t *testing.T) {
	equalURLDecodeUni(t,
		"%",
		"%",
	)
}

func Test_URLDecodeUni_Invalid5(t *testing.T) {
	equalURLDecodeUni(t,
		"%0",
		"%0",
	)
}

func Test_URLDecodeUni_Invalid6(t *testing.T) {
	equalURLDecodeUni(t,
		"%%",
		"%%",
	)
}

func Test_URLDecodeUni_Invalid7(t *testing.T) {
	equalURLDecodeUni(t,
		"%0g",
		"%0g",
	)
}

func Test_URLDecodeUni_Invalid8(t *testing.T) {
	equalURLDecodeUni(t,
		"%gg",
		"%gg",
	)
}

//The following tests were added to get 100% code coverage
func Test_URLDecodeUni_To100Percent1(t *testing.T) {
	//Test capital conversion in unicode
	equalURLDecodeUni(t,
		"%uff01%uFF02%uFf03",
		"\x21\x22\x23",
	)
}

func Test_URLDecodeUni_To100Percent2(t *testing.T) {
	//Test incompleat unicode
	equalURLDecodeUni(t,
		"%uff0",
		"%uff0",
	)
}
