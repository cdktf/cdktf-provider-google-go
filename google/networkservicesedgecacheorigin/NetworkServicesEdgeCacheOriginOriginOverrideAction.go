package networkservicesedgecacheorigin


type NetworkServicesEdgeCacheOriginOriginOverrideAction struct {
	// header_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/network_services_edge_cache_origin#header_action NetworkServicesEdgeCacheOrigin#header_action}
	HeaderAction *NetworkServicesEdgeCacheOriginOriginOverrideActionHeaderAction `field:"optional" json:"headerAction" yaml:"headerAction"`
	// url_rewrite block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/network_services_edge_cache_origin#url_rewrite NetworkServicesEdgeCacheOrigin#url_rewrite}
	UrlRewrite *NetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewrite `field:"optional" json:"urlRewrite" yaml:"urlRewrite"`
}
