package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary struct {
	// cloud_storage_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#cloud_storage_path DataLossPreventionJobTrigger#cloud_storage_path}
	CloudStoragePath *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath `field:"optional" json:"cloudStoragePath" yaml:"cloudStoragePath"`
	// word_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#word_list DataLossPreventionJobTrigger#word_list}
	WordList *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList `field:"optional" json:"wordList" yaml:"wordList"`
}

