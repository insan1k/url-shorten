package database

import (
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var Driver neo4j.Driver

func LoadDB() (err error) {
	jConf := func(conf *neo4j.Config) {
		conf.Encrypted = configuration.C.Neo4JSecure
	}
	auth := neo4j.BasicAuth(configuration.C.Neo4JUser, configuration.C.Neo4JPassword, configuration.C.Neo4JRealm)

	Driver, err = neo4j.NewDriver(configuration.C.Neo4JPassword, auth, jConf)
	return
}

func StopDB() (err error) {
	Driver.Close()
	return
}

//todo: attempt to create a parser for structs Neo4j driver is a mess
