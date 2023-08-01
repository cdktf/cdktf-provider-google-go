package datastreamstream


type DatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset struct {
	// Dataset ID in the format projects/{project}/datasets/{dataset_id} or {project}:{dataset_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/datastream_stream#dataset_id DatastreamStream#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
}

