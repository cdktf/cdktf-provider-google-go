//go:build no_runtime_type_checking

package cloudbuildtrigger

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudbuildTriggerBuildStepList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudbuildTriggerBuildStepList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudbuildTriggerBuildStepList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudbuildTriggerBuildStepList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudbuildTriggerBuildStepList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudbuildTriggerBuildStepList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudbuildTriggerBuildStepListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

