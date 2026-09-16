package cluster

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func Sign(data string) string {

	secret :=
		os.Getenv("KINGDNS_CLUSTER_TOKEN")

	mac := hmac.New(
		sha256.New,
		[]byte(secret),
	)

	mac.Write(
		[]byte(data),
	)

	return hex.EncodeToString(
		mac.Sum(nil),
	)
}

func Verify(
	data string,
	signature string,
) bool {

	expected := Sign(data)

	return hmac.Equal(
		[]byte(expected),
		[]byte(signature),
	)
}
