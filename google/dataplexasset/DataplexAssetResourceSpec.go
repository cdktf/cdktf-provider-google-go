package dataplexasset


type DataplexAssetResourceSpec struct {
	// Required. Immutable. Type of resource. Possible values: STORAGE_BUCKET, BIGQUERY_DATASET.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_asset#type DataplexAsset#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Immutable.
	//
	// Relative name of the cloud resource that contains the data that is being managed within a lake. For example: `projects/{project_number}/buckets/{bucket_id}` `projects/{project_number}/datasets/{dataset_id}`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_asset#name DataplexAsset#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

