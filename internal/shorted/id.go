package shorted

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
)

// ID defines how we can find a shortened urlShorten in our system
// todo: look for something smaller and still unique to shorted to
type ID uuid.UUID

// NewIDFromAPI creates and parses an ID from the API endpoint
func NewIDFromAPI(val string) (i ID, err error) {
	err = i.decodeID(val)
	return
}

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

//encodeID encodes the ShortURL.ShortID as base58
func (i ID) encodeID() string {
	b, _ := uuid.UUID(i).MarshalBinary()
	encoded := base58.Encode(b)
	return encoded
}

//decodeID decodes the ShortURL.ShortID from base58
func (i *ID) decodeID(b58 string) (err error) {
	var uID uuid.UUID
	err = uID.UnmarshalBinary(base58.Decode(b58))
	*i = ID(uID)
	return
}
