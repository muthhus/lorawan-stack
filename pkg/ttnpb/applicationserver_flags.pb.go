// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v0.0.0-dev
// - protoc              v3.17.3
// source: lorawan-stack/api/applicationserver.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	types "github.com/gogo/protobuf/types"
	pflag "github.com/spf13/pflag"
)

// AddSelectFlagsForApplicationLink adds flags to select fields in ApplicationLink.
func AddSelectFlagsForApplicationLink(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("default-formatters", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("default-formatters", prefix), true), hidden))
	AddSelectFlagsForMessagePayloadFormatters(flags, flagsplugin.Prefix("default-formatters", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("skip-payload-crypto", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("skip-payload-crypto", prefix), false), hidden))
}

// SelectFromFlags outputs the fieldmask paths forApplicationLink message from select flags.
func PathsFromSelectFlagsForApplicationLink(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("default_formatters", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("default_formatters", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForMessagePayloadFormatters(flags, flagsplugin.Prefix("default_formatters", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("skip_payload_crypto", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("skip_payload_crypto", prefix))
	}
	return paths, nil
}

// AddSetFlagsForApplicationLink adds flags to select fields in ApplicationLink.
func AddSetFlagsForApplicationLink(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForMessagePayloadFormatters(flags, flagsplugin.Prefix("default-formatters", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("skip-payload-crypto.value", prefix), "", hidden))
	flagsplugin.AddAlias(flags, flagsplugin.Prefix("skip-payload-crypto.value", prefix), flagsplugin.Prefix("skip-payload-crypto", prefix), hidden)
}

// SetFromFlags sets the ApplicationLink message from flags.
func (m *ApplicationLink) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("default_formatters", prefix)); changed {
		m.DefaultFormatters = &MessagePayloadFormatters{}
		if setPaths, err := m.DefaultFormatters.SetFromFlags(flags, flagsplugin.Prefix("default_formatters", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("skip_payload_crypto.value", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.SkipPayloadCrypto = &types.BoolValue{Value: val}
		paths = append(paths, flagsplugin.Prefix("skip_payload_crypto", prefix))
	}
	return paths, nil
}
