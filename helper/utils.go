package helper

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

func SubString(Text string, SubText string) string {
	return strings.TrimSpace(Text)[strings.Index(strings.TrimSpace(Text), SubText)+1 : len([]rune(strings.TrimSpace(Text)))]
}

func HashSha1(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	return string(h.Sum(nil))
}

func HashSha256(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func StringConvertToFloat64(text string) (*float64, error) {
	var number *float64
	res, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return nil, err
	}
	number = &res
	return number, nil
}
