package main

import (
	"encoding/json"
	"errors"
)

type DellErrorList struct {
	GetAssetWarrantyResponse struct {
		_Xmlns                 string `json:"@xmlns"`
		GetAssetWarrantyResult struct {
			_A     string `json:"@a"`
			_I     string `json:"@i"`
			Faults struct {
				FaultException []struct {
					Code    string `json:"Code"`
					Message string `json:"Message"`
				} `json:"FaultException"`
			} `json:"Faults"`
			Response interface{} `json:"Response"`
		} `json:"GetAssetWarrantyResult"`
	} `json:"GetAssetWarrantyResponse"`
}

type DellError struct {
	Code    string
	Message string
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
				} `json:"DellAsset"`
			} `json:"Response"`
		} `json:"GetAssetWarrantyResult"`
	} `json:"GetAssetWarrantyResponse"`
}

type Asset struct {
	CountryLookupCode  string
	CustomerNumber     string
	IsDuplicate        string
	ItemClassCode      string
	LocalChannel       string
	MachineDescription string
	OrderNumber        string
	ParentServiceTag   string
	ServiceTag         string
	ShipDate           string
	Warranties         []Warranty
}

type Warranty struct {
	EndDate                 string
	EntitlementType         string
	ItemNumber              string
	ServiceLevelCode        string
	ServiceLevelDescription string
	ServiceLevelGroup       string
	ServiceProvider         string
	StartDate               string
}

func createDellErrorList(errors DellErrorList) []DellError {
	items := errors.GetAssetWarrantyResponse.GetAssetWarrantyResult.Faults.FaultException
	faultList := []DellError{}
	for _, item := range items {
		f := DellError{
			Code:    item.Code,
			Message: item.Message,
		}
		faultList = append(faultList, f)
	}
	return faultList
}

func createDellAssetFromList(warranties DellWarrantyList) []Asset {
	items := warranties.GetAssetWarrantyResponse.GetAssetWarrantyResult.Response.DellAsset
	assetList := []Asset{}
	for _, item := range items {
		warranties := []Warranty{}
		for _, warranty := range item.Warranties.Warranty {
			w := Warranty{
				StartDate:               warranty.StartDate,
				EndDate:                 warranty.EndDate,
				EntitlementType:         warranty.EntitlementType,
				ItemNumber:              warranty.ItemNumber,
				ServiceLevelCode:        warranty.ServiceLevelCode,
				ServiceLevelDescription: warranty.ServiceLevelDescription,
				ServiceLevelGroup:       warranty.ServiceLevelGroup,
				ServiceProvider:         warranty.ServiceProvider,
			}
			warranties = append(warranties, w)
		}
		asset := Asset{
			CountryLookupCode:  item.CountryLookupCode,
			CustomerNumber:     item.CustomerNumber,
			IsDuplicate:        item.IsDuplicate,
			ItemClassCode:      item.ItemClassCode,
			LocalChannel:       item.LocalChannel,
			MachineDescription: item.MachineDescription,
			OrderNumber:        item.OrderNumber,
			ParentServiceTag:   item.ParentServiceTag._Nil,
			ServiceTag:         item.ServiceTag,
			ShipDate:           item.ShipDate,
			Warranties:         warranties,
		}
		assetList = append(assetList, asset)
	}
	return assetList
}

func getWarrantyInformation(content []byte) ([]Asset, error) {
	dellAssets := []Asset{}
	warranty := DellWarrantyList{}
	err := json.Unmarshal(content, &warranty)
	if err != nil {
		return dellAssets, errors.New("failed to unmarshal assets")
	}
	dellAssets = createDellAssetFromList(warranty)
	if len(dellAssets) == 0 {
		return []Asset{}, errors.New("asset list was empty")
	}
	return dellAssets, nil
}

func getErrorInformation(content []byte) ([]DellError, error) {
	dellErrors := []DellError{}
	faults := DellErrorList{}
	err := json.Unmarshal(content, &faults)
	if err != nil {
		return dellErrors, errors.New("failed to unmarshal faults")
	}
	dellErrors = createDellErrorList(faults)
	return dellErrors, nil
}
