package accesscontextmanageraccesslevel


type AccessContextManagerAccessLevelCustom struct {
	// expr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.72.1/docs/resources/access_context_manager_access_level#expr AccessContextManagerAccessLevel#expr}
	Expr *AccessContextManagerAccessLevelCustomExpr `field:"required" json:"expr" yaml:"expr"`
}

