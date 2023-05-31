package bigtablegcpolicy


type BigtableGcPolicyMaxVersion struct {
	// Number of version before applying the GC policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/bigtable_gc_policy#number BigtableGcPolicy#number}
	Number *float64 `field:"required" json:"number" yaml:"number"`
}

