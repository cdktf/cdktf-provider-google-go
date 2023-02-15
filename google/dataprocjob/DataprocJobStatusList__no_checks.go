//go:build no_runtime_type_checking

package dataprocjob

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataprocJobStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataprocJobStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataprocJobStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataprocJobStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataprocJobStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataprocJobStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

