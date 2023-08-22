package datastreamstream


type DatastreamStreamDestinationConfigBigqueryDestinationConfig struct {
	// The guaranteed data freshness (in seconds) when querying tables created by the stream.
	//
	// Editing this field will only affect new tables created in the future, but existing tables
	// will not be impacted. Lower values mean that queries will return fresher data, but may result in higher cost.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". Defaults to 900s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/datastream_stream#data_freshness DatastreamStream#data_freshness}
	DataFreshness *string `field:"optional" json:"dataFreshness" yaml:"dataFreshness"`
	// single_target_dataset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/datastream_stream#single_target_dataset DatastreamStream#single_target_dataset}
	SingleTargetDataset *DatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset `field:"optional" json:"singleTargetDataset" yaml:"singleTargetDataset"`
	// source_hierarchy_datasets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/datastream_stream#source_hierarchy_datasets DatastreamStream#source_hierarchy_datasets}
	SourceHierarchyDatasets *DatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets `field:"optional" json:"sourceHierarchyDatasets" yaml:"sourceHierarchyDatasets"`
}

