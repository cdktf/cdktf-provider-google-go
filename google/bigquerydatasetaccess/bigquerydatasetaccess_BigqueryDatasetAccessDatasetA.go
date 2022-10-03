package bigquerydatasetaccess


type BigqueryDatasetAccessDatasetA struct {
	// dataset block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset_access#dataset BigqueryDatasetAccessA#dataset}
	Dataset *BigqueryDatasetAccessDatasetDatasetA `field:"required" json:"dataset" yaml:"dataset"`
	// Which resources in the dataset this entry applies to.
	//
	// Currently, only views are supported,
	// but additional target types may be added in the future. Possible values: VIEWS
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset_access#target_types BigqueryDatasetAccessA#target_types}
	TargetTypes *[]*string `field:"required" json:"targetTypes" yaml:"targetTypes"`
}

