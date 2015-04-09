package main

import (
	"encoding/json"
	"errors"
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

type Asset struct {
	CountryLookupCode  float64
	CustomerNumber     float64
	IsDuplicate        string
	ItemClassCode      string
	LocalChannel       float64
	MachineDescription string
	OrderNumber        float64
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
	ServiceLevelGroup       float64
	ServiceProvider         string
	StartDate               string
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

// there has to be a better way...

func createDellAssetFromSingle(warranties DellWarrantySingle) []Asset {
	item := warranties.GetAssetWarrantyResponse.GetAssetWarrantyResult.Response.DellAsset
	assetList := make([]Asset, 1)
	warrantyList := make([]Warranty, len(item.Warranties.Warranty))
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
		warrantyList = append(warrantyList, w)
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
		Warranties:         warrantyList,
	}
	assetList = append(assetList, asset)
	return assetList
}

func getWarrantyInformation(content []byte, multiple bool) ([]Asset, error) {
	dellAssets := []Asset{}
	// a lack of generics and not knowing better is causing this
	// schism right now. look into a better solution.
	// it also doesn't help that the API returns this subtle difference.
	if multiple {
		warranty := DellWarrantyList{}
		err := json.Unmarshal(content, &warranty)
		if err != nil {
			return dellAssets, errors.New("failed to unmarshal")
		}
		dellAssets = createDellAssetFromList(warranty)
	} else {
		warranty := DellWarrantySingle{}
		err := json.Unmarshal(content, &warranty)
		if err != nil {
			return dellAssets, errors.New("failed to unmarshal")
		}
		dellAssets = createDellAssetFromSingle(warranty)
	}
	return dellAssets, nil
}
