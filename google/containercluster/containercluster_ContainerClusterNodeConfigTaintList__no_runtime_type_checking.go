//go:build no_runtime_type_checking

package containercluster

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerClusterNodeConfigTaintList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerClusterNodeConfigTaintList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerClusterNodeConfigTaintList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerClusterNodeConfigTaintList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerClusterNodeConfigTaintList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerClusterNodeConfigTaintList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerClusterNodeConfigTaintListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

