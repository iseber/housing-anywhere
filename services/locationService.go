package services

import (
	"github.com/spf13/viper"
	"housing-anywhere/models"
	"log"
	"strconv"
)

func Calculate(location models.Location) float64{
	config := getConfiguration()

	SectorID := parseStringToFloat(config.SectorId)
	dimensionX := parseStringToFloat(location.X)
	dimensionY := parseStringToFloat(location.Y)
	dimensionZ := parseStringToFloat(location.Z)
	Velocity := parseStringToFloat(location.Vel)

	calculation := dimensionX * SectorID + dimensionY * SectorID + dimensionZ * SectorID + Velocity

	return calculation
}

func parseStringToFloat(dimension string) float64{
	dimensionFloat, err := strconv.ParseFloat(dimension, 64)
	if err != nil{
		log.Fatal(err)
	}

	return dimensionFloat
}

func getConfiguration() models.Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration models.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return configuration
}
