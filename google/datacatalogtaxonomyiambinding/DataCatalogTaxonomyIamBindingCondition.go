package datacatalogtaxonomyiambinding


type DataCatalogTaxonomyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_taxonomy_iam_binding#expression DataCatalogTaxonomyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_taxonomy_iam_binding#title DataCatalogTaxonomyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_taxonomy_iam_binding#description DataCatalogTaxonomyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}
