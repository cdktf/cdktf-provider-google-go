package bigtablegcpolicy


type BigtableGcPolicyMaxVersion struct {
	// Number of version before applying the GC policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.75.1/docs/resources/bigtable_gc_policy#number BigtableGcPolicy#number}
	Number *float64 `field:"required" json:"number" yaml:"number"`
}

