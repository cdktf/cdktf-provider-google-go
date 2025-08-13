// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedkafkacluster


type ManagedKafkaClusterGcpConfig struct {
	// access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/managed_kafka_cluster#access_config ManagedKafkaCluster#access_config}
	AccessConfig *ManagedKafkaClusterGcpConfigAccessConfig `field:"required" json:"accessConfig" yaml:"accessConfig"`
	// The Cloud KMS Key name to use for encryption.
	//
	// The key must be located in the same region as the cluster and cannot be changed. Must be in the format 'projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/managed_kafka_cluster#kms_key ManagedKafkaCluster#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

