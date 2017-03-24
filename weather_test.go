//Stupid test just to try how it works

package main

import "testing"

func TestConverToCel(t *testing.T){
  expected := 0.0
  actual := ConverToCel(273.15)
  if actual != expected {
   t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual) 
  }
}
