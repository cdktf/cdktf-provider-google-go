//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt google Provider for Terraform CDK (cdktf)
package google

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGoogleCloudRunServiceTemplateMetadataList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGoogleCloudRunServiceTemplateMetadataList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGoogleCloudRunServiceTemplateMetadataList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGoogleCloudRunServiceTemplateMetadataList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGoogleCloudRunServiceTemplateMetadataList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGoogleCloudRunServiceTemplateMetadataListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

