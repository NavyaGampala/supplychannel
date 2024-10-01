package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type Asset struct {
	ProductID         string `json:"ProductID"`
	Name              string `json:"Name"`
	Description       string `json:"Description"`
	ManufacturingDate string `json:"ManufacturingDate"`
	BatchNo           string `json:"BatchNo"`
	SupplyDate        string `json:"SupplyDate"`
	WareLocation      string `json:"WareLocation"`
	WholesaleDate     string  `json:"WholesaleDate"`
	WholesaleLocation string  `json:"WholesaleLocation"`
	WholesaleQuantity string `json:"WholesaleQuantity"`
	Status            string `json:"Status"`
}

//InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
		{ProductID: "asset1", Description: "NA", BatchNo: "1", Name: "apple", ManufacturingDate:"1/1/24",SupplyDate:"1/2/24",WareLocation:"W1",WholesaleDate:"1/3/24",WholesaleLocation:"WH1",WholesaleQuantity:"10",Status:"NA"},
		{ProductID: "asset2", Description: "NA", BatchNo: "2", Name: "banana", ManufacturingDate:"1/2/24",SupplyDate:"1/2/24",WareLocation:"W1",WholesaleDate:"1/3/24",WholesaleLocation:"WH1",WholesaleQuantity:"10",Status:"NA"},
		{ProductID: "asset3", Description: "NA", BatchNo: "3", Name: "grape", ManufacturingDate:"1/3/24",SupplyDate:"1/2/24",WareLocation:"W1",WholesaleDate:"1/3/24",WholesaleLocation:"WH1",WholesaleQuantity:"10",Status:"NA"},
		{ProductID: "asset4", Description: "NA", BatchNo: "4", Name: "guava", ManufacturingDate:"1/4/24",SupplyDate:"1/2/24",WareLocation:"W1",WholesaleDate:"1/3/24",WholesaleLocation:"WH1",WholesaleQuantity:"10",Status:"NA"},
		{ProductID: "asset5", Description: "NA", BatchNo: "5", Name: "potato", ManufacturingDate:"1/5/24",SupplyDate:"1/2/24",WareLocation:"W1",WholesaleDate:"1/3/24",WholesaleLocation:"WH1",WholesaleQuantity:"10",Status:"NA"},
		{ProductID: "asset6", Description: "NA", BatchNo: "6", Name: "tomato", ManufacturingDate:"1/6/24",SupplyDate:"1/2/24",WareLocation:"W1",WholesaleDate:"1/3/24",WholesaleLocation:"WH1",WholesaleQuantity:"10",Status:"NA"},
		{ProductID: "asset7", Description: "NA", BatchNo: "7", Name: "carrot", ManufacturingDate:"1/7/24",SupplyDate:"1/2/24",WareLocation:"W1",WholesaleDate:"1/3/24",WholesaleLocation:"WH1",WholesaleQuantity:"10",Status:"NA"},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ProductID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateProduct(ctx contractapi.TransactionContextInterface, id string, name string, description string, manufdate string, batch string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	asset := Asset{
		ProductID:              id,
		Name:                   name,
		Description:            description,
		ManufacturingDate:      manufdate,
		BatchNo:                batch,
		SupplyDate:         "NA",
        WareLocation:   "NA",
        WholesaleDate:  "NA",
        WholesaleLocation: "NA",
        WholesaleQuantity:  "NA",
        Status:  "created",
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}
func (s *SmartContract)  SupplyProduct(ctx contractapi.TransactionContextInterface, id string, supplyDate string, warehouseLocation string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}
	asset.SupplyDate = supplyDate
	asset.WareLocation = warehouseLocation
	asset.Status = "Supplied"
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, assetJSON)

}
func (s *SmartContract)  WholesaleProduct(ctx contractapi.TransactionContextInterface, id string, wholesaleDate string , wholeSaleLocation string , wholesaleQuantity string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}
	asset.WholesaleDate = wholesaleDate
	asset.WholesaleLocation = wholeSaleLocation
	asset.WholesaleQuantity = wholesaleQuantity
	asset.Status = "Wholesaled"
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, assetJSON)
}
func (s *SmartContract) QueryProduct(ctx contractapi.TransactionContextInterface, id string ) (*Asset, error) {
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return nil,err
	}

	return asset , nil

}
func (s *SmartContract) UpdateProductStatus(ctx contractapi.TransactionContextInterface, id string , status string ) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.Status = status
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(id, assetJSON)
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, status string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	// overwriting original asset with new asset
	asset := Asset{
		ProductID:             id,
		Status:			status,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}


// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// TransferAsset updates the owner field of asset with given id in world state, and returns the old owner.
// func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string) (string, error) {
// 	asset, err := s.ReadAsset(ctx, id)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	oldOwner := asset.Owner
// 	asset.Owner = newOwner
//
// 	assetJSON, err := json.Marshal(asset)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	err = ctx.GetStub().PutState(id, assetJSON)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	return oldOwner, nil
// }

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
// package chaincode
//
// import (
// 	"encoding/json"
// 	"fmt"
//
// 	"github.com/hyperledger/fabric-contract-api-go/contractapi"
// )
//
// // SmartContract provides functions for managing an Asset
// type SmartContract struct {
// 	contractapi.Contract
// }
//
// // Asset describes basic details of what makes up a simple asset
// // Insert struct field in alphabetic order => to achieve determinism across languages
// // golang keeps the order when marshal to json but doesn't order automatically
// type Boat struct {
//
// 	ID             string `json:"ID"`
// 	Color          string `json:"Color"`
//
// 	Owner          string `json:"Owner"`
// 	Value           int    `json:"Value"`
// }
//
// // InitLedger adds a base set of assets to the ledger
// func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
// 	assets := []Boat{
// 		{ID: "asset1", Color: "blue",  Owner: "Tomoko", Value: 300},
// 		{ID: "asset2", Color: "red", Owner: "Brad", Value: 400},
// 		{ID: "asset3", Color: "green",  Owner: "Jin Soo", Value: 500},
// 		{ID: "asset4", Color: "yellow",  Owner: "Max", Value: 600},
// 		{ID: "asset5", Color: "black", Owner: "Adriana", Value: 700},
// 		{ID: "asset6", Color: "white", Owner: "Michel", Value: 800},
// 	}
//
// 	for _, asset := range assets {
// 		assetJSON, err := json.Marshal(asset)
// 		if err != nil {
// 			return err
// 		}
//
// 		err = ctx.GetStub().PutState(asset.ID, assetJSON)
// 		if err != nil {
// 			return fmt.Errorf("failed to put to world state. %v", err)
// 		}
// 	}
//
// 	return nil
// }
//
// // CreateAsset issues a new asset to the world state with given details.
// func (s *SmartContract) CreateBoat(ctx contractapi.TransactionContextInterface, id string, color string,  owner string, value int) error {
// 	exists, err := s.BoatExists(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	if exists {
// 		return fmt.Errorf("the asset %s already exists", id)
// 	}
//
// 	asset := Boat{
// 		ID:             id,
// 		Color:          color,
//
// 		Owner:          owner,
// 		Value: value,
// 	}
// 	assetJSON, err := json.Marshal(asset)
// 	if err != nil {
// 		return err
// 	}
//
// 	return ctx.GetStub().PutState(id, assetJSON)
// }
//
// // ReadAsset returns the asset stored in the world state with given id.
// func (s *SmartContract) ReadBoat(ctx contractapi.TransactionContextInterface, id string) (*Boat, error) {
// 	assetJSON, err := ctx.GetStub().GetState(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read from world state: %v", err)
// 	}
// 	if assetJSON == nil {
// 		return nil, fmt.Errorf("the asset %s does not exist", id)
// 	}
//
// 	var asset Boat
// 	err = json.Unmarshal(assetJSON, &asset)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return &asset, nil
// }
//
// // UpdateAsset updates an existing asset in the world state with provided parameters.
// func (s *SmartContract) UpdateBoat(ctx contractapi.TransactionContextInterface, id string, color string, owner string, value int) error {
// 	exists, err := s.BoatExists(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	if !exists {
// 		return fmt.Errorf("the asset %s does not exist", id)
// 	}
//
// 	// overwriting original asset with new asset
// 	asset := Boat{
// 		ID:             id,
// 		Color:          color,
//
// 		Owner:          owner,
// 		Value: value,
// 	}
// 	assetJSON, err := json.Marshal(asset)
// 	if err != nil {
// 		return err
// 	}
//
// 	return ctx.GetStub().PutState(id, assetJSON)
// }
//
// // DeleteAsset deletes an given asset from the world state.
// func (s *SmartContract) DeleteBoat(ctx contractapi.TransactionContextInterface, id string) error {
// 	exists, err := s.BoatExists(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	if !exists {
// 		return fmt.Errorf("the asset %s does not exist", id)
// 	}
//
// 	return ctx.GetStub().DelState(id)
// }
//
// // AssetExists returns true when asset with given ID exists in world state
// func (s *SmartContract) BoatExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
// 	assetJSON, err := ctx.GetStub().GetState(id)
// 	if err != nil {
// 		return false, fmt.Errorf("failed to read from world state: %v", err)
// 	}
//
// 	return assetJSON != nil, nil
// }
//
// // TransferAsset updates the owner field of asset with given id in world state, and returns the old owner.
// func (s *SmartContract) TransferBoat(ctx contractapi.TransactionContextInterface, id string, newOwner string) (string, error) {
// 	asset, err := s.ReadBoat(ctx, id)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	oldOwner := asset.Owner
// 	asset.Owner = newOwner
//
// 	assetJSON, err := json.Marshal(asset)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	err = ctx.GetStub().PutState(id, assetJSON)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	return oldOwner, nil
// }
//
// // GetAllAssets returns all assets found in world state
// func (s *SmartContract) GetAllBoats(ctx contractapi.TransactionContextInterface) ([]*Boat, error) {
// 	// range query with empty string for startKey and endKey does an
// 	// open-ended query of all assets in the chaincode namespace.
// 	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resultsIterator.Close()
//
// 	var assets []*Boat
// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		var asset Boat./
// 		err = json.Unmarshal(queryResponse.Value, &asset)
// 		if err != nil {
// 			return nil, err
// 		}
// 		assets = append(assets, &asset)
// 	}
//
// 	return assets, nil
// }
