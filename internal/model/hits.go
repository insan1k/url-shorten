package model

import (
	"github.com/insan1k/one-qr-dot-me/internal/database"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"time"
)

type Hits struct {
	Count int   `json:"count"`
	Hits  []Hit `json:"hits"`
}

func HitsFromDb(shortID string, start, end time.Time) (h Hits, err error) {
	session, err := database.Driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return
	}
	defer func() {
		err = session.Close()
	}()
	data := map[string]interface{}{
		"short_id": shortID,
		"start":    start.Format(time.RFC3339Nano),
		"end":      end.Format(time.RFC3339Nano),
	}
	result, err := session.Run("MATCH (n:Hit) "+
		"where date(datetime(n:timestamp))>$start"+
		"or "+
		"date(datetime(n:timestamp))<$end"+
		"and short_id=$short_id`;", data)
	if err != nil {
		return
	}
	for result.Next() {
		var hit Hit
		hit, err = NewHitFromDB(result.Record())
		if err != nil {
			return
		}
		h.Hits = append(h.Hits, hit)
	}
	if _, err = result.Consume(); err != nil {
		return
	}
	h.Count = len(h.Hits)
	return
}
