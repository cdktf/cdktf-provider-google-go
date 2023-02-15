package alloydbinstance


type AlloydbInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_instance#create AlloydbInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_instance#delete AlloydbInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_instance#update AlloydbInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

