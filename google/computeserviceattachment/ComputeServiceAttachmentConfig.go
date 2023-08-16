package computeserviceattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeServiceAttachmentConfig struct {
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
	// The connection preference to use for this service attachment. Valid values include "ACCEPT_AUTOMATIC", "ACCEPT_MANUAL".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_service_attachment#connection_preference ComputeServiceAttachment#connection_preference}
	ConnectionPreference *string `field:"required" json:"connectionPreference" yaml:"connectionPreference"`
	// If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_service_attachment#enable_proxy_protocol ComputeServiceAttachment#enable_proxy_protocol}
	EnableProxyProtocol interface{} `field:"required" json:"enableProxyProtocol" yaml:"enableProxyProtocol"`
	// Name of the resource.
	//
	// The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_service_attachment#name ComputeServiceAttachment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of subnets that is provided for NAT in this service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_service_attachment#nat_subnets ComputeServiceAttachment#nat_subnets}
	NatSubnets *[]*string `field:"required" json:"natSubnets" yaml:"natSubnets"`
	// The URL of a forwarding rule that represents the service identified by this service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_service_attachment#target_service ComputeServiceAttachment#target_service}
	TargetService *string `field:"required" json:"targetService" yaml:"targetService"`
	// consumer_accept_lists block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_service_attachment#consumer_accept_lists ComputeServiceAttachment#consumer_accept_lists}
	ConsumerAcceptLists interface{} `field:"optional" json:"consumerAcceptLists" yaml:"consumerAcceptLists"`
	// An array of projects that are not allowed to connect to this service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_service_attachment#consumer_reject_lists ComputeServiceAttachment#consumer_reject_lists}
	ConsumerRejectLists *[]*string `field:"optional" json:"consumerRejectLists" yaml:"consumerRejectLists"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_service_attachment#description ComputeServiceAttachment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If specified, the domain name will be used during the integration between the PSC connected endpoints and the Cloud DNS.
	//
	// For example, this is a
	// valid domain name: "p.mycompany.com.". Current max number of domain names
	// supported is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_service_attachment#domain_names ComputeServiceAttachment#domain_names}
	DomainNames *[]*string `field:"optional" json:"domainNames" yaml:"domainNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_service_attachment#id ComputeServiceAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_service_attachment#project ComputeServiceAttachment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// This flag determines whether a consumer accept/reject list change can reconcile the statuses of existing ACCEPTED or REJECTED PSC endpoints.
	//
	// If false, connection policy update will only affect existing PENDING PSC endpoints. Existing ACCEPTED/REJECTED endpoints will remain untouched regardless how the connection policy is modified .
	// If true, update will affect both PENDING and ACCEPTED/REJECTED PSC endpoints. For example, an ACCEPTED PSC endpoint will be moved to REJECTED if its project is added to the reject list.
	//
	// For newly created service attachment, this boolean defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_service_attachment#reconcile_connections ComputeServiceAttachment#reconcile_connections}
	ReconcileConnections interface{} `field:"optional" json:"reconcileConnections" yaml:"reconcileConnections"`
	// URL of the region where the resource resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_service_attachment#region ComputeServiceAttachment#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_service_attachment#timeouts ComputeServiceAttachment#timeouts}
	Timeouts *ComputeServiceAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

