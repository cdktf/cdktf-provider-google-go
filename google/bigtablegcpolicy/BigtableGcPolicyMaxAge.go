package bigtablegcpolicy


type BigtableGcPolicyMaxAge struct {
	// Number of days before applying GC policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_gc_policy#days BigtableGcPolicy#days}
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Duration before applying GC policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_gc_policy#duration BigtableGcPolicy#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
}

