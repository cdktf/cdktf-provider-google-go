package bigtableinstanceiammember


type BigtableInstanceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_instance_iam_member#expression BigtableInstanceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_instance_iam_member#title BigtableInstanceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_instance_iam_member#description BigtableInstanceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

