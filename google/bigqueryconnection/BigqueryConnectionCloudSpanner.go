package bigqueryconnection


type BigqueryConnectionCloudSpanner struct {
	// Cloud Spanner database in the form 'project/instance/database'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_connection#database BigqueryConnection#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// If parallelism should be used when reading from Cloud Spanner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_connection#use_parallelism BigqueryConnection#use_parallelism}
	UseParallelism interface{} `field:"optional" json:"useParallelism" yaml:"useParallelism"`
	// If the serverless analytics service should be used to read data from Cloud Spanner.
	//
	// useParallelism must be set when using serverless analytics
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_connection#use_serverless_analytics BigqueryConnection#use_serverless_analytics}
	UseServerlessAnalytics interface{} `field:"optional" json:"useServerlessAnalytics" yaml:"useServerlessAnalytics"`
}

