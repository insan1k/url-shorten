package model

import (
	"github.com/insan1k/one-qr-dot-me/internal/database"
	"github.com/insan1k/one-qr-dot-me/internal/logger"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"time"
)

// Hits represents a aggregate of hits
type Hits struct {
	Count int   `json:"count"`
	Hits  []Hit `json:"hits"`
}

// HitsFromDb retrieves a group of hits from the database
func HitsFromDb(shortID string, start, end time.Time) (h Hits, err error) {
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
		"short_id": shortID,
		"start":    start.Format(time.RFC3339Nano),
		"end":      end.Format(time.RFC3339Nano),
	}
	result, err := session.Run("MATCH (n:ShortURL)--(m:Hit) "+
		"WHERE n.short_id = $short_id "+
		"AND ( datetime(m.timestamp) > datetime($start) "+
		"OR "+
		"datetime(m.timestamp) < datetime($end) ) "+
		"RETURN m", data)
	if err != nil {
		return
	}
	for result.Next() {
		var hit Hit
		hit, err = NewHitFromDB(result.Record().GetByIndex(0).(neo4j.Node))
		if err != nil {
			return
		}
		h.Hits = append(h.Hits, hit)
	}
	h.Count = len(h.Hits)
	return
}
