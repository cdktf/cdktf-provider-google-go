package bigqueryjob


type BigqueryJobQueryUserDefinedFunctionResources struct {
	// An inline resource that contains code for a user-defined function (UDF).
	//
	// Providing a inline code resource is equivalent to providing a URI for a file containing the same code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_job#inline_code BigqueryJob#inline_code}
	InlineCode *string `field:"optional" json:"inlineCode" yaml:"inlineCode"`
	// A code resource to load from a Google Cloud Storage URI (gs://bucket/path).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_job#resource_uri BigqueryJob#resource_uri}
	ResourceUri *string `field:"optional" json:"resourceUri" yaml:"resourceUri"`
}

