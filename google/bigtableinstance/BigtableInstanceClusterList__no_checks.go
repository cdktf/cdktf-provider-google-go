//go:build no_runtime_type_checking

package bigtableinstance

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BigtableInstanceClusterList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BigtableInstanceClusterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BigtableInstanceClusterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BigtableInstanceClusterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BigtableInstanceClusterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BigtableInstanceClusterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBigtableInstanceClusterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

