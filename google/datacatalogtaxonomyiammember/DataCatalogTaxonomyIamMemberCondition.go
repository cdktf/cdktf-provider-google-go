package datacatalogtaxonomyiammember


type DataCatalogTaxonomyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_taxonomy_iam_member#expression DataCatalogTaxonomyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_taxonomy_iam_member#title DataCatalogTaxonomyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_taxonomy_iam_member#description DataCatalogTaxonomyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

