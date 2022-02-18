package main

import "testing"

//testing function for Validate. Tests all Validate methods.
func TestValidate(t *testing.T) {
	var v Validate
	r := root(RootDir())
	v = r
	characterResult := v.CharacterScan()
	licenseResult := v.LicenseScan()
	readMeResult := v.ReadMeScan()
	tabResult := v.TabScan()

	if characterResult != true {
		t.Error("Character Scan did not run correctly.")
	} else if licenseResult != true {
		t.Error("License Scan did not run correctly.")
	} else if readMeResult != true {
		t.Error("ReadMe Scan did not run correctly")
	} else if tabResult != true {
		t.Error("Tab Scan did not run correctly.")
	}
}
