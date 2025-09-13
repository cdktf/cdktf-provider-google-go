// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedkafkacluster


type ManagedKafkaClusterTlsConfigTrustConfigCasConfigs struct {
	// The name of the CA pool to pull CA certificates from.
	//
	// The CA pool does not need to be in the same project or location as the Kafka cluster. Must be in the format 'projects/PROJECT_ID/locations/LOCATION/caPools/CA_POOL_ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/managed_kafka_cluster#ca_pool ManagedKafkaCluster#ca_pool}
	CaPool *string `field:"required" json:"caPool" yaml:"caPool"`
}

