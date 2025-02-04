// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploytarget


type ClouddeployTargetCustomTarget struct {
	// Required. The name of the CustomTargetType. Format must be `projects/{project}/locations/{location}/customTargetTypes/{custom_target_type}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/clouddeploy_target#custom_target_type ClouddeployTarget#custom_target_type}
	CustomTargetType *string `field:"required" json:"customTargetType" yaml:"customTargetType"`
}

