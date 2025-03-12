// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqldatabaseinstance


type SqlDatabaseInstanceReplicationCluster struct {
	// If the instance is a primary instance, then this field identifies the disaster recovery (DR) replica.
	//
	// The standard format of this field is "your-project:your-instance". You can also set this field to "your-instance", but cloud SQL backend will convert it to the aforementioned standard format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/sql_database_instance#failover_dr_replica_name SqlDatabaseInstance#failover_dr_replica_name}
	FailoverDrReplicaName *string `field:"optional" json:"failoverDrReplicaName" yaml:"failoverDrReplicaName"`
}

