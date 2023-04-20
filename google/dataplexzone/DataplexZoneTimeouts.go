package dataplexzone


type DataplexZoneTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.1/docs/resources/dataplex_zone#create DataplexZone#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.1/docs/resources/dataplex_zone#delete DataplexZone#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.1/docs/resources/dataplex_zone#update DataplexZone#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

