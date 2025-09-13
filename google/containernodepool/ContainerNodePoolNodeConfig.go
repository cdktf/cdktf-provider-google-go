// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeConfig struct {
	// advanced_machine_features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#advanced_machine_features ContainerNodePool#advanced_machine_features}
	AdvancedMachineFeatures *ContainerNodePoolNodeConfigAdvancedMachineFeatures `field:"optional" json:"advancedMachineFeatures" yaml:"advancedMachineFeatures"`
	// boot_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#boot_disk ContainerNodePool#boot_disk}
	BootDisk *ContainerNodePoolNodeConfigBootDisk `field:"optional" json:"bootDisk" yaml:"bootDisk"`
	// The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#boot_disk_kms_key ContainerNodePool#boot_disk_kms_key}
	BootDiskKmsKey *string `field:"optional" json:"bootDiskKmsKey" yaml:"bootDiskKmsKey"`
	// confidential_nodes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#confidential_nodes ContainerNodePool#confidential_nodes}
	ConfidentialNodes *ContainerNodePoolNodeConfigConfidentialNodes `field:"optional" json:"confidentialNodes" yaml:"confidentialNodes"`
	// containerd_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#containerd_config ContainerNodePool#containerd_config}
	ContainerdConfig *ContainerNodePoolNodeConfigContainerdConfig `field:"optional" json:"containerdConfig" yaml:"containerdConfig"`
	// Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#disk_size_gb ContainerNodePool#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#disk_type ContainerNodePool#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// If enabled boot disks are configured with confidential mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#enable_confidential_storage ContainerNodePool#enable_confidential_storage}
	EnableConfidentialStorage interface{} `field:"optional" json:"enableConfidentialStorage" yaml:"enableConfidentialStorage"`
	// ephemeral_storage_local_ssd_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#ephemeral_storage_local_ssd_config ContainerNodePool#ephemeral_storage_local_ssd_config}
	EphemeralStorageLocalSsdConfig *ContainerNodePoolNodeConfigEphemeralStorageLocalSsdConfig `field:"optional" json:"ephemeralStorageLocalSsdConfig" yaml:"ephemeralStorageLocalSsdConfig"`
	// fast_socket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#fast_socket ContainerNodePool#fast_socket}
	FastSocket *ContainerNodePoolNodeConfigFastSocket `field:"optional" json:"fastSocket" yaml:"fastSocket"`
	// Enables Flex Start provisioning model for the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#flex_start ContainerNodePool#flex_start}
	FlexStart interface{} `field:"optional" json:"flexStart" yaml:"flexStart"`
	// gcfs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#gcfs_config ContainerNodePool#gcfs_config}
	GcfsConfig *ContainerNodePoolNodeConfigGcfsConfig `field:"optional" json:"gcfsConfig" yaml:"gcfsConfig"`
	// guest_accelerator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#guest_accelerator ContainerNodePool#guest_accelerator}
	GuestAccelerator interface{} `field:"optional" json:"guestAccelerator" yaml:"guestAccelerator"`
	// gvnic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#gvnic ContainerNodePool#gvnic}
	Gvnic *ContainerNodePoolNodeConfigGvnic `field:"optional" json:"gvnic" yaml:"gvnic"`
	// host_maintenance_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#host_maintenance_policy ContainerNodePool#host_maintenance_policy}
	HostMaintenancePolicy *ContainerNodePoolNodeConfigHostMaintenancePolicy `field:"optional" json:"hostMaintenancePolicy" yaml:"hostMaintenancePolicy"`
	// The image type to use for this node.
	//
	// Note that for a given image type, the latest version of it will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#image_type ContainerNodePool#image_type}
	ImageType *string `field:"optional" json:"imageType" yaml:"imageType"`
	// kubelet_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#kubelet_config ContainerNodePool#kubelet_config}
	KubeletConfig *ContainerNodePoolNodeConfigKubeletConfig `field:"optional" json:"kubeletConfig" yaml:"kubeletConfig"`
	// The map of Kubernetes labels (key/value pairs) to be applied to each node.
	//
	// These will added in addition to any default label(s) that Kubernetes may apply to the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#labels ContainerNodePool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// linux_node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#linux_node_config ContainerNodePool#linux_node_config}
	LinuxNodeConfig *ContainerNodePoolNodeConfigLinuxNodeConfig `field:"optional" json:"linuxNodeConfig" yaml:"linuxNodeConfig"`
	// local_nvme_ssd_block_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#local_nvme_ssd_block_config ContainerNodePool#local_nvme_ssd_block_config}
	LocalNvmeSsdBlockConfig *ContainerNodePoolNodeConfigLocalNvmeSsdBlockConfig `field:"optional" json:"localNvmeSsdBlockConfig" yaml:"localNvmeSsdBlockConfig"`
	// The number of local SSD disks to be attached to the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#local_ssd_count ContainerNodePool#local_ssd_count}
	LocalSsdCount *float64 `field:"optional" json:"localSsdCount" yaml:"localSsdCount"`
	// LocalSsdEncryptionMode specified the method used for encrypting the local SSDs attached to the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#local_ssd_encryption_mode ContainerNodePool#local_ssd_encryption_mode}
	LocalSsdEncryptionMode *string `field:"optional" json:"localSsdEncryptionMode" yaml:"localSsdEncryptionMode"`
	// Type of logging agent that is used as the default value for node pools in the cluster.
	//
	// Valid values include DEFAULT and MAX_THROUGHPUT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#logging_variant ContainerNodePool#logging_variant}
	LoggingVariant *string `field:"optional" json:"loggingVariant" yaml:"loggingVariant"`
	// The name of a Google Compute Engine machine type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#machine_type ContainerNodePool#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// The runtime of each node in the node pool in seconds, terminated by 's'. Example: "3600s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#max_run_duration ContainerNodePool#max_run_duration}
	MaxRunDuration *string `field:"optional" json:"maxRunDuration" yaml:"maxRunDuration"`
	// The metadata key/value pairs assigned to instances in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#metadata ContainerNodePool#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Minimum CPU platform to be used by this instance.
	//
	// The instance may be scheduled on the specified or newer CPU platform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#min_cpu_platform ContainerNodePool#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
	// Setting this field will assign instances of this pool to run on the specified node group.
	//
	// This is useful for running workloads on sole tenant nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#node_group ContainerNodePool#node_group}
	NodeGroup *string `field:"optional" json:"nodeGroup" yaml:"nodeGroup"`
	// The set of Google API scopes to be made available on all of the node VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#oauth_scopes ContainerNodePool#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
	// Whether the nodes are created as preemptible VM instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#preemptible ContainerNodePool#preemptible}
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// reservation_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#reservation_affinity ContainerNodePool#reservation_affinity}
	ReservationAffinity *ContainerNodePoolNodeConfigReservationAffinity `field:"optional" json:"reservationAffinity" yaml:"reservationAffinity"`
	// The GCE resource labels (a map of key/value pairs) to be applied to the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#resource_labels ContainerNodePool#resource_labels}
	ResourceLabels *map[string]*string `field:"optional" json:"resourceLabels" yaml:"resourceLabels"`
	// A map of resource manager tags.
	//
	// Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#resource_manager_tags ContainerNodePool#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
	// secondary_boot_disks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#secondary_boot_disks ContainerNodePool#secondary_boot_disks}
	SecondaryBootDisks interface{} `field:"optional" json:"secondaryBootDisks" yaml:"secondaryBootDisks"`
	// The Google Cloud Platform Service Account to be used by the node VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#service_account ContainerNodePool#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#shielded_instance_config ContainerNodePool#shielded_instance_config}
	ShieldedInstanceConfig *ContainerNodePoolNodeConfigShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// sole_tenant_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#sole_tenant_config ContainerNodePool#sole_tenant_config}
	SoleTenantConfig *ContainerNodePoolNodeConfigSoleTenantConfig `field:"optional" json:"soleTenantConfig" yaml:"soleTenantConfig"`
	// Whether the nodes are created as spot VM instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#spot ContainerNodePool#spot}
	Spot interface{} `field:"optional" json:"spot" yaml:"spot"`
	// The list of Storage Pools where boot disks are provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#storage_pools ContainerNodePool#storage_pools}
	StoragePools *[]*string `field:"optional" json:"storagePools" yaml:"storagePools"`
	// The list of instance tags applied to all nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#tags ContainerNodePool#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// taint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#taint ContainerNodePool#taint}
	Taint interface{} `field:"optional" json:"taint" yaml:"taint"`
	// windows_node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#windows_node_config ContainerNodePool#windows_node_config}
	WindowsNodeConfig *ContainerNodePoolNodeConfigWindowsNodeConfig `field:"optional" json:"windowsNodeConfig" yaml:"windowsNodeConfig"`
	// workload_metadata_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_node_pool#workload_metadata_config ContainerNodePool#workload_metadata_config}
	WorkloadMetadataConfig *ContainerNodePoolNodeConfigWorkloadMetadataConfig `field:"optional" json:"workloadMetadataConfig" yaml:"workloadMetadataConfig"`
}

