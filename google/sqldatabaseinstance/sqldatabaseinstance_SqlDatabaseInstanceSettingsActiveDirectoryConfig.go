package sqldatabaseinstance


type SqlDatabaseInstanceSettingsActiveDirectoryConfig struct {
	// Domain name of the Active Directory for SQL Server (e.g., mydomain.com).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#domain SqlDatabaseInstance#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

