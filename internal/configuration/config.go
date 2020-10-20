package configuration

import (
	"github.com/spf13/cast"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"
)

var C Configuration

type Configuration struct {
	Scheme string
	Port   string
	Domain string
	// CacheConfiguration arguments
	CacheShards             int
	CacheLifeWindow         time.Duration
	CacheCleanWindow        time.Duration
	CacheMaxEntriesInWindow int
	CacheMaxEntrySize       int
	CacheHardMaxCacheSize   int
	// Neo4J arguments
	Neo4JSecure   bool
	Neo4JTarget   string
	Neo4JUser     string
	Neo4JPassword string
	Neo4JRealm    string
	// HTTP arguments
	ServerCertPath   string
	ServerKeyPath    string
	HttpHostname     string
	HttpProtocol     string
	HttpReadTimeout  time.Duration
	HttpWriteTimeout time.Duration
	HttpIdleTimeout  time.Duration
	pFlag            *pflag.FlagSet
	viperEnvAndFile  *viper.Viper
	viperFlag        *viper.Viper
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
		Usage:    "set the configuration file location\" e.g.:\" ./config.yml \"",
		castType: castString,
	},
	envCacheShards: {
		Flag:     "cShards",
		Env:      envCacheShards,
		File:     "",
		Default:  "",
		Usage:    "\" e.g.:\" \"",
		castType: castInt,
	},
}

func loadConf(c *Configuration) (got Configuration) {
	// don't forget to add new configurations here
	return Configuration{
		CacheShards:             cast.ToInt(getConf(c, envCacheShards)),
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
	registerFlags(c)
	parseFlags(c)
	registerEnvs(c)
	parseFile(c)
	C = loadConf(c)
}
