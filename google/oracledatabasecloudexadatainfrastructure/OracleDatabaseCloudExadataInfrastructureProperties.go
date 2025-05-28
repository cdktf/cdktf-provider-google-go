// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracledatabasecloudexadatainfrastructure


type OracleDatabaseCloudExadataInfrastructureProperties struct {
	// The shape of the Exadata Infrastructure.
	//
	// The shape determines the
	// amount of CPU, storage, and memory resources allocated to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/oracle_database_cloud_exadata_infrastructure#shape OracleDatabaseCloudExadataInfrastructure#shape}
	Shape *string `field:"required" json:"shape" yaml:"shape"`
	// The number of compute servers for the Exadata Infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/oracle_database_cloud_exadata_infrastructure#compute_count OracleDatabaseCloudExadataInfrastructure#compute_count}
	ComputeCount *float64 `field:"optional" json:"computeCount" yaml:"computeCount"`
	// customer_contacts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/oracle_database_cloud_exadata_infrastructure#customer_contacts OracleDatabaseCloudExadataInfrastructure#customer_contacts}
	CustomerContacts interface{} `field:"optional" json:"customerContacts" yaml:"customerContacts"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/oracle_database_cloud_exadata_infrastructure#maintenance_window OracleDatabaseCloudExadataInfrastructure#maintenance_window}
	MaintenanceWindow *OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The number of Cloud Exadata storage servers for the Exadata Infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/oracle_database_cloud_exadata_infrastructure#storage_count OracleDatabaseCloudExadataInfrastructure#storage_count}
	StorageCount *float64 `field:"optional" json:"storageCount" yaml:"storageCount"`
	// The total storage allocated to the Exadata Infrastructure resource, in gigabytes (GB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/oracle_database_cloud_exadata_infrastructure#total_storage_size_gb OracleDatabaseCloudExadataInfrastructure#total_storage_size_gb}
	TotalStorageSizeGb *float64 `field:"optional" json:"totalStorageSizeGb" yaml:"totalStorageSizeGb"`
}

