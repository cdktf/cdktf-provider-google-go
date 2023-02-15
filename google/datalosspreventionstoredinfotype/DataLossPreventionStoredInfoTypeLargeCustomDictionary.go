package datalosspreventionstoredinfotype


type DataLossPreventionStoredInfoTypeLargeCustomDictionary struct {
	// output_path block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_stored_info_type#output_path DataLossPreventionStoredInfoType#output_path}
	OutputPath *DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath `field:"required" json:"outputPath" yaml:"outputPath"`
	// big_query_field block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_stored_info_type#big_query_field DataLossPreventionStoredInfoType#big_query_field}
	BigQueryField *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField `field:"optional" json:"bigQueryField" yaml:"bigQueryField"`
	// cloud_storage_file_set block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_stored_info_type#cloud_storage_file_set DataLossPreventionStoredInfoType#cloud_storage_file_set}
	CloudStorageFileSet *DataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet `field:"optional" json:"cloudStorageFileSet" yaml:"cloudStorageFileSet"`
}

