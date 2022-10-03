package bigqueryconnection


type BigqueryConnectionAws struct {
	// access_role block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_connection#access_role BigqueryConnection#access_role}
	AccessRole *BigqueryConnectionAwsAccessRole `field:"required" json:"accessRole" yaml:"accessRole"`
}

