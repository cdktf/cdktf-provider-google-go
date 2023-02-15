//go:build no_runtime_type_checking

package endpointsservice

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EndpointsServiceApisMethodsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EndpointsServiceApisMethodsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EndpointsServiceApisMethodsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EndpointsServiceApisMethodsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EndpointsServiceApisMethodsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEndpointsServiceApisMethodsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

