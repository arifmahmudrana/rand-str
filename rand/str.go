package rand

import (
	"bytes"
	"crypto/rand"
)

func randomBytes(siz int) ([]byte, error) {
	b := make([]byte, siz)
	_, err := rand.Read(b)

	return b, err
}

func generateRandomString(siz int, str string, b []byte) string {
	charsLength := len(str)
	cursor := 0
	var buffer bytes.Buffer
	for i := 0; i < siz; i++ {
		cursor += int(b[i])
		buffer.WriteByte(str[cursor%charsLength])
	}

	return buffer.String()
}

func RandomString(siz int, str string) (string, error) {
	b, err := randomBytes(siz)
	if err != nil {
		return "", err
	}

	return generateRandomString(siz, str, b), nil
}
