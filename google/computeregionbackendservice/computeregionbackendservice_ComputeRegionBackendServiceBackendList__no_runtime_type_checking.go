//go:build no_runtime_type_checking

package computeregionbackendservice

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ComputeRegionBackendServiceBackendList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ComputeRegionBackendServiceBackendList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewComputeRegionBackendServiceBackendListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

