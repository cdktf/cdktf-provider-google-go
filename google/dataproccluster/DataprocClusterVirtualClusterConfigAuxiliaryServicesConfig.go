package dataproccluster


type DataprocClusterVirtualClusterConfigAuxiliaryServicesConfig struct {
	// metastore_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster#metastore_config DataprocCluster#metastore_config}
	MetastoreConfig *DataprocClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig `field:"optional" json:"metastoreConfig" yaml:"metastoreConfig"`
	// spark_history_server_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster#spark_history_server_config DataprocCluster#spark_history_server_config}
	SparkHistoryServerConfig *DataprocClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig `field:"optional" json:"sparkHistoryServerConfig" yaml:"sparkHistoryServerConfig"`
}

