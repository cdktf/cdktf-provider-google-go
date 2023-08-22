package datacatalogentrygroupiambinding


type DataCatalogEntryGroupIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_catalog_entry_group_iam_binding#expression DataCatalogEntryGroupIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_catalog_entry_group_iam_binding#title DataCatalogEntryGroupIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_catalog_entry_group_iam_binding#description DataCatalogEntryGroupIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

