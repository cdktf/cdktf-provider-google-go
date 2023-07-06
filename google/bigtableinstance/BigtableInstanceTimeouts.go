package bigtableinstance


type BigtableInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.72.1/docs/resources/bigtable_instance#create BigtableInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.72.1/docs/resources/bigtable_instance#update BigtableInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

