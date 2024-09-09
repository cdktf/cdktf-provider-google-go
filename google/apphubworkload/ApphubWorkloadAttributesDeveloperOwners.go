// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apphubworkload


type ApphubWorkloadAttributesDeveloperOwners struct {
	// Email address of the contacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/apphub_workload#email ApphubWorkload#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Contact's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/apphub_workload#display_name ApphubWorkload#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

