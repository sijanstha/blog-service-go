package stringutils

import (
	"log"
	"strconv"

	"github.com/sony/sonyflake"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func IsEmptyOrNull(s string) bool {
	return s == "" || len(s) == 0
}

func GenerateUniqueId() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	return strconv.FormatUint(id, 10)
}

func ParseInteger(str string) int64 {
	data, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return data
}
