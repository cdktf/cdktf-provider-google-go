package datastreamstream


type DatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsDatasetTemplate struct {
	// The geographic location where the dataset should reside. See https://cloud.google.com/bigquery/docs/locations for supported locations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#location DatastreamStream#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// If supplied, every created dataset will have its name prefixed by the provided value.
	//
	// The prefix and name will be separated by an underscore. i.e. _.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#dataset_id_prefix DatastreamStream#dataset_id_prefix}
	DatasetIdPrefix *string `field:"optional" json:"datasetIdPrefix" yaml:"datasetIdPrefix"`
}

