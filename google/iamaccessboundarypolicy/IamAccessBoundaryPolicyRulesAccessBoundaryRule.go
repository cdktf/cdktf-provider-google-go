package iamaccessboundarypolicy


type IamAccessBoundaryPolicyRulesAccessBoundaryRule struct {
	// availability_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_access_boundary_policy#availability_condition IamAccessBoundaryPolicy#availability_condition}
	AvailabilityCondition *IamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition `field:"optional" json:"availabilityCondition" yaml:"availabilityCondition"`
	// A list of permissions that may be allowed for use on the specified resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_access_boundary_policy#available_permissions IamAccessBoundaryPolicy#available_permissions}
	AvailablePermissions *[]*string `field:"optional" json:"availablePermissions" yaml:"availablePermissions"`
	// The full resource name of a Google Cloud resource entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iam_access_boundary_policy#available_resource IamAccessBoundaryPolicy#available_resource}
	AvailableResource *string `field:"optional" json:"availableResource" yaml:"availableResource"`
}

