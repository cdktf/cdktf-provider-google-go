package bigtablegcpolicy


type BigtableGcPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigtable_gc_policy#create BigtableGcPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigtable_gc_policy#delete BigtableGcPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

