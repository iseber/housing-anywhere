package services

import (
	"github.com/spf13/viper"
	"housing-anywhere/models"
	"testing"
)

func init(){
	viper.AddConfigPath("../")
}

func TestCalculate(t *testing.T) {
	location := models.Location{ "10.25", "12.23", "15.12", "40" }
	expected := 77.60
	actual := Calculate(location)

	if actual != expected {
		t.Errorf("calculate returned unexpected value: got %v want %v",
			actual, expected)
	}
}
