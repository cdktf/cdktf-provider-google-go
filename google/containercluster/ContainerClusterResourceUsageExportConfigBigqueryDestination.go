package containercluster


type ContainerClusterResourceUsageExportConfigBigqueryDestination struct {
	// The ID of a BigQuery Dataset.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#dataset_id ContainerCluster#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
}

