package accesscontextmanageraccesslevel


type AccessContextManagerAccessLevelCustomExpr struct {
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_access_level#expression AccessContextManagerAccessLevel#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Description of the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_access_level#description AccessContextManagerAccessLevel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_access_level#location AccessContextManagerAccessLevel#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Title for the expression, i.e. a short string describing its purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_access_level#title AccessContextManagerAccessLevel#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

