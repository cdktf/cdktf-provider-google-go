package bigqueryconnection


type BigqueryConnectionAzure struct {
	// The id of customer's directory that host the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_connection#customer_tenant_id BigqueryConnection#customer_tenant_id}
	CustomerTenantId *string `field:"required" json:"customerTenantId" yaml:"customerTenantId"`
	// The Azure Application (client) ID where the federated credentials will be hosted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_connection#federated_application_client_id BigqueryConnection#federated_application_client_id}
	FederatedApplicationClientId *string `field:"optional" json:"federatedApplicationClientId" yaml:"federatedApplicationClientId"`
}

