package cloudbuildv2connectioniambinding


type Cloudbuildv2ConnectionIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuildv2_connection_iam_binding#expression Cloudbuildv2ConnectionIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuildv2_connection_iam_binding#title Cloudbuildv2ConnectionIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuildv2_connection_iam_binding#description Cloudbuildv2ConnectionIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

