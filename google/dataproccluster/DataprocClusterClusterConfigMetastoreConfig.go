package dataproccluster


type DataprocClusterClusterConfigMetastoreConfig struct {
	// Resource name of an existing Dataproc Metastore service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster#dataproc_metastore_service DataprocCluster#dataproc_metastore_service}
	DataprocMetastoreService *string `field:"required" json:"dataprocMetastoreService" yaml:"dataprocMetastoreService"`
}

