// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package scripting

import (
	"context"
)

// Engine represents a scripting engine.
type Engine interface {
	Run(ctx context.Context, script, fn string, params ...interface{}) (func(target interface{}) error, error)
}

// AheadOfTimeEngine extends Engine with the capability of compiling the script ahead of time.
type AheadOfTimeEngine interface {
	Engine

	Compile(ctx context.Context, script string) (run func(context.Context, string, ...interface{}) (func(interface{}) error, error), err error)
}
