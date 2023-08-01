package dnsrecordset


type DnsRecordSetRoutingPolicyPrimaryBackupPrimaryInternalLoadBalancers struct {
	// The frontend IP address of the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_record_set#ip_address DnsRecordSet#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: ["tcp", "udp"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_record_set#ip_protocol DnsRecordSet#ip_protocol}
	IpProtocol *string `field:"required" json:"ipProtocol" yaml:"ipProtocol"`
	// The type of load balancer. This value is case-sensitive. Possible values: ["regionalL4ilb", "regionalL7ilb].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_record_set#load_balancer_type DnsRecordSet#load_balancer_type}
	LoadBalancerType *string `field:"required" json:"loadBalancerType" yaml:"loadBalancerType"`
	// The fully qualified url of the network in which the load balancer belongs. This should be formatted like `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_record_set#network_url DnsRecordSet#network_url}
	NetworkUrl *string `field:"required" json:"networkUrl" yaml:"networkUrl"`
	// The configured port of the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_record_set#port DnsRecordSet#port}
	Port *string `field:"required" json:"port" yaml:"port"`
	// The ID of the project in which the load balancer belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_record_set#project DnsRecordSet#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The region of the load balancer. Only needed for regional load balancers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_record_set#region DnsRecordSet#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

