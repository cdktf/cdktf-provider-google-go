package sqldatabaseinstance


type SqlDatabaseInstanceSettingsSqlServerAuditConfig struct {
	// The name of the destination bucket (e.g., gs://mybucket).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#bucket SqlDatabaseInstance#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// How long to keep generated audit files.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"..
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#retention_interval SqlDatabaseInstance#retention_interval}
	RetentionInterval *string `field:"optional" json:"retentionInterval" yaml:"retentionInterval"`
	// How often to upload generated audit files.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#upload_interval SqlDatabaseInstance#upload_interval}
	UploadInterval *string `field:"optional" json:"uploadInterval" yaml:"uploadInterval"`
}

