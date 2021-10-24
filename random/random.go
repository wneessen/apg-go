package random

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"
	"strings"
)

// Bitmask sizes for the string generators (based on 93 chars total)
const (
	letterIdxBits = 7                    // 7 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// GetChar generates random characters based on given character range
// and password length
func GetChar(cr string, l int) (string, error) {
	if l < 1 {
		return "", fmt.Errorf("length is negative")
	}
	rs := strings.Builder{}
	rs.Grow(l)
	crl := len(cr)

	rp := make([]byte, 8)
	_, err := rand.Read(rp)
	if err != nil {
		return rs.String(), err
	}
	for i, c, r := l-1, binary.BigEndian.Uint64(rp), letterIdxMax; i >= 0; {
		if r == 0 {
			_, err := rand.Read(rp)
			if err != nil {
				return rs.String(), err
			}
			c, r = binary.BigEndian.Uint64(rp), letterIdxMax
		}
		if idx := int(c & letterIdxMask); idx < crl {
			rs.WriteByte(cr[idx])
			i--
		}
		c >>= letterIdxBits
		r--
	}
	return rs.String(), nil
}

// GetNum generates a random number with given maximum value
func GetNum(maxNum int) (int, error) {
	if maxNum <= 0 {
		err := fmt.Errorf("provided maxNum is <= 0: %v", maxNum)
		return 0, err
	}
	maxNumBigInt := big.NewInt(int64(maxNum))
	if !maxNumBigInt.IsUint64() {
		err := fmt.Errorf("big.NewInt() generation returned negative value: %v", maxNumBigInt)
		return 0, err
	}
	randNum64, err := rand.Int(rand.Reader, maxNumBigInt)
	if err != nil {
		return 0, err
	}
	randNum := int(randNum64.Int64())
	if randNum < 0 {
		err := fmt.Errorf("generated random number does not fit as int64: %v", randNum64)
		return 0, err
	}
	return randNum, nil
}

// CoinFlip performs a simple coinflip based on the rand library and returns true or false
func CoinFlip() bool {
	num := big.NewInt(2)
	cf, _ := rand.Int(rand.Reader, num)
	r := int(cf.Int64())
	return r == 1
}
