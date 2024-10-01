// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume


type NetappVolumeExportPolicyRules struct {
	// Defines the access type for clients matching the 'allowedClients' specification. Possible values: ["READ_ONLY", "READ_WRITE", "READ_NONE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/netapp_volume#access_type NetappVolume#access_type}
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// Defines the client ingress specification (allowed clients) as a comma seperated list with IPv4 CIDRs or IPv4 host addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/netapp_volume#allowed_clients NetappVolume#allowed_clients}
	AllowedClients *string `field:"optional" json:"allowedClients" yaml:"allowedClients"`
	// If enabled, the root user (UID = 0) of the specified clients doesn't get mapped to nobody (UID = 65534).
	//
	// This is also known as no_root_squash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/netapp_volume#has_root_access NetappVolume#has_root_access}
	HasRootAccess *string `field:"optional" json:"hasRootAccess" yaml:"hasRootAccess"`
	// If enabled (true) the rule defines a read only access for clients matching the 'allowedClients' specification.
	//
	// It enables nfs clients to mount using 'integrity' kerberos security mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/netapp_volume#kerberos5i_read_only NetappVolume#kerberos5i_read_only}
	Kerberos5IReadOnly interface{} `field:"optional" json:"kerberos5IReadOnly" yaml:"kerberos5IReadOnly"`
	// If enabled (true) the rule defines read and write access for clients matching the 'allowedClients' specification.
	//
	// It enables nfs clients to mount using 'integrity' kerberos security mode. The 'kerberos5iReadOnly' value is ignored if this is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/netapp_volume#kerberos5i_read_write NetappVolume#kerberos5i_read_write}
	Kerberos5IReadWrite interface{} `field:"optional" json:"kerberos5IReadWrite" yaml:"kerberos5IReadWrite"`
	// If enabled (true) the rule defines a read only access for clients matching the 'allowedClients' specification.
	//
	// It enables nfs clients to mount using 'privacy' kerberos security mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/netapp_volume#kerberos5p_read_only NetappVolume#kerberos5p_read_only}
	Kerberos5PReadOnly interface{} `field:"optional" json:"kerberos5PReadOnly" yaml:"kerberos5PReadOnly"`
	// If enabled (true) the rule defines read and write access for clients matching the 'allowedClients' specification.
	//
	// It enables nfs clients to mount using 'privacy' kerberos security mode. The 'kerberos5pReadOnly' value is ignored if this is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/netapp_volume#kerberos5p_read_write NetappVolume#kerberos5p_read_write}
	Kerberos5PReadWrite interface{} `field:"optional" json:"kerberos5PReadWrite" yaml:"kerberos5PReadWrite"`
	// If enabled (true) the rule defines a read only access for clients matching the 'allowedClients' specification.
	//
	// It enables nfs clients to mount using 'authentication' kerberos security mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/netapp_volume#kerberos5_read_only NetappVolume#kerberos5_read_only}
	Kerberos5ReadOnly interface{} `field:"optional" json:"kerberos5ReadOnly" yaml:"kerberos5ReadOnly"`
	// If enabled (true) the rule defines read and write access for clients matching the 'allowedClients' specification.
	//
	// It enables nfs clients to mount using 'authentication' kerberos security mode. The 'kerberos5ReadOnly' value is ignored if this is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/netapp_volume#kerberos5_read_write NetappVolume#kerberos5_read_write}
	Kerberos5ReadWrite interface{} `field:"optional" json:"kerberos5ReadWrite" yaml:"kerberos5ReadWrite"`
	// Enable to apply the export rule to NFSV3 clients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/netapp_volume#nfsv3 NetappVolume#nfsv3}
	Nfsv3 interface{} `field:"optional" json:"nfsv3" yaml:"nfsv3"`
	// Enable to apply the export rule to NFSV4.1 clients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/netapp_volume#nfsv4 NetappVolume#nfsv4}
	Nfsv4 interface{} `field:"optional" json:"nfsv4" yaml:"nfsv4"`
}

