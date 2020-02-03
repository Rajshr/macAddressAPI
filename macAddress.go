package main

import(
	"os"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/Sirupsen/logrus"
)

type macApiResponse struct {
	VendorDetails struct {
		Oui            string `json:"oui"`
		IsPrivate      bool   `json:"isPrivate"`
		CompanyName    string `json:"companyName"`
		CompanyAddress string `json:"companyAddress"`
		CountryCode    string `json:"countryCode"`
	} `json:"vendorDetails"`
}

func main(){
	getMacAddressDetails()
}

func  getMacAddressDetails()  {
	macAddress := os.Getenv("MAC_ADDRESS")
	if macAddress == "" {
		logrus.Fatal("Please enter the mac address for which details are to be retrieved.")
	}

	macApiKey := os.Getenv("MAC_API_KEY")	
	if macApiKey == "" {
		logrus.Fatal("Mac Api key is required to fetch the details.")
	}
	
	macApiURL := "https://api.macaddress.io/v1?output=json&search=" + macAddress
	request, err := http.NewRequest("GET", macApiURL, nil)
	if err != nil {
		logrus.Fatal(err)
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Authentication-Token", macApiKey)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		logrus.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	macApiResponse := macApiResponse{}
	json.Unmarshal(body, &macApiResponse)
	logrus.Info("Company Name for provided MAC Address is : ", macApiResponse.VendorDetails.CompanyName)

}
