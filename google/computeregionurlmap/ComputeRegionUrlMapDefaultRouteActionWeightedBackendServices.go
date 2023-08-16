package computeregionurlmap


type ComputeRegionUrlMapDefaultRouteActionWeightedBackendServices struct {
	// The full or partial URL to the default BackendService resource.
	//
	// Before forwarding the request to backendService, the load balancer applies any relevant headerActions specified as part of this backendServiceWeight.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_url_map#backend_service ComputeRegionUrlMap#backend_service}
	BackendService *string `field:"optional" json:"backendService" yaml:"backendService"`
	// header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_url_map#header_action ComputeRegionUrlMap#header_action}
	HeaderAction *ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderAction `field:"optional" json:"headerAction" yaml:"headerAction"`
	// Specifies the fraction of traffic sent to a backend service, computed as weight / (sum of all weightedBackendService weights in routeAction) .
	//
	// The selection of a backend service is determined only for new traffic. Once a user's request has been directed to a backend service, subsequent requests are sent to the same backend service as determined by the backend service's session affinity policy.
	// The value must be from 0 to 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_url_map#weight ComputeRegionUrlMap#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

