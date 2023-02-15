//go:build no_runtime_type_checking

package monitoringservice

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MonitoringServiceTelemetryList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MonitoringServiceTelemetryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MonitoringServiceTelemetryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MonitoringServiceTelemetryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MonitoringServiceTelemetryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMonitoringServiceTelemetryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

