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

package provider

import (
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"fmc": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.

	username := os.Getenv("FMC_USERNAME")
	password := os.Getenv("FMC_PASSWORD")
	token := os.Getenv("FMC_TOKEN")

	if password == "" && token == "" {
		t.Fatal("For acceptance tests FMC (FMC_USERNAME and FMC_PASSWORD) or cdFMC (FMC_TOKEN) env variables must be set")
	}

	if password != "" && token != "" {
		t.Fatal("For acceptance tests only one of FMC (FMC_USERNAME and FMC_PASSWORD) or cdFMC (FMC_TOKEN) credentials must be set")
	}

	if password != "" && username == "" {
		t.Fatal("For acceptance tests FMC_USERNAME env variable must be set")
	}

	if v := os.Getenv("FMC_URL"); v == "" {
		t.Fatal("FMC_URL env variable must be set for acceptance tests")
	}
}

func testAccErrorCheck(t *testing.T, err error) error {
	if strings.Contains(err.Error(), "UnsupportedVersion") {
		t.Skip("UnsupportedVersion: Skipping check - resource not supported in the detected FMC version.")
		return nil
	}
	return err
}
