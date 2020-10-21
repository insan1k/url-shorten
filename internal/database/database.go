package database

import (
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Driver is the singleton of our neo4j driver
var Driver neo4j.Driver

// Load loads database connection into the singleton
func Load() (err error) {
	jConf := func(conf *neo4j.Config) {
		conf.Encrypted = configuration.C.Neo4JSecure
	}
	auth := neo4j.BasicAuth(configuration.C.Neo4JUser, configuration.C.Neo4JPassword, configuration.C.Neo4JRealm)

	Driver, err = neo4j.NewDriver(configuration.C.Neo4JTarget, auth, jConf)
	return
}

// Stop our database connection
func Stop() (err error) {
	return Driver.Close()
}

//todo: attempt to create a parser for structs Neo4j driver is a mess
