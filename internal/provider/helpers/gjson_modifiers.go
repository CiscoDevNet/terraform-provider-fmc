// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package helpers

import (
	"encoding/json"
	"strings"

	"github.com/tidwall/gjson"
)

func init() {
	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "lower" {
			return strings.ToLower(json)
		}
		return json
	})

	// Usage: `value|@normalize_ip`, see NormalizeIP
	gjson.AddModifier("normalize_ip", func(jsonStr, arg string) string {
		res := gjson.Parse(jsonStr)
		if res.Type != gjson.String {
			return jsonStr
		}
		out, err := json.Marshal(NormalizeIP(res.String()))
		if err != nil {
			return jsonStr
		}
		return string(out)
	})
}
