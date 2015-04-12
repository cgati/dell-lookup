package main

import (
	"encoding/json"
	"errors"
)

type DellResult struct {
	Assets []DellAsset
	Faults []DellError
}

type DellError struct {
	Code    string
	Message string
}

type DellWarrantyList struct {
	GetAssetWarrantyResponse struct {
		_Xmlns                 string `json:"@xmlns"`
		GetAssetWarrantyResult struct {
			_A     string `json:"@a"`
			_I     string `json:"@i"`
			Faults *struct {
				FaultException json.RawMessage
			} `json:"Faults"`
			Response *struct {
				DellAsset json.RawMessage
			} `json:"Response"`
		} `json:"GetAssetWarrantyResult"`
	} `json:"GetAssetWarrantyResponse"`
}

type DellAsset struct {
	AssetParts struct {
		_Nil string `json:"@nil"`
	} `json:"AssetParts"`
	CountryLookupCode  string `json:"CountryLookupCode"`
	CustomerNumber     string `json:"CustomerNumber"`
	IsDuplicate        string `json:"IsDuplicate"`
	ItemClassCode      string `json:"ItemClassCode"`
	LocalChannel       string `json:"LocalChannel"`
	MachineDescription string `json:"MachineDescription"`
	OrderNumber        string `json:"OrderNumber"`
	ParentServiceTag   struct {
		_Nil string `json:"@nil"`
	} `json:"ParentServiceTag"`
	ServiceTag string `json:"ServiceTag"`
	ShipDate   string `json:"ShipDate"`
	Warranties struct {
		Warranty []struct {
			EndDate                 string `json:"EndDate"`
			EntitlementType         string `json:"EntitlementType"`
			ItemNumber              string `json:"ItemNumber"`
			ServiceLevelCode        string `json:"ServiceLevelCode"`
			ServiceLevelDescription string `json:"ServiceLevelDescription"`
			ServiceLevelGroup       string `json:"ServiceLevelGroup"`
			ServiceProvider         string `json:"ServiceProvider"`
			StartDate               string `json:"StartDate"`
		} `json:"Warranty"`
	} `json:"Warranties"`
}

func getWarrantyInformation(content []byte) (DellResult, error) {
	dellAssets := []DellAsset{}
	dellAsset := DellAsset{}
	fault := DellError{}
	faults := []DellError{}
	warranty := DellWarrantyList{}

	err := json.Unmarshal(content, &warranty)
	if err != nil {
		return DellResult{}, errors.New("failed to unmarshal assets")
	}
	result := warranty.GetAssetWarrantyResponse.GetAssetWarrantyResult
	if result.Response != nil {
		if result.Response.DellAsset[0] == '[' {
			_ = json.Unmarshal(result.Response.DellAsset, &dellAssets)
		} else {
			_ = json.Unmarshal(result.Response.DellAsset, &dellAsset)
			dellAssets = []DellAsset{dellAsset}
		}
	}
	if result.Faults != nil {
		if result.Faults.FaultException[0] == '[' {
			_ = json.Unmarshal(result.Faults.FaultException, &faults)
		} else {
			_ = json.Unmarshal(result.Faults.FaultException, &fault)
			faults = []DellError{fault}
		}
	}
	dellResult := DellResult{dellAssets, faults}
	return dellResult, nil
}
