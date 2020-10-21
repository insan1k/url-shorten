package configuration

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

//handle errors in configuration package
func handleError(err error) {
	log.Fatal(fmt.Errorf("fatal config caught an error:%v", err))
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
func registerEnvs(config *Configuration) {
	config.viperEnvAndFile.SetEnvPrefix(envPrefix)
	for _, k := range configurationMap {
		if k.Env != "" {
			if err := config.viperEnvAndFile.BindEnv(k.Env); err != nil {
				handleError(err)
			}
		}
	}
}

//parseFlags
//parses registered flags that can be passed as arguments
func parseFlags(config *Configuration) {
	if err := config.pFlag.Parse(config.pFlag.Args()); err != nil {
		handleError(err)
	}
	if err := config.viperFlag.BindPFlags(pflag.CommandLine); err != nil {
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
			return
		}
	}
	configFromFlag := c.viperFlag.GetString(configurationMap[configFile].Flag)
	if configFromFlag != "" {
		c.viperEnvAndFile.SetConfigFile(configFromFlag)
		err := c.viperEnvAndFile.ReadInConfig()
		if err == nil {
			return
		}
	}
	c.viperEnvAndFile.SetConfigFile(configurationMap[configFile].Default)
	err := c.viperEnvAndFile.ReadInConfig()
	if err == nil {
		return
	}
	c.viperEnvAndFile.SetConfigFile("./" + fileName + "." + fileExtension)
	err = c.viperEnvAndFile.ReadInConfig()
	if err != nil {
		return
	}
}

func getConf(c *Configuration, name string) (got interface{}) {
	return getter(c.viperEnvAndFile, c.viperFlag, configurationMap[name])
}

func getter(env *viper.Viper, flag *viper.Viper, c config) (asserted interface{}) {
	asserted = c.Default
	assertYaml := getConfFromYaml(env, c)
	assertEnv := getConfFromEnv(env, c)
	assertFlag := getConfFromFlag(flag, c)
	if assertYaml != "" {
		asserted = assertYaml
	}
	if assertEnv != "" {
		asserted = assertEnv
	}
	if assertFlag != "" {
		asserted = assertFlag
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
