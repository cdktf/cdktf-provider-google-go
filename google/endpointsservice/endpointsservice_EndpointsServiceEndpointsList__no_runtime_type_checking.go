//go:build no_runtime_type_checking

package endpointsservice

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EndpointsServiceEndpointsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EndpointsServiceEndpointsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EndpointsServiceEndpointsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EndpointsServiceEndpointsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EndpointsServiceEndpointsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEndpointsServiceEndpointsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

