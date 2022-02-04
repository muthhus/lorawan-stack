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

func New16BytesFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &ExactBytesValue{length: 16},
	}
}

func GetAESKey(fs *pflag.FlagSet, name string) (value types.AES128Key, set bool, err error) {
	flag := fs.Lookup(name)
	if flag == nil {
		return types.AES128Key{}, false, &flagsplugin.ErrFlagNotFound{FlagName: name}
	}
	var aes types.AES128Key
	if !flag.Changed {
		return aes, flag.Changed, nil
	}
	if err := aes.Unmarshal(flag.Value.(*ExactBytesValue).Value); err != nil {
		return aes, false, err
	}
	return aes, flag.Changed, nil
}
