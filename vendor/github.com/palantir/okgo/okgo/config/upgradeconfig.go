// Copyright 2016 Palantir Technologies, Inc.
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

package config

import (
	"github.com/palantir/godel/pkg/versionedconfig"
	"github.com/pkg/errors"

	"github.com/palantir/okgo/okgo"
	"github.com/palantir/okgo/okgo/config/internal/legacy"
	"github.com/palantir/okgo/okgo/config/internal/v0"
)

func UpgradeConfig(cfgBytes []byte, factory okgo.CheckerFactory) ([]byte, error) {
	if versionedconfig.IsLegacyConfig(cfgBytes) {
		v0Bytes, err := legacy.UpgradeConfig(cfgBytes, factory)
		if err != nil {
			return nil, err
		}
		cfgBytes = v0Bytes
	}

	version, err := versionedconfig.ConfigVersion(cfgBytes)
	if err != nil {
		return nil, err
	}
	switch version {
	case "", "0":
		return v0.UpgradeConfig(cfgBytes, factory)
	default:
		return nil, errors.Errorf("unsupported version: %s", version)
	}
}
