package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobStorageConfigHybridOptions struct {
	// A short description of where the data is coming from.
	//
	// Will be stored once in the job. 256 max length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#description DataLossPreventionJobTrigger#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// To organize findings, these labels will be added to each finding.
	//
	// Label keys must be between 1 and 63 characters long and must conform to the following regular expression: '[a-z]([-a-z0-9]*[a-z0-9])?'.
	//
	// Label values must be between 0 and 63 characters long and must conform to the regular expression '([a-z]([-a-z0-9]*[a-z0-9])?)?'.
	//
	// No more than 10 labels can be associated with a given finding.
	//
	// Examples:
	// '"environment" : "production"'
	// '"pipeline" : "etl"'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#labels DataLossPreventionJobTrigger#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// These are labels that each inspection request must include within their 'finding_labels' map.
	//
	// Request
	// may contain others, but any missing one of these will be rejected.
	//
	// Label keys must be between 1 and 63 characters long and must conform to the following regular expression: '[a-z]([-a-z0-9]*[a-z0-9])?'.
	//
	// No more than 10 keys can be required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#required_finding_label_keys DataLossPreventionJobTrigger#required_finding_label_keys}
	RequiredFindingLabelKeys *[]*string `field:"optional" json:"requiredFindingLabelKeys" yaml:"requiredFindingLabelKeys"`
	// table_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#table_options DataLossPreventionJobTrigger#table_options}
	TableOptions *DataLossPreventionJobTriggerInspectJobStorageConfigHybridOptionsTableOptions `field:"optional" json:"tableOptions" yaml:"tableOptions"`
}

