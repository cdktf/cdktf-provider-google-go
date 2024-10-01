// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancefromtemplate


type ComputeInstanceFromTemplateServiceAccount struct {
	// A list of service scopes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_instance_from_template#scopes ComputeInstanceFromTemplate#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The service account e-mail address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_instance_from_template#email ComputeInstanceFromTemplate#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
}

