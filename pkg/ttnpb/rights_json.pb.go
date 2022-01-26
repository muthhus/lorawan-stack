// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.3.1
// - protoc             v3.9.1
// source: lorawan-stack/api/rights.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the Right to JSON.
func (x Right) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	s.WriteEnumString(int32(x), Right_name)
}

// MarshalText marshals the Right to text.
func (x Right) MarshalText() ([]byte, error) {
	return []byte(jsonplugin.GetEnumString(int32(x), Right_name)), nil
}

// MarshalJSON marshals the Right to JSON.
func (x Right) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// Right_customvalue contains custom string values that extend Right_value.
var Right_customvalue = map[string]int32{
	"USER_INFO":                          1,
	"USER_SETTINGS_BASIC":                2,
	"USER_SETTINGS_API_KEYS":             3,
	"USER_DELETE":                        4,
	"USER_AUTHORIZED_CLIENTS":            5,
	"USER_APPLICATIONS_LIST":             6,
	"USER_APPLICATIONS_CREATE":           7,
	"USER_GATEWAYS_LIST":                 8,
	"USER_GATEWAYS_CREATE":               9,
	"USER_CLIENTS_LIST":                  10,
	"USER_CLIENTS_CREATE":                11,
	"USER_ORGANIZATIONS_LIST":            12,
	"USER_ORGANIZATIONS_CREATE":          13,
	"USERS_NOTIFICATIONS_READ":           59,
	"USER_ALL":                           14,
	"APPLICATION_INFO":                   15,
	"APPLICATION_SETTINGS_BASIC":         16,
	"APPLICATION_SETTINGS_API_KEYS":      17,
	"APPLICATION_SETTINGS_COLLABORATORS": 18,
	"APPLICATION_SETTINGS_PACKAGES":      56,
	"APPLICATION_DELETE":                 19,
	"APPLICATION_DEVICES_READ":           20,
	"APPLICATION_DEVICES_WRITE":          21,
	"APPLICATION_DEVICES_READ_KEYS":      22,
	"APPLICATION_DEVICES_WRITE_KEYS":     23,
	"APPLICATION_TRAFFIC_READ":           24,
	"APPLICATION_TRAFFIC_UP_WRITE":       25,
	"APPLICATION_TRAFFIC_DOWN_WRITE":     26,
	"APPLICATION_LINK":                   27,
	"APPLICATION_ALL":                    28,
	"CLIENT_ALL":                         29,
	"GATEWAY_INFO":                       30,
	"GATEWAY_SETTINGS_BASIC":             31,
	"GATEWAY_SETTINGS_API_KEYS":          32,
	"GATEWAY_SETTINGS_COLLABORATORS":     33,
	"GATEWAY_DELETE":                     34,
	"GATEWAY_TRAFFIC_READ":               35,
	"GATEWAY_TRAFFIC_DOWN_WRITE":         36,
	"GATEWAY_LINK":                       37,
	"GATEWAY_STATUS_READ":                38,
	"GATEWAY_LOCATION_READ":              39,
	"GATEWAY_WRITE_SECRETS":              57,
	"GATEWAY_READ_SECRETS":               58,
	"GATEWAY_ALL":                        40,
	"ORGANIZATION_INFO":                  41,
	"ORGANIZATION_SETTINGS_BASIC":        42,
	"ORGANIZATION_SETTINGS_API_KEYS":     43,
	"ORGANIZATION_SETTINGS_MEMBERS":      44,
	"ORGANIZATION_DELETE":                45,
	"ORGANIZATION_APPLICATIONS_LIST":     46,
	"ORGANIZATION_APPLICATIONS_CREATE":   47,
	"ORGANIZATION_GATEWAYS_LIST":         48,
	"ORGANIZATION_GATEWAYS_CREATE":       49,
	"ORGANIZATION_CLIENTS_LIST":          50,
	"ORGANIZATION_CLIENTS_CREATE":        51,
	"ORGANIZATION_ADD_AS_COLLABORATOR":   52,
	"ORGANIZATION_ALL":                   53,
	"SEND_INVITES":                       54,
	"ALL":                                55,
}

// UnmarshalProtoJSON unmarshals the Right from JSON.
func (x *Right) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	v := s.ReadEnum(Right_value, Right_customvalue)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read Right enum: %v", err)
		return
	}
	*x = Right(v)
}

// UnmarshalText unmarshals the Right from text.
func (x *Right) UnmarshalText(b []byte) error {
	i, err := jsonplugin.ParseEnumString(string(b), Right_value)
	if err != nil {
		return err
	}
	*x = Right(i)
	return nil
}

// UnmarshalJSON unmarshals the Right from JSON.
func (x *Right) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the Rights message to JSON.
func (x *Rights) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.Rights) > 0 || s.HasField("rights") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("rights")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Rights {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the Rights to JSON.
func (x Rights) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the Rights message from JSON.
func (x *Rights) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "rights":
			s.AddField("rights")
			if s.ReadNil() {
				x.Rights = nil
				return
			}
			s.ReadArray(func() {
				var v Right
				v.UnmarshalProtoJSON(s)
				x.Rights = append(x.Rights, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the Rights from JSON.
func (x *Rights) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the APIKey message to JSON.
func (x *APIKey) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Id != "" || s.HasField("id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("id")
		s.WriteString(x.Id)
	}
	if x.Key != "" || s.HasField("key") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("key")
		s.WriteString(x.Key)
	}
	if x.Name != "" || s.HasField("name") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("name")
		s.WriteString(x.Name)
	}
	if len(x.Rights) > 0 || s.HasField("rights") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("rights")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Rights {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
	}
	if x.CreatedAt != nil || s.HasField("created_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("created_at")
		if x.CreatedAt == nil {
			s.WriteNil()
		} else {
			gogo.MarshalTimestamp(s, x.CreatedAt)
		}
	}
	if x.UpdatedAt != nil || s.HasField("updated_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("updated_at")
		if x.UpdatedAt == nil {
			s.WriteNil()
		} else {
			gogo.MarshalTimestamp(s, x.UpdatedAt)
		}
	}
	if x.ExpiresAt != nil || s.HasField("expires_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("expires_at")
		if x.ExpiresAt == nil {
			s.WriteNil()
		} else {
			gogo.MarshalTimestamp(s, x.ExpiresAt)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the APIKey to JSON.
func (x APIKey) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the APIKey message from JSON.
func (x *APIKey) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "id":
			s.AddField("id")
			x.Id = s.ReadString()
		case "key":
			s.AddField("key")
			x.Key = s.ReadString()
		case "name":
			s.AddField("name")
			x.Name = s.ReadString()
		case "rights":
			s.AddField("rights")
			if s.ReadNil() {
				x.Rights = nil
				return
			}
			s.ReadArray(func() {
				var v Right
				v.UnmarshalProtoJSON(s)
				x.Rights = append(x.Rights, v)
			})
		case "created_at", "createdAt":
			s.AddField("created_at")
			if s.ReadNil() {
				x.CreatedAt = nil
				return
			}
			v := gogo.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.CreatedAt = v
		case "updated_at", "updatedAt":
			s.AddField("updated_at")
			if s.ReadNil() {
				x.UpdatedAt = nil
				return
			}
			v := gogo.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.UpdatedAt = v
		case "expires_at", "expiresAt":
			s.AddField("expires_at")
			if s.ReadNil() {
				x.ExpiresAt = nil
				return
			}
			v := gogo.UnmarshalTimestamp(s)
			if s.Err() != nil {
				return
			}
			x.ExpiresAt = v
		}
	})
}

// UnmarshalJSON unmarshals the APIKey from JSON.
func (x *APIKey) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the APIKeys message to JSON.
func (x *APIKeys) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.ApiKeys) > 0 || s.HasField("api_keys") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("api_keys")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.ApiKeys {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("api_keys"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the APIKeys to JSON.
func (x APIKeys) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the APIKeys message from JSON.
func (x *APIKeys) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "api_keys", "apiKeys":
			s.AddField("api_keys")
			if s.ReadNil() {
				x.ApiKeys = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.ApiKeys = append(x.ApiKeys, nil)
					return
				}
				v := &APIKey{}
				v.UnmarshalProtoJSON(s.WithField("api_keys", false))
				if s.Err() != nil {
					return
				}
				x.ApiKeys = append(x.ApiKeys, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the APIKeys from JSON.
func (x *APIKeys) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the Collaborator message to JSON.
func (x *Collaborator) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Ids != nil || s.HasField("ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ids")
		// NOTE: OrganizationOrUserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.Ids)
	}
	if len(x.Rights) > 0 || s.HasField("rights") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("rights")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Rights {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the Collaborator to JSON.
func (x Collaborator) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the Collaborator message from JSON.
func (x *Collaborator) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "ids":
			s.AddField("ids")
			if s.ReadNil() {
				x.Ids = nil
				return
			}
			// NOTE: OrganizationOrUserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationOrUserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.Ids = &v
		case "rights":
			s.AddField("rights")
			if s.ReadNil() {
				x.Rights = nil
				return
			}
			s.ReadArray(func() {
				var v Right
				v.UnmarshalProtoJSON(s)
				x.Rights = append(x.Rights, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the Collaborator from JSON.
func (x *Collaborator) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the GetCollaboratorResponse message to JSON.
func (x *GetCollaboratorResponse) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Ids != nil || s.HasField("ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ids")
		// NOTE: OrganizationOrUserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.Ids)
	}
	if len(x.Rights) > 0 || s.HasField("rights") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("rights")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Rights {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GetCollaboratorResponse to JSON.
func (x GetCollaboratorResponse) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the GetCollaboratorResponse message from JSON.
func (x *GetCollaboratorResponse) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "ids":
			s.AddField("ids")
			if s.ReadNil() {
				x.Ids = nil
				return
			}
			// NOTE: OrganizationOrUserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationOrUserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.Ids = &v
		case "rights":
			s.AddField("rights")
			if s.ReadNil() {
				x.Rights = nil
				return
			}
			s.ReadArray(func() {
				var v Right
				v.UnmarshalProtoJSON(s)
				x.Rights = append(x.Rights, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the GetCollaboratorResponse from JSON.
func (x *GetCollaboratorResponse) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the Collaborators message to JSON.
func (x *Collaborators) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.Collaborators) > 0 || s.HasField("collaborators") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("collaborators")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Collaborators {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("collaborators"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the Collaborators to JSON.
func (x Collaborators) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the Collaborators message from JSON.
func (x *Collaborators) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "collaborators":
			s.AddField("collaborators")
			if s.ReadNil() {
				x.Collaborators = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.Collaborators = append(x.Collaborators, nil)
					return
				}
				v := &Collaborator{}
				v.UnmarshalProtoJSON(s.WithField("collaborators", false))
				if s.Err() != nil {
					return
				}
				x.Collaborators = append(x.Collaborators, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the Collaborators from JSON.
func (x *Collaborators) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
