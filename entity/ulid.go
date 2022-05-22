package entity

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func GenerateUlid() ulid.ULID {
	t := time.Now()

	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)

	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	return id
}
