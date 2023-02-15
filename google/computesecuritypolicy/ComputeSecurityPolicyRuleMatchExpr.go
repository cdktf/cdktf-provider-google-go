package computesecuritypolicy


type ComputeSecurityPolicyRuleMatchExpr struct {
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// The application context of the containing message determines which well-known feature set of CEL is supported.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_security_policy#expression ComputeSecurityPolicy#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

