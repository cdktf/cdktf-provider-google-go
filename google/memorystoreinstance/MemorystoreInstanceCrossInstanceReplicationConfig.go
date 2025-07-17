// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorystoreinstance


type MemorystoreInstanceCrossInstanceReplicationConfig struct {
	// The instance role supports the following values: 1.
	//
	// 'INSTANCE_ROLE_UNSPECIFIED': This is an independent instance that has never participated in cross instance replication. It allows both reads and writes.
	// 2. 'NONE': This is an independent instance that previously participated in cross instance replication(either as a 'PRIMARY' or 'SECONDARY' cluster). It allows both reads and writes.
	// 3. 'PRIMARY': This instance serves as the replication source for secondary instance that are replicating from it. Any data written to it is automatically replicated to its secondary clusters. It allows both reads and writes.
	// 4. 'SECONDARY': This instance replicates data from the primary instance. It allows only reads. Possible values: ["INSTANCE_ROLE_UNSPECIFIED", "NONE", "PRIMARY", "SECONDARY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/memorystore_instance#instance_role MemorystoreInstance#instance_role}
	InstanceRole *string `field:"optional" json:"instanceRole" yaml:"instanceRole"`
	// primary_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/memorystore_instance#primary_instance MemorystoreInstance#primary_instance}
	PrimaryInstance *MemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstance `field:"optional" json:"primaryInstance" yaml:"primaryInstance"`
	// secondary_instances block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/memorystore_instance#secondary_instances MemorystoreInstance#secondary_instances}
	SecondaryInstances interface{} `field:"optional" json:"secondaryInstances" yaml:"secondaryInstances"`
}

