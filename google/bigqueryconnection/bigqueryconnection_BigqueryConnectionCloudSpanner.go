package bigqueryconnection


type BigqueryConnectionCloudSpanner struct {
	// Cloud Spanner database in the form 'project/instance/database'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_connection#database BigqueryConnection#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// If parallelism should be used when reading from Cloud Spanner.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_connection#use_parallelism BigqueryConnection#use_parallelism}
	UseParallelism interface{} `field:"optional" json:"useParallelism" yaml:"useParallelism"`
}

