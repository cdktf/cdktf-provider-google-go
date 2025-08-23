// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubplugininstance


type ApihubPluginInstanceActionsCurationConfigCustomCuration struct {
	// The unique name of the curation resource. This will be the name of the curation resource in the format: 'projects/{project}/locations/{location}/curations/{curation}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apihub_plugin_instance#curation ApihubPluginInstance#curation}
	Curation *string `field:"required" json:"curation" yaml:"curation"`
}

