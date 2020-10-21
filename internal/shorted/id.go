package shorted

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
)

// ID defines how we can find a shortened urlShorten in our system
// todo: look for something smaller and still unique to shorted to
type ID uuid.UUID

// NewID creates and return a new v1 UUID
func NewID() (i ID, err error) {
	if id, err := uuid.NewUUID(); err == nil {
		i = ID(id)
	}
	return
}

//string returns the string form of the UUID
func (i ID) string() string {
	return uuid.UUID(i).String()
}

//encodeID encodes the ShortURL.ID as base58
func (i ID) encodeID() string {
	b, _ := uuid.UUID(i).MarshalBinary()
	return base58.Encode(b)
}

//decodeID decodes the ShortURL.ID from base58
func (i *ID) decodeID(b58 string) (err error) {
	decoded, err := uuid.ParseBytes(base58.Decode(b58))
	if err != nil {
		return err
	}
	*i = ID(decoded)
	return
}
