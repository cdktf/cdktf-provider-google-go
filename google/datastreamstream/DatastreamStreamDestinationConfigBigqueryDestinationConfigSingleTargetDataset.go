package datastreamstream


type DatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset struct {
	// Dataset ID in the format projects/{project}/datasets/{dataset_id}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#dataset_id DatastreamStream#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
}

