package datacatalogentrygroupiammember


type DataCatalogEntryGroupIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/data_catalog_entry_group_iam_member#expression DataCatalogEntryGroupIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/data_catalog_entry_group_iam_member#title DataCatalogEntryGroupIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/data_catalog_entry_group_iam_member#description DataCatalogEntryGroupIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

