package monitoringcustomservice


type MonitoringCustomServiceTelemetry struct {
	// The full name of the resource that defines this service. Formatted as described in https://cloud.google.com/apis/design/resource_names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/monitoring_custom_service#resource_name MonitoringCustomService#resource_name}
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
}

