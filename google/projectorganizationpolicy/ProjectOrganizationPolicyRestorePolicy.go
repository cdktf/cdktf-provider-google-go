// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectorganizationpolicy


type ProjectOrganizationPolicyRestorePolicy struct {
	// May only be set to true. If set, then the default Policy is restored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.81.0/docs/resources/project_organization_policy#default ProjectOrganizationPolicy#default}
	Default interface{} `field:"required" json:"default" yaml:"default"`
}

