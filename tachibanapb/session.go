package tachibanapb

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func (x *LoginRequest) GetKey(today time.Time) string {
	b := sha256.Sum256([]byte(today.Format("2006-01-02") + ":" + x.UserId + ":" + x.Password))
	return hex.EncodeToString(b[:])
}
