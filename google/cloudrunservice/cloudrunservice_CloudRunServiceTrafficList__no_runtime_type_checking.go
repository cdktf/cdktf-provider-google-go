//go:build no_runtime_type_checking

package cloudrunservice

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudRunServiceTrafficList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudRunServiceTrafficList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceTrafficList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceTrafficList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceTrafficList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceTrafficList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudRunServiceTrafficListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

