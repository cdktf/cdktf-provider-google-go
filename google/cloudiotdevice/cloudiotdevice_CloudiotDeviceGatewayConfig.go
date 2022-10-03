package cloudiotdevice


type CloudiotDeviceGatewayConfig struct {
	// Indicates whether the device is a gateway. Possible values: ["ASSOCIATION_ONLY", "DEVICE_AUTH_TOKEN_ONLY", "ASSOCIATION_AND_DEVICE_AUTH_TOKEN"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudiot_device#gateway_auth_method CloudiotDevice#gateway_auth_method}
	GatewayAuthMethod *string `field:"optional" json:"gatewayAuthMethod" yaml:"gatewayAuthMethod"`
	// Indicates whether the device is a gateway. Default value: "NON_GATEWAY" Possible values: ["GATEWAY", "NON_GATEWAY"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudiot_device#gateway_type CloudiotDevice#gateway_type}
	GatewayType *string `field:"optional" json:"gatewayType" yaml:"gatewayType"`
}

