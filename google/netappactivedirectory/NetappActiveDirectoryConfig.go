// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappactivedirectory

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappActiveDirectoryConfig struct {
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
	// Comma separated list of DNS server IP addresses for the Active Directory domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#dns NetappActiveDirectory#dns}
	Dns *string `field:"required" json:"dns" yaml:"dns"`
	// Fully qualified domain name for the Active Directory domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#domain NetappActiveDirectory#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Name of the region for the policy to apply to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#location NetappActiveDirectory#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name of the Active Directory pool. Needs to be unique per location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#name NetappActiveDirectory#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// NetBIOS name prefix of the server to be created.
	//
	// A five-character random ID is generated automatically, for example, -6f9a, and appended to the prefix. The full UNC share path will have the following format:
	// '\\NetBIOS_PREFIX-ABCD.DOMAIN_NAME\SHARE_NAME'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#net_bios_prefix NetappActiveDirectory#net_bios_prefix}
	NetBiosPrefix *string `field:"required" json:"netBiosPrefix" yaml:"netBiosPrefix"`
	// Password for specified username.
	//
	// Note - Manual changes done to the password will not be detected. Terraform will not re-apply the password, unless you use a new password in Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#password NetappActiveDirectory#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Username for the Active Directory account with permissions to create the compute account within the specified organizational unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#username NetappActiveDirectory#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Domain user accounts to be added to the local Administrators group of the SMB service.
	//
	// Comma-separated list of domain users or groups. The Domain Admin group is automatically added when the service joins your domain as a hidden group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#administrators NetappActiveDirectory#administrators}
	Administrators *[]*string `field:"optional" json:"administrators" yaml:"administrators"`
	// Enables AES-128 and AES-256 encryption for Kerberos-based communication with Active Directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#aes_encryption NetappActiveDirectory#aes_encryption}
	AesEncryption interface{} `field:"optional" json:"aesEncryption" yaml:"aesEncryption"`
	// Domain user/group accounts to be added to the Backup Operators group of the SMB service.
	//
	// The Backup Operators group allows members to backup and restore files regardless of whether they have read or write access to the files. Comma-separated list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#backup_operators NetappActiveDirectory#backup_operators}
	BackupOperators *[]*string `field:"optional" json:"backupOperators" yaml:"backupOperators"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#description NetappActiveDirectory#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If enabled, traffic between the SMB server to Domain Controller (DC) will be encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#encrypt_dc_connections NetappActiveDirectory#encrypt_dc_connections}
	EncryptDcConnections interface{} `field:"optional" json:"encryptDcConnections" yaml:"encryptDcConnections"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#id NetappActiveDirectory#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Hostname of the Active Directory server used as Kerberos Key Distribution Center. Only required for volumes using kerberized NFSv4.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#kdc_hostname NetappActiveDirectory#kdc_hostname}
	KdcHostname *string `field:"optional" json:"kdcHostname" yaml:"kdcHostname"`
	// IP address of the Active Directory server used as Kerberos Key Distribution Center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#kdc_ip NetappActiveDirectory#kdc_ip}
	KdcIp *string `field:"optional" json:"kdcIp" yaml:"kdcIp"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#labels NetappActiveDirectory#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Specifies whether or not the LDAP traffic needs to be signed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#ldap_signing NetappActiveDirectory#ldap_signing}
	LdapSigning interface{} `field:"optional" json:"ldapSigning" yaml:"ldapSigning"`
	// Local UNIX users on clients without valid user information in Active Directory are blocked from access to LDAP enabled volumes.
	//
	// This option can be used to temporarily switch such volumes to AUTH_SYS authentication (user ID + 1-16 groups).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#nfs_users_with_ldap NetappActiveDirectory#nfs_users_with_ldap}
	NfsUsersWithLdap interface{} `field:"optional" json:"nfsUsersWithLdap" yaml:"nfsUsersWithLdap"`
	// Name of the Organizational Unit where you intend to create the computer account for NetApp Volumes.
	//
	// Defaults to 'CN=Computers' if left empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#organizational_unit NetappActiveDirectory#organizational_unit}
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#project NetappActiveDirectory#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Domain accounts that require elevated privileges such as 'SeSecurityPrivilege' to manage security logs. Comma-separated list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#security_operators NetappActiveDirectory#security_operators}
	SecurityOperators *[]*string `field:"optional" json:"securityOperators" yaml:"securityOperators"`
	// Specifies an Active Directory site to manage domain controller selection.
	//
	// Use when Active Directory domain controllers in multiple regions are configured. Defaults to 'Default-First-Site-Name' if left empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#site NetappActiveDirectory#site}
	Site *string `field:"optional" json:"site" yaml:"site"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/netapp_active_directory#timeouts NetappActiveDirectory#timeouts}
	Timeouts *NetappActiveDirectoryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

