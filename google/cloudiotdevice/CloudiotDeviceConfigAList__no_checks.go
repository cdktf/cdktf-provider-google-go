//go:build no_runtime_type_checking

package cloudiotdevice

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudiotDeviceConfigAList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudiotDeviceConfigAList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudiotDeviceConfigAList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudiotDeviceConfigAList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudiotDeviceConfigAList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudiotDeviceConfigAListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

