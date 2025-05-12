package safaricom

import (
	"encoding/base64"
	"time"
)

const passKey = "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
const shortCode = "174379"

func GenerateTimestamp() string {
	return time.Now().Format("20060102150405")
}

func GeneratePassword(timestamp string) string {
	raw := shortCode + passKey + timestamp
	return base64.StdEncoding.EncodeToString([]byte(raw))
}
