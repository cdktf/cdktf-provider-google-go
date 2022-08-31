// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type BigqueryDataTransferConfigEmailPreferences struct {
	// If true, email notifications will be sent on transfer run failures.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_data_transfer_config#enable_failure_email BigqueryDataTransferConfig#enable_failure_email}
	EnableFailureEmail interface{} `field:"required" json:"enableFailureEmail" yaml:"enableFailureEmail"`
}

