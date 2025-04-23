// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracledatabaseautonomousdatabase


type OracleDatabaseAutonomousDatabaseProperties struct {
	// Possible values:  DB_WORKLOAD_UNSPECIFIED OLTP DW AJD APEX.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#db_workload OracleDatabaseAutonomousDatabase#db_workload}
	DbWorkload *string `field:"required" json:"dbWorkload" yaml:"dbWorkload"`
	// The license type used for the Autonomous Database.   Possible values:  LICENSE_TYPE_UNSPECIFIED LICENSE_INCLUDED BRING_YOUR_OWN_LICENSE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#license_type OracleDatabaseAutonomousDatabase#license_type}
	LicenseType *string `field:"required" json:"licenseType" yaml:"licenseType"`
	// The retention period for the Autonomous Database.
	//
	// This field is specified
	// in days, can range from 1 day to 60 days, and has a default value of
	// 60 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#backup_retention_period_days OracleDatabaseAutonomousDatabase#backup_retention_period_days}
	BackupRetentionPeriodDays *float64 `field:"optional" json:"backupRetentionPeriodDays" yaml:"backupRetentionPeriodDays"`
	// The character set for the Autonomous Database. The default is AL32UTF8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#character_set OracleDatabaseAutonomousDatabase#character_set}
	CharacterSet *string `field:"optional" json:"characterSet" yaml:"characterSet"`
	// The number of compute servers for the Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#compute_count OracleDatabaseAutonomousDatabase#compute_count}
	ComputeCount *float64 `field:"optional" json:"computeCount" yaml:"computeCount"`
	// customer_contacts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#customer_contacts OracleDatabaseAutonomousDatabase#customer_contacts}
	CustomerContacts interface{} `field:"optional" json:"customerContacts" yaml:"customerContacts"`
	// The size of the data stored in the database, in gigabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#data_storage_size_gb OracleDatabaseAutonomousDatabase#data_storage_size_gb}
	DataStorageSizeGb *float64 `field:"optional" json:"dataStorageSizeGb" yaml:"dataStorageSizeGb"`
	// The size of the data stored in the database, in terabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#data_storage_size_tb OracleDatabaseAutonomousDatabase#data_storage_size_tb}
	DataStorageSizeTb *float64 `field:"optional" json:"dataStorageSizeTb" yaml:"dataStorageSizeTb"`
	// The edition of the Autonomous Databases.   Possible values:  DATABASE_EDITION_UNSPECIFIED STANDARD_EDITION ENTERPRISE_EDITION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#db_edition OracleDatabaseAutonomousDatabase#db_edition}
	DbEdition *string `field:"optional" json:"dbEdition" yaml:"dbEdition"`
	// The Oracle Database version for the Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#db_version OracleDatabaseAutonomousDatabase#db_version}
	DbVersion *string `field:"optional" json:"dbVersion" yaml:"dbVersion"`
	// This field indicates if auto scaling is enabled for the Autonomous Database CPU core count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#is_auto_scaling_enabled OracleDatabaseAutonomousDatabase#is_auto_scaling_enabled}
	IsAutoScalingEnabled interface{} `field:"optional" json:"isAutoScalingEnabled" yaml:"isAutoScalingEnabled"`
	// This field indicates if auto scaling is enabled for the Autonomous Database storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#is_storage_auto_scaling_enabled OracleDatabaseAutonomousDatabase#is_storage_auto_scaling_enabled}
	IsStorageAutoScalingEnabled interface{} `field:"optional" json:"isStorageAutoScalingEnabled" yaml:"isStorageAutoScalingEnabled"`
	// The maintenance schedule of the Autonomous Database.   Possible values:  MAINTENANCE_SCHEDULE_TYPE_UNSPECIFIED EARLY REGULAR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#maintenance_schedule_type OracleDatabaseAutonomousDatabase#maintenance_schedule_type}
	MaintenanceScheduleType *string `field:"optional" json:"maintenanceScheduleType" yaml:"maintenanceScheduleType"`
	// This field specifies if the Autonomous Database requires mTLS connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#mtls_connection_required OracleDatabaseAutonomousDatabase#mtls_connection_required}
	MtlsConnectionRequired interface{} `field:"optional" json:"mtlsConnectionRequired" yaml:"mtlsConnectionRequired"`
	// The national character set for the Autonomous Database. The default is AL16UTF16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#n_character_set OracleDatabaseAutonomousDatabase#n_character_set}
	NCharacterSet *string `field:"optional" json:"nCharacterSet" yaml:"nCharacterSet"`
	// Possible values:  OPERATIONS_INSIGHTS_STATE_UNSPECIFIED ENABLING ENABLED DISABLING NOT_ENABLED FAILED_ENABLING FAILED_DISABLING.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#operations_insights_state OracleDatabaseAutonomousDatabase#operations_insights_state}
	OperationsInsightsState *string `field:"optional" json:"operationsInsightsState" yaml:"operationsInsightsState"`
	// The private endpoint IP address for the Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#private_endpoint_ip OracleDatabaseAutonomousDatabase#private_endpoint_ip}
	PrivateEndpointIp *string `field:"optional" json:"privateEndpointIp" yaml:"privateEndpointIp"`
	// The private endpoint label for the Autonomous Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/oracle_database_autonomous_database#private_endpoint_label OracleDatabaseAutonomousDatabase#private_endpoint_label}
	PrivateEndpointLabel *string `field:"optional" json:"privateEndpointLabel" yaml:"privateEndpointLabel"`
}

