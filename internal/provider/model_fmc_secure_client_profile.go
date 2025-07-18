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

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SecureClientProfile struct {
	Id          types.String `tfsdk:"id"`
	Domain      types.String `tfsdk:"domain"`
	Name        types.String `tfsdk:"name"`
	FileName    types.String `tfsdk:"file_name"`
	Type        types.String `tfsdk:"type"`
	Description types.String `tfsdk:"description"`
	FileType    types.String `tfsdk:"file_type"`
	Path        types.String `tfsdk:"path"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SecureClientProfile) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/anyconnectprofiles"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SecureClientProfile) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("fileName"); value.Exists() {
		data.FileName = types.StringValue(value.String())
	} else {
		data.FileName = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("fileType"); value.Exists() {
		data.FileType = types.StringValue(value.String())
	} else {
		data.FileType = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SecureClientProfile) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("fileName"); value.Exists() && !data.FileName.IsNull() {
		data.FileName = types.StringValue(value.String())
	} else {
		data.FileName = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("fileType"); value.Exists() && !data.FileType.IsNull() {
		data.FileType = types.StringValue(value.String())
	} else {
		data.FileType = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SecureClientProfile) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.FileName.IsUnknown() {
		if value := res.Get("fileName"); value.Exists() {
			data.FileName = types.StringValue(value.String())
		} else {
			data.FileName = types.StringNull()
		}
	}
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// toBodyMultiPartUpload opens a file located at the Path, formats the contents as a MIME multipart/form-data
// body of a HTTP POST request (so, a file upload) and returns them as a stream.
// The second return is the corresponding Content-Type header value, and the third an error.
func (data *SecureClientProfile) toBodyMultiPartUpload(ctx context.Context) (io.Reader, string, error) {
	const dummyContentType = "application/octet-stream"

	source, err := os.Open(data.Path.ValueString())
	if err != nil {
		return nil, dummyContentType, err
	}

	tflog.Info(ctx, "swim file opened", map[string]any{"sourcePath": data.Path.ValueString()})

	readPipe, writePipe := io.Pipe()

	mw := multipart.NewWriter(writePipe)

	// data is copied in parallel, as soon as remote accepts it
	go func() {
		defer source.Close()

		// Create a form field for the user-defined name of the secure client image
		if err := mw.WriteField("name", data.Name.ValueString()); err != nil {
			panic(err)
		}

		// Create a form field for the file type
		if err := mw.WriteField("fileType", data.FileType.ValueString()); err != nil {
			panic(err)
		}

		if data.Id.ValueString() != "" {
			// Create a form field for the ID of the secure client image
			if err := mw.WriteField("id", data.Id.ValueString()); err != nil {
				panic(err)
			}
		}

		if data.Description.ValueString() != "" {
			// Create a form field for the description of the secure client image
			if err := mw.WriteField("description", data.Description.ValueString()); err != nil {
				panic(err)
			}
		}

		// Create a form file writer for the payloadFile field
		partWriter, err := mw.CreateFormFile("payloadFile", filepath.Base(data.Path.ValueString()))
		if err != nil {
			panic(err)
		}

		tflog.Info(ctx, "swim mw created")

		// Copy the file data to the form file writer
		var count int64
		if count, err = io.Copy(partWriter, source); err != nil {
			panic(err)
		}

		tflog.Info(ctx, "swim copied through pipe", map[string]any{"countBytes": count})

		// Close the multipart writer to get the terminating boundary

		if err := mw.Close(); err != nil {
			return
		}

		if err := writePipe.Close(); err != nil {
			return
		}

		tflog.Info(ctx, "swim: pipe closed for writing")
	}()

	return readPipe, mw.FormDataContentType(), nil
}
