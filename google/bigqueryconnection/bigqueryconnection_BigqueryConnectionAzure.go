package bigqueryconnection


type BigqueryConnectionAzure struct {
	// The id of customer's directory that host the data.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_connection#customer_tenant_id BigqueryConnection#customer_tenant_id}
	CustomerTenantId *string `field:"required" json:"customerTenantId" yaml:"customerTenantId"`
}

