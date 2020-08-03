package services

import (
	"github.com/spf13/viper"
	"housing-anywhere/models"
	"strconv"
)

func Calculate(location models.Location) (float64, error){
	config, err := getConfiguration()
	if err != nil {
		return 0, err
	}

	SectorID, err := parseStringToFloat(config.SectorId)
	dimensionX, err := parseStringToFloat(location.X)
	dimensionY, err := parseStringToFloat(location.Y)
	dimensionZ, err := parseStringToFloat(location.Z)
	Velocity, err := parseStringToFloat(location.Vel)

	calculation := dimensionX * SectorID + dimensionY * SectorID + dimensionZ * SectorID + Velocity

	return calculation, nil
}

func parseStringToFloat(dimension string) (float64, error){
	dimensionFloat, err := strconv.ParseFloat(dimension, 64)
	if err != nil{
		return 0, err
	}

	return dimensionFloat, nil
}

func getConfiguration() (models.Configuration, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration models.Configuration

	if err := viper.ReadInConfig(); err != nil {
		return models.Configuration{}, err
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		return models.Configuration{}, err
	}

	return configuration, nil
}
