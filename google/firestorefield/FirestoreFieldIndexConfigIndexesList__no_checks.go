//go:build no_runtime_type_checking

package firestorefield

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirestoreFieldIndexConfigIndexesList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FirestoreFieldIndexConfigIndexesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FirestoreFieldIndexConfigIndexesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FirestoreFieldIndexConfigIndexesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FirestoreFieldIndexConfigIndexesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FirestoreFieldIndexConfigIndexesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFirestoreFieldIndexConfigIndexesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

