package datalosspreventionstoredinfotype


type DataLossPreventionStoredInfoTypeDictionaryWordList struct {
	// Words or phrases defining the dictionary.
	//
	// The dictionary must contain at least one
	// phrase and every phrase must contain at least 2 characters that are letters or digits.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_stored_info_type#words DataLossPreventionStoredInfoType#words}
	Words *[]*string `field:"required" json:"words" yaml:"words"`
}

