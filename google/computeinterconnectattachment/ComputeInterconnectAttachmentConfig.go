package computeinterconnectattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInterconnectAttachmentConfig struct {
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
	// Name of the resource.
	//
	// Provided by the client when the resource is created. The
	// name must be 1-63 characters long, and comply with RFC1035. Specifically, the
	// name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a
	// lowercase letter, and all following characters must be a dash, lowercase
	// letter, or digit, except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#name ComputeInterconnectAttachment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// URL of the cloud router to be used for dynamic routing.
	//
	// This router must be in
	// the same region as this InterconnectAttachment. The InterconnectAttachment will
	// automatically connect the Interconnect to the network & region within which the
	// Cloud Router is configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#router ComputeInterconnectAttachment#router}
	Router *string `field:"required" json:"router" yaml:"router"`
	// Whether the VLAN attachment is enabled or disabled.  When using PARTNER type this will Pre-Activate the interconnect attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#admin_enabled ComputeInterconnectAttachment#admin_enabled}
	AdminEnabled interface{} `field:"optional" json:"adminEnabled" yaml:"adminEnabled"`
	// Provisioned bandwidth capacity for the interconnect attachment.
	//
	// For attachments of type DEDICATED, the user can set the bandwidth.
	// For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
	// Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
	// Defaults to BPS_10G Possible values: ["BPS_50M", "BPS_100M", "BPS_200M", "BPS_300M", "BPS_400M", "BPS_500M", "BPS_1G", "BPS_2G", "BPS_5G", "BPS_10G", "BPS_20G", "BPS_50G"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#bandwidth ComputeInterconnectAttachment#bandwidth}
	Bandwidth *string `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// Up to 16 candidate prefixes that can be used to restrict the allocation of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
	//
	// All prefixes must be within link-local address space (169.254.0.0/16)
	// and must be /29 or shorter (/28, /27, etc). Google will attempt to select
	// an unused /29 from the supplied candidate prefix(es). The request will
	// fail if all possible /29s are in use on Google's edge. If not supplied,
	// Google will randomly select an unused /29 from all of link-local space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#candidate_subnets ComputeInterconnectAttachment#candidate_subnets}
	CandidateSubnets *[]*string `field:"optional" json:"candidateSubnets" yaml:"candidateSubnets"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#description ComputeInterconnectAttachment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Desired availability domain for the attachment.
	//
	// Only available for type
	// PARTNER, at creation time. For improved reliability, customers should
	// configure a pair of attachments with one per availability domain. The
	// selected availability domain will be provided to the Partner via the
	// pairing key so that the provisioned circuit will lie in the specified
	// domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#edge_availability_domain ComputeInterconnectAttachment#edge_availability_domain}
	EdgeAvailabilityDomain *string `field:"optional" json:"edgeAvailabilityDomain" yaml:"edgeAvailabilityDomain"`
	// Indicates the user-supplied encryption option of this interconnect attachment. Can only be specified at attachment creation for PARTNER or DEDICATED attachments.
	//
	// NONE - This is the default value, which means that the VLAN attachment
	// carries unencrypted traffic. VMs are able to send traffic to, or receive
	// traffic from, such a VLAN attachment.
	//
	// IPSEC - The VLAN attachment carries only encrypted traffic that is
	// encrypted by an IPsec device, such as an HA VPN gateway or third-party
	// IPsec VPN. VMs cannot directly send traffic to, or receive traffic from,
	// such a VLAN attachment. To use HA VPN over Cloud Interconnect, the VLAN
	// attachment must be created with this option. Default value: "NONE" Possible values: ["NONE", "IPSEC"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#encryption ComputeInterconnectAttachment#encryption}
	Encryption *string `field:"optional" json:"encryption" yaml:"encryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#id ComputeInterconnectAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// URL of the underlying Interconnect object that this attachment's traffic will traverse through.
	//
	// Required if type is DEDICATED, must not
	// be set if type is PARTNER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#interconnect ComputeInterconnectAttachment#interconnect}
	Interconnect *string `field:"optional" json:"interconnect" yaml:"interconnect"`
	// URL of addresses that have been reserved for the interconnect attachment, Used only for interconnect attachment that has the encryption option as IPSEC.
	//
	// The addresses must be RFC 1918 IP address ranges. When creating HA VPN
	// gateway over the interconnect attachment, if the attachment is configured
	// to use an RFC 1918 IP address, then the VPN gateway's IP address will be
	// allocated from the IP address range specified here.
	//
	// For example, if the HA VPN gateway's interface 0 is paired to this
	// interconnect attachment, then an RFC 1918 IP address for the VPN gateway
	// interface 0 will be allocated from the IP address specified for this
	// interconnect attachment.
	//
	// If this field is not specified for interconnect attachment that has
	// encryption option as IPSEC, later on when creating HA VPN gateway on this
	// interconnect attachment, the HA VPN gateway's IP address will be
	// allocated from regional external IP address pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#ipsec_internal_addresses ComputeInterconnectAttachment#ipsec_internal_addresses}
	IpsecInternalAddresses *[]*string `field:"optional" json:"ipsecInternalAddresses" yaml:"ipsecInternalAddresses"`
	// Maximum Transmission Unit (MTU), in bytes, of packets passing through this interconnect attachment.
	//
	// Currently, only 1440 and 1500 are allowed. If not specified, the value will default to 1440.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#mtu ComputeInterconnectAttachment#mtu}
	Mtu *string `field:"optional" json:"mtu" yaml:"mtu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#project ComputeInterconnectAttachment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Region where the regional interconnect attachment resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#region ComputeInterconnectAttachment#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#timeouts ComputeInterconnectAttachment#timeouts}
	Timeouts *ComputeInterconnectAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The type of InterconnectAttachment you wish to create. Defaults to DEDICATED. Possible values: ["DEDICATED", "PARTNER", "PARTNER_PROVIDER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#type ComputeInterconnectAttachment#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When using PARTNER type this will be managed upstream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_interconnect_attachment#vlan_tag8021q ComputeInterconnectAttachment#vlan_tag8021q}
	VlanTag8021Q *float64 `field:"optional" json:"vlanTag8021Q" yaml:"vlanTag8021Q"`
}

