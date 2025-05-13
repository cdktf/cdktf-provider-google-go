// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubhostprojectregistration


type ApihubHostProjectRegistrationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/apihub_host_project_registration#create ApihubHostProjectRegistration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/apihub_host_project_registration#delete ApihubHostProjectRegistration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

