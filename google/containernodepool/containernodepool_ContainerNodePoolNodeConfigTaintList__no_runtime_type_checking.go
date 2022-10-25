//go:build no_runtime_type_checking

package containernodepool

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerNodePoolNodeConfigTaintList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigTaintList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigTaintList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigTaintList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigTaintList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigTaintList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerNodePoolNodeConfigTaintListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

