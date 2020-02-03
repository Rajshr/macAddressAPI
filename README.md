### About 

* Utility developed in Golang ,to query the [MAC address lookup API](https://macaddress.io/) to fetch the system details 

### Prerequisite
* Docker is required

### Execution
* Docker image can be found [here](https://hub.docker.com/repository/docker/chashruti/macaddresspi/) , which takes MAC Address and [MAC API Key](https://macaddress.io/api/documentation/making-requests) as an input

* Command for execution :
  - `docker run -e MAC_ADDRESS=<mac address> -e MAC_API_KEY=<mac api key> chashruti/macaddressapi:latest`
