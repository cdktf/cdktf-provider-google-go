package dataproccluster


type DataprocClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig struct {
	// The Hive Metastore configuration for this workload.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster#dataproc_metastore_service DataprocCluster#dataproc_metastore_service}
	DataprocMetastoreService *string `field:"optional" json:"dataprocMetastoreService" yaml:"dataprocMetastoreService"`
}

