package heap

import "unicode/utf16"

var internedStrings = map[string]*Object{}

func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	chars := stringToUtf16(goStr)
	jChars := &Object{loader.LoadClass("[C"), chars, nil}

	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)

	internedStrings[goStr] = jStr
	return jStr
}

func stringToUtf16(str string) []uint16 {
	runes := []rune(str)
	return utf16.Encode(runes)
}
func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

func utf16ToString(chars []uint16) string {
	runes := utf16.Decode(chars)
	return string(runes)
}
func InternString(jStr *Object) *Object {
	goStr := GoString(jStr)
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	internedStrings[goStr] = jStr
	return jStr
}
