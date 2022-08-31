// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type AppEngineServiceNetworkSettingsNetworkSettings struct {
	// The ingress settings for version or service. Default value: "INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED" Possible values: ["INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED", "INGRESS_TRAFFIC_ALLOWED_ALL", "INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY", "INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_service_network_settings#ingress_traffic_allowed AppEngineServiceNetworkSettings#ingress_traffic_allowed}
	IngressTrafficAllowed *string `field:"optional" json:"ingressTrafficAllowed" yaml:"ingressTrafficAllowed"`
}

