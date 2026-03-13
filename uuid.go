package helper

import (
	"fmt"
	"crypto/rand"
)


// Returns a 32 bit random key based on the input key.
func UUID16(key string) (string, error) {
	keylen := len(key)
    if keylen < 1 {
        return "", fmt.Errorf("Error: Empty key.")
    }

	if keylen > 16 {
		keylen = 16
	}

	keybuff:= make([]byte, keylen)
    keybuff = []byte(key)
    if _, err := rand.Read(keybuff); err != nil {
		return "", fmt.Errorf("Error: Error while reading key. Actual error: %s", err.Error())
    }

	uuidbuff := make([]byte, uuid_len_32)
	if _, err := rand.Read(uuidbuff); err != nil {
		return "", fmt.Errorf("Error: Error while reading uuid. Actual error: %s", err.Error())
    }

	for i, j := 0, RandomInt(0, uuid_len_32 - keylen); i < keylen; i++ {
		uuidbuff[j] = keybuff[i]
	}

	uuid := fmt.Sprintf("%x%x%x%x%x", uuidbuff[0:6], uuidbuff[6:8], uuidbuff[8:10], uuidbuff[10:12], uuidbuff[12:])

	return uuid, nil
}


func UUID32(key string) (string, error) {
	return "", nil
}


func UUID64(key string) (string, error) {
	return "", nil
}
