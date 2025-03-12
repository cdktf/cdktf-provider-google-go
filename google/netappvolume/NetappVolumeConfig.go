// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappVolumeConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Capacity of the volume (in GiB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#capacity_gib NetappVolume#capacity_gib}
	CapacityGib *string `field:"required" json:"capacityGib" yaml:"capacityGib"`
	// Name of the pool location.
	//
	// Usually a region name, expect for some STANDARD service level pools which require a zone name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#location NetappVolume#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the volume. Needs to be unique per location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#name NetappVolume#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The protocol of the volume.
	//
	// Allowed combinations are '['NFSV3']', '['NFSV4']', '['SMB']', '['NFSV3', 'NFSV4']', '['SMB', 'NFSV3']' and '['SMB', 'NFSV4']'. Possible values: ["NFSV3", "NFSV4", "SMB"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#protocols NetappVolume#protocols}
	Protocols *[]*string `field:"required" json:"protocols" yaml:"protocols"`
	// Share name (SMB) or export path (NFS) of the volume. Needs to be unique per location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#share_name NetappVolume#share_name}
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	// Name of the storage pool to create the volume in. Pool needs enough spare capacity to accommodate the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#storage_pool NetappVolume#storage_pool}
	StoragePool *string `field:"required" json:"storagePool" yaml:"storagePool"`
	// backup_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#backup_config NetappVolume#backup_config}
	BackupConfig *NetappVolumeBackupConfig `field:"optional" json:"backupConfig" yaml:"backupConfig"`
	// Policy to determine if the volume should be deleted forcefully.
	//
	// Volumes may have nested snapshot resources. Deleting such a volume will fail.
	// Setting this parameter to FORCE will delete volumes including nested snapshots.
	// Possible values: DEFAULT, FORCE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#deletion_policy NetappVolume#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#description NetappVolume#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// export_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#export_policy NetappVolume#export_policy}
	ExportPolicy *NetappVolumeExportPolicy `field:"optional" json:"exportPolicy" yaml:"exportPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#id NetappVolume#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Flag indicating if the volume is a kerberos volume or not, export policy rules control kerberos security modes (krb5, krb5i, krb5p).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#kerberos_enabled NetappVolume#kerberos_enabled}
	KerberosEnabled interface{} `field:"optional" json:"kerberosEnabled" yaml:"kerberosEnabled"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#labels NetappVolume#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Optional. Flag indicating if the volume will be a large capacity volume or a regular volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#large_capacity NetappVolume#large_capacity}
	LargeCapacity interface{} `field:"optional" json:"largeCapacity" yaml:"largeCapacity"`
	// Optional.
	//
	// Flag indicating if the volume will have an IP address per node for volumes supporting multiple IP endpoints.
	// Only the volume with largeCapacity will be allowed to have multiple endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#multiple_endpoints NetappVolume#multiple_endpoints}
	MultipleEndpoints interface{} `field:"optional" json:"multipleEndpoints" yaml:"multipleEndpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#project NetappVolume#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// restore_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#restore_parameters NetappVolume#restore_parameters}
	RestoreParameters *NetappVolumeRestoreParameters `field:"optional" json:"restoreParameters" yaml:"restoreParameters"`
	// List of actions that are restricted on this volume. Possible values: ["DELETE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#restricted_actions NetappVolume#restricted_actions}
	RestrictedActions *[]*string `field:"optional" json:"restrictedActions" yaml:"restrictedActions"`
	// Security Style of the Volume.
	//
	// Use UNIX to use UNIX or NFSV4 ACLs for file permissions.
	// Use NTFS to use NTFS ACLs for file permissions. Can only be set for volumes which use SMB together with NFS as protocol. Possible values: ["NTFS", "UNIX"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#security_style NetappVolume#security_style}
	SecurityStyle *string `field:"optional" json:"securityStyle" yaml:"securityStyle"`
	// Settings for volumes with SMB access. Possible values: ["ENCRYPT_DATA", "BROWSABLE", "CHANGE_NOTIFY", "NON_BROWSABLE", "OPLOCKS", "SHOW_SNAPSHOT", "SHOW_PREVIOUS_VERSIONS", "ACCESS_BASED_ENUMERATION", "CONTINUOUSLY_AVAILABLE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#smb_settings NetappVolume#smb_settings}
	SmbSettings *[]*string `field:"optional" json:"smbSettings" yaml:"smbSettings"`
	// If enabled, a NFS volume will contain a read-only .snapshot directory which provides access to each of the volume's snapshots. Will enable "Previous Versions" support for SMB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#snapshot_directory NetappVolume#snapshot_directory}
	SnapshotDirectory interface{} `field:"optional" json:"snapshotDirectory" yaml:"snapshotDirectory"`
	// snapshot_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#snapshot_policy NetappVolume#snapshot_policy}
	SnapshotPolicy *NetappVolumeSnapshotPolicy `field:"optional" json:"snapshotPolicy" yaml:"snapshotPolicy"`
	// tiering_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#tiering_policy NetappVolume#tiering_policy}
	TieringPolicy *NetappVolumeTieringPolicy `field:"optional" json:"tieringPolicy" yaml:"tieringPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#timeouts NetappVolume#timeouts}
	Timeouts *NetappVolumeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Unix permission the mount point will be created with. Default is 0770. Applicable for UNIX security style volumes only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/netapp_volume#unix_permissions NetappVolume#unix_permissions}
	UnixPermissions *string `field:"optional" json:"unixPermissions" yaml:"unixPermissions"`
}

