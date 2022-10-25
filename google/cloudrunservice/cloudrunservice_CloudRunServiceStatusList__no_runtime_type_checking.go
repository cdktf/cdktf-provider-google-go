//go:build no_runtime_type_checking

package cloudrunservice

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudRunServiceStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudRunServiceStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudRunServiceStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

