// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecValidateFileRemote struct {
	// URI from which to fetch the object. It should contain both the protocol and path following the format '{protocol}://{location}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/os_config_os_policy_assignment#uri OsConfigOsPolicyAssignment#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// SHA256 checksum of the remote file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/os_config_os_policy_assignment#sha256_checksum OsConfigOsPolicyAssignment#sha256_checksum}
	Sha256Checksum *string `field:"optional" json:"sha256Checksum" yaml:"sha256Checksum"`
}

