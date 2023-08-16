package alloydbinstance


type AlloydbInstanceReadPoolConfig struct {
	// Read capacity, i.e. number of nodes in a read pool instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/alloydb_instance#node_count AlloydbInstance#node_count}
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
}

