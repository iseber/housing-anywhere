package services

import (
	"github.com/tkanos/gonfig"
	"housing-anywhere/models"
	"log"
	"os"
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
	configuration := models.Configuration{}
	if os.Getenv("ENV") == "development" {
		err := gonfig.GetConf("configs/config.development.json", &configuration)
		if err != nil {
			log.Fatal(err)
		}
	}

	return configuration
}
