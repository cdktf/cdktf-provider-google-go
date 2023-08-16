package dataplexlake


type DataplexLakeMetastore struct {
	// Optional. A relative reference to the Dataproc Metastore (https://cloud.google.com/dataproc-metastore/docs) service associated with the lake: `projects/{project_id}/locations/{location_id}/services/{service_id}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_lake#service DataplexLake#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

