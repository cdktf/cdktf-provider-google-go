//go:build no_runtime_type_checking

package cloudiotdevice

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudiotDeviceCredentialsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudiotDeviceCredentialsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudiotDeviceCredentialsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudiotDeviceCredentialsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudiotDeviceCredentialsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudiotDeviceCredentialsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudiotDeviceCredentialsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

