// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package custom_flags

import (
	"github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	"github.com/spf13/pflag"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

func New8BytesFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &ExactBytesValue{length: 8},
	}
}

func GetEUI64(fs *pflag.FlagSet, name string) (value types.EUI64, set bool, err error) {
	flag := fs.Lookup(name)
	if flag == nil {
		return types.EUI64{}, false, &flagsplugin.ErrFlagNotFound{FlagName: name}
	}
	var eui64 types.EUI64
	if !flag.Changed {
		return eui64, flag.Changed, nil
	}
	if err := eui64.Unmarshal(flag.Value.(*ExactBytesValue).Value); err != nil {
		return eui64, false, err
	}
	return eui64, flag.Changed, nil
}

func New8BytesSliceFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &ExactBytesSliceValue{length: 8},
	}
}

func GetEUI64Slice(fs *pflag.FlagSet, name string) (value []types.EUI64, set bool, err error) {
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &flagsplugin.ErrFlagNotFound{FlagName: name}
	}
	value = make([]types.EUI64, len(flag.Value.(*ExactBytesSliceValue).Values))
	for i, v := range flag.Value.(*ExactBytesSliceValue).Values {
		var eui64 types.EUI64
		if err := eui64.Unmarshal(v.Value); err != nil {
			return nil, false, err
		}
		value[i] = eui64
	}
	return value, flag.Changed, nil
}
