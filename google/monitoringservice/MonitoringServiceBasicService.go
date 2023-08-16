package monitoringservice


type MonitoringServiceBasicService struct {
	// Labels that specify the resource that emits the monitoring data which is used for SLO reporting of this 'Service'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_service#service_labels MonitoringService#service_labels}
	ServiceLabels *map[string]*string `field:"optional" json:"serviceLabels" yaml:"serviceLabels"`
	// The type of service that this basic service defines, e.g. APP_ENGINE service type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_service#service_type MonitoringService#service_type}
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
}

