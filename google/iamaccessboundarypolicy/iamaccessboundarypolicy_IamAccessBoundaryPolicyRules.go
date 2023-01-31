package iamaccessboundarypolicy


type IamAccessBoundaryPolicyRules struct {
	// access_boundary_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iam_access_boundary_policy#access_boundary_rule IamAccessBoundaryPolicy#access_boundary_rule}
	AccessBoundaryRule *IamAccessBoundaryPolicyRulesAccessBoundaryRule `field:"optional" json:"accessBoundaryRule" yaml:"accessBoundaryRule"`
	// The description of the rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iam_access_boundary_policy#description IamAccessBoundaryPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

