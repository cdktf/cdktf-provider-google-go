//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt google Provider for Terraform CDK (cdktf)
package google

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MonitoringMetricDescriptorLabelsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MonitoringMetricDescriptorLabelsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MonitoringMetricDescriptorLabelsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MonitoringMetricDescriptorLabelsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MonitoringMetricDescriptorLabelsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MonitoringMetricDescriptorLabelsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMonitoringMetricDescriptorLabelsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

