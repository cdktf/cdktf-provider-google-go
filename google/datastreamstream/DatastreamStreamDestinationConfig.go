package datastreamstream


type DatastreamStreamDestinationConfig struct {
	// Destination connection profile resource. Format: projects/{project}/locations/{location}/connectionProfiles/{name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/datastream_stream#destination_connection_profile DatastreamStream#destination_connection_profile}
	DestinationConnectionProfile *string `field:"required" json:"destinationConnectionProfile" yaml:"destinationConnectionProfile"`
	// bigquery_destination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/datastream_stream#bigquery_destination_config DatastreamStream#bigquery_destination_config}
	BigqueryDestinationConfig *DatastreamStreamDestinationConfigBigqueryDestinationConfig `field:"optional" json:"bigqueryDestinationConfig" yaml:"bigqueryDestinationConfig"`
	// gcs_destination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/datastream_stream#gcs_destination_config DatastreamStream#gcs_destination_config}
	GcsDestinationConfig *DatastreamStreamDestinationConfigGcsDestinationConfig `field:"optional" json:"gcsDestinationConfig" yaml:"gcsDestinationConfig"`
}

