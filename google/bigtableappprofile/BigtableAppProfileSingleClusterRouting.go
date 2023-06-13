package bigtableappprofile


type BigtableAppProfileSingleClusterRouting struct {
	// The cluster to which read/write requests should be routed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigtable_app_profile#cluster_id BigtableAppProfile#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// If true, CheckAndMutateRow and ReadModifyWriteRow requests are allowed by this app profile.
	//
	// It is unsafe to send these requests to the same table/row/column in multiple clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigtable_app_profile#allow_transactional_writes BigtableAppProfile#allow_transactional_writes}
	AllowTransactionalWrites interface{} `field:"optional" json:"allowTransactionalWrites" yaml:"allowTransactionalWrites"`
}

