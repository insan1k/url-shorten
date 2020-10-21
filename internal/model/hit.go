package model

import (
	"errors"
	"fmt"
	"github.com/insan1k/one-qr-dot-me/internal/database"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/spf13/cast"
	"time"
)

type Hit struct {
	ID        string        `json:"id"`
	ShortID   string        `json:"short_id"`
	From      string        `json:"from"`
	To        string        `json:"to"`
	Address   string        `json:"address"`
	WasCached bool          `json:"was_cached"`
	Took      time.Duration `json:"took"`
	Timestamp time.Time     `json:"timestamp"`
}

func (h Hit) PersistDB() (err error) {
	session, err := database.Driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return
	}
	defer func() {
		err = session.Close()
	}()
	data := map[string]interface{}{
		"id":         h.ID,
		"short_id":   h.ShortID,
		"from":       h.From,
		"to":         h.From,
		"address":    h.From,
		"was_cached": h.WasCached,
		"took":       h.Took,
		"timestamp":  h.Timestamp.Format(time.RFC3339Nano),
	}
	result, err := session.Run(""+
		"CREATE (n:Hit {"+
		"id: $id, "+
		"short_id: $short_id, "+
		"from: $from, "+
		"to: $to, "+
		"address: $address, "+
		"was_cached: $was_cached, "+
		"took: $took, "+
		"timestamp: timestamp})", data)
	if err != nil {
		return
	}
	if _, err = result.Consume(); err != nil {
		return err
	}
	return
}

func GetHitFromDB(id string) (h Hit, err error) {
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
	result, err := session.Run("MATCH (n:Hit) WHERE id(n) = $id", data)
	if err != nil {
		return
	}
	if result.Next() {
		err = h.parseFromDB(result.Record())
		if err != nil {
			return
		}
	} else {
		err = result.Err()
		return
	}
	return
}

func NewHitFromDB(record neo4j.Record) (h Hit, err error) {
	err = h.parseFromDB(record)
	return
}

func (h *Hit) parseFromDB(record neo4j.Record) (err error) {
	jID, ok := record.Get("id")
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jShortID, ok := record.Get("short_id")
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jFrom, ok := record.Get("from")
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jTo, ok := record.Get("to")
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jAddress, ok := record.Get("id")
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jWasCached, ok := record.Get("was_cached")
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jTook, ok := record.Get("took")
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jTimestamp, ok := record.Get("timestamp")
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	h.ID, err = cast.ToStringE(jID)
	if err != nil {
		err = fmt.Errorf(couldNotParseError, err)
		return
	}
	h.ShortID, err = cast.ToStringE(jShortID)
	if err != nil {
		err = fmt.Errorf(couldNotParseError, err)
		return
	}
	h.From, err = cast.ToStringE(jFrom)
	if err != nil {
		err = fmt.Errorf(couldNotParseError, err)
		return
	}
	h.To, err = cast.ToStringE(jTo)
	if err != nil {
		err = fmt.Errorf(couldNotParseError, err)
		return
	}
	h.Address, err = cast.ToStringE(jAddress)
	if err != nil {
		err = fmt.Errorf(couldNotParseError, err)
		return
	}
	h.WasCached, err = cast.ToBoolE(jWasCached)
	if err != nil {
		err = fmt.Errorf(couldNotParseError, err)
		return
	}
	h.Took, err = cast.ToDurationE(jTook)
	if err != nil {
		err = fmt.Errorf(couldNotParseError, err)
		return
	}
	h.Timestamp, err = cast.ToTimeE(jTimestamp)
	if err != nil {
		err = fmt.Errorf(couldNotParseError, err)
		return
	}
	return
}
