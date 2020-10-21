package model

import (
	"errors"
	"fmt"
	"github.com/insan1k/one-qr-dot-me/internal/cache"
	"github.com/insan1k/one-qr-dot-me/internal/database"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shamaton/msgpack"
	"github.com/spf13/cast"
	"time"
)

const (
	notFoundError      = "parse from db failed, required field not found"
	couldNotParseError = "parse from db failed invalid data %v "
)

// ShortURL is the model struct for ShortURL
type ShortURL struct {
	ID        string    `json:"id" msgpack:"id"`
	Original  string    `json:"original" msgpack:"original"`
	Short     string    `json:"short" msgpack:"short"`
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp"`
}

// PersistCache saves a ShortURL to the cache
func (s ShortURL) PersistCache() (err error) {
	packed, err := s.toMsgPack()
	if err != nil {
		return
	}
	err = cache.C.Set(s.ID, packed)
	return
}

// PersistDB saves a ShortURL to the database
func (s ShortURL) PersistDB() (err error) {
	session, err := database.Driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return
	}
	defer func() {
		err = session.Close()
	}()
	data := map[string]interface{}{
		"id":        s.ID,
		"original":  s.Original,
		"short":     s.Short,
		"timestamp": s.Timestamp.Format(time.RFC3339Nano),
	}
	result, err := session.Run("CREATE (n:ShortURL {id: $id, original: $original, short: $short, timestamp: $time})", data)
	if err != nil {
		return
	}
	if _, err = result.Consume(); err != nil {
		return err
	}
	return
}

// FindShortURL retrieves looks for and retrieves a ShortURL from cache or database
func FindShortURL(id string) (s ShortURL, cached bool, err error) {
	s, err = shortURLFromCache(id)
	if err == nil {
		cached = true
		return
	}
	s, err = shortURLFromDB(id)
	return
}

//shortURLFromCache retrieves a ShortURL from the cache
func shortURLFromCache(id string) (s ShortURL, err error) {
	packed, err := cache.C.Get(id)
	if err != nil {
		return
	}
	err = s.fromMsgPack(packed)
	return
}

//shortURLFromDB retrieves a ShortURL from the database
func shortURLFromDB(id string) (s ShortURL, err error) {
	session, err := database.Driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return
	}
	defer func() {
		err = session.Close()
	}()
	data := map[string]interface{}{
		"id": id,
	}
	result, err := session.Run("MATCH (n:ShortURL) WHERE id(n) = $id", data)
	if err != nil {
		return
	}
	if result.Next() {
		err = s.parseFromDB(result.Record())
		if err != nil {
			return
		}
	} else {
		err = result.Err()
		return
	}
	return
}

func (s *ShortURL) parseFromDB(record neo4j.Record) (err error) {
	jID, ok := record.Get("id")
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jOriginal, ok := record.Get("original")
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jShort, ok := record.Get("short")
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jTimestamp, ok := record.Get("timestamp")
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	s.ID, err = cast.ToStringE(jID)
	if err != nil {
		err = fmt.Errorf(couldNotParseError, err)
		return
	}
	s.ID, err = cast.ToStringE(jShort)
	if err != nil {
		err = fmt.Errorf(couldNotParseError, err)
		return
	}
	s.ID, err = cast.ToStringE(jOriginal)
	if err != nil {
		err = fmt.Errorf(couldNotParseError, err)
		return
	}
	s.ID, err = cast.ToStringE(jTimestamp)
	if err != nil {
		err = fmt.Errorf(couldNotParseError, err)
		return
	}
	return
}

func (s ShortURL) toMsgPack() (packed []byte, err error) {
	packed, err = msgpack.Encode(s)
	return
}

func (s *ShortURL) fromMsgPack(packed []byte) (err error) {
	err = msgpack.Decode(packed, &s)
	return
}
