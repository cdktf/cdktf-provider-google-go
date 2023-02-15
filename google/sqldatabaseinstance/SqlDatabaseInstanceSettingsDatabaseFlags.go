package sqldatabaseinstance


type SqlDatabaseInstanceSettingsDatabaseFlags struct {
	// Name of the flag.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#name SqlDatabaseInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the flag.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#value SqlDatabaseInstance#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

