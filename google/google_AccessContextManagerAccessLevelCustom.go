// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type AccessContextManagerAccessLevelCustom struct {
	// expr block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_access_level#expr AccessContextManagerAccessLevel#expr}
	Expr *AccessContextManagerAccessLevelCustomExpr `field:"required" json:"expr" yaml:"expr"`
}

