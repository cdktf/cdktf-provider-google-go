package databasemigrationserviceconnectionprofile


type DatabaseMigrationServiceConnectionProfileCloudsqlSettings struct {
	// The Database Migration Service source connection profile ID, in the format: projects/my_project_name/locations/us-central1/connectionProfiles/connection_profile_ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#source_id DatabaseMigrationServiceConnectionProfile#source_id}
	SourceId *string `field:"required" json:"sourceId" yaml:"sourceId"`
	// The activation policy specifies when the instance is activated;
	//
	// it is applicable only when the instance state is 'RUNNABLE'. Possible values: ["ALWAYS", "NEVER"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#activation_policy DatabaseMigrationServiceConnectionProfile#activation_policy}
	ActivationPolicy *string `field:"optional" json:"activationPolicy" yaml:"activationPolicy"`
	// If you enable this setting, Cloud SQL checks your available storage every 30 seconds.
	//
	// If the available storage falls below a threshold size, Cloud SQL automatically adds additional storage capacity.
	// If the available storage repeatedly falls below the threshold size, Cloud SQL continues to add storage until it reaches the maximum of 30 TB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#auto_storage_increase DatabaseMigrationServiceConnectionProfile#auto_storage_increase}
	AutoStorageIncrease interface{} `field:"optional" json:"autoStorageIncrease" yaml:"autoStorageIncrease"`
	// The KMS key name used for the csql instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#cmek_key_name DatabaseMigrationServiceConnectionProfile#cmek_key_name}
	CmekKeyName *string `field:"optional" json:"cmekKeyName" yaml:"cmekKeyName"`
	// The Cloud SQL default instance level collation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#collation DatabaseMigrationServiceConnectionProfile#collation}
	Collation *string `field:"optional" json:"collation" yaml:"collation"`
	// The database flags passed to the Cloud SQL instance at startup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#database_flags DatabaseMigrationServiceConnectionProfile#database_flags}
	DatabaseFlags *map[string]*string `field:"optional" json:"databaseFlags" yaml:"databaseFlags"`
	// The database engine type and version. Currently supported values located at https://cloud.google.com/database-migration/docs/reference/rest/v1/projects.locations.connectionProfiles#sqldatabaseversion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#database_version DatabaseMigrationServiceConnectionProfile#database_version}
	DatabaseVersion *string `field:"optional" json:"databaseVersion" yaml:"databaseVersion"`
	// The storage capacity available to the database, in GB. The minimum (and default) size is 10GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#data_disk_size_gb DatabaseMigrationServiceConnectionProfile#data_disk_size_gb}
	DataDiskSizeGb *string `field:"optional" json:"dataDiskSizeGb" yaml:"dataDiskSizeGb"`
	// The type of storage. Possible values: ["PD_SSD", "PD_HDD"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#data_disk_type DatabaseMigrationServiceConnectionProfile#data_disk_type}
	DataDiskType *string `field:"optional" json:"dataDiskType" yaml:"dataDiskType"`
	// ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#ip_config DatabaseMigrationServiceConnectionProfile#ip_config}
	IpConfig *DatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfig `field:"optional" json:"ipConfig" yaml:"ipConfig"`
	// Input only. Initial root password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#root_password DatabaseMigrationServiceConnectionProfile#root_password}
	RootPassword *string `field:"optional" json:"rootPassword" yaml:"rootPassword"`
	// The maximum size to which storage capacity can be automatically increased.
	//
	// The default value is 0, which specifies that there is no limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#storage_auto_resize_limit DatabaseMigrationServiceConnectionProfile#storage_auto_resize_limit}
	StorageAutoResizeLimit *string `field:"optional" json:"storageAutoResizeLimit" yaml:"storageAutoResizeLimit"`
	// The tier (or machine type) for this instance, for example: db-n1-standard-1 (MySQL instances) or db-custom-1-3840 (PostgreSQL instances).
	//
	// For more information, see https://cloud.google.com/sql/docs/mysql/instance-settings
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#tier DatabaseMigrationServiceConnectionProfile#tier}
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// The resource labels for a Cloud SQL instance to use to annotate any related underlying resources such as Compute Engine VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#user_labels DatabaseMigrationServiceConnectionProfile#user_labels}
	UserLabels *map[string]*string `field:"optional" json:"userLabels" yaml:"userLabels"`
	// The Google Cloud Platform zone where your Cloud SQL datdabse instance is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/database_migration_service_connection_profile#zone DatabaseMigrationServiceConnectionProfile#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

