// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancefromtemplate


type ComputeInstanceFromTemplateServiceAccount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.23.0/docs/resources/compute_instance_from_template#email ComputeInstanceFromTemplate#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.23.0/docs/resources/compute_instance_from_template#scopes ComputeInstanceFromTemplate#scopes}.
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

