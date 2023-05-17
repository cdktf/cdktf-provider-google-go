package databasemigrationserviceconnectionprofile


type DatabaseMigrationServiceConnectionProfileMysql struct {
	// Required. The IP or hostname of the source MySQL database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#host DatabaseMigrationServiceConnectionProfile#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Required.
	//
	// Input only. The password for the user that Database Migration Service will be using to connect to the database.
	// This field is not returned on request, and the value is encrypted when stored in Database Migration Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#password DatabaseMigrationServiceConnectionProfile#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Required. The network port of the source MySQL database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#port DatabaseMigrationServiceConnectionProfile#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Required.
	//
	// The username that Database Migration Service will use to connect to the database. The value is encrypted when stored in Database Migration Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#username DatabaseMigrationServiceConnectionProfile#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// If the source is a Cloud SQL database, use this field to provide the Cloud SQL instance ID of the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#cloud_sql_id DatabaseMigrationServiceConnectionProfile#cloud_sql_id}
	CloudSqlId *string `field:"optional" json:"cloudSqlId" yaml:"cloudSqlId"`
	// ssl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#ssl DatabaseMigrationServiceConnectionProfile#ssl}
	Ssl *DatabaseMigrationServiceConnectionProfileMysqlSsl `field:"optional" json:"ssl" yaml:"ssl"`
}

