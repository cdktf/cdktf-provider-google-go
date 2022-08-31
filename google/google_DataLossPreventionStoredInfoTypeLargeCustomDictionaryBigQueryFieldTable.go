// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable struct {
	// The dataset ID of the table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_stored_info_type#dataset_id DataLossPreventionStoredInfoType#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// The Google Cloud Platform project ID of the project containing the table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_stored_info_type#project_id DataLossPreventionStoredInfoType#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The name of the table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_stored_info_type#table_id DataLossPreventionStoredInfoType#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
}

