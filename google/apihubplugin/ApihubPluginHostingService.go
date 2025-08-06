// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubplugin


type ApihubPluginHostingService struct {
	// The URI of the service implemented by the plugin developer, used to invoke the plugin's functionality.
	//
	// This information is only required for
	// user defined plugins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/apihub_plugin#service_uri ApihubPlugin#service_uri}
	ServiceUri *string `field:"optional" json:"serviceUri" yaml:"serviceUri"`
}

