package configuration

import (
	"github.com/insan1k/one-qr-dot-me/internal/logger"
	"github.com/spf13/viper"
	"os"
)

//handle errors in configuration package
func handleError(err error) {
	logger.L.Fatalf("config error %v", err.Error())
}

//registerFlags
//register the flags with pFlag
func registerFlags(c *Configuration) {
	for _, k := range configurationMap {
		if k.Flag != "" {
			flagRegister(k, c)
		}
	}
}

//registerEnvs
//register environment variables
func registerEnvs(c *Configuration) {
	c.viperEnvAndFile.SetEnvPrefix(envPrefix)
	for _, k := range configurationMap {
		if k.Env != "" {
			if err := c.viperEnvAndFile.BindEnv(k.Env); err != nil {
				handleError(err)
			}
		}
	}
}

//parseFlags
//parses registered flags that can be passed as arguments
func parseFlags(c *Configuration) {
	if err := c.pFlag.Parse(os.Args[1:]); err != nil {
		handleError(err)
	}
	if err := c.viperFlag.BindPFlags(c.pFlag); err != nil {
		handleError(err)
	}
}

//parseFile
//parses the yaml file
func parseFile(c *Configuration) {
	c.viperEnvAndFile.SetConfigType(fileType)
	configFromEnv := c.viperEnvAndFile.GetString(configurationMap[configFile].Env)
	if configFromEnv != "" {
		c.viperEnvAndFile.SetConfigFile(configFromEnv)
		err := c.viperEnvAndFile.ReadInConfig()
		if err == nil {
			logger.L.Debug("config file loaded from from env")
			return
		}
	}
	configFromFlag := c.viperFlag.GetString(configurationMap[configFile].Flag)
	if configFromFlag != "" {
		c.viperEnvAndFile.SetConfigFile(configFromFlag)
		err := c.viperEnvAndFile.ReadInConfig()
		if err == nil {
			logger.L.Debug("config file loaded from from flag")
			return
		}
	}
	c.viperEnvAndFile.SetConfigFile(configurationMap[configFile].Default)
	err := c.viperEnvAndFile.ReadInConfig()
	if err == nil {
		logger.L.Debug("config file loaded from defaults")
		return
	}
	c.viperEnvAndFile.SetConfigFile("./" + fileName + "." + fileExtension)
	err = c.viperEnvAndFile.ReadInConfig()
	if err != nil {
		logger.L.Info("config file was not found in pwd")
		return
	}
}

func getConf(c *Configuration, name string) (got interface{}) {
	return getter(c.viperEnvAndFile, c.viperFlag, configurationMap[name])
}

func getter(env *viper.Viper, flag *viper.Viper, c config) (asserted interface{}) {
	asserted = c.Default
	inFile := getConfFromYaml(env, c)
	inEnv := getConfFromEnv(env, c)
	inFlag := getConfFromFlag(flag, c)
	if inFile != "" {
		asserted = inFile
	}
	if inEnv != "" {
		asserted = inEnv
	}
	if inFlag != "" {
		asserted = inFlag
	}
	return
}

func getConfFromYaml(snake *viper.Viper, c config) (asserted string) {
	if c.File != "" {
		asserted = snake.GetString(c.File)
	}
	return
}

func getConfFromEnv(snake *viper.Viper, c config) (asserted string) {
	if c.Env != "" {
		asserted = snake.GetString(c.Env)
	}
	return
}

func getConfFromFlag(snake *viper.Viper, c config) (asserted string) {
	if c.Flag != "" {
		if got := snake.GetString(c.Flag); got != c.Default {
			asserted = got
		}
	}
	return
}
