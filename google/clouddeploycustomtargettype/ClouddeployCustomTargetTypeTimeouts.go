// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploycustomtargettype


type ClouddeployCustomTargetTypeTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/clouddeploy_custom_target_type#create ClouddeployCustomTargetType#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/clouddeploy_custom_target_type#delete ClouddeployCustomTargetType#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/clouddeploy_custom_target_type#update ClouddeployCustomTargetType#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

