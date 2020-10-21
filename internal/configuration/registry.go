package configuration

import "github.com/spf13/cast"

type registerFlagType func(conf config, c *Configuration)

// all types accepted by cast package
const (
	castBool          = iota
	castDuration      = iota
	castFloat64       = iota
	castFloat32       = iota
	castInt64         = iota
	castInt32         = iota
	castInt16         = iota
	castInt8          = iota
	castInt           = iota
	castUint          = iota
	castUint64        = iota
	castUint32        = iota
	castUint16        = iota
	castUint8         = iota
	castString        = iota
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
		castBoolSlice:     registerBoolSlice,
		castStringSlice:   registerStringSlice,
		castIntSlice:      registerIntSlice,
		castDurationSlice: registerDurationSlice,
	}
	typeFunc[conf.castType](conf, c)
}

func registerBool(conf config, c *Configuration) {
	c.pFlag.Bool(
		conf.Flag,
		cast.ToBool(conf.Default),
		conf.Usage,
	)
}

func registerDuration(conf config, c *Configuration) {
	c.pFlag.Duration(
		conf.Flag,
		cast.ToDuration(conf.Default),
		conf.Usage,
	)
}

func registerFloat64(conf config, c *Configuration) {
	c.pFlag.Float64(
		conf.Flag,
		cast.ToFloat64(conf.Default),
		conf.Usage,
	)
}

func registerFloat32(conf config, c *Configuration) {
	c.pFlag.Float32(
		conf.Flag,
		cast.ToFloat32(conf.Default),
		conf.Usage,
	)
}

func registerInt64(conf config, c *Configuration) {
	c.pFlag.Int64(
		conf.Flag,
		cast.ToInt64(conf.Default),
		conf.Usage,
	)
}

func registerInt32(conf config, c *Configuration) {
	c.pFlag.Int32(
		conf.Flag,
		cast.ToInt32(conf.Default),
		conf.Usage,
	)
}

func registerInt16(conf config, c *Configuration) {
	c.pFlag.Int16(
		conf.Flag,
		cast.ToInt16(conf.Default),
		conf.Usage,
	)
}

func registerInt8(conf config, c *Configuration) {
	c.pFlag.Int8(
		conf.Flag,
		cast.ToInt8(conf.Default),
		conf.Usage,
	)
}

func registerInt(conf config, c *Configuration) {
	c.pFlag.Int(
		conf.Flag,
		cast.ToInt(conf.Default),
		conf.Usage,
	)
}

func registerUint(conf config, c *Configuration) {
	c.pFlag.Uint(
		conf.Flag,
		cast.ToUint(conf.Default),
		conf.Usage,
	)
}

func registerUint64(conf config, c *Configuration) {
	c.pFlag.Uint64(
		conf.Flag,
		cast.ToUint64(conf.Default),
		conf.Usage,
	)
}

func registerUint32(conf config, c *Configuration) {
	c.pFlag.Uint32(
		conf.Flag,
		cast.ToUint32(conf.Default),
		conf.Usage,
	)
}

func registerUint16(conf config, c *Configuration) {
	c.pFlag.Uint16(
		conf.Flag,
		cast.ToUint16(conf.Default),
		conf.Usage,
	)
}

func registerUint8(conf config, c *Configuration) {
	c.pFlag.Uint8(
		conf.Flag,
		cast.ToUint8(conf.Default),
		conf.Usage,
	)
}

func registerString(conf config, c *Configuration) {
	c.pFlag.String(
		conf.Flag,
		cast.ToString(conf.Default),
		conf.Usage,
	)
}

func registerBoolSlice(conf config, c *Configuration) {
	c.pFlag.BoolSlice(
		conf.Flag,
		cast.ToBoolSlice(conf.Default),
		conf.Usage,
	)
}

func registerStringSlice(conf config, c *Configuration) {
	c.pFlag.StringSlice(
		conf.Flag,
		cast.ToStringSlice(conf.Default),
		conf.Usage,
	)
}

func registerIntSlice(conf config, c *Configuration) {
	c.pFlag.IntSlice(
		conf.Flag,
		cast.ToIntSlice(conf.Default),
		conf.Usage,
	)
}

func registerDurationSlice(conf config, c *Configuration) {
	c.pFlag.DurationSlice(
		conf.Flag,
		cast.ToDurationSlice(conf.Default),
		conf.Usage,
	)
}


