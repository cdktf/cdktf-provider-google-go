//go:build no_runtime_type_checking

package cloudrunservice

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudRunServiceStatusConditionsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudRunServiceStatusConditionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceStatusConditionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceStatusConditionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceStatusConditionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudRunServiceStatusConditionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

