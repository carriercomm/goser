// Code generated by protoc-gen-go.
// source: pb/log.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/log.proto

It has these top-level messages:
	HTTP
	Origin
	Log
*/
package pb

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type CacheStatus int32

const (
	CacheStatus_CACHESTATUS_UNKNOWN CacheStatus = 0
	CacheStatus_MISS                CacheStatus = 1
	CacheStatus_EXPIRED             CacheStatus = 2
	CacheStatus_HIT                 CacheStatus = 3
)

var CacheStatus_name = map[int32]string{
	0: "CACHESTATUS_UNKNOWN",
	1: "MISS",
	2: "EXPIRED",
	3: "HIT",
}
var CacheStatus_value = map[string]int32{
	"CACHESTATUS_UNKNOWN": 0,
	"MISS":                1,
	"EXPIRED":             2,
	"HIT":                 3,
}

func (x CacheStatus) Enum() *CacheStatus {
	p := new(CacheStatus)
	*p = x
	return p
}
func (x CacheStatus) String() string {
	return proto.EnumName(CacheStatus_name, int32(x))
}
func (x *CacheStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CacheStatus_value, data, "CacheStatus")
	if err != nil {
		return err
	}
	*x = CacheStatus(value)
	return nil
}

type ZonePlan int32

const (
	ZonePlan_ZONEPLAN_UNKNOWN ZonePlan = 0
	ZonePlan_FREE             ZonePlan = 1
	ZonePlan_PRO              ZonePlan = 2
	ZonePlan_BIZ              ZonePlan = 3
	ZonePlan_ENT              ZonePlan = 4
)

var ZonePlan_name = map[int32]string{
	0: "ZONEPLAN_UNKNOWN",
	1: "FREE",
	2: "PRO",
	3: "BIZ",
	4: "ENT",
}
var ZonePlan_value = map[string]int32{
	"ZONEPLAN_UNKNOWN": 0,
	"FREE":             1,
	"PRO":              2,
	"BIZ":              3,
	"ENT":              4,
}

func (x ZonePlan) Enum() *ZonePlan {
	p := new(ZonePlan)
	*p = x
	return p
}
func (x ZonePlan) String() string {
	return proto.EnumName(ZonePlan_name, int32(x))
}
func (x *ZonePlan) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZonePlan_value, data, "ZonePlan")
	if err != nil {
		return err
	}
	*x = ZonePlan(value)
	return nil
}

type Country int32

const (
	Country_UNKNOWN Country = 0
	Country_A1      Country = 1
	Country_A2      Country = 2
	Country_O1      Country = 3
	Country_AD      Country = 4
	Country_AE      Country = 5
	Country_AF      Country = 6
	Country_AG      Country = 7
	Country_AI      Country = 8
	Country_AL      Country = 9
	Country_AM      Country = 10
	Country_AO      Country = 11
	Country_AP      Country = 12
	Country_AQ      Country = 13
	Country_AR      Country = 14
	Country_AS      Country = 15
	Country_AT      Country = 16
	Country_AU      Country = 17
	Country_AW      Country = 18
	Country_AX      Country = 19
	Country_AZ      Country = 20
	Country_BA      Country = 21
	Country_BB      Country = 22
	Country_BD      Country = 23
	Country_BE      Country = 24
	Country_BF      Country = 25
	Country_BG      Country = 26
	Country_BH      Country = 27
	Country_BI      Country = 28
	Country_BJ      Country = 29
	Country_BL      Country = 30
	Country_BM      Country = 31
	Country_BN      Country = 32
	Country_BO      Country = 33
	Country_BQ      Country = 34
	Country_BR      Country = 35
	Country_BS      Country = 36
	Country_BT      Country = 37
	Country_BV      Country = 38
	Country_BW      Country = 39
	Country_BY      Country = 40
	Country_BZ      Country = 41
	Country_CA      Country = 42
	Country_CC      Country = 43
	Country_CD      Country = 44
	Country_CF      Country = 45
	Country_CG      Country = 46
	Country_CH      Country = 47
	Country_CI      Country = 48
	Country_CK      Country = 49
	Country_CL      Country = 50
	Country_CM      Country = 51
	Country_CN      Country = 52
	Country_CO      Country = 53
	Country_CR      Country = 54
	Country_CU      Country = 55
	Country_CV      Country = 56
	Country_CW      Country = 57
	Country_CX      Country = 58
	Country_CY      Country = 59
	Country_CZ      Country = 60
	Country_DE      Country = 61
	Country_DJ      Country = 62
	Country_DK      Country = 63
	Country_DM      Country = 64
	Country_DO      Country = 65
	Country_DZ      Country = 66
	Country_EC      Country = 67
	Country_EE      Country = 68
	Country_EG      Country = 69
	Country_EH      Country = 70
	Country_ER      Country = 71
	Country_ES      Country = 72
	Country_ET      Country = 73
	Country_EU      Country = 74
	Country_FI      Country = 75
	Country_FJ      Country = 76
	Country_FK      Country = 77
	Country_FM      Country = 78
	Country_FO      Country = 79
	Country_FR      Country = 80
	Country_GA      Country = 81
	Country_GB      Country = 82
	Country_GD      Country = 83
	Country_GE      Country = 84
	Country_GF      Country = 85
	Country_GG      Country = 86
	Country_GH      Country = 87
	Country_GI      Country = 88
	Country_GL      Country = 89
	Country_GM      Country = 90
	Country_GN      Country = 91
	Country_GP      Country = 92
	Country_GQ      Country = 93
	Country_GR      Country = 94
	Country_GS      Country = 95
	Country_GT      Country = 96
	Country_GU      Country = 97
	Country_GW      Country = 98
	Country_GY      Country = 99
	Country_HK      Country = 100
	Country_HM      Country = 101
	Country_HN      Country = 102
	Country_HR      Country = 103
	Country_HT      Country = 104
	Country_HU      Country = 105
	Country_ID      Country = 106
	Country_IE      Country = 107
	Country_IL      Country = 108
	Country_IM      Country = 109
	Country_IN      Country = 110
	Country_IO      Country = 111
	Country_IQ      Country = 112
	Country_IR      Country = 113
	Country_IS      Country = 114
	Country_IT      Country = 115
	Country_JE      Country = 116
	Country_JM      Country = 117
	Country_JO      Country = 118
	Country_JP      Country = 119
	Country_KE      Country = 120
	Country_KG      Country = 121
	Country_KH      Country = 122
	Country_KI      Country = 123
	Country_KM      Country = 124
	Country_KN      Country = 125
	Country_KP      Country = 126
	Country_KR      Country = 127
	Country_KW      Country = 128
	Country_KY      Country = 129
	Country_KZ      Country = 130
	Country_LA      Country = 131
	Country_LB      Country = 132
	Country_LC      Country = 133
	Country_LI      Country = 134
	Country_LK      Country = 135
	Country_LR      Country = 136
	Country_LS      Country = 137
	Country_LT      Country = 138
	Country_LU      Country = 139
	Country_LV      Country = 140
	Country_LY      Country = 141
	Country_MA      Country = 142
	Country_MC      Country = 143
	Country_MD      Country = 144
	Country_ME      Country = 145
	Country_MF      Country = 146
	Country_MG      Country = 147
	Country_MH      Country = 148
	Country_MK      Country = 149
	Country_ML      Country = 150
	Country_MM      Country = 151
	Country_MN      Country = 152
	Country_MO      Country = 153
	Country_MP      Country = 154
	Country_MQ      Country = 155
	Country_MR      Country = 156
	Country_MS      Country = 157
	Country_MT      Country = 158
	Country_MU      Country = 159
	Country_MV      Country = 160
	Country_MW      Country = 161
	Country_MX      Country = 162
	Country_MY      Country = 163
	Country_MZ      Country = 164
	Country_NA      Country = 165
	Country_NC      Country = 166
	Country_NE      Country = 167
	Country_NF      Country = 168
	Country_NG      Country = 169
	Country_NI      Country = 170
	Country_NL      Country = 171
	Country_NO      Country = 172
	Country_NP      Country = 173
	Country_NR      Country = 174
	Country_NU      Country = 175
	Country_NZ      Country = 176
	Country_OM      Country = 177
	Country_PA      Country = 178
	Country_PE      Country = 179
	Country_PF      Country = 180
	Country_PG      Country = 181
	Country_PH      Country = 182
	Country_PK      Country = 183
	Country_PL      Country = 184
	Country_PM      Country = 185
	Country_PN      Country = 186
	Country_PR      Country = 187
	Country_PS      Country = 188
	Country_PT      Country = 189
	Country_PW      Country = 190
	Country_PY      Country = 191
	Country_QA      Country = 192
	Country_RE      Country = 193
	Country_RO      Country = 194
	Country_RS      Country = 195
	Country_RU      Country = 196
	Country_RW      Country = 197
	Country_SA      Country = 198
	Country_SB      Country = 199
	Country_SC      Country = 200
	Country_SD      Country = 201
	Country_SE      Country = 202
	Country_SG      Country = 203
	Country_SH      Country = 204
	Country_SI      Country = 205
	Country_SJ      Country = 206
	Country_SK      Country = 207
	Country_SL      Country = 208
	Country_SM      Country = 209
	Country_SN      Country = 210
	Country_SO      Country = 211
	Country_SR      Country = 212
	Country_SS      Country = 213
	Country_ST      Country = 214
	Country_SV      Country = 215
	Country_SX      Country = 216
	Country_SY      Country = 217
	Country_SZ      Country = 218
	Country_TC      Country = 219
	Country_TD      Country = 220
	Country_TF      Country = 221
	Country_TG      Country = 222
	Country_TH      Country = 223
	Country_TJ      Country = 224
	Country_TK      Country = 225
	Country_TL      Country = 226
	Country_TM      Country = 227
	Country_TN      Country = 228
	Country_TO      Country = 229
	Country_TR      Country = 230
	Country_TT      Country = 231
	Country_TV      Country = 232
	Country_TW      Country = 233
	Country_TZ      Country = 234
	Country_UA      Country = 235
	Country_UG      Country = 236
	Country_UM      Country = 237
	Country_US      Country = 238
	Country_UY      Country = 239
	Country_UZ      Country = 240
	Country_VA      Country = 241
	Country_VC      Country = 242
	Country_VE      Country = 243
	Country_VG      Country = 244
	Country_VI      Country = 245
	Country_VN      Country = 246
	Country_VU      Country = 247
	Country_WF      Country = 248
	Country_WS      Country = 249
	Country_XX      Country = 250
	Country_YE      Country = 251
	Country_YT      Country = 252
	Country_ZA      Country = 253
	Country_ZM      Country = 254
	Country_ZW      Country = 255
)

var Country_name = map[int32]string{
	0:   "UNKNOWN",
	1:   "A1",
	2:   "A2",
	3:   "O1",
	4:   "AD",
	5:   "AE",
	6:   "AF",
	7:   "AG",
	8:   "AI",
	9:   "AL",
	10:  "AM",
	11:  "AO",
	12:  "AP",
	13:  "AQ",
	14:  "AR",
	15:  "AS",
	16:  "AT",
	17:  "AU",
	18:  "AW",
	19:  "AX",
	20:  "AZ",
	21:  "BA",
	22:  "BB",
	23:  "BD",
	24:  "BE",
	25:  "BF",
	26:  "BG",
	27:  "BH",
	28:  "BI",
	29:  "BJ",
	30:  "BL",
	31:  "BM",
	32:  "BN",
	33:  "BO",
	34:  "BQ",
	35:  "BR",
	36:  "BS",
	37:  "BT",
	38:  "BV",
	39:  "BW",
	40:  "BY",
	41:  "BZ",
	42:  "CA",
	43:  "CC",
	44:  "CD",
	45:  "CF",
	46:  "CG",
	47:  "CH",
	48:  "CI",
	49:  "CK",
	50:  "CL",
	51:  "CM",
	52:  "CN",
	53:  "CO",
	54:  "CR",
	55:  "CU",
	56:  "CV",
	57:  "CW",
	58:  "CX",
	59:  "CY",
	60:  "CZ",
	61:  "DE",
	62:  "DJ",
	63:  "DK",
	64:  "DM",
	65:  "DO",
	66:  "DZ",
	67:  "EC",
	68:  "EE",
	69:  "EG",
	70:  "EH",
	71:  "ER",
	72:  "ES",
	73:  "ET",
	74:  "EU",
	75:  "FI",
	76:  "FJ",
	77:  "FK",
	78:  "FM",
	79:  "FO",
	80:  "FR",
	81:  "GA",
	82:  "GB",
	83:  "GD",
	84:  "GE",
	85:  "GF",
	86:  "GG",
	87:  "GH",
	88:  "GI",
	89:  "GL",
	90:  "GM",
	91:  "GN",
	92:  "GP",
	93:  "GQ",
	94:  "GR",
	95:  "GS",
	96:  "GT",
	97:  "GU",
	98:  "GW",
	99:  "GY",
	100: "HK",
	101: "HM",
	102: "HN",
	103: "HR",
	104: "HT",
	105: "HU",
	106: "ID",
	107: "IE",
	108: "IL",
	109: "IM",
	110: "IN",
	111: "IO",
	112: "IQ",
	113: "IR",
	114: "IS",
	115: "IT",
	116: "JE",
	117: "JM",
	118: "JO",
	119: "JP",
	120: "KE",
	121: "KG",
	122: "KH",
	123: "KI",
	124: "KM",
	125: "KN",
	126: "KP",
	127: "KR",
	128: "KW",
	129: "KY",
	130: "KZ",
	131: "LA",
	132: "LB",
	133: "LC",
	134: "LI",
	135: "LK",
	136: "LR",
	137: "LS",
	138: "LT",
	139: "LU",
	140: "LV",
	141: "LY",
	142: "MA",
	143: "MC",
	144: "MD",
	145: "ME",
	146: "MF",
	147: "MG",
	148: "MH",
	149: "MK",
	150: "ML",
	151: "MM",
	152: "MN",
	153: "MO",
	154: "MP",
	155: "MQ",
	156: "MR",
	157: "MS",
	158: "MT",
	159: "MU",
	160: "MV",
	161: "MW",
	162: "MX",
	163: "MY",
	164: "MZ",
	165: "NA",
	166: "NC",
	167: "NE",
	168: "NF",
	169: "NG",
	170: "NI",
	171: "NL",
	172: "NO",
	173: "NP",
	174: "NR",
	175: "NU",
	176: "NZ",
	177: "OM",
	178: "PA",
	179: "PE",
	180: "PF",
	181: "PG",
	182: "PH",
	183: "PK",
	184: "PL",
	185: "PM",
	186: "PN",
	187: "PR",
	188: "PS",
	189: "PT",
	190: "PW",
	191: "PY",
	192: "QA",
	193: "RE",
	194: "RO",
	195: "RS",
	196: "RU",
	197: "RW",
	198: "SA",
	199: "SB",
	200: "SC",
	201: "SD",
	202: "SE",
	203: "SG",
	204: "SH",
	205: "SI",
	206: "SJ",
	207: "SK",
	208: "SL",
	209: "SM",
	210: "SN",
	211: "SO",
	212: "SR",
	213: "SS",
	214: "ST",
	215: "SV",
	216: "SX",
	217: "SY",
	218: "SZ",
	219: "TC",
	220: "TD",
	221: "TF",
	222: "TG",
	223: "TH",
	224: "TJ",
	225: "TK",
	226: "TL",
	227: "TM",
	228: "TN",
	229: "TO",
	230: "TR",
	231: "TT",
	232: "TV",
	233: "TW",
	234: "TZ",
	235: "UA",
	236: "UG",
	237: "UM",
	238: "US",
	239: "UY",
	240: "UZ",
	241: "VA",
	242: "VC",
	243: "VE",
	244: "VG",
	245: "VI",
	246: "VN",
	247: "VU",
	248: "WF",
	249: "WS",
	250: "XX",
	251: "YE",
	252: "YT",
	253: "ZA",
	254: "ZM",
	255: "ZW",
}
var Country_value = map[string]int32{
	"UNKNOWN": 0,
	"A1":      1,
	"A2":      2,
	"O1":      3,
	"AD":      4,
	"AE":      5,
	"AF":      6,
	"AG":      7,
	"AI":      8,
	"AL":      9,
	"AM":      10,
	"AO":      11,
	"AP":      12,
	"AQ":      13,
	"AR":      14,
	"AS":      15,
	"AT":      16,
	"AU":      17,
	"AW":      18,
	"AX":      19,
	"AZ":      20,
	"BA":      21,
	"BB":      22,
	"BD":      23,
	"BE":      24,
	"BF":      25,
	"BG":      26,
	"BH":      27,
	"BI":      28,
	"BJ":      29,
	"BL":      30,
	"BM":      31,
	"BN":      32,
	"BO":      33,
	"BQ":      34,
	"BR":      35,
	"BS":      36,
	"BT":      37,
	"BV":      38,
	"BW":      39,
	"BY":      40,
	"BZ":      41,
	"CA":      42,
	"CC":      43,
	"CD":      44,
	"CF":      45,
	"CG":      46,
	"CH":      47,
	"CI":      48,
	"CK":      49,
	"CL":      50,
	"CM":      51,
	"CN":      52,
	"CO":      53,
	"CR":      54,
	"CU":      55,
	"CV":      56,
	"CW":      57,
	"CX":      58,
	"CY":      59,
	"CZ":      60,
	"DE":      61,
	"DJ":      62,
	"DK":      63,
	"DM":      64,
	"DO":      65,
	"DZ":      66,
	"EC":      67,
	"EE":      68,
	"EG":      69,
	"EH":      70,
	"ER":      71,
	"ES":      72,
	"ET":      73,
	"EU":      74,
	"FI":      75,
	"FJ":      76,
	"FK":      77,
	"FM":      78,
	"FO":      79,
	"FR":      80,
	"GA":      81,
	"GB":      82,
	"GD":      83,
	"GE":      84,
	"GF":      85,
	"GG":      86,
	"GH":      87,
	"GI":      88,
	"GL":      89,
	"GM":      90,
	"GN":      91,
	"GP":      92,
	"GQ":      93,
	"GR":      94,
	"GS":      95,
	"GT":      96,
	"GU":      97,
	"GW":      98,
	"GY":      99,
	"HK":      100,
	"HM":      101,
	"HN":      102,
	"HR":      103,
	"HT":      104,
	"HU":      105,
	"ID":      106,
	"IE":      107,
	"IL":      108,
	"IM":      109,
	"IN":      110,
	"IO":      111,
	"IQ":      112,
	"IR":      113,
	"IS":      114,
	"IT":      115,
	"JE":      116,
	"JM":      117,
	"JO":      118,
	"JP":      119,
	"KE":      120,
	"KG":      121,
	"KH":      122,
	"KI":      123,
	"KM":      124,
	"KN":      125,
	"KP":      126,
	"KR":      127,
	"KW":      128,
	"KY":      129,
	"KZ":      130,
	"LA":      131,
	"LB":      132,
	"LC":      133,
	"LI":      134,
	"LK":      135,
	"LR":      136,
	"LS":      137,
	"LT":      138,
	"LU":      139,
	"LV":      140,
	"LY":      141,
	"MA":      142,
	"MC":      143,
	"MD":      144,
	"ME":      145,
	"MF":      146,
	"MG":      147,
	"MH":      148,
	"MK":      149,
	"ML":      150,
	"MM":      151,
	"MN":      152,
	"MO":      153,
	"MP":      154,
	"MQ":      155,
	"MR":      156,
	"MS":      157,
	"MT":      158,
	"MU":      159,
	"MV":      160,
	"MW":      161,
	"MX":      162,
	"MY":      163,
	"MZ":      164,
	"NA":      165,
	"NC":      166,
	"NE":      167,
	"NF":      168,
	"NG":      169,
	"NI":      170,
	"NL":      171,
	"NO":      172,
	"NP":      173,
	"NR":      174,
	"NU":      175,
	"NZ":      176,
	"OM":      177,
	"PA":      178,
	"PE":      179,
	"PF":      180,
	"PG":      181,
	"PH":      182,
	"PK":      183,
	"PL":      184,
	"PM":      185,
	"PN":      186,
	"PR":      187,
	"PS":      188,
	"PT":      189,
	"PW":      190,
	"PY":      191,
	"QA":      192,
	"RE":      193,
	"RO":      194,
	"RS":      195,
	"RU":      196,
	"RW":      197,
	"SA":      198,
	"SB":      199,
	"SC":      200,
	"SD":      201,
	"SE":      202,
	"SG":      203,
	"SH":      204,
	"SI":      205,
	"SJ":      206,
	"SK":      207,
	"SL":      208,
	"SM":      209,
	"SN":      210,
	"SO":      211,
	"SR":      212,
	"SS":      213,
	"ST":      214,
	"SV":      215,
	"SX":      216,
	"SY":      217,
	"SZ":      218,
	"TC":      219,
	"TD":      220,
	"TF":      221,
	"TG":      222,
	"TH":      223,
	"TJ":      224,
	"TK":      225,
	"TL":      226,
	"TM":      227,
	"TN":      228,
	"TO":      229,
	"TR":      230,
	"TT":      231,
	"TV":      232,
	"TW":      233,
	"TZ":      234,
	"UA":      235,
	"UG":      236,
	"UM":      237,
	"US":      238,
	"UY":      239,
	"UZ":      240,
	"VA":      241,
	"VC":      242,
	"VE":      243,
	"VG":      244,
	"VI":      245,
	"VN":      246,
	"VU":      247,
	"WF":      248,
	"WS":      249,
	"XX":      250,
	"YE":      251,
	"YT":      252,
	"ZA":      253,
	"ZM":      254,
	"ZW":      255,
}

func (x Country) Enum() *Country {
	p := new(Country)
	*p = x
	return p
}
func (x Country) String() string {
	return proto.EnumName(Country_name, int32(x))
}
func (x *Country) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Country_value, data, "Country")
	if err != nil {
		return err
	}
	*x = Country(value)
	return nil
}

type HTTP_Protocol int32

const (
	HTTP_HTTP_PROTOCOL_UNKNOWN HTTP_Protocol = 0
	HTTP_HTTP10                HTTP_Protocol = 1
	HTTP_HTTP11                HTTP_Protocol = 2
)

var HTTP_Protocol_name = map[int32]string{
	0: "HTTP_PROTOCOL_UNKNOWN",
	1: "HTTP10",
	2: "HTTP11",
}
var HTTP_Protocol_value = map[string]int32{
	"HTTP_PROTOCOL_UNKNOWN": 0,
	"HTTP10":                1,
	"HTTP11":                2,
}

func (x HTTP_Protocol) Enum() *HTTP_Protocol {
	p := new(HTTP_Protocol)
	*p = x
	return p
}
func (x HTTP_Protocol) String() string {
	return proto.EnumName(HTTP_Protocol_name, int32(x))
}
func (x *HTTP_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HTTP_Protocol_value, data, "HTTP_Protocol")
	if err != nil {
		return err
	}
	*x = HTTP_Protocol(value)
	return nil
}

type HTTP_Method int32

const (
	HTTP_METHOD_UNKNOWN HTTP_Method = 0
	HTTP_GET            HTTP_Method = 1
	HTTP_POST           HTTP_Method = 2
	HTTP_DELETE         HTTP_Method = 3
	HTTP_PUT            HTTP_Method = 4
	HTTP_HEAD           HTTP_Method = 5
	HTTP_PURGE          HTTP_Method = 6
	HTTP_OPTIONS        HTTP_Method = 7
	HTTP_PROPFIND       HTTP_Method = 8
	HTTP_MKCOL          HTTP_Method = 9
	HTTP_PATCH          HTTP_Method = 10
)

var HTTP_Method_name = map[int32]string{
	0:  "METHOD_UNKNOWN",
	1:  "GET",
	2:  "POST",
	3:  "DELETE",
	4:  "PUT",
	5:  "HEAD",
	6:  "PURGE",
	7:  "OPTIONS",
	8:  "PROPFIND",
	9:  "MKCOL",
	10: "PATCH",
}
var HTTP_Method_value = map[string]int32{
	"METHOD_UNKNOWN": 0,
	"GET":            1,
	"POST":           2,
	"DELETE":         3,
	"PUT":            4,
	"HEAD":           5,
	"PURGE":          6,
	"OPTIONS":        7,
	"PROPFIND":       8,
	"MKCOL":          9,
	"PATCH":          10,
}

func (x HTTP_Method) Enum() *HTTP_Method {
	p := new(HTTP_Method)
	*p = x
	return p
}
func (x HTTP_Method) String() string {
	return proto.EnumName(HTTP_Method_name, int32(x))
}
func (x *HTTP_Method) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HTTP_Method_value, data, "HTTP_Method")
	if err != nil {
		return err
	}
	*x = HTTP_Method(value)
	return nil
}

type Origin_Protocol int32

const (
	Origin_ORIGIN_PROTOCOL_UNKNOWN Origin_Protocol = 0
	Origin_HTTP                    Origin_Protocol = 1
	Origin_HTTPS                   Origin_Protocol = 2
)

var Origin_Protocol_name = map[int32]string{
	0: "ORIGIN_PROTOCOL_UNKNOWN",
	1: "HTTP",
	2: "HTTPS",
}
var Origin_Protocol_value = map[string]int32{
	"ORIGIN_PROTOCOL_UNKNOWN": 0,
	"HTTP":  1,
	"HTTPS": 2,
}

func (x Origin_Protocol) Enum() *Origin_Protocol {
	p := new(Origin_Protocol)
	*p = x
	return p
}
func (x Origin_Protocol) String() string {
	return proto.EnumName(Origin_Protocol_name, int32(x))
}
func (x *Origin_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Origin_Protocol_value, data, "Origin_Protocol")
	if err != nil {
		return err
	}
	*x = Origin_Protocol(value)
	return nil
}

type HTTP struct {
	Protocol         *HTTP_Protocol `protobuf:"varint,1,opt,name=protocol,enum=proto.HTTP_Protocol" json:"protocol,omitempty"`
	Status           *uint32        `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	HostStatus       *uint32        `protobuf:"varint,3,opt,name=hostStatus" json:"hostStatus,omitempty"`
	UpStatus         *uint32        `protobuf:"varint,4,opt,name=upStatus" json:"upStatus,omitempty"`
	Method           *HTTP_Method   `protobuf:"varint,5,opt,name=method,enum=proto.HTTP_Method" json:"method,omitempty"`
	ContentType      *string        `protobuf:"bytes,6,opt,name=contentType" json:"contentType,omitempty"`
	UserAgent        *string        `protobuf:"bytes,7,opt,name=userAgent" json:"userAgent,omitempty"`
	Referer          *string        `protobuf:"bytes,8,opt,name=referer" json:"referer,omitempty"`
	RequestURI       *string        `protobuf:"bytes,9,opt,name=requestURI" json:"requestURI,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *HTTP) Reset()         { *m = HTTP{} }
func (m *HTTP) String() string { return proto.CompactTextString(m) }
func (*HTTP) ProtoMessage()    {}

func (m *HTTP) GetProtocol() HTTP_Protocol {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return HTTP_HTTP_PROTOCOL_UNKNOWN
}

func (m *HTTP) GetStatus() uint32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *HTTP) GetHostStatus() uint32 {
	if m != nil && m.HostStatus != nil {
		return *m.HostStatus
	}
	return 0
}

func (m *HTTP) GetUpStatus() uint32 {
	if m != nil && m.UpStatus != nil {
		return *m.UpStatus
	}
	return 0
}

func (m *HTTP) GetMethod() HTTP_Method {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return HTTP_METHOD_UNKNOWN
}

func (m *HTTP) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *HTTP) GetUserAgent() string {
	if m != nil && m.UserAgent != nil {
		return *m.UserAgent
	}
	return ""
}

func (m *HTTP) GetReferer() string {
	if m != nil && m.Referer != nil {
		return *m.Referer
	}
	return ""
}

func (m *HTTP) GetRequestURI() string {
	if m != nil && m.RequestURI != nil {
		return *m.RequestURI
	}
	return ""
}

type Origin struct {
	Ip               []byte           `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port             *uint32          `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	Hostname         *string          `protobuf:"bytes,3,opt,name=hostname" json:"hostname,omitempty"`
	Protocol         *Origin_Protocol `protobuf:"varint,4,opt,name=protocol,enum=proto.Origin_Protocol" json:"protocol,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Origin) Reset()         { *m = Origin{} }
func (m *Origin) String() string { return proto.CompactTextString(m) }
func (*Origin) ProtoMessage()    {}

func (m *Origin) GetIp() []byte {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *Origin) GetPort() uint32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *Origin) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *Origin) GetProtocol() Origin_Protocol {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return Origin_ORIGIN_PROTOCOL_UNKNOWN
}

type Log struct {
	Timestamp        *int64       `protobuf:"fixed64,1,opt,name=timestamp" json:"timestamp,omitempty"`
	ZoneId           *uint32      `protobuf:"varint,2,opt,name=zoneId" json:"zoneId,omitempty"`
	ZonePlan         *ZonePlan    `protobuf:"varint,3,opt,name=zonePlan,enum=proto.ZonePlan" json:"zonePlan,omitempty"`
	Http             *HTTP        `protobuf:"bytes,4,opt,name=http" json:"http,omitempty"`
	Origin           *Origin      `protobuf:"bytes,5,opt,name=origin" json:"origin,omitempty"`
	Country          *Country     `protobuf:"varint,6,opt,name=country,enum=proto.Country" json:"country,omitempty"`
	CacheStatus      *CacheStatus `protobuf:"varint,7,opt,name=cacheStatus,enum=proto.CacheStatus" json:"cacheStatus,omitempty"`
	ServerIp         []byte       `protobuf:"bytes,8,opt,name=serverIp" json:"serverIp,omitempty"`
	ServerName       *string      `protobuf:"bytes,9,opt,name=serverName" json:"serverName,omitempty"`
	RemoteIp         []byte       `protobuf:"bytes,10,opt,name=remoteIp" json:"remoteIp,omitempty"`
	BytesDlv         *uint64      `protobuf:"varint,11,opt,name=bytesDlv" json:"bytesDlv,omitempty"`
	RayId            *string      `protobuf:"bytes,12,opt,name=rayId" json:"rayId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}

func (m *Log) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Log) GetZoneId() uint32 {
	if m != nil && m.ZoneId != nil {
		return *m.ZoneId
	}
	return 0
}

func (m *Log) GetZonePlan() ZonePlan {
	if m != nil && m.ZonePlan != nil {
		return *m.ZonePlan
	}
	return ZonePlan_ZONEPLAN_UNKNOWN
}

func (m *Log) GetHttp() *HTTP {
	if m != nil {
		return m.Http
	}
	return nil
}

func (m *Log) GetOrigin() *Origin {
	if m != nil {
		return m.Origin
	}
	return nil
}

func (m *Log) GetCountry() Country {
	if m != nil && m.Country != nil {
		return *m.Country
	}
	return Country_UNKNOWN
}

func (m *Log) GetCacheStatus() CacheStatus {
	if m != nil && m.CacheStatus != nil {
		return *m.CacheStatus
	}
	return CacheStatus_CACHESTATUS_UNKNOWN
}

func (m *Log) GetServerIp() []byte {
	if m != nil {
		return m.ServerIp
	}
	return nil
}

func (m *Log) GetServerName() string {
	if m != nil && m.ServerName != nil {
		return *m.ServerName
	}
	return ""
}

func (m *Log) GetRemoteIp() []byte {
	if m != nil {
		return m.RemoteIp
	}
	return nil
}

func (m *Log) GetBytesDlv() uint64 {
	if m != nil && m.BytesDlv != nil {
		return *m.BytesDlv
	}
	return 0
}

func (m *Log) GetRayId() string {
	if m != nil && m.RayId != nil {
		return *m.RayId
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.CacheStatus", CacheStatus_name, CacheStatus_value)
	proto.RegisterEnum("proto.ZonePlan", ZonePlan_name, ZonePlan_value)
	proto.RegisterEnum("proto.Country", Country_name, Country_value)
	proto.RegisterEnum("proto.HTTP_Protocol", HTTP_Protocol_name, HTTP_Protocol_value)
	proto.RegisterEnum("proto.HTTP_Method", HTTP_Method_name, HTTP_Method_value)
	proto.RegisterEnum("proto.Origin_Protocol", Origin_Protocol_name, Origin_Protocol_value)
}
