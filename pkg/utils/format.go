package utils

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"strings"
)

func StructToJsonString(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func IsEthAddr(addr string) bool {
	addr = strings.ToLower(addr)
	if !strings.HasPrefix(addr, "0x") || len(addr) != 42 {
		return false
	}
	_, err := hex.DecodeString(strings.TrimLeft(addr, "0x"))
	return err == nil
}

func FormatEthAddr(addr string) (string, error) {
	if IsEthAddr(addr) {
		return strings.ToLower(addr), nil
	}
	if len(addr) == 40 && IsEthAddr("0x"+addr) {
		return strings.ToLower("0x" + addr), nil
	}
	return "", errors.New("invalid eth address")
}

func UnsafeFormatEthAddr(addr string) string {
	addr, _ = FormatEthAddr(addr)
	return addr
}
