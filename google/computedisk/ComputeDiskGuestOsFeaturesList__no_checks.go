//go:build no_runtime_type_checking

package computedisk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ComputeDiskGuestOsFeaturesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ComputeDiskGuestOsFeaturesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ComputeDiskGuestOsFeaturesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ComputeDiskGuestOsFeaturesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ComputeDiskGuestOsFeaturesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ComputeDiskGuestOsFeaturesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewComputeDiskGuestOsFeaturesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

