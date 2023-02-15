//go:build no_runtime_type_checking

package containernodepool

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerNodePoolNodeConfigGuestAcceleratorList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigGuestAcceleratorList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigGuestAcceleratorList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigGuestAcceleratorList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigGuestAcceleratorList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigGuestAcceleratorList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerNodePoolNodeConfigGuestAcceleratorListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

