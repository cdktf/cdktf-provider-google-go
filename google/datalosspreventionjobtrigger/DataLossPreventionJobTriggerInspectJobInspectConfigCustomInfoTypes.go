package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypes struct {
	// info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_job_trigger#info_type DataLossPreventionJobTrigger#info_type}
	InfoType *DataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesInfoType `field:"required" json:"infoType" yaml:"infoType"`
	// dictionary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_job_trigger#dictionary DataLossPreventionJobTrigger#dictionary}
	Dictionary *DataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary `field:"optional" json:"dictionary" yaml:"dictionary"`
	// If set to EXCLUSION_TYPE_EXCLUDE this infoType will not cause a finding to be returned.
	//
	// It still can be used for rules matching. Possible values: ["EXCLUSION_TYPE_EXCLUDE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_job_trigger#exclusion_type DataLossPreventionJobTrigger#exclusion_type}
	ExclusionType *string `field:"optional" json:"exclusionType" yaml:"exclusionType"`
	// Likelihood to return for this CustomInfoType.
	//
	// This base value can be altered by a detection rule if the finding meets the criteria
	// specified by the rule. Default value: "VERY_LIKELY" Possible values: ["VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_job_trigger#likelihood DataLossPreventionJobTrigger#likelihood}
	Likelihood *string `field:"optional" json:"likelihood" yaml:"likelihood"`
	// regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_job_trigger#regex DataLossPreventionJobTrigger#regex}
	Regex *DataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesRegex `field:"optional" json:"regex" yaml:"regex"`
	// sensitivity_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_job_trigger#sensitivity_score DataLossPreventionJobTrigger#sensitivity_score}
	SensitivityScore *DataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesSensitivityScore `field:"optional" json:"sensitivityScore" yaml:"sensitivityScore"`
	// stored_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_job_trigger#stored_type DataLossPreventionJobTrigger#stored_type}
	StoredType *DataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType `field:"optional" json:"storedType" yaml:"storedType"`
	// surrogate_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_job_trigger#surrogate_type DataLossPreventionJobTrigger#surrogate_type}
	SurrogateType *DataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType `field:"optional" json:"surrogateType" yaml:"surrogateType"`
}

