package model

import (
	"errors"
	"fmt"
	"github.com/insan1k/one-qr-dot-me/internal/database"
	"github.com/insan1k/one-qr-dot-me/internal/logger"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/spf13/cast"
	"time"
)

// Hit represents the model of a redirection event
type Hit struct {
	HitID     string        `json:"hit_id"`
	ShortID   string        `json:"short_id"`
	From      string        `json:"from"`
	To        string        `json:"to"`
	Address   string        `json:"address"`
	WasCached bool          `json:"was_cached"`
	Took      time.Duration `json:"took"`
	Timestamp time.Time     `json:"timestamp"`
}

// Persist writes a single Hit to the database
func (h Hit) Persist() (err error) {
	err = h.perisistDB()
	if err != nil {
		return
	}
	err = h.persistRelationshipDB()
	return
}

func (h Hit) perisistDB() (err error) {
	session, err := database.Driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return
	}
	defer func() {
		if deferErr := session.Close(); deferErr != deferErr {
			logger.L.Errorf("session close error %v", deferErr)
		}
	}()
	data := map[string]interface{}{
		"hit_id":     h.HitID,
		"short_id":   h.ShortID,
		"from":       h.From,
		"to":         h.To,
		"address":    h.Address,
		"was_cached": h.WasCached,
		"took":       h.Took,
		"timestamp":  h.Timestamp.Format(time.RFC3339Nano),
	}
	result, err := session.Run(""+
		"CREATE (n:Hit {"+
		"hit_id: $hit_id, "+
		"short_id: $short_id, "+
		"from: $from, "+
		"to: $to, "+
		"address: $address, "+
		"was_cached: $was_cached, "+
		"took: $took, "+
		"timestamp: $timestamp "+
		"})", data)
	if err != nil {
		return
	}
	if _, err = result.Consume(); err != nil {
		return err
	}
	return
}

// persistRelationshipDB persists the relationship of model.ShortURL and model.Hit
func (h Hit) persistRelationshipDB() (err error) {
	session, err := database.Driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return
	}
	defer func() {
		if deferErr := session.Close(); deferErr != deferErr {
			logger.L.Errorf("session close error %v", deferErr)
		}
	}()
	data := map[string]interface{}{
		"hit_id":   h.HitID,
		"short_id": h.ShortID,
	}
	result, err := session.Run(""+
		"MATCH (a:ShortURL),(b:Hit) "+
		"WHERE a.short_id = $short_id AND b.short_id = $short_id AND b.hit_id = $hit_id "+
		"CREATE (a)-[r:RELTYPE]->(b) "+
		"RETURN type(r)", data)
	if err != nil {
		return
	}
	if _, err = result.Consume(); err != nil {
		return err
	}
	return
}

// GetHitFromDB gets a single Hit from the database
func GetHitFromDB(id string) (h Hit, err error) {
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
		"id": id,
	}
	result, err := session.Run("MATCH (n:Hit) WHERE n.id = $id RETURN n", data)
	if err != nil {
		return
	}
	if result.Next() {
		err = h.parseFromDB(result.Record().GetByIndex(0).(neo4j.Node))
		if err != nil {
			return
		}
	} else {
		err = result.Err()
		return
	}
	return
}

// NewHitFromDB creates a Hit from a database record
func NewHitFromDB(n neo4j.Node) (h Hit, err error) {
	err = h.parseFromDB(n)
	return
}

//todo: create a JSON parser from neo4j.Node the complexity here is just unacceptable
func (h *Hit) parseFromDB(n neo4j.Node) (err error) {
	jID, ok := n.Props()["hit_id"]
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jShortID, ok := n.Props()["short_id"]
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jFrom, ok := n.Props()["from"]
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jTo, ok := n.Props()["to"]
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jAddress, ok := n.Props()["address"]
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jWasCached, ok := n.Props()["was_cached"]
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jTook, ok := n.Props()["took"]
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	jTimestamp, ok := n.Props()["timestamp"]
	if !ok {
		err = errors.New(notFoundError)
		return
	}
	h.HitID, err = cast.ToStringE(jID)
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
