package alloydbcluster


type AlloydbClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_cluster#create AlloydbCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_cluster#delete AlloydbCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_cluster#update AlloydbCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
