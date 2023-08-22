package bigquerydatatransferconfig


type BigqueryDataTransferConfigEmailPreferences struct {
	// If true, email notifications will be sent on transfer run failures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_data_transfer_config#enable_failure_email BigqueryDataTransferConfig#enable_failure_email}
	EnableFailureEmail interface{} `field:"required" json:"enableFailureEmail" yaml:"enableFailureEmail"`
}

