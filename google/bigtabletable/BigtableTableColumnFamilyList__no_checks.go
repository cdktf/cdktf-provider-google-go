//go:build no_runtime_type_checking

package bigtabletable

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BigtableTableColumnFamilyList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BigtableTableColumnFamilyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BigtableTableColumnFamilyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BigtableTableColumnFamilyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BigtableTableColumnFamilyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BigtableTableColumnFamilyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBigtableTableColumnFamilyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

