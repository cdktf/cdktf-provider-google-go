//go:build no_runtime_type_checking

package cloudrunservice

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudRunServiceTemplateSpecVolumesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecVolumesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecVolumesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecVolumesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecVolumesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecVolumesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudRunServiceTemplateSpecVolumesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

