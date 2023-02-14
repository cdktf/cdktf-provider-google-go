//go:build no_runtime_type_checking

package alloydbcluster

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlloydbClusterBackupSourceList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlloydbClusterBackupSourceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlloydbClusterBackupSourceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlloydbClusterBackupSourceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlloydbClusterBackupSourceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlloydbClusterBackupSourceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

