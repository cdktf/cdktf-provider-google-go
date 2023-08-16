package bigqueryjob


type BigqueryJobLoadTimePartitioning struct {
	// The only type supported is DAY, which will generate one partition per day.
	//
	// Providing an empty string used to cause an error,
	// but in OnePlatform the field will be treated as unset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_job#type BigqueryJob#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Number of milliseconds for which to keep the storage for a partition.
	//
	// A wrapper is used here because 0 is an invalid value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_job#expiration_ms BigqueryJob#expiration_ms}
	ExpirationMs *string `field:"optional" json:"expirationMs" yaml:"expirationMs"`
	// If not set, the table is partitioned by pseudo column '_PARTITIONTIME';
	//
	// if set, the table is partitioned by this field.
	// The field must be a top-level TIMESTAMP or DATE field. Its mode must be NULLABLE or REQUIRED.
	// A wrapper is used here because an empty string is an invalid value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_job#field BigqueryJob#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
}

