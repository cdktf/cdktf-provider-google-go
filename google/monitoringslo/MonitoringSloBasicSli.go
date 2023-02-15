package monitoringslo


type MonitoringSloBasicSli struct {
	// availability block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_slo#availability MonitoringSlo#availability}
	Availability *MonitoringSloBasicSliAvailability `field:"optional" json:"availability" yaml:"availability"`
	// latency block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_slo#latency MonitoringSlo#latency}
	Latency *MonitoringSloBasicSliLatency `field:"optional" json:"latency" yaml:"latency"`
	// An optional set of locations to which this SLI is relevant.
	//
	// Telemetry from other locations will not be used to calculate
	// performance for this SLI. If omitted, this SLI applies to all
	// locations in which the Service has activity. For service types
	// that don't support breaking down by location, setting this
	// field will result in an error.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_slo#location MonitoringSlo#location}
	Location *[]*string `field:"optional" json:"location" yaml:"location"`
	// An optional set of RPCs to which this SLI is relevant.
	//
	// Telemetry from other methods will not be used to calculate
	// performance for this SLI. If omitted, this SLI applies to all
	// the Service's methods. For service types that don't support
	// breaking down by method, setting this field will result in an
	// error.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_slo#method MonitoringSlo#method}
	Method *[]*string `field:"optional" json:"method" yaml:"method"`
	// The set of API versions to which this SLI is relevant.
	//
	// Telemetry from other API versions will not be used to
	// calculate performance for this SLI. If omitted,
	// this SLI applies to all API versions. For service types
	// that don't support breaking down by version, setting this
	// field will result in an error.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_slo#version MonitoringSlo#version}
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}

