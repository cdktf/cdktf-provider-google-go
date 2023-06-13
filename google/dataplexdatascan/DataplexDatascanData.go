package dataplexdatascan


type DataplexDatascanData struct {
	// The Dataplex entity that represents the data source(e.g. BigQuery table) for Datascan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan#entity DataplexDatascan#entity}
	Entity *string `field:"optional" json:"entity" yaml:"entity"`
	// The service-qualified full resource name of the cloud resource for a DataScan job to scan against.
	//
	// The field could be:
	// (Cloud Storage bucket for DataDiscoveryScan)BigQuery table of type "TABLE" for DataProfileScan/DataQualityScan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan#resource DataplexDatascan#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

