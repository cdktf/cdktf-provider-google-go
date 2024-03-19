// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarefhirstore


type HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigLastUpdatedPartitionConfig struct {
	// Type of partitioning. Possible values: ["PARTITION_TYPE_UNSPECIFIED", "HOUR", "DAY", "MONTH", "YEAR"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/healthcare_fhir_store#type HealthcareFhirStore#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Number of milliseconds for which to keep the storage for a partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/healthcare_fhir_store#expiration_ms HealthcareFhirStore#expiration_ms}
	ExpirationMs *string `field:"optional" json:"expirationMs" yaml:"expirationMs"`
}

