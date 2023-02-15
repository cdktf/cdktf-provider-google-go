package bigtablegcpolicy


type BigtableGcPolicyMaxVersion struct {
	// Number of version before applying the GC policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_gc_policy#number BigtableGcPolicy#number}
	Number *float64 `field:"required" json:"number" yaml:"number"`
}

