package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper struct {
	// The location of the repository directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/os_config_os_policy_assignment#base_url OsConfigOsPolicyAssignment#base_url}
	BaseUrl *string `field:"required" json:"baseUrl" yaml:"baseUrl"`
	// A one word, unique name for this repository.
	//
	// This is the 'repo id' in the zypper config file and also the 'display_name' if 'display_name' is omitted. This id is also used as the unique identifier when checking for GuestPolicy conflicts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/os_config_os_policy_assignment#id OsConfigOsPolicyAssignment#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The display name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/os_config_os_policy_assignment#display_name OsConfigOsPolicyAssignment#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// URIs of GPG keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/os_config_os_policy_assignment#gpg_keys OsConfigOsPolicyAssignment#gpg_keys}
	GpgKeys *[]*string `field:"optional" json:"gpgKeys" yaml:"gpgKeys"`
}

