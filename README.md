About 

Utility developed in Golang ,to query the `MAC address lookup API`[https://macaddress.io/] to fetch the system details 

Prerequisite
Docker is required

Execution
Docker image can be found here , which takes MAC Address and `MAC API Key`[https://macaddress.io/api/documentation/making-requests] as an input

command to run the docker image :

docker run -e MAC_ADDRESS=<mac address> -e MAC_API_KEY=<mac api key> fetchmacdetails:latest