// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubplugin


type ApihubPluginDocumentation struct {
	// The uri of the externally hosted documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/apihub_plugin#external_uri ApihubPlugin#external_uri}
	ExternalUri *string `field:"optional" json:"externalUri" yaml:"externalUri"`
}

