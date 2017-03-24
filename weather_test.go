package main

import (
  "testing"
  "net/http"
  "net/http/httptest" 
)
func TestConverToCel(t *testing.T){
  expected := 0.0
  actual := ConverToCel(273.15)
  if actual != expected {
   t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual) 
  }
}

func TestWeatherAPIOK(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)

if r.Method != "GET" {
      t.Errorf("Expected 'GET' request, got '%s'", r.Method)
    }

r.ParseForm()
    location := r.Form.Get("q")
    if location != "Valencia,es" {
      t.Errorf("Expected request to have 'Valencia,es', got: '%s'", location)
    }

}))
 
 defer ts.Close()
  TestServerBaseUrl := ts.URL
  err := Temp(TestServerBaseUrl)
   if err != nil {
    t.Errorf("Temp() returned an error: %s", err)
  }
}
