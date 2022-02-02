package mongodb

import (
	"encoding/binary"
	"sync/atomic"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ts              = uint32(time.Date(2021, 9, 1, 0, 0, 0, 0, time.UTC).Unix())
	objectIDCounter uint32
)

// NewObjectID generates stable BSON ObjectID to make conversion results stable.
func NewObjectID(id, c uint32) primitive.ObjectID {
	processUnique := make([]byte, 5)
	binary.BigEndian.PutUint32(processUnique, id)

	var b [12]byte

	binary.BigEndian.PutUint32(b[0:4], ts)
	copy(b[4:9], processUnique)

	if c == 0 {
		c = atomic.AddUint32(&objectIDCounter, 1)
	}
	b[9] = byte(c >> 16)
	b[10] = byte(c >> 8)
	b[11] = byte(c)

	return b
}
