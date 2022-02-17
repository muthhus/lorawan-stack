// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v0.0.0-dev
// - protoc              v3.17.3
// source: lorawan-stack/api/search_services.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	gogo "github.com/TheThingsIndustries/protoc-gen-go-flags/gogo"
	pflag "github.com/spf13/pflag"
)

// AddSetFlagsForSearchApplicationsRequest adds flags to select fields in SearchApplicationsRequest.
func AddSetFlagsForSearchApplicationsRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("id-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("description-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringStringMapFlag(flagsplugin.Prefix("attributes-contain", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted", prefix), "", hidden))
}

// SetFromFlags sets the SearchApplicationsRequest message from flags.
func (m *SearchApplicationsRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("id_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.IdContains = val
		paths = append(paths, flagsplugin.Prefix("id_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.NameContains = val
		paths = append(paths, flagsplugin.Prefix("name_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("description_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.DescriptionContains = val
		paths = append(paths, flagsplugin.Prefix("description_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetStringStringMap(flags, flagsplugin.Prefix("attributes_contain", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.AttributesContain = val
		paths = append(paths, flagsplugin.Prefix("attributes_contain", prefix))
	}
	if val, selected, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.FieldMask = gogo.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Deleted = val
		paths = append(paths, flagsplugin.Prefix("deleted", prefix))
	}
	return paths, nil
}

// AddSetFlagsForSearchClientsRequest adds flags to select fields in SearchClientsRequest.
func AddSetFlagsForSearchClientsRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("id-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("description-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringStringMapFlag(flagsplugin.Prefix("attributes-contain", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("state", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted", prefix), "", hidden))
}

// SetFromFlags sets the SearchClientsRequest message from flags.
func (m *SearchClientsRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("id_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.IdContains = val
		paths = append(paths, flagsplugin.Prefix("id_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.NameContains = val
		paths = append(paths, flagsplugin.Prefix("name_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("description_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.DescriptionContains = val
		paths = append(paths, flagsplugin.Prefix("description_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetStringStringMap(flags, flagsplugin.Prefix("attributes_contain", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.AttributesContain = val
		paths = append(paths, flagsplugin.Prefix("attributes_contain", prefix))
	}
	if val, selected, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("state", prefix)); err != nil {
		return nil, err
	} else if selected {
		for _, v := range val {
			enumValue, err := flagsplugin.SetEnumString(v, State_value)
			if err != nil {
				return nil, err
			}
			m.State = append(m.State, State(enumValue))
		}
		paths = append(paths, flagsplugin.Prefix("state", prefix))
	}
	if val, selected, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.FieldMask = gogo.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Deleted = val
		paths = append(paths, flagsplugin.Prefix("deleted", prefix))
	}
	return paths, nil
}

// AddSetFlagsForSearchGatewaysRequest adds flags to select fields in SearchGatewaysRequest.
func AddSetFlagsForSearchGatewaysRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("id-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("description-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringStringMapFlag(flagsplugin.Prefix("attributes-contain", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("eui-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted", prefix), "", hidden))
}

// SetFromFlags sets the SearchGatewaysRequest message from flags.
func (m *SearchGatewaysRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("id_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.IdContains = val
		paths = append(paths, flagsplugin.Prefix("id_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.NameContains = val
		paths = append(paths, flagsplugin.Prefix("name_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("description_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.DescriptionContains = val
		paths = append(paths, flagsplugin.Prefix("description_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetStringStringMap(flags, flagsplugin.Prefix("attributes_contain", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.AttributesContain = val
		paths = append(paths, flagsplugin.Prefix("attributes_contain", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("eui_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.EuiContains = val
		paths = append(paths, flagsplugin.Prefix("eui_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.FieldMask = gogo.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Deleted = val
		paths = append(paths, flagsplugin.Prefix("deleted", prefix))
	}
	return paths, nil
}

// AddSetFlagsForSearchOrganizationsRequest adds flags to select fields in SearchOrganizationsRequest.
func AddSetFlagsForSearchOrganizationsRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("id-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("description-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringStringMapFlag(flagsplugin.Prefix("attributes-contain", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted", prefix), "", hidden))
}

// SetFromFlags sets the SearchOrganizationsRequest message from flags.
func (m *SearchOrganizationsRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("id_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.IdContains = val
		paths = append(paths, flagsplugin.Prefix("id_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.NameContains = val
		paths = append(paths, flagsplugin.Prefix("name_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("description_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.DescriptionContains = val
		paths = append(paths, flagsplugin.Prefix("description_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetStringStringMap(flags, flagsplugin.Prefix("attributes_contain", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.AttributesContain = val
		paths = append(paths, flagsplugin.Prefix("attributes_contain", prefix))
	}
	if val, selected, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.FieldMask = gogo.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Deleted = val
		paths = append(paths, flagsplugin.Prefix("deleted", prefix))
	}
	return paths, nil
}

// AddSetFlagsForSearchUsersRequest adds flags to select fields in SearchUsersRequest.
func AddSetFlagsForSearchUsersRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("id-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("description-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringStringMapFlag(flagsplugin.Prefix("attributes-contain", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("state", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted", prefix), "", hidden))
}

// SetFromFlags sets the SearchUsersRequest message from flags.
func (m *SearchUsersRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("id_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.IdContains = val
		paths = append(paths, flagsplugin.Prefix("id_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.NameContains = val
		paths = append(paths, flagsplugin.Prefix("name_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("description_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.DescriptionContains = val
		paths = append(paths, flagsplugin.Prefix("description_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetStringStringMap(flags, flagsplugin.Prefix("attributes_contain", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.AttributesContain = val
		paths = append(paths, flagsplugin.Prefix("attributes_contain", prefix))
	}
	if val, selected, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("state", prefix)); err != nil {
		return nil, err
	} else if selected {
		for _, v := range val {
			enumValue, err := flagsplugin.SetEnumString(v, State_value)
			if err != nil {
				return nil, err
			}
			m.State = append(m.State, State(enumValue))
		}
		paths = append(paths, flagsplugin.Prefix("state", prefix))
	}
	if val, selected, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.FieldMask = gogo.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Deleted = val
		paths = append(paths, flagsplugin.Prefix("deleted", prefix))
	}
	return paths, nil
}

// AddSetFlagsForSearchEndDevicesRequest adds flags to select fields in SearchEndDevicesRequest.
func AddSetFlagsForSearchEndDevicesRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application-ids", prefix), hidden)
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("id-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("description-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringStringMapFlag(flagsplugin.Prefix("attributes-contain", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("dev-eui-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("join-eui-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("dev-addr-contains", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", hidden))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", hidden))
}

// SetFromFlags sets the SearchEndDevicesRequest message from flags.
func (m *SearchEndDevicesRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if selected := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("application_ids", prefix)); selected {
		m.ApplicationIds = &ApplicationIdentifiers{}
		if setPaths, err := m.ApplicationIds.SetFromFlags(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("id_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.IdContains = val
		paths = append(paths, flagsplugin.Prefix("id_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.NameContains = val
		paths = append(paths, flagsplugin.Prefix("name_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("description_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.DescriptionContains = val
		paths = append(paths, flagsplugin.Prefix("description_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetStringStringMap(flags, flagsplugin.Prefix("attributes_contain", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.AttributesContain = val
		paths = append(paths, flagsplugin.Prefix("attributes_contain", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("dev_eui_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.DevEuiContains = val
		paths = append(paths, flagsplugin.Prefix("dev_eui_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("join_eui_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.JoinEuiContains = val
		paths = append(paths, flagsplugin.Prefix("join_eui_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("dev_addr_contains", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.DevAddrContains = val
		paths = append(paths, flagsplugin.Prefix("dev_addr_contains", prefix))
	}
	if val, selected, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.FieldMask = gogo.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, selected, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, selected, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if selected {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	return paths, nil
}
