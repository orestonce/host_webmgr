package ymdUuid

import (
	"github.com/satori/go.uuid"
)

func MustNewUUID() string {
	u1 := uuid.Must(uuid.NewV4())
	return u1.String()
}
