//go:build no_runtime_type_checking

package appengineapplication

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppEngineApplicationUrlDispatchRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppEngineApplicationUrlDispatchRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppEngineApplicationUrlDispatchRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppEngineApplicationUrlDispatchRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppEngineApplicationUrlDispatchRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppEngineApplicationUrlDispatchRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

