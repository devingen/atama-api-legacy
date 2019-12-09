package util

import "hash/fnv"

var lastUsedHashInt int

func SequentialHashInt() int {
	lastUsedHashInt++
	return lastUsedHashInt
}

var lastUsedHashUInt16 uint16

func SequentialHashUInt16() uint16 {
	lastUsedHashUInt16++
	return lastUsedHashUInt16
}

var lastUsedHashUInt32 uint32

func SequentialHashUInt32() uint32 {
	lastUsedHashUInt32++
	return lastUsedHashUInt32
}

func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
