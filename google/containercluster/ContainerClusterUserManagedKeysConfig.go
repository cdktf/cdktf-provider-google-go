// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterUserManagedKeysConfig struct {
	// The Certificate Authority Service caPool to use for the aggreation CA in this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#aggregation_ca ContainerCluster#aggregation_ca}
	AggregationCa *string `field:"optional" json:"aggregationCa" yaml:"aggregationCa"`
	// The Certificate Authority Service caPool to use for the cluster CA in this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#cluster_ca ContainerCluster#cluster_ca}
	ClusterCa *string `field:"optional" json:"clusterCa" yaml:"clusterCa"`
	// The Cloud KMS cryptoKey to use for Confidential Hyperdisk on the control plane nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#control_plane_disk_encryption_key ContainerCluster#control_plane_disk_encryption_key}
	ControlPlaneDiskEncryptionKey *string `field:"optional" json:"controlPlaneDiskEncryptionKey" yaml:"controlPlaneDiskEncryptionKey"`
	// The Certificate Authority Service caPool to use for the etcd API CA in this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#etcd_api_ca ContainerCluster#etcd_api_ca}
	EtcdApiCa *string `field:"optional" json:"etcdApiCa" yaml:"etcdApiCa"`
	// The Certificate Authority Service caPool to use for the etcd peer CA in this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#etcd_peer_ca ContainerCluster#etcd_peer_ca}
	EtcdPeerCa *string `field:"optional" json:"etcdPeerCa" yaml:"etcdPeerCa"`
	// Resource path of the Cloud KMS cryptoKey to use for encryption of internal etcd backups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#gkeops_etcd_backup_encryption_key ContainerCluster#gkeops_etcd_backup_encryption_key}
	GkeopsEtcdBackupEncryptionKey *string `field:"optional" json:"gkeopsEtcdBackupEncryptionKey" yaml:"gkeopsEtcdBackupEncryptionKey"`
	// The Cloud KMS cryptoKeyVersions to use for signing service account JWTs issued by this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#service_account_signing_keys ContainerCluster#service_account_signing_keys}
	ServiceAccountSigningKeys *[]*string `field:"optional" json:"serviceAccountSigningKeys" yaml:"serviceAccountSigningKeys"`
	// The Cloud KMS cryptoKeyVersions to use for verifying service account JWTs issued by this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#service_account_verification_keys ContainerCluster#service_account_verification_keys}
	ServiceAccountVerificationKeys *[]*string `field:"optional" json:"serviceAccountVerificationKeys" yaml:"serviceAccountVerificationKeys"`
}

