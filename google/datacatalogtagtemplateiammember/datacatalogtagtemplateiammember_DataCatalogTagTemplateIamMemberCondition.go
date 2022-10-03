package datacatalogtagtemplateiammember


type DataCatalogTagTemplateIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_tag_template_iam_member#expression DataCatalogTagTemplateIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_tag_template_iam_member#title DataCatalogTagTemplateIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_tag_template_iam_member#description DataCatalogTagTemplateIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

