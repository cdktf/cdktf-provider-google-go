// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodeConfig struct {
	// advanced_machine_features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#advanced_machine_features ContainerCluster#advanced_machine_features}
	AdvancedMachineFeatures *ContainerClusterNodeConfigAdvancedMachineFeatures `field:"optional" json:"advancedMachineFeatures" yaml:"advancedMachineFeatures"`
	// boot_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#boot_disk ContainerCluster#boot_disk}
	BootDisk *ContainerClusterNodeConfigBootDisk `field:"optional" json:"bootDisk" yaml:"bootDisk"`
	// The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#boot_disk_kms_key ContainerCluster#boot_disk_kms_key}
	BootDiskKmsKey *string `field:"optional" json:"bootDiskKmsKey" yaml:"bootDiskKmsKey"`
	// confidential_nodes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#confidential_nodes ContainerCluster#confidential_nodes}
	ConfidentialNodes *ContainerClusterNodeConfigConfidentialNodes `field:"optional" json:"confidentialNodes" yaml:"confidentialNodes"`
	// containerd_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#containerd_config ContainerCluster#containerd_config}
	ContainerdConfig *ContainerClusterNodeConfigContainerdConfig `field:"optional" json:"containerdConfig" yaml:"containerdConfig"`
	// Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#disk_size_gb ContainerCluster#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#disk_type ContainerCluster#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// If enabled boot disks are configured with confidential mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#enable_confidential_storage ContainerCluster#enable_confidential_storage}
	EnableConfidentialStorage interface{} `field:"optional" json:"enableConfidentialStorage" yaml:"enableConfidentialStorage"`
	// ephemeral_storage_local_ssd_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#ephemeral_storage_local_ssd_config ContainerCluster#ephemeral_storage_local_ssd_config}
	EphemeralStorageLocalSsdConfig *ContainerClusterNodeConfigEphemeralStorageLocalSsdConfig `field:"optional" json:"ephemeralStorageLocalSsdConfig" yaml:"ephemeralStorageLocalSsdConfig"`
	// fast_socket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#fast_socket ContainerCluster#fast_socket}
	FastSocket *ContainerClusterNodeConfigFastSocket `field:"optional" json:"fastSocket" yaml:"fastSocket"`
	// Enables Flex Start provisioning model for the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#flex_start ContainerCluster#flex_start}
	FlexStart interface{} `field:"optional" json:"flexStart" yaml:"flexStart"`
	// gcfs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#gcfs_config ContainerCluster#gcfs_config}
	GcfsConfig *ContainerClusterNodeConfigGcfsConfig `field:"optional" json:"gcfsConfig" yaml:"gcfsConfig"`
	// guest_accelerator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#guest_accelerator ContainerCluster#guest_accelerator}
	GuestAccelerator interface{} `field:"optional" json:"guestAccelerator" yaml:"guestAccelerator"`
	// gvnic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#gvnic ContainerCluster#gvnic}
	Gvnic *ContainerClusterNodeConfigGvnic `field:"optional" json:"gvnic" yaml:"gvnic"`
	// host_maintenance_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#host_maintenance_policy ContainerCluster#host_maintenance_policy}
	HostMaintenancePolicy *ContainerClusterNodeConfigHostMaintenancePolicy `field:"optional" json:"hostMaintenancePolicy" yaml:"hostMaintenancePolicy"`
	// The image type to use for this node.
	//
	// Note that for a given image type, the latest version of it will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#image_type ContainerCluster#image_type}
	ImageType *string `field:"optional" json:"imageType" yaml:"imageType"`
	// kubelet_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#kubelet_config ContainerCluster#kubelet_config}
	KubeletConfig *ContainerClusterNodeConfigKubeletConfig `field:"optional" json:"kubeletConfig" yaml:"kubeletConfig"`
	// The map of Kubernetes labels (key/value pairs) to be applied to each node.
	//
	// These will added in addition to any default label(s) that Kubernetes may apply to the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#labels ContainerCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// linux_node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#linux_node_config ContainerCluster#linux_node_config}
	LinuxNodeConfig *ContainerClusterNodeConfigLinuxNodeConfig `field:"optional" json:"linuxNodeConfig" yaml:"linuxNodeConfig"`
	// local_nvme_ssd_block_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#local_nvme_ssd_block_config ContainerCluster#local_nvme_ssd_block_config}
	LocalNvmeSsdBlockConfig *ContainerClusterNodeConfigLocalNvmeSsdBlockConfig `field:"optional" json:"localNvmeSsdBlockConfig" yaml:"localNvmeSsdBlockConfig"`
	// The number of local SSD disks to be attached to the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#local_ssd_count ContainerCluster#local_ssd_count}
	LocalSsdCount *float64 `field:"optional" json:"localSsdCount" yaml:"localSsdCount"`
	// LocalSsdEncryptionMode specified the method used for encrypting the local SSDs attached to the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#local_ssd_encryption_mode ContainerCluster#local_ssd_encryption_mode}
	LocalSsdEncryptionMode *string `field:"optional" json:"localSsdEncryptionMode" yaml:"localSsdEncryptionMode"`
	// Type of logging agent that is used as the default value for node pools in the cluster.
	//
	// Valid values include DEFAULT and MAX_THROUGHPUT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#logging_variant ContainerCluster#logging_variant}
	LoggingVariant *string `field:"optional" json:"loggingVariant" yaml:"loggingVariant"`
	// The name of a Google Compute Engine machine type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#machine_type ContainerCluster#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// The runtime of each node in the node pool in seconds, terminated by 's'. Example: "3600s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#max_run_duration ContainerCluster#max_run_duration}
	MaxRunDuration *string `field:"optional" json:"maxRunDuration" yaml:"maxRunDuration"`
	// The metadata key/value pairs assigned to instances in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#metadata ContainerCluster#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Minimum CPU platform to be used by this instance.
	//
	// The instance may be scheduled on the specified or newer CPU platform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#min_cpu_platform ContainerCluster#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
	// Setting this field will assign instances of this pool to run on the specified node group.
	//
	// This is useful for running workloads on sole tenant nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#node_group ContainerCluster#node_group}
	NodeGroup *string `field:"optional" json:"nodeGroup" yaml:"nodeGroup"`
	// The set of Google API scopes to be made available on all of the node VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#oauth_scopes ContainerCluster#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
	// Whether the nodes are created as preemptible VM instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#preemptible ContainerCluster#preemptible}
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// reservation_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#reservation_affinity ContainerCluster#reservation_affinity}
	ReservationAffinity *ContainerClusterNodeConfigReservationAffinity `field:"optional" json:"reservationAffinity" yaml:"reservationAffinity"`
	// The GCE resource labels (a map of key/value pairs) to be applied to the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#resource_labels ContainerCluster#resource_labels}
	ResourceLabels *map[string]*string `field:"optional" json:"resourceLabels" yaml:"resourceLabels"`
	// A map of resource manager tags.
	//
	// Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#resource_manager_tags ContainerCluster#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
	// secondary_boot_disks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#secondary_boot_disks ContainerCluster#secondary_boot_disks}
	SecondaryBootDisks interface{} `field:"optional" json:"secondaryBootDisks" yaml:"secondaryBootDisks"`
	// The Google Cloud Platform Service Account to be used by the node VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#service_account ContainerCluster#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#shielded_instance_config ContainerCluster#shielded_instance_config}
	ShieldedInstanceConfig *ContainerClusterNodeConfigShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// sole_tenant_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#sole_tenant_config ContainerCluster#sole_tenant_config}
	SoleTenantConfig *ContainerClusterNodeConfigSoleTenantConfig `field:"optional" json:"soleTenantConfig" yaml:"soleTenantConfig"`
	// Whether the nodes are created as spot VM instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#spot ContainerCluster#spot}
	Spot interface{} `field:"optional" json:"spot" yaml:"spot"`
	// The list of Storage Pools where boot disks are provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#storage_pools ContainerCluster#storage_pools}
	StoragePools *[]*string `field:"optional" json:"storagePools" yaml:"storagePools"`
	// The list of instance tags applied to all nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#tags ContainerCluster#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// taint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#taint ContainerCluster#taint}
	Taint interface{} `field:"optional" json:"taint" yaml:"taint"`
	// windows_node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#windows_node_config ContainerCluster#windows_node_config}
	WindowsNodeConfig *ContainerClusterNodeConfigWindowsNodeConfig `field:"optional" json:"windowsNodeConfig" yaml:"windowsNodeConfig"`
	// workload_metadata_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#workload_metadata_config ContainerCluster#workload_metadata_config}
	WorkloadMetadataConfig *ContainerClusterNodeConfigWorkloadMetadataConfig `field:"optional" json:"workloadMetadataConfig" yaml:"workloadMetadataConfig"`
}

