//go:build no_runtime_type_checking

package bigqueryjob

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BigqueryJobStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BigqueryJobStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BigqueryJobStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BigqueryJobStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BigqueryJobStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBigqueryJobStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

