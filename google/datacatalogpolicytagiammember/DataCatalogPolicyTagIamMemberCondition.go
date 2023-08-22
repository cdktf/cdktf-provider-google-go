package datacatalogpolicytagiammember


type DataCatalogPolicyTagIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_catalog_policy_tag_iam_member#expression DataCatalogPolicyTagIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_catalog_policy_tag_iam_member#title DataCatalogPolicyTagIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_catalog_policy_tag_iam_member#description DataCatalogPolicyTagIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

