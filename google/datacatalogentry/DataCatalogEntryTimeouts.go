package datacatalogentry


type DataCatalogEntryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_entry#create DataCatalogEntry#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_entry#delete DataCatalogEntry#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_entry#update DataCatalogEntry#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

