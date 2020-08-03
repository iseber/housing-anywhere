#Housing-Anywhere
##Go code challenge
This application returns calculated location for the drones for locating data store point.

##Notes

* Go modules used when building the app
* App has 1 post endpoint to calculate and return location and 1 healthcheck to ensure if it is up and ready
* Default port is 5000
* Application handles most of the errors by checking posted values and logs accordingly.
* `SectorID` for observed sector has been kept in configuration and does not change at the runtime.   
It can be editable through: `config.json` 

##How-to-run-app
`docker build -t dns .`  
`docker run -p 5000:5000 -d dns:latest`

Location endpoint serves under http://localhost:5000/api/v1  
Healtcheck endpoint serves under http://localhost:5000/api/v1/health