package computeregionurlmap


type ComputeRegionUrlMapPathMatcher struct {
	// The name to which this PathMatcher is referred by the HostRule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_url_map#name ComputeRegionUrlMap#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A reference to a RegionBackendService resource.
	//
	// This will be used if
	// none of the pathRules defined by this PathMatcher is matched by
	// the URL's path portion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_url_map#default_service ComputeRegionUrlMap#default_service}
	DefaultService *string `field:"optional" json:"defaultService" yaml:"defaultService"`
	// default_url_redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_url_map#default_url_redirect ComputeRegionUrlMap#default_url_redirect}
	DefaultUrlRedirect *ComputeRegionUrlMapPathMatcherDefaultUrlRedirect `field:"optional" json:"defaultUrlRedirect" yaml:"defaultUrlRedirect"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_url_map#description ComputeRegionUrlMap#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// path_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_url_map#path_rule ComputeRegionUrlMap#path_rule}
	PathRule interface{} `field:"optional" json:"pathRule" yaml:"pathRule"`
	// route_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_url_map#route_rules ComputeRegionUrlMap#route_rules}
	RouteRules interface{} `field:"optional" json:"routeRules" yaml:"routeRules"`
}

