package dataprocmetastoreservice


type DataprocMetastoreServiceTelemetryConfig struct {
	// The output format of the Dataproc Metastore service's logs. Default value: "JSON" Possible values: ["LEGACY", "JSON"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_metastore_service#log_format DataprocMetastoreService#log_format}
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
}
