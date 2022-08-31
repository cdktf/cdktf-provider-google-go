//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt google Provider for Terraform CDK (cdktf)
package google

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGoogleContainerClusterClusterAutoscalingList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGoogleContainerClusterClusterAutoscalingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGoogleContainerClusterClusterAutoscalingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGoogleContainerClusterClusterAutoscalingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGoogleContainerClusterClusterAutoscalingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGoogleContainerClusterClusterAutoscalingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

