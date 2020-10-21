package configuration

import (
	"github.com/spf13/cast"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"
)

var C Configuration

type Configuration struct {
	// shorten service configuration
	Scheme string
	Port   string
	Domain string
	// HTTP arguments
	ServerCertPath   string
	ServerKeyPath    string
	HttpHostname     string
	HttpProtocol     string
	HttpReadTimeout  time.Duration
	HttpWriteTimeout time.Duration
	HttpIdleTimeout  time.Duration
	// CacheConfiguration arguments
	CacheShards             int
	CacheLifeWindow         time.Duration
	CacheCleanWindow        time.Duration
	CacheMaxEntriesInWindow int
	CacheMaxEntrySize       int
	CacheHardMaxCacheSize   int
	// Neo4J arguments
	Neo4JSecure     bool
	Neo4JTarget     string
	Neo4JUser       string
	Neo4JPassword   string
	Neo4JRealm      string
	pFlag           *pflag.FlagSet
	viperEnvAndFile *viper.Viper
	viperFlag       *viper.Viper
}

type config struct {
	Flag     string
	Env      string
	File     string
	Default  string
	Usage    string
	castType int
}

const (
	// config file stuff
	fileType      = "yaml"
	fileExtension = "yml"
	fileName      = "conf"
	// don't forget to add new configurations here
	envPrefix                  = "SHORT_URL"
	configFile                 = "CONFIG_FILE"
	envShortenScheme           = "SHORTEN_SCHEME"
	envShortenPort             = "SHORTEN_PORT"
	envShortenDomain           = "SHORTEN_DOMAIN"
	envServerTLSCertPath       = "TLS_CERT_PATH"
	envServerTLSKeyPath        = "TLS_KEY_PATH"
	envHttpHostname            = "HTTP_HOSTNAME"
	envHttpProtocol            = "HTTP_BIND_PROTOCOL"
	envHttpReadTimeout         = "HTTP_READ_TIMEOUT"
	envHttpWriteTimeout        = "HTTP_WRITE_TIMEOUT"
	envHttpIdleTimeout         = "HTTP_IDLE_TIMEOUT"
	envCacheShards             = "CACHE_SHARDS"
	envCacheLifeWindow         = "CACHE_LIFE_WINDOW"
	envCacheCleanWindow        = "CACHE_CLEAN_WINDOW"
	envCacheMaxEntriesInWindow = "CACHE_ENTRIES_IN_WINDOW"
	envCacheMaxEntrySize       = "CACHE_MAX_ENTRY_SIZE"
	envCacheHardMaxCacheSize   = "CACHE_HARD_MAX_CACHE_SIZE"
	envNeo4JSecure             = "NEO4J_SECURE"
	envNeo4JTarget             = "NEO4J_TARGET"
	envNeo4JUser               = "NEO4J_USER"
	envNeo4JPassword           = "NEO4J_PASSWORD"
	envNeo4JRealm              = "NEO4J_REALM"
)

var configurationMap = map[string]config{
	// don't forget to add new configurations here
	configFile: {
		Flag:     "configFile",
		Env:      configFile,
		File:     "",
		Default:  "one-qr-dot-me/" + fileName + "." + fileExtension,
		Usage:    "set the configuration file location e.g.:\" ./config.yml\"",
		castType: castString,
	},
	envShortenScheme: {
		Flag:     "shortenScheme",
		Env:      "",
		File:     "ShortenScheme",
		Default:  "http",
		Usage:    "set either http or http as the shorten url service external protocol e.g.:\"https\" defaults to http",
		castType: castString,
	},
	envShortenPort: {
		Flag:     "shortenPort",
		Env:      "",
		File:     "ShortenPort",
		Default:  "",
		Usage:    "set the url service external port e.g.:\":8081\" defaults to empty",
		castType: castString,
	},
	envShortenDomain: {
		Flag:     "shortenDomain",
		Env:      "",
		File:     "ShortenDomain",
		Default:  "localhost",
		Usage:    "set the url service external domain e.g.:\"example.com\" defaults to localhost",
		castType: castString,
	},
	envServerTLSCertPath: {
		Flag:     "tlsCertPath",
		Env:      envServerTLSCertPath,
		File:     "ServerTLSCertPath",
		Default:  "",
		Usage:    "set the http server tls cert path e.g.:\"\" defaults to empty",
		castType: castString,
	},
	envServerTLSKeyPath: {
		Flag:     "tlsKeyPath",
		Env:      envServerTLSKeyPath,
		File:     "ServerTLSKeyPath",
		Default:  "",
		Usage:    "set the http server tls key path  e.g.:\"\" defaults to empty",
		castType: castString,
	},
	envHttpHostname: {
		Flag:     "httpHostname",
		Env:      envHttpHostname,
		File:     "HttpHostname",
		Default:  "localhost",
		Usage:    "set the http server bind hostname e.g.:\"localhost:8081\" defaults to localhost",
		castType: castString,
	},
	envHttpProtocol: {
		Flag:     "httpProtocol",
		Env:      envHttpProtocol,
		File:     "",
		Default:  "tcp",
		Usage:    "set the http server bind protocol e.g.:\"tcp4\" defaults to tcp, supported protocols are tcp4, tcp6 and tcp",
		castType: castString,
	},
	envHttpReadTimeout: {
		Flag:     "httpReadTimeout",
		Env:      "",
		File:     "HTTPReadTimeout",
		Default:  "1s",
		Usage:    "set the http server read timeout e.g.:\"5s\" defaults to 1 second",
		castType: castDuration,
	},
	envHttpWriteTimeout: {
		Flag:     "httpWriteTimeout",
		Env:      "",
		File:     "HTTPWriteTimeout",
		Default:  "1s",
		Usage:    "set the http server write timeout e.g.:\"5s\" defaults to 1 second",
		castType: castDuration,
	},
	envHttpIdleTimeout: {
		Flag:     "httpIdleTimeout",
		Env:      "",
		File:     "HTTPIdleTimeout",
		Default:  "1s",
		Usage:    "set the http server idle timeout e.g.:\"5s\" defaults to 1 second",
		castType: castDuration,
	},
	envCacheShards: {
		Flag:     "cShards",
		Env:      "",
		File:     "CacheShards",
		Default:  "2048",
		Usage:    "number of cache shards, value must be a power of two e.g.:\"4096\" defaults to 2048",
		castType: castInt,
	},
	envCacheLifeWindow: {
		Flag:     "cLifeWindow",
		Env:      "",
		File:     "CacheLifeWindow",
		Default:  "30s",
		Usage:    "time after which a cache entry can be evicted e.g.:\"30s\" defaults to 30s",
		castType: castDuration,
	},
	envCacheCleanWindow: {
		Flag:     "cCleanWindow",
		Env:      "",
		File:     "CacheCleanWindow",
		Default:  "60s",
		Usage:    "interval between removing expired cache entries e.g.:\"120s\" defaults to 60s",
		castType: castDuration,
	},
	envCacheMaxEntriesInWindow: {
		Flag:     "cMaxEntriesInWindow",
		Env:      "",
		File:     "CacheMaxEntriesInWindow",
		Default:  "0",
		Usage:    "max number of cache entries in life window. this is used only to calculate initial size for cache shards. e.g.:\"128\" defaults to 0",
		castType: castInt,
	},
	envCacheMaxEntrySize: {
		Flag:     "cMaxEntrySize",
		Env:      "",
		File:     "CacheMaxEntrySize",
		Default:  "32",
		Usage:    "max size size of entry in bytes. this is used only to calculate initial size for cache shards. e.g.:\"128\" defaults to 32",
		castType: castInt,
	},
	envCacheHardMaxCacheSize: {
		Flag:     "cHardMaxCacheSize",
		Env:      "",
		File:     "CacheHardMaxCacheSize",
		Default:  "128",
		Usage:    "hard max cache size in megabytes. e.g.:\"512\" defaults to 128",
		castType: castInt,
	},
	envNeo4JSecure: {
		Flag:     "neo4jSecure",
		Env:      envNeo4JSecure,
		File:     "Neo4JSecure",
		Default:  "false",
		Usage:    "sets whether to turn on/off TLS encryption for neo4j. e.g.:\"true\" defaults to false",
		castType: castBool,
	},
	envNeo4JTarget: {
		Flag:     "neo4jTarget",
		Env:      envNeo4JTarget,
		File:     "Neo4JTarget",
		Default:  "bolt://localhost:7687",
		Usage:    "sets the neo4j url. e.g.:\"bolt://db.server:7687\" defaults to bolt://localhost:7687",
		castType: castString,
	},
	envNeo4JUser: {
		Flag:     "neo4jUser",
		Env:      envNeo4JUser,
		File:     "Neo4JUser",
		Default:  "neo4j-admin",
		Usage:    "sets the neo4j user. e.g.:\"my-application-user\" defaults to neo4j-admin",
		castType: castString,
	},
	envNeo4JPassword: {
		Flag:     "neo4jPassword",
		Env:      envNeo4JPassword,
		File:     "Neo4JPassword",
		Default:  "secret",
		Usage:    "sets the neo4j password. e.g.:\"my-application-secure-password\" defaults to secret",
		castType: castString,
	},
	envNeo4JRealm: {
		Flag:     "Neo4JRealm",
		Env:      envNeo4JRealm,
		File:     "neo4jRealm",
		Default:  "Neo4JRealm",
		Usage:    "sets the neo4j realm. e.g.:\"my-neo4j-realm\" defaults to Neo4J",
		castType: castString,
	},
}

func loadConf(c *Configuration) (got Configuration) {
	// don't forget to add new configurations here
	return Configuration{
		CacheShards:             cast.ToInt(getConf(c, envCacheShards)),
		Scheme:                  cast.ToString(getConf(c, envShortenScheme)),
		Port:                    cast.ToString(getConf(c, envShortenPort)),
		Domain:                  cast.ToString(getConf(c, envShortenDomain)),
		ServerCertPath:          cast.ToString(getConf(c, envServerTLSCertPath)),
		ServerKeyPath:           cast.ToString(getConf(c, envServerTLSKeyPath)),
		HttpHostname:            cast.ToString(getConf(c, envHttpHostname)),
		HttpProtocol:            cast.ToString(getConf(c, envHttpProtocol)),
		HttpReadTimeout:         cast.ToDuration(getConf(c, envHttpReadTimeout)),
		HttpWriteTimeout:        cast.ToDuration(getConf(c, envHttpWriteTimeout)),
		HttpIdleTimeout:         cast.ToDuration(getConf(c, envHttpIdleTimeout)),
		CacheLifeWindow:         cast.ToDuration(getConf(c, envCacheLifeWindow)),
		CacheCleanWindow:        cast.ToDuration(getConf(c, envCacheCleanWindow)),
		CacheMaxEntriesInWindow: cast.ToInt(getConf(c, envCacheMaxEntriesInWindow)),
		CacheMaxEntrySize:       cast.ToInt(getConf(c, envCacheMaxEntrySize)),
		CacheHardMaxCacheSize:   cast.ToInt(getConf(c, envCacheHardMaxCacheSize)),
		Neo4JSecure:             cast.ToBool(getConf(c, envNeo4JSecure)),
		Neo4JTarget:             cast.ToString(getConf(c, envNeo4JTarget)),
		Neo4JUser:               cast.ToString(getConf(c, envNeo4JUser)),
		Neo4JPassword:           cast.ToString(getConf(c, envNeo4JPassword)),
		Neo4JRealm:              cast.ToString(getConf(c, envNeo4JRealm)),
	}
}

// Load settings into Configuration in the following order of precedence
// - flag
// - environment variables
// - file
func (c *Configuration) Load() {
	c.pFlag = pflag.NewFlagSet("url-shortener", pflag.PanicOnError)
	c.viperEnvAndFile = viper.New()
	c.viperFlag = viper.New()
	registerFlags(c)
	parseFlags(c)
	registerEnvs(c)
	parseFile(c)
	C = loadConf(c)
}
