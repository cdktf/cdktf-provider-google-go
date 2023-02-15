package dataproccluster


type DataprocClusterClusterConfigEncryptionConfig struct {
	// The Cloud KMS key name to use for PD disk encryption for all instances in the cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster#kms_key_name DataprocCluster#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

