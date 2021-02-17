/*
2021-02-10

Written by wowlsh93
*/

package util

import (
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

func hexToBigInt(val string) *big.Int {
	if val[:2] == "0x" {
		val = val[2:]
	}
	i := new(big.Int)
	i.SetString(val, 16)
	return i
}

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{Data: bh.Data, Len: bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

//func BytesToString2(b []byte) string {
//	s := make([]string, len(b))
//	for i := range b {
//		s[i] = strconv.Itoa(int(b[i]))
//	}
//	return strings.Join(s, ".")
//}

// ParseInt parse hex string value to int
func ParseInt(value string) (int, error) {
	i, err := strconv.ParseInt(strings.TrimPrefix(value, "0x"), 16, 64)
	if err != nil {
		return 0, err
	}

	return int(i), nil
}

// ParseBigInt parse hex string value to big.Int
func ParseBigInt(value string) (big.Int, error) {
	i := big.Int{}
	_, err := fmt.Sscan(value, &i)

	return i, err
}

// IntToHex convert int to hexadecimal representation
func IntToHex(i int) string {
	return fmt.Sprintf("0x%x", i)
}

// BigToHex covert big.Int to hexadecimal representation
func BigToHex(bigInt big.Int) string {
	if bigInt.BitLen() == 0 {
		return "0x0"
	}

	return "0x" + strings.TrimPrefix(fmt.Sprintf("%x", bigInt.Bytes()), "0")
}
