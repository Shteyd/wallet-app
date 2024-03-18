package byteconv

import "unsafe"

func ToString(slice []byte) string {
	return unsafe.String(unsafe.SliceData(slice), len(slice))
}

func ToSlice(str string) []byte {
	return unsafe.Slice(unsafe.StringData(str), len(str))
}
