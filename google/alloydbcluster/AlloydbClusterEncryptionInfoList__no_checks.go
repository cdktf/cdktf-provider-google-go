//go:build no_runtime_type_checking

package alloydbcluster

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlloydbClusterEncryptionInfoList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlloydbClusterEncryptionInfoList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlloydbClusterEncryptionInfoList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlloydbClusterEncryptionInfoList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlloydbClusterEncryptionInfoList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlloydbClusterEncryptionInfoListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

