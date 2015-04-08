package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type DellWarrantySingle struct {
	GetAssetWarrantyResponse struct {
		_Xmlns                 string `json:"@xmlns"`
		GetAssetWarrantyResult struct {
			_A       string      `json:"@a"`
			_I       string      `json:"@i"`
			Faults   interface{} `json:"Faults"`
			Response struct {
				DellAsset struct {
					AssetParts struct {
						_Nil string `json:"@nil"`
					} `json:"AssetParts"`
					CountryLookupCode  float64 `json:"CountryLookupCode"`
					CustomerNumber     float64 `json:"CustomerNumber"`
					IsDuplicate        string  `json:"IsDuplicate"`
					ItemClassCode      string  `json:"ItemClassCode"`
					LocalChannel       float64 `json:"LocalChannel"`
					MachineDescription string  `json:"MachineDescription"`
					OrderNumber        float64 `json:"OrderNumber"`
					ParentServiceTag   struct {
						_Nil string `json:"@nil"`
					} `json:"ParentServiceTag"`
					ServiceTag string `json:"ServiceTag"`
					ShipDate   string `json:"ShipDate"`
					Warranties struct {
						Warranty []struct {
							EndDate                 string  `json:"EndDate"`
							EntitlementType         string  `json:"EntitlementType"`
							ItemNumber              string  `json:"ItemNumber"`
							ServiceLevelCode        string  `json:"ServiceLevelCode"`
							ServiceLevelDescription string  `json:"ServiceLevelDescription"`
							ServiceLevelGroup       float64 `json:"ServiceLevelGroup"`
							ServiceProvider         string  `json:"ServiceProvider"`
							StartDate               string  `json:"StartDate"`
						} `json:"Warranty"`
					} `json:"Warranties"`
				} `json:"DellAsset"`
			} `json:"Response"`
		} `json:"GetAssetWarrantyResult"`
	} `json:"GetAssetWarrantyResponse"`
}

type DellWarrantyList struct {
	GetAssetWarrantyResponse struct {
		_Xmlns                 string `json:"@xmlns"`
		GetAssetWarrantyResult struct {
			_A       string      `json:"@a"`
			_I       string      `json:"@i"`
			Faults   interface{} `json:"Faults"`
			Response struct {
				DellAsset []struct {
					AssetParts struct {
						_Nil string `json:"@nil"`
					} `json:"AssetParts"`
					CountryLookupCode  float64 `json:"CountryLookupCode"`
					CustomerNumber     float64 `json:"CustomerNumber"`
					IsDuplicate        string  `json:"IsDuplicate"`
					ItemClassCode      string  `json:"ItemClassCode"`
					LocalChannel       float64 `json:"LocalChannel"`
					MachineDescription string  `json:"MachineDescription"`
					OrderNumber        float64 `json:"OrderNumber"`
					ParentServiceTag   struct {
						_Nil string `json:"@nil"`
					} `json:"ParentServiceTag"`
					ServiceTag string `json:"ServiceTag"`
					ShipDate   string `json:"ShipDate"`
					Warranties struct {
						Warranty []struct {
							EndDate                 string  `json:"EndDate"`
							EntitlementType         string  `json:"EntitlementType"`
							ItemNumber              string  `json:"ItemNumber"`
							ServiceLevelCode        string  `json:"ServiceLevelCode"`
							ServiceLevelDescription string  `json:"ServiceLevelDescription"`
							ServiceLevelGroup       float64 `json:"ServiceLevelGroup"`
							ServiceProvider         string  `json:"ServiceProvider"`
							StartDate               string  `json:"StartDate"`
						} `json:"Warranty"`
					} `json:"Warranties"`
				} `json:"DellAsset"`
			} `json:"Response"`
		} `json:"GetAssetWarrantyResult"`
	} `json:"GetAssetWarrantyResponse"`
}

func getWarrantyInformation(content []byte, multiple bool) {
	// a lack of generics and not knowing better is causing this
	// schism right now. look into a better solution.
	// it also doesn't help that the API returns this subtle difference.
	if multiple {
		warranty := DellWarrantyList{}
		err := json.Unmarshal(content, &warranty)
		if err != nil {
			log.Fatal(err)
		}
		assets := warranty.GetAssetWarrantyResponse.GetAssetWarrantyResult.Response.DellAsset
		for _, asset := range assets {
			fmt.Println(asset.MachineDescription + " - " + asset.ServiceTag)
		}
	} else {
		warranty := DellWarrantySingle{}
		err := json.Unmarshal(content, &warranty)
		if err != nil {
			log.Fatal(err)
		}
		asset := warranty.GetAssetWarrantyResponse.GetAssetWarrantyResult.Response.DellAsset
		fmt.Println(asset.MachineDescription + " - " + asset.ServiceTag)
	}

	// quick and dirty testing
	// will fix soon
}
