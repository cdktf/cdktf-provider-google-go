// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiSourceRemote struct {
	// Required.
	//
	// URI from which to fetch the object. It should contain both the protocol and path following the format `{protocol}://{location}`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_os_policy_assignment#uri OsConfigOsPolicyAssignment#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// SHA256 checksum of the remote file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_os_policy_assignment#sha256_checksum OsConfigOsPolicyAssignment#sha256_checksum}
	Sha256Checksum *string `field:"optional" json:"sha256Checksum" yaml:"sha256Checksum"`
}

