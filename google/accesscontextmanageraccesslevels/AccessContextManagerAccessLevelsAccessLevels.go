package accesscontextmanageraccesslevels


type AccessContextManagerAccessLevelsAccessLevels struct {
	// Resource name for the Access Level.
	//
	// The short_name component must begin
	// with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_access_levels#name AccessContextManagerAccessLevels#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Human readable title. Must be unique within the Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_access_levels#title AccessContextManagerAccessLevels#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// basic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_access_levels#basic AccessContextManagerAccessLevels#basic}
	Basic *AccessContextManagerAccessLevelsAccessLevelsBasic `field:"optional" json:"basic" yaml:"basic"`
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_access_levels#custom AccessContextManagerAccessLevels#custom}
	Custom *AccessContextManagerAccessLevelsAccessLevelsCustom `field:"optional" json:"custom" yaml:"custom"`
	// Description of the AccessLevel and its use. Does not affect behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_access_levels#description AccessContextManagerAccessLevels#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

