package sqldatabaseinstance


type SqlDatabaseInstanceSettings struct {
	// The machine type to use.
	//
	// See tiers for more details and supported versions. Postgres supports only shared-core machine types, and custom machine types such as db-custom-2-13312. See the Custom Machine Type Documentation to learn about specifying custom machine types.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#tier SqlDatabaseInstance#tier}
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// This specifies when the instance should be active. Can be either ALWAYS, NEVER or ON_DEMAND.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#activation_policy SqlDatabaseInstance#activation_policy}
	ActivationPolicy *string `field:"optional" json:"activationPolicy" yaml:"activationPolicy"`
	// active_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#active_directory_config SqlDatabaseInstance#active_directory_config}
	ActiveDirectoryConfig *SqlDatabaseInstanceSettingsActiveDirectoryConfig `field:"optional" json:"activeDirectoryConfig" yaml:"activeDirectoryConfig"`
	// The availability type of the Cloud SQL instance, high availability (REGIONAL) or single zone (ZONAL).
	//
	// For all instances, ensure that
	// settings.backup_configuration.enabled is set to true.
	// For MySQL instances, ensure that settings.backup_configuration.binary_log_enabled is set to true.
	// For Postgres instances, ensure that settings.backup_configuration.point_in_time_recovery_enabled
	// is set to true.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#availability_type SqlDatabaseInstance#availability_type}
	AvailabilityType *string `field:"optional" json:"availabilityType" yaml:"availabilityType"`
	// backup_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#backup_configuration SqlDatabaseInstance#backup_configuration}
	BackupConfiguration *SqlDatabaseInstanceSettingsBackupConfiguration `field:"optional" json:"backupConfiguration" yaml:"backupConfiguration"`
	// The name of server instance collation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#collation SqlDatabaseInstance#collation}
	Collation *string `field:"optional" json:"collation" yaml:"collation"`
	// database_flags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#database_flags SqlDatabaseInstance#database_flags}
	DatabaseFlags interface{} `field:"optional" json:"databaseFlags" yaml:"databaseFlags"`
	// Enables auto-resizing of the storage size. Defaults to true. Set to false if you want to set disk_size.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#disk_autoresize SqlDatabaseInstance#disk_autoresize}
	DiskAutoresize interface{} `field:"optional" json:"diskAutoresize" yaml:"diskAutoresize"`
	// The maximum size, in GB, to which storage capacity can be automatically increased.
	//
	// The default value is 0, which specifies that there is no limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#disk_autoresize_limit SqlDatabaseInstance#disk_autoresize_limit}
	DiskAutoresizeLimit *float64 `field:"optional" json:"diskAutoresizeLimit" yaml:"diskAutoresizeLimit"`
	// The size of data disk, in GB.
	//
	// Size of a running instance cannot be reduced but can be increased. If you want to set this field, set disk_autoresize to false.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#disk_size SqlDatabaseInstance#disk_size}
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// The type of data disk: PD_SSD or PD_HDD.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#disk_type SqlDatabaseInstance#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// insights_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#insights_config SqlDatabaseInstance#insights_config}
	InsightsConfig *SqlDatabaseInstanceSettingsInsightsConfig `field:"optional" json:"insightsConfig" yaml:"insightsConfig"`
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#ip_configuration SqlDatabaseInstance#ip_configuration}
	IpConfiguration *SqlDatabaseInstanceSettingsIpConfiguration `field:"optional" json:"ipConfiguration" yaml:"ipConfiguration"`
	// location_preference block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#location_preference SqlDatabaseInstance#location_preference}
	LocationPreference *SqlDatabaseInstanceSettingsLocationPreference `field:"optional" json:"locationPreference" yaml:"locationPreference"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#maintenance_window SqlDatabaseInstance#maintenance_window}
	MaintenanceWindow *SqlDatabaseInstanceSettingsMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// password_validation_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#password_validation_policy SqlDatabaseInstance#password_validation_policy}
	PasswordValidationPolicy *SqlDatabaseInstanceSettingsPasswordValidationPolicy `field:"optional" json:"passwordValidationPolicy" yaml:"passwordValidationPolicy"`
	// Pricing plan for this instance, can only be PER_USE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#pricing_plan SqlDatabaseInstance#pricing_plan}
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
	// sql_server_audit_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#sql_server_audit_config SqlDatabaseInstance#sql_server_audit_config}
	SqlServerAuditConfig *SqlDatabaseInstanceSettingsSqlServerAuditConfig `field:"optional" json:"sqlServerAuditConfig" yaml:"sqlServerAuditConfig"`
	// A set of key/value user label pairs to assign to the instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#user_labels SqlDatabaseInstance#user_labels}
	UserLabels *map[string]*string `field:"optional" json:"userLabels" yaml:"userLabels"`
}

