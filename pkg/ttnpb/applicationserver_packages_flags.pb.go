// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v0.0.0-dev
// - protoc              v3.17.3
// source: lorawan-stack/api/applicationserver_packages.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	gogo "github.com/TheThingsIndustries/protoc-gen-go-flags/gogo"
	pflag "github.com/spf13/pflag"
)

// AddSelectFlagsForApplicationPackageAssociationIdentifiers adds flags to select fields in ApplicationPackageAssociationIdentifiers.
func AddSelectFlagsForApplicationPackageAssociationIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("end-device-ids", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("end-device-ids", prefix), true), hidden))
	// NOTE: end_device_ids (EndDeviceIdentifiers) does not seem to have select flags.
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("f-port", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("f-port", prefix), false), hidden))
}

// SelectFromFlags outputs the fieldmask paths forApplicationPackageAssociationIdentifiers message from select flags.
func PathsFromSelectFlagsForApplicationPackageAssociationIdentifiers(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("end_device_ids", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("end_device_ids", prefix))
	}
	// NOTE: end_device_ids (EndDeviceIdentifiers) does not seem to have select flags.
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("f_port", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("f_port", prefix))
	}
	return paths, nil
}

// AddSetFlagsForApplicationPackageAssociationIdentifiers adds flags to select fields in ApplicationPackageAssociationIdentifiers.
func AddSetFlagsForApplicationPackageAssociationIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForEndDeviceIdentifiers(flags, flagsplugin.Prefix("end-device-ids", prefix), hidden)
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("f-port", prefix), "", hidden))
}

// SetFromFlags sets the ApplicationPackageAssociationIdentifiers message from flags.
func (m *ApplicationPackageAssociationIdentifiers) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if selected := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("end_device_ids", prefix)); selected {
		m.EndDeviceIds = &EndDeviceIdentifiers{}
		if setPaths, err := m.EndDeviceIds.SetFromFlags(flags, flagsplugin.Prefix("end_device_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("f_port", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.FPort = val
		paths = append(paths, flagsplugin.Prefix("f_port", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForApplicationPackageAssociation adds flags to select fields in ApplicationPackageAssociation.
func AddSelectFlagsForApplicationPackageAssociation(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("ids", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("ids", prefix), true), hidden))
	// NOTE: ids (ApplicationPackageAssociationIdentifiers) does not seem to have select flags.
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("created-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("created-at", prefix), false), hidden))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("updated-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("updated-at", prefix), false), hidden))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("package-name", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("package-name", prefix), false), hidden))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("data", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("data", prefix), false), hidden))
}

// SelectFromFlags outputs the fieldmask paths forApplicationPackageAssociation message from select flags.
func PathsFromSelectFlagsForApplicationPackageAssociation(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("ids", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("ids", prefix))
	}
	// NOTE: ids (ApplicationPackageAssociationIdentifiers) does not seem to have select flags.
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("created_at", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("created_at", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("updated_at", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("updated_at", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("package_name", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("package_name", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("data", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("data", prefix))
	}
	return paths, nil
}

// AddSetFlagsForApplicationPackageAssociation adds flags to select fields in ApplicationPackageAssociation.
func AddSetFlagsForApplicationPackageAssociation(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForApplicationPackageAssociationIdentifiers(flags, flagsplugin.Prefix("ids", prefix), hidden)
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("created-at", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("updated-at", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("package-name", prefix), "", hidden))
	// FIXME: Skipping Data because this WKT is currently not supported.
}

// SetFromFlags sets the ApplicationPackageAssociation message from flags.
func (m *ApplicationPackageAssociation) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if selected := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("ids", prefix)); selected {
		m.Ids = &ApplicationPackageAssociationIdentifiers{}
		if setPaths, err := m.Ids.SetFromFlags(flags, flagsplugin.Prefix("ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, selected, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("created_at", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.CreatedAt = gogo.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("created_at", prefix))
	}
	if val, selected, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("updated_at", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.UpdatedAt = gogo.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("updated_at", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("package_name", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.PackageName = val
		paths = append(paths, flagsplugin.Prefix("package_name", prefix))
	}
	// FIXME: Skipping Data because this WKT is not supported.
	return paths, nil
}

// AddSelectFlagsForApplicationPackageDefaultAssociationIdentifiers adds flags to select fields in ApplicationPackageDefaultAssociationIdentifiers.
func AddSelectFlagsForApplicationPackageDefaultAssociationIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("application-ids", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("application-ids", prefix), true), hidden))
	AddSelectFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application-ids", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("f-port", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("f-port", prefix), false), hidden))
}

// SelectFromFlags outputs the fieldmask paths forApplicationPackageDefaultAssociationIdentifiers message from select flags.
func PathsFromSelectFlagsForApplicationPackageDefaultAssociationIdentifiers(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("application_ids", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("f_port", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("f_port", prefix))
	}
	return paths, nil
}

// AddSetFlagsForApplicationPackageDefaultAssociationIdentifiers adds flags to select fields in ApplicationPackageDefaultAssociationIdentifiers.
func AddSetFlagsForApplicationPackageDefaultAssociationIdentifiers(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application-ids", prefix), hidden)
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("f-port", prefix), "", hidden))
}

// SetFromFlags sets the ApplicationPackageDefaultAssociationIdentifiers message from flags.
func (m *ApplicationPackageDefaultAssociationIdentifiers) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if selected := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("application_ids", prefix)); selected {
		m.ApplicationIds = &ApplicationIdentifiers{}
		if setPaths, err := m.ApplicationIds.SetFromFlags(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("f_port", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.FPort = val
		paths = append(paths, flagsplugin.Prefix("f_port", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForApplicationPackageDefaultAssociation adds flags to select fields in ApplicationPackageDefaultAssociation.
func AddSelectFlagsForApplicationPackageDefaultAssociation(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("ids", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("ids", prefix), true), hidden))
	// NOTE: ids (ApplicationPackageDefaultAssociationIdentifiers) does not seem to have select flags.
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("created-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("created-at", prefix), false), hidden))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("updated-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("updated-at", prefix), false), hidden))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("package-name", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("package-name", prefix), false), hidden))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("data", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("data", prefix), false), hidden))
}

// SelectFromFlags outputs the fieldmask paths forApplicationPackageDefaultAssociation message from select flags.
func PathsFromSelectFlagsForApplicationPackageDefaultAssociation(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("ids", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("ids", prefix))
	}
	// NOTE: ids (ApplicationPackageDefaultAssociationIdentifiers) does not seem to have select flags.
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("created_at", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("created_at", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("updated_at", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("updated_at", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("package_name", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("package_name", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("data", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("data", prefix))
	}
	return paths, nil
}

// AddSetFlagsForApplicationPackageDefaultAssociation adds flags to select fields in ApplicationPackageDefaultAssociation.
func AddSetFlagsForApplicationPackageDefaultAssociation(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForApplicationPackageDefaultAssociationIdentifiers(flags, flagsplugin.Prefix("ids", prefix), hidden)
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("created-at", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("updated-at", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("package-name", prefix), "", hidden))
	// FIXME: Skipping Data because this WKT is currently not supported.
}

// SetFromFlags sets the ApplicationPackageDefaultAssociation message from flags.
func (m *ApplicationPackageDefaultAssociation) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if selected := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("ids", prefix)); selected {
		m.Ids = &ApplicationPackageDefaultAssociationIdentifiers{}
		if setPaths, err := m.Ids.SetFromFlags(flags, flagsplugin.Prefix("ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, selected, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("created_at", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.CreatedAt = gogo.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("created_at", prefix))
	}
	if val, selected, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("updated_at", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.UpdatedAt = gogo.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("updated_at", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("package_name", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.PackageName = val
		paths = append(paths, flagsplugin.Prefix("package_name", prefix))
	}
	// FIXME: Skipping Data because this WKT is not supported.
	return paths, nil
}
