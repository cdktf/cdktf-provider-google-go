package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionary struct {
	// cloud_storage_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.0/docs/resources/data_loss_prevention_job_trigger#cloud_storage_path DataLossPreventionJobTrigger#cloud_storage_path}
	CloudStoragePath *DataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath `field:"optional" json:"cloudStoragePath" yaml:"cloudStoragePath"`
	// word_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.0/docs/resources/data_loss_prevention_job_trigger#word_list DataLossPreventionJobTrigger#word_list}
	WordList *DataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList `field:"optional" json:"wordList" yaml:"wordList"`
}
