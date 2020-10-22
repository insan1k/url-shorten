package model

import (
	"errors"
	"fmt"
	"github.com/insan1k/one-qr-dot-me/internal/cache"
	"github.com/insan1k/one-qr-dot-me/internal/database"
	"github.com/insan1k/one-qr-dot-me/internal/logger"
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
	ShortID   string    `json:"short_id" msgpack:"short_id"`
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
	err = cache.C.Set(s.ShortID, packed)
	return
}

// PersistDB saves a ShortURL to the database
func (s ShortURL) PersistDB() (err error) {
	session, err := database.Driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return
	}
	defer func() {
		if deferErr := session.Close(); deferErr != nil {
			logger.L.Errorf("session close error %v", deferErr)
		}
	}()
	data := map[string]interface{}{
		"short_id":  s.ShortID,
		"original":  s.Original,
		"short":     s.Short,
		"timestamp": s.Timestamp.Format(time.RFC3339Nano),
	}
	result, err := session.Run("CREATE (n:ShortURL {short_id: $short_id, original: $original, short: $short, timestamp: $timestamp})", data)
	if err != nil {
		return
	}
	if result.Err() != nil {
		err = result.Err()
		return
	}
	if _, err = result.Consume(); err != nil {
		return err
	}
	return
}

// FindShortURL retrieves looks for and retrieves a ShortURL from cache or database
func FindShortURL(id string) (s ShortURL, cached bool, err error) {
	if s, err = shortURLFromCache(id); err == nil {
		cached = true
		return
	}
	if s, err = shortURLFromDB(id); err != nil {
		return
	}
	// we're going to suppress this error and log it here
	if errCache := s.PersistCache(); errCache != nil {
		logger.L.Errorf("error persisting the short url from DB in Cache %v", errCache)
	}
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
		if deferErr := session.Close(); deferErr != nil {
			logger.L.Errorf("session close error %v", deferErr)
		}
	}()
	data := map[string]interface{}{
		"short_id": id,
	}
	result, err := session.Run("MATCH (n:ShortURL) WHERE n.short_id = $short_id RETURN n", data)
	if err != nil {
		return
	}
	if result.Err() != nil {
		err = result.Err()
		return
	}
	if result.Next() {
		err = s.parseFromDB(result.Record().GetByIndex(0).(neo4j.Node))
		if err != nil {
			return
		}
	} else {
		err = errors.New("node not found")
		return
	}
	return
}

func (s *ShortURL) parseFromDB(node neo4j.Node) (err error) {
	jID, ok := node.Props()["short_id"]
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jOriginal, ok := node.Props()["original"]
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jShort, ok := node.Props()["short"]
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jTimestamp, ok := node.Props()["timestamp"]
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	s.ShortID, err = cast.ToStringE(jID)
	if err != nil {
		err = fmt.Errorf(couldNotParseError, err)
		return
	}
	s.Short, err = cast.ToStringE(jShort)
	if err != nil {
		err = fmt.Errorf(couldNotParseError, err)
		return
	}
	s.Original, err = cast.ToStringE(jOriginal)
	if err != nil {
		err = fmt.Errorf(couldNotParseError, err)
		return
	}
	s.Timestamp, err = cast.ToTimeE(jTimestamp)
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
