//go:build no_runtime_type_checking

package cloudiotdevice

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudiotDeviceStateList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudiotDeviceStateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudiotDeviceStateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudiotDeviceStateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudiotDeviceStateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudiotDeviceStateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

