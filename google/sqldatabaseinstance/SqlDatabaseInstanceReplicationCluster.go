// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqldatabaseinstance


type SqlDatabaseInstanceReplicationCluster struct {
	// If the instance is a primary instance, then this field identifies the disaster recovery (DR) replica.
	//
	// The standard format of this field is "your-project:your-instance". You can also set this field to "your-instance", but cloud SQL backend will convert it to the aforementioned standard format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/sql_database_instance#failover_dr_replica_name SqlDatabaseInstance#failover_dr_replica_name}
	FailoverDrReplicaName *string `field:"optional" json:"failoverDrReplicaName" yaml:"failoverDrReplicaName"`
	// If set, this field indicates this instance has a private service access (PSA) DNS endpoint that is pointing to the primary instance of the cluster.
	//
	// If this instance is the primary, then the DNS endpoint points to this instance. After a switchover or replica failover operation, this DNS endpoint points to the promoted instance. This is a read-only field, returned to the user as information. This field can exist even if a standalone instance doesn't have a DR replica yet or the DR replica is deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/sql_database_instance#psa_write_endpoint SqlDatabaseInstance#psa_write_endpoint}
	PsaWriteEndpoint *string `field:"optional" json:"psaWriteEndpoint" yaml:"psaWriteEndpoint"`
}

