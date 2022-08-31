//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt google Provider for Terraform CDK (cdktf)
package google

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGoogleContainerClusterNodePoolListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

