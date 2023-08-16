package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionary struct {
	// cloud_storage_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_inspect_template#cloud_storage_path DataLossPreventionInspectTemplate#cloud_storage_path}
	CloudStoragePath *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath `field:"optional" json:"cloudStoragePath" yaml:"cloudStoragePath"`
	// word_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_inspect_template#word_list DataLossPreventionInspectTemplate#word_list}
	WordList *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListStruct `field:"optional" json:"wordList" yaml:"wordList"`
}

