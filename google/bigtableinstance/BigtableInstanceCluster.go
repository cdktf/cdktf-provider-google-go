package bigtableinstance


type BigtableInstanceCluster struct {
	// The ID of the Cloud Bigtable cluster.
	//
	// Must be 6-30 characters and must only contain hyphens, lowercase letters and numbers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigtable_instance#cluster_id BigtableInstance#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// autoscaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigtable_instance#autoscaling_config BigtableInstance#autoscaling_config}
	AutoscalingConfig *BigtableInstanceClusterAutoscalingConfig `field:"optional" json:"autoscalingConfig" yaml:"autoscalingConfig"`
	// Describes the Cloud KMS encryption key that will be used to protect the destination Bigtable cluster.
	//
	// The requirements for this key are: 1) The Cloud Bigtable service account associated with the project that contains this cluster must be granted the cloudkms.cryptoKeyEncrypterDecrypter role on the CMEK key. 2) Only regional keys can be used and the region of the CMEK key must match the region of the cluster. 3) All clusters within an instance must use the same CMEK key. Values are of the form projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{key}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigtable_instance#kms_key_name BigtableInstance#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// The number of nodes in the cluster.
	//
	// If no value is set, Cloud Bigtable automatically allocates nodes based on your data footprint and optimized for 50% storage utilization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigtable_instance#num_nodes BigtableInstance#num_nodes}
	NumNodes *float64 `field:"optional" json:"numNodes" yaml:"numNodes"`
	// The storage type to use. One of "SSD" or "HDD". Defaults to "SSD".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigtable_instance#storage_type BigtableInstance#storage_type}
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// The zone to create the Cloud Bigtable cluster in.
	//
	// Each cluster must have a different zone in the same region. Zones that support Bigtable instances are noted on the Cloud Bigtable locations page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigtable_instance#zone BigtableInstance#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

