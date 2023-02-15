package bigtableinstanceiambinding


type BigtableInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_instance_iam_binding#expression BigtableInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_instance_iam_binding#title BigtableInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_instance_iam_binding#description BigtableInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

