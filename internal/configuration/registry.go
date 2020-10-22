package configuration

import (
	"errors"
	"github.com/spf13/cast"
	"strings"
)

type registerFlagType func(conf config, c *Configuration)

// all types accepted by cast package
const (
	castBool     = iota
	castDuration = iota
	castFloat64  = iota
	castFloat32  = iota
	castInt64    = iota
	castInt32    = iota
	castInt16    = iota
	castInt8     = iota
	castInt      = iota
	castUint     = iota
	castUint64   = iota
	castUint32   = iota
	castUint16   = iota
	castUint8    = iota
	castString   = iota
	//implement parser methods for these types if need arises
	castBoolSlice     = iota
	castStringSlice   = iota
	castIntSlice      = iota
	castDurationSlice = iota
)

func flagRegister(conf config, c *Configuration) {
	typeFunc := map[int]registerFlagType{
		castBool:          registerBool,
		castDuration:      registerDuration,
		castFloat64:       registerFloat64,
		castFloat32:       registerFloat32,
		castInt64:         registerInt64,
		castInt32:         registerInt32,
		castInt16:         registerInt16,
		castInt8:          registerInt8,
		castInt:           registerInt,
		castUint:          registerUint,
		castUint64:        registerUint64,
		castUint32:        registerUint32,
		castUint16:        registerUint16,
		castUint8:         registerUint8,
		castString:        registerString,
		castBoolSlice:     notImplemented,
		castStringSlice:   registerStringSlice,
		castIntSlice:      notImplemented,
		castDurationSlice: notImplemented,
	}
	typeFunc[conf.castType](conf, c)
}

func registerBool(conf config, c *Configuration) {
	val, err := cast.ToBoolE(conf.Default)
	if err != nil {
		handleError(err)
	}
	c.pFlag.Bool(
		conf.Flag,
		val,
		conf.Usage,
	)
}

func registerDuration(conf config, c *Configuration) {
	val, err := cast.ToDurationE(conf.Default)
	if err != nil {
		handleError(err)
	}
	c.pFlag.Duration(
		conf.Flag,
		val,
		conf.Usage,
	)
}

func registerFloat64(conf config, c *Configuration) {
	val, err := cast.ToFloat64E(conf.Default)
	if err != nil {
		handleError(err)
	}
	c.pFlag.Float64(
		conf.Flag,
		val,
		conf.Usage,
	)
}

func registerFloat32(conf config, c *Configuration) {
	val, err := cast.ToFloat32E(conf.Default)
	if err != nil {
		handleError(err)
	}
	c.pFlag.Float32(
		conf.Flag,
		val,
		conf.Usage,
	)
}

func registerInt64(conf config, c *Configuration) {
	val, err := cast.ToInt64E(conf.Default)
	if err != nil {
		handleError(err)
	}
	c.pFlag.Int64(
		conf.Flag,
		val,
		conf.Usage,
	)
}

func registerInt32(conf config, c *Configuration) {
	val, err := cast.ToInt32E(conf.Default)
	if err != nil {
		handleError(err)
	}
	c.pFlag.Int32(
		conf.Flag,
		val,
		conf.Usage,
	)
}

func registerInt16(conf config, c *Configuration) {
	val, err := cast.ToInt16E(conf.Default)
	if err != nil {
		handleError(err)
	}
	c.pFlag.Int16(
		conf.Flag,
		val,
		conf.Usage,
	)
}

func registerInt8(conf config, c *Configuration) {
	val, err := cast.ToInt8E(conf.Default)
	if err != nil {
		handleError(err)
	}
	c.pFlag.Int8(
		conf.Flag,
		val,
		conf.Usage,
	)
}

func registerInt(conf config, c *Configuration) {
	val, err := cast.ToIntE(conf.Default)
	if err != nil {
		handleError(err)
	}
	c.pFlag.Int(
		conf.Flag,
		val,
		conf.Usage,
	)
}

func registerUint(conf config, c *Configuration) {
	val, err := cast.ToUintE(conf.Default)
	if err != nil {
		handleError(err)
	}
	c.pFlag.Uint(
		conf.Flag,
		val,
		conf.Usage,
	)
}

func registerUint64(conf config, c *Configuration) {
	val, err := cast.ToUint64E(conf.Default)
	if err != nil {
		handleError(err)
	}
	c.pFlag.Uint64(
		conf.Flag,
		val,
		conf.Usage,
	)
}

func registerUint32(conf config, c *Configuration) {
	val, err := cast.ToUint32E(conf.Default)
	if err != nil {
		handleError(err)
	}
	c.pFlag.Uint32(
		conf.Flag,
		val,
		conf.Usage,
	)
}

func registerUint16(conf config, c *Configuration) {
	val, err := cast.ToUint16E(conf.Default)
	if err != nil {
		handleError(err)
	}
	c.pFlag.Uint16(
		conf.Flag,
		val,
		conf.Usage,
	)
}

func registerUint8(conf config, c *Configuration) {
	val, err := cast.ToUint8E(conf.Default)
	if err != nil {
		handleError(err)
	}
	c.pFlag.Uint8(
		conf.Flag,
		val,
		conf.Usage,
	)
}

func registerString(conf config, c *Configuration) {
	c.pFlag.String(
		conf.Flag,
		conf.Default,
		conf.Usage,
	)
}

func registerStringSlice(conf config, c *Configuration) {
	ss, err := cast.ToStringSliceE(strings.ReplaceAll(conf.Default, ",", " "))
	if err != nil {
		handleError(err)
	}
	c.pFlag.StringSlice(
		conf.Flag,
		ss,
		conf.Usage,
	)
}

func notImplemented(config, *Configuration) {
	handleError(errors.New("not implemented"))
}
