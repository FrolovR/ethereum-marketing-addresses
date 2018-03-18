package main

import (
	"fmt"
	"math/big"
	"strconv"
	"bytes"
	"strings"
	"./sha3"
	"./secp256k1"
	"encoding/hex"
)

func main() {
	var P string = "580e9e8c13aa4da3040962f57273ed5f1759c4edb32c6705310362bbf52f2c3680c811172f177b2c3ed1d53ac7c521b8b1813284fbe0279557401bfe20d256da"
	var prefix string = "face"
	Curve := secp256k1.S256()
	x1, y1 := pubKey(P)
	var Kprefix int = 1
	var addr string

	for !isPrefix(prefix, addr) {
		Knew := pad64(Hex(int64(Kprefix)))
		x2, y2 := Curve.Privtopub(decode_hex(Knew))
		x, y := Curve.Add(x1, y1, x2, y2)
		Pnew := pad128(concat(hexBig(x), hexBig(y)))
		Pnew = pad64(encode_hex(Keccak256(decode_hex(Pnew))))
		addr = Pnew[24:]
		Kprefix++
	}

	fmt.Print(addr)
}

func pubKey(P string) (*big.Int, *big.Int) {
	x, _ := new(big.Int).SetString(P[:64], 16)
	y, _ := new(big.Int).SetString(P[64:], 16)
	return x, y
}

func isPrefix(prefix string, addr string) bool {
	return strings.HasPrefix(addr, prefix)
}

func pad64(str string) string {
	var buffer bytes.Buffer

	for i := 0; i < 64 - len(str); i++ {
		buffer.WriteString("0")
	}
	strs := []string{buffer.String(), str}

	return strings.Join(strs, "")
}

func pad128(str string) string {
	var buffer bytes.Buffer
	for i := 0; i < 128 - len(str); i++ {
		buffer.WriteString("0")
	}
	strs := []string{buffer.String(), str}

	return strings.Join(strs, "")
}

func Hex(number int64) string {
	return strconv.FormatInt(number, 16)
}

func Keccak256(data ...[]byte) []byte {
	d := sha3.NewKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

func hexBig(number *big.Int) string {
	return fmt.Sprintf("%x", number)
}

func decode_hex(str string) []byte {
	result, _ := hex.DecodeString(str)
	return result
}

func encode_hex(bytes []byte) string {
	result := hex.EncodeToString(bytes)
	return result
}

func concat(str1 string, str2 string) string {
	strs := []string{str1, str2}
	return strings.Join(strs, "")
}