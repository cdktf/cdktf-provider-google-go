//go:build no_runtime_type_checking

package gkehubfeature

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GkeHubFeatureStateList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GkeHubFeatureStateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GkeHubFeatureStateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GkeHubFeatureStateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GkeHubFeatureStateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGkeHubFeatureStateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

