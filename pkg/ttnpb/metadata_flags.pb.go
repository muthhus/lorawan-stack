// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v0.0.0-dev
// - protoc              v3.17.3
// source: lorawan-stack/api/metadata.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	pflag "github.com/spf13/pflag"
)

// AddSelectFlagsForLocation adds flags to select fields in Location.
func AddSelectFlagsForLocation(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("latitude", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("latitude", prefix), false), hidden))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("longitude", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("longitude", prefix), false), hidden))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("altitude", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("altitude", prefix), false), hidden))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("accuracy", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("accuracy", prefix), false), hidden))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("source", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("source", prefix), false), hidden))
}

// SelectFromFlags outputs the fieldmask paths forLocation message from select flags.
func PathsFromSelectFlagsForLocation(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("latitude", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("latitude", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("longitude", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("longitude", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("altitude", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("altitude", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("accuracy", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("accuracy", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("source", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("source", prefix))
	}
	return paths, nil
}

// AddSetFlagsForLocation adds flags to select fields in Location.
func AddSetFlagsForLocation(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewFloat64Flag(flagsplugin.Prefix("latitude", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewFloat64Flag(flagsplugin.Prefix("longitude", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewInt32Flag(flagsplugin.Prefix("altitude", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewInt32Flag(flagsplugin.Prefix("accuracy", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("source", prefix), flagsplugin.EnumValueDesc(LocationSource_value), hidden))
}

// SetFromFlags sets the Location message from flags.
func (m *Location) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetFloat64(flags, flagsplugin.Prefix("latitude", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Latitude = val
		paths = append(paths, flagsplugin.Prefix("latitude", prefix))
	}
	if val, selected, err := flagsplugin.GetFloat64(flags, flagsplugin.Prefix("longitude", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Longitude = val
		paths = append(paths, flagsplugin.Prefix("longitude", prefix))
	}
	if val, selected, err := flagsplugin.GetInt32(flags, flagsplugin.Prefix("altitude", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Altitude = val
		paths = append(paths, flagsplugin.Prefix("altitude", prefix))
	}
	if val, selected, err := flagsplugin.GetInt32(flags, flagsplugin.Prefix("accuracy", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Accuracy = val
		paths = append(paths, flagsplugin.Prefix("accuracy", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("source", prefix)); err != nil {
		return nil, err
	} else if selected {
		enumValue, err := flagsplugin.SetEnumString(val, LocationSource_value)
		if err != nil {
			return nil, err
		}
		m.Source = LocationSource(enumValue)
		paths = append(paths, flagsplugin.Prefix("source", prefix))
	}
	return paths, nil
}
