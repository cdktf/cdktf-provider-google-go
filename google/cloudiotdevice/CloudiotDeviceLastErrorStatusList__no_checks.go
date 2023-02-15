//go:build no_runtime_type_checking

package cloudiotdevice

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudiotDeviceLastErrorStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudiotDeviceLastErrorStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudiotDeviceLastErrorStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudiotDeviceLastErrorStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudiotDeviceLastErrorStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudiotDeviceLastErrorStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

