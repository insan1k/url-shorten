package configuration

import (
	"github.com/spf13/cast"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

const (
	//URLShortenerPath defines in which endpoint we will perform redirects
	URLShortenerPath = "/"
)

// C is our configuration singleton
var C Configuration

// Configuration holds all the configuration parameters for our program
type Configuration struct {
	// shorten service configuration
	ShortenScheme string
	ShortenPort   string
	ShortenDomain string
	// HTTP arguments
	HTTPTLSCertPath  string
	HTTPTLSKeyPath   string
	HTTPHostname     string
	HTTPBindProtocol string
	HTTPReadTimeout  time.Duration
	HTTPWriteTimeout time.Duration
	HTTPIdleTimeout  time.Duration
	// CacheConfiguration arguments
	CacheShards             int
	CacheLifeWindow         time.Duration
	CacheCleanWindow        time.Duration
	CacheMaxEntriesInWindow int
	CacheMaxEntrySize       int
	CacheHardMaxSize        int
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

func envToCamelCase(val string) string {
	ss := strings.Split(val, "_")
	ss[0] = strings.ToLower(ss[0])
	if len(ss) == 1 {
		return ss[0]
	}
	var camelcase string
	for i, s := range ss {
		if i > 0 {
			firstLetter := s[0:1]
			firstLetter = strings.ToUpper(firstLetter)
			s = firstLetter + strings.ToLower(s[1:])
		}
		camelcase += s
	}
	return camelcase
}

func envToFlag(val string) string {
	return envToCamelCase(val)
}

const (
	// config file stuff
	fileType      = "yaml"
	fileExtension = "yml"
	fileName      = "config"
	// don't forget to add new configurations here
	envPrefix                  = "S"
	configFile                 = "CONFIG_FILE"
	envShortenScheme           = "SHORTEN_SCHEME"
	envShortenPort             = "SHORTEN_PORT"
	envShortenDomain           = "SHORTEN_DOMAIN"
	envHTTPTLSCertPath         = "HTTP_TLS_CERT_PATH"
	envHTTPTLSKeyPath          = "HTTP_TLS_KEY_PATH"
	envHTTPHostname            = "HTTP_HOSTNAME"
	envHTTPProtocol            = "HTTP_BIND_PROTOCOL"
	envHTTPReadTimeout         = "HTTP_READ_TIMEOUT"
	envHTTPWriteTimeout        = "HTTP_WRITE_TIMEOUT"
	envHTTPIdleTimeout         = "HTTP_IDLE_TIMEOUT"
	envCacheShards             = "CACHE_SHARDS"
	envCacheLifeWindow         = "CACHE_LIFE_WINDOW"
	envCacheCleanWindow        = "CACHE_CLEAN_WINDOW"
	envCacheMaxEntriesInWindow = "CACHE_ENTRIES_IN_WINDOW"
	envCacheMaxEntrySize       = "CACHE_MAX_ENTRY_SIZE"
	envCacheHardMaxSize        = "CACHE_HARD_MAX_SIZE"
	envNeo4JSecure             = "NEO4J_SECURE"
	envNeo4JTarget             = "NEO4J_TARGET"
	envNeo4JUser               = "NEO4J_USER"
	envNeo4JPassword           = "NEO4J_PASSWORD"
	envNeo4JRealm              = "NEO4J_REALM"
)

var configurationMap = map[string]config{
	// don't forget to add new configurations here
	configFile: {
		Flag:     envToFlag(configFile),
		Env:      configFile,
		File:     "",
		Default:  "./cmd/url-shorten/" + fileName + "." + fileExtension,
		Usage:    "set the configuration file location e.g.:\" ./config.yml\"",
		castType: castString,
	},
	envShortenScheme: {
		Flag:     "",
		Env:      envShortenScheme,
		File:     envToCamelCase(envShortenScheme),
		Default:  "http",
		Usage:    "set either http or http as the shorten url service external protocol e.g.:\"https\" defaults to http",
		castType: castString,
	},
	envShortenPort: {
		Flag:     "",
		Env:      envShortenPort,
		File:     envToCamelCase(envShortenPort),
		Default:  "",
		Usage:    "set the url service external port e.g.:\":8081\" defaults to empty",
		castType: castString,
	},
	envShortenDomain: {
		Flag:     "",
		Env:      envShortenDomain,
		File:     envToCamelCase(envShortenDomain),
		Default:  "localhost",
		Usage:    "set the url service external domain e.g.:\"example.com\" defaults to localhost",
		castType: castString,
	},
	envHTTPTLSCertPath: {
		Flag:     envToFlag(envHTTPTLSCertPath),
		Env:      envHTTPTLSCertPath,
		File:     envToCamelCase(envHTTPTLSCertPath),
		Default:  "",
		Usage:    "set the http server tls cert path e.g.:\"\" defaults to empty",
		castType: castString,
	},
	envHTTPTLSKeyPath: {
		Flag:     envToFlag(envHTTPTLSKeyPath),
		Env:      envHTTPTLSKeyPath,
		File:     envToCamelCase(envHTTPTLSKeyPath),
		Default:  "",
		Usage:    "set the http server tls key path  e.g.:\"\" defaults to empty",
		castType: castString,
	},
	envHTTPHostname: {
		Flag:     envToFlag(envHTTPHostname),
		Env:      envHTTPHostname,
		File:     envToCamelCase(envHTTPHostname),
		Default:  "localhost:8080",
		Usage:    "set the http server bind hostname e.g.:\"localhost:8081\" defaults to localhost:8080",
		castType: castString,
	},
	envHTTPProtocol: {
		Flag:     envToFlag(envHTTPProtocol),
		Env:      envHTTPProtocol,
		File:     envToCamelCase(envHTTPProtocol),
		Default:  "tcp",
		Usage:    "set the http server bind protocol e.g.:\"tcp4\" defaults to tcp, supported protocols are tcp4, tcp6 and tcp",
		castType: castString,
	},
	envHTTPReadTimeout: {
		Flag:     "",
		Env:      envHTTPReadTimeout,
		File:     envToCamelCase(envHTTPReadTimeout),
		Default:  "1s",
		Usage:    "set the http server read timeout e.g.:\"5s\" defaults to 1 second",
		castType: castDuration,
	},
	envHTTPWriteTimeout: {
		Flag:     "",
		Env:      envHTTPWriteTimeout,
		File:     envToCamelCase(envHTTPWriteTimeout),
		Default:  "1s",
		Usage:    "set the http server write timeout e.g.:\"5s\" defaults to 1 second",
		castType: castDuration,
	},
	envHTTPIdleTimeout: {
		Flag:     "",
		Env:      envHTTPIdleTimeout,
		File:     envToCamelCase(envHTTPIdleTimeout),
		Default:  "1s",
		Usage:    "set the http server idle timeout e.g.:\"5s\" defaults to 1 second",
		castType: castDuration,
	},
	envCacheShards: {
		Flag:     "",
		Env:      envCacheShards,
		File:     envToCamelCase(envCacheShards),
		Default:  "2048",
		Usage:    "number of cache shards, value must be a power of two e.g.:\"4096\" defaults to 2048",
		castType: castInt,
	},
	envCacheLifeWindow: {
		Flag:     "",
		Env:      envCacheLifeWindow,
		File:     envToCamelCase(envCacheLifeWindow),
		Default:  "30s",
		Usage:    "time after which a cache entry can be evicted e.g.:\"30s\" defaults to 30s",
		castType: castDuration,
	},
	envCacheCleanWindow: {
		Flag:     "",
		Env:      envCacheCleanWindow,
		File:     envToCamelCase(envCacheCleanWindow),
		Default:  "60s",
		Usage:    "interval between removing expired cache entries e.g.:\"120s\" defaults to 60s",
		castType: castDuration,
	},
	envCacheMaxEntriesInWindow: {
		Flag:     "",
		Env:      envCacheMaxEntriesInWindow,
		File:     envToCamelCase(envCacheMaxEntriesInWindow),
		Default:  "0",
		Usage:    "max number of cache entries in life window. this is used only to calculate initial size for cache shards. e.g.:\"128\" defaults to 0",
		castType: castInt,
	},
	envCacheMaxEntrySize: {
		Flag:     "",
		Env:      envCacheMaxEntrySize,
		File:     envToCamelCase(envCacheMaxEntrySize),
		Default:  "32",
		Usage:    "max size size of entry in bytes. this is used only to calculate initial size for cache shards. e.g.:\"128\" defaults to 32",
		castType: castInt,
	},
	envCacheHardMaxSize: {
		Flag:     "",
		Env:      envCacheHardMaxSize,
		File:     envToCamelCase(envCacheHardMaxSize),
		Default:  "128",
		Usage:    "hard max cache size in megabytes. e.g.:\"512\" defaults to 128",
		castType: castInt,
	},
	envNeo4JSecure: {
		Flag:     envToFlag(envNeo4JSecure),
		Env:      envNeo4JSecure,
		File:     envToCamelCase(envNeo4JSecure),
		Default:  "false",
		Usage:    "sets whether to turn on/off TLS encryption for neo4j. e.g.:\"true\" defaults to false",
		castType: castBool,
	},
	envNeo4JTarget: {
		Flag:     envToFlag(envNeo4JTarget),
		Env:      envNeo4JTarget,
		File:     envToCamelCase(envNeo4JTarget),
		Default:  "bolt://localhost:7687",
		Usage:    "sets the neo4j url. e.g.:\"bolt://db.server:7687\" defaults to bolt://localhost:7687",
		castType: castString,
	},
	envNeo4JUser: {
		Flag:     envToFlag(envNeo4JUser),
		Env:      envNeo4JUser,
		File:     envToCamelCase(envNeo4JUser),
		Default:  "",
		Usage:    "sets the neo4j user. e.g.:\"my-application-user\" defaults to neo4j-admin",
		castType: castString,
	},
	envNeo4JPassword: {
		Flag:     envToFlag(envNeo4JPassword),
		Env:      envNeo4JPassword,
		File:     envToCamelCase(envNeo4JPassword),
		Default:  "",
		Usage:    "sets the neo4j password. e.g.:\"my-application-secure-password\" defaults to secret",
		castType: castString,
	},
	envNeo4JRealm: {
		Flag:     envToFlag(envNeo4JRealm),
		Env:      envNeo4JRealm,
		File:     envToCamelCase(envNeo4JRealm),
		Default:  "",
		Usage:    "sets the neo4j realm. e.g.:\"my-neo4j-realm\" defaults to Neo4J",
		castType: castString,
	},
}

func loadConf(c *Configuration) (got Configuration) {
	// don't forget to add new configurations here
	return Configuration{
		CacheShards:             cast.ToInt(getConf(c, envCacheShards)),
		ShortenScheme:           cast.ToString(getConf(c, envShortenScheme)),
		ShortenPort:             cast.ToString(getConf(c, envShortenPort)),
		ShortenDomain:           cast.ToString(getConf(c, envShortenDomain)),
		HTTPTLSCertPath:         cast.ToString(getConf(c, envHTTPTLSCertPath)),
		HTTPTLSKeyPath:          cast.ToString(getConf(c, envHTTPTLSKeyPath)),
		HTTPHostname:            cast.ToString(getConf(c, envHTTPHostname)),
		HTTPBindProtocol:        cast.ToString(getConf(c, envHTTPProtocol)),
		HTTPReadTimeout:         cast.ToDuration(getConf(c, envHTTPReadTimeout)),
		HTTPWriteTimeout:        cast.ToDuration(getConf(c, envHTTPWriteTimeout)),
		HTTPIdleTimeout:         cast.ToDuration(getConf(c, envHTTPIdleTimeout)),
		CacheLifeWindow:         cast.ToDuration(getConf(c, envCacheLifeWindow)),
		CacheCleanWindow:        cast.ToDuration(getConf(c, envCacheCleanWindow)),
		CacheMaxEntriesInWindow: cast.ToInt(getConf(c, envCacheMaxEntriesInWindow)),
		CacheMaxEntrySize:       cast.ToInt(getConf(c, envCacheMaxEntrySize)),
		CacheHardMaxSize:        cast.ToInt(getConf(c, envCacheHardMaxSize)),
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
	c.pFlag = pflag.NewFlagSet(os.Args[0], pflag.PanicOnError)
	c.viperEnvAndFile = viper.New()
	c.viperFlag = viper.New()
	registerFlags(c)
	parseFlags(c)
	registerEnvs(c)
	parseFile(c)
	C = loadConf(c)
}
