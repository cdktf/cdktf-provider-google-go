package bigqueryconnection


type BigqueryConnectionCloudSql struct {
	// credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_connection#credential BigqueryConnection#credential}
	Credential *BigqueryConnectionCloudSqlCredential `field:"required" json:"credential" yaml:"credential"`
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_connection#database BigqueryConnection#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Cloud SQL instance ID in the form project:location:instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_connection#instance_id BigqueryConnection#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Type of the Cloud SQL database. Possible values: ["DATABASE_TYPE_UNSPECIFIED", "POSTGRES", "MYSQL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_connection#type BigqueryConnection#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

