// Code generated by "stringer -type Enum consts.go"; DO NOT EDIT.

package rmeta

import "strconv"

const _Enum_name = "BaseCharShortIntLongFloatCounterCharStarDoubleDouble32LegacyCharUCharUShortUIntULongBitsLong64ULong64BoolFloat16OffsetLOffsetPObjectAnyObjectpObjectPTStringTObjectTNamedAnypAnyPAnyPnoVTSTLpSkipSkipLSkipPConvConvLConvPSTLSTLstringStreamerStreamLoopCacheArtificialCacheNewCacheDeleteNeedObjectForVirtualBaseClassMissing"

var _Enum_map = map[Enum]string{
	0:     _Enum_name[0:4],
	1:     _Enum_name[4:8],
	2:     _Enum_name[8:13],
	3:     _Enum_name[13:16],
	4:     _Enum_name[16:20],
	5:     _Enum_name[20:25],
	6:     _Enum_name[25:32],
	7:     _Enum_name[32:40],
	8:     _Enum_name[40:46],
	9:     _Enum_name[46:54],
	10:    _Enum_name[54:64],
	11:    _Enum_name[64:69],
	12:    _Enum_name[69:75],
	13:    _Enum_name[75:79],
	14:    _Enum_name[79:84],
	15:    _Enum_name[84:88],
	16:    _Enum_name[88:94],
	17:    _Enum_name[94:101],
	18:    _Enum_name[101:105],
	19:    _Enum_name[105:112],
	20:    _Enum_name[112:119],
	40:    _Enum_name[119:126],
	61:    _Enum_name[126:132],
	62:    _Enum_name[132:135],
	63:    _Enum_name[135:142],
	64:    _Enum_name[142:149],
	65:    _Enum_name[149:156],
	66:    _Enum_name[156:163],
	67:    _Enum_name[163:169],
	68:    _Enum_name[169:173],
	69:    _Enum_name[173:177],
	70:    _Enum_name[177:185],
	71:    _Enum_name[185:189],
	100:   _Enum_name[189:193],
	120:   _Enum_name[193:198],
	140:   _Enum_name[198:203],
	200:   _Enum_name[203:207],
	220:   _Enum_name[207:212],
	240:   _Enum_name[212:217],
	300:   _Enum_name[217:220],
	365:   _Enum_name[220:229],
	500:   _Enum_name[229:237],
	501:   _Enum_name[237:247],
	600:   _Enum_name[247:252],
	1000:  _Enum_name[252:262],
	1001:  _Enum_name[262:270],
	1002:  _Enum_name[270:281],
	99997: _Enum_name[281:310],
	99999: _Enum_name[310:317],
}

func (i Enum) String() string {
	if str, ok := _Enum_map[i]; ok {
		return str
	}
	return "Enum(" + strconv.FormatInt(int64(i), 10) + ")"
}
