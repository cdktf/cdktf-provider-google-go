package notebooksinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotebooksInstanceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A reference to the zone where the machine resides.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#location NotebooksInstance#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// A reference to a machine type which defines VM kind.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#machine_type NotebooksInstance#machine_type}
	MachineType *string `field:"required" json:"machineType" yaml:"machineType"`
	// The name specified for the Notebook instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#name NotebooksInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// accelerator_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#accelerator_config NotebooksInstance#accelerator_config}
	AcceleratorConfig *NotebooksInstanceAcceleratorConfig `field:"optional" json:"acceleratorConfig" yaml:"acceleratorConfig"`
	// The size of the boot disk in GB attached to this instance, up to a maximum of 64000 GB (64 TB).
	//
	// The minimum recommended value is 100 GB.
	// If not specified, this defaults to 100.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#boot_disk_size_gb NotebooksInstance#boot_disk_size_gb}
	BootDiskSizeGb *float64 `field:"optional" json:"bootDiskSizeGb" yaml:"bootDiskSizeGb"`
	// Possible disk types for notebook instances. Possible values: ["DISK_TYPE_UNSPECIFIED", "PD_STANDARD", "PD_SSD", "PD_BALANCED", "PD_EXTREME"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#boot_disk_type NotebooksInstance#boot_disk_type}
	BootDiskType *string `field:"optional" json:"bootDiskType" yaml:"bootDiskType"`
	// container_image block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#container_image NotebooksInstance#container_image}
	ContainerImage *NotebooksInstanceContainerImage `field:"optional" json:"containerImage" yaml:"containerImage"`
	// Instance creation time.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#create_time NotebooksInstance#create_time}
	CreateTime *string `field:"optional" json:"createTime" yaml:"createTime"`
	// Specify a custom Cloud Storage path where the GPU driver is stored.
	//
	// If not specified, we'll automatically choose from official GPU drivers.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#custom_gpu_driver_path NotebooksInstance#custom_gpu_driver_path}
	CustomGpuDriverPath *string `field:"optional" json:"customGpuDriverPath" yaml:"customGpuDriverPath"`
	// The size of the data disk in GB attached to this instance, up to a maximum of 64000 GB (64 TB).
	//
	// You can choose the size of the data disk based on how big your notebooks and data are.
	// If not specified, this defaults to 100.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#data_disk_size_gb NotebooksInstance#data_disk_size_gb}
	DataDiskSizeGb *float64 `field:"optional" json:"dataDiskSizeGb" yaml:"dataDiskSizeGb"`
	// Possible disk types for notebook instances. Possible values: ["DISK_TYPE_UNSPECIFIED", "PD_STANDARD", "PD_SSD", "PD_BALANCED", "PD_EXTREME"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#data_disk_type NotebooksInstance#data_disk_type}
	DataDiskType *string `field:"optional" json:"dataDiskType" yaml:"dataDiskType"`
	// Disk encryption method used on the boot and data disks, defaults to GMEK. Possible values: ["DISK_ENCRYPTION_UNSPECIFIED", "GMEK", "CMEK"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#disk_encryption NotebooksInstance#disk_encryption}
	DiskEncryption *string `field:"optional" json:"diskEncryption" yaml:"diskEncryption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#id NotebooksInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the end user authorizes Google Cloud to install GPU driver on this instance.
	//
	// If this field is empty or set to false, the GPU driver
	// won't be installed. Only applicable to instances with GPUs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#install_gpu_driver NotebooksInstance#install_gpu_driver}
	InstallGpuDriver interface{} `field:"optional" json:"installGpuDriver" yaml:"installGpuDriver"`
	// The list of owners of this instance after creation.
	//
	// Format: alias@example.com.
	// Currently supports one owner only.
	// If not specified, all of the service account users of
	// your VM instance's service account can use the instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#instance_owners NotebooksInstance#instance_owners}
	InstanceOwners *[]*string `field:"optional" json:"instanceOwners" yaml:"instanceOwners"`
	// The KMS key used to encrypt the disks, only applicable if diskEncryption is CMEK. Format: projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#kms_key NotebooksInstance#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Labels to apply to this instance.
	//
	// These can be later modified by the setLabels method.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#labels NotebooksInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Custom metadata to apply to this instance.
	//
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#metadata NotebooksInstance#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// The name of the VPC that this instance is in. Format: projects/{project_id}/global/networks/{network_id}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#network NotebooksInstance#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The type of vNIC driver. Possible values: ["UNSPECIFIED_NIC_TYPE", "VIRTIO_NET", "GVNIC"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#nic_type NotebooksInstance#nic_type}
	NicType *string `field:"optional" json:"nicType" yaml:"nicType"`
	// The notebook instance will not register with the proxy..
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#no_proxy_access NotebooksInstance#no_proxy_access}
	NoProxyAccess interface{} `field:"optional" json:"noProxyAccess" yaml:"noProxyAccess"`
	// No public IP will be assigned to this instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#no_public_ip NotebooksInstance#no_public_ip}
	NoPublicIp interface{} `field:"optional" json:"noPublicIp" yaml:"noPublicIp"`
	// If true, the data disk will not be auto deleted when deleting the instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#no_remove_data_disk NotebooksInstance#no_remove_data_disk}
	NoRemoveDataDisk interface{} `field:"optional" json:"noRemoveDataDisk" yaml:"noRemoveDataDisk"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up.
	//
	// The path must be a URL
	// or Cloud Storage path (gs://path-to-file/file-name).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#post_startup_script NotebooksInstance#post_startup_script}
	PostStartupScript *string `field:"optional" json:"postStartupScript" yaml:"postStartupScript"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#project NotebooksInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// reservation_affinity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#reservation_affinity NotebooksInstance#reservation_affinity}
	ReservationAffinity *NotebooksInstanceReservationAffinity `field:"optional" json:"reservationAffinity" yaml:"reservationAffinity"`
	// The service account on this instance, giving access to other Google Cloud services.
	//
	// You can use any service account within
	// the same project, but you must have the service account user
	// permission to use the instance. If not specified,
	// the Compute Engine default service account is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#service_account NotebooksInstance#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Optional.
	//
	// The URIs of service account scopes to be included in Compute Engine instances.
	// If not specified, the following scopes are defined:
	// - https://www.googleapis.com/auth/cloud-platform
	// - https://www.googleapis.com/auth/userinfo.email
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#service_account_scopes NotebooksInstance#service_account_scopes}
	ServiceAccountScopes *[]*string `field:"optional" json:"serviceAccountScopes" yaml:"serviceAccountScopes"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#shielded_instance_config NotebooksInstance#shielded_instance_config}
	ShieldedInstanceConfig *NotebooksInstanceShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// The name of the subnet that this instance is in. Format: projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#subnet NotebooksInstance#subnet}
	Subnet *string `field:"optional" json:"subnet" yaml:"subnet"`
	// The Compute Engine tags to add to instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#tags NotebooksInstance#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#timeouts NotebooksInstance#timeouts}
	Timeouts *NotebooksInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Instance update time.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#update_time NotebooksInstance#update_time}
	UpdateTime *string `field:"optional" json:"updateTime" yaml:"updateTime"`
	// vm_image block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/notebooks_instance#vm_image NotebooksInstance#vm_image}
	VmImage *NotebooksInstanceVmImage `field:"optional" json:"vmImage" yaml:"vmImage"`
}

