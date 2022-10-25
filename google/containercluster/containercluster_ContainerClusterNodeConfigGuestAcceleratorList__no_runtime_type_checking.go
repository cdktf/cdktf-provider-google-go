//go:build no_runtime_type_checking

package containercluster

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerClusterNodeConfigGuestAcceleratorList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerClusterNodeConfigGuestAcceleratorList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerClusterNodeConfigGuestAcceleratorList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerClusterNodeConfigGuestAcceleratorList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerClusterNodeConfigGuestAcceleratorList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerClusterNodeConfigGuestAcceleratorList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerClusterNodeConfigGuestAcceleratorListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

