//go:build no_runtime_type_checking

package alloydbcluster

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlloydbClusterMigrationSourceList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlloydbClusterMigrationSourceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlloydbClusterMigrationSourceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlloydbClusterMigrationSourceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlloydbClusterMigrationSourceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlloydbClusterMigrationSourceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

