package services

import (
	"github.com/spf13/viper"
	"housing-anywhere/models"
	"testing"
)

func init(){
	viper.AddConfigPath("../")
}

func TestCalculate_ShouldReturnCorrectResult_WhenDimensionsAreCorrect(t *testing.T) {
	location := models.Location{ "10.25", "12.23", "15.12", "40" }
	expected := 77.60
	actual, err := Calculate(location)

	if err != nil{
		t.Errorf("calculate returned error: got %v want %v",
			actual, expected)
	}

	if actual != expected {
		t.Errorf("calculate returned unexpected value: got %v want %v",
			actual, expected)
	}
}
