package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputConfig struct {
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#table DataLossPreventionJobTrigger#table}
	Table *DataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputConfigTable `field:"required" json:"table" yaml:"table"`
	// Schema used for writing the findings for Inspect jobs.
	//
	// This field is only used for
	// Inspect and must be unspecified for Risk jobs. Columns are derived from the Finding
	// object. If appending to an existing table, any columns from the predefined schema
	// that are missing will be added. No columns in the existing table will be deleted.
	//
	// If unspecified, then all available columns will be used for a new table or an (existing)
	// table with no schema, and no changes will be made to an existing table that has a schema.
	// Only for use with external storage. Possible values: ["BASIC_COLUMNS", "GCS_COLUMNS", "DATASTORE_COLUMNS", "BIG_QUERY_COLUMNS", "ALL_COLUMNS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#output_schema DataLossPreventionJobTrigger#output_schema}
	OutputSchema *string `field:"optional" json:"outputSchema" yaml:"outputSchema"`
}

