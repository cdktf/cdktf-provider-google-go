package bigtabletableiammember


type BigtableTableIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.68.0/docs/resources/bigtable_table_iam_member#expression BigtableTableIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.68.0/docs/resources/bigtable_table_iam_member#title BigtableTableIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.68.0/docs/resources/bigtable_table_iam_member#description BigtableTableIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

