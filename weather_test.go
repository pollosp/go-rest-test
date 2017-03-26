package main

import (
  "testing"
  "net/http"
  "net/http/httptest" 
  "os"
)
func TestConverToCel(t *testing.T){
  expected := 0.0
  actual := ConverToCel(273.15)
  if actual != expected {
   t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual) 
  }
}

func TestWeatherAPIOK(t *testing.T) {
  
  location_mock := os.Getenv("W_LOCATION")
  if location_mock == "" {
    location_mock = "Valencia,es"
  }
  token := os.Getenv("W_TOKEN")

  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)

if r.Method != "GET" {
      t.Errorf("Expected 'GET' request, got '%s'", r.Method)
    }

r.ParseForm()
    location := r.Form.Get("q")
    if location != location_mock {
      t.Errorf("Expected request to have %s, got: '%s'",location_mock, location)
    }

}))
 
 defer ts.Close()
  TestServerBaseUrl := ts.URL
  err := Temp(TestServerBaseUrl, location_mock, token)
   if err != nil {
    t.Errorf("Temp() returned an error: %s", err)
  }
}
