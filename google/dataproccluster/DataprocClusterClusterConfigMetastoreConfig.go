package dataproccluster


type DataprocClusterClusterConfigMetastoreConfig struct {
	// Resource name of an existing Dataproc Metastore service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_cluster#dataproc_metastore_service DataprocCluster#dataproc_metastore_service}
	DataprocMetastoreService *string `field:"required" json:"dataprocMetastoreService" yaml:"dataprocMetastoreService"`
}

