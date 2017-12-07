package zDatabase

import (
	"encoding/json"
	"eveKillmailCrawler/killmail"
	"sort"
)

type ZDatabase struct {
	SolarSystems map[int]*SolarSystem `json:"SolarSystems"`
}

func (d *ZDatabase) AddSolarSystem(id int) {
	if d.SolarSystems[id] == nil {
		d.SolarSystems[id] = &SolarSystem{
			Id:          id,
			KillmailIDs: make([]int, 0),
			Minmatar:    make(map[int]*Ship),
			Amarr:       make(map[int]*Ship),
			Neutral:     make(map[int]*Ship),
		}
	}
}

func (d ZDatabase) AddKillmails(solarSystemID int, kill killmail.ZKillmail) {
	d.AddSolarSystem(solarSystemID)
	d.SolarSystems[solarSystemID].AddVictim(&kill)
	d.SolarSystems[solarSystemID].AddAttacker(&kill)
}

func (d ZDatabase) ExportToJson() []byte {
	b, _ := json.Marshal(d)
	return b
}

func (d *ZDatabase) ImportFromJson(b []byte) {
	json.Unmarshal(b, d)
}

//todo deze functies komen te veel voor PLS DRY!
//ik snap echt niet helemaal hoe deze pointer shiz werkt
func (d *ZDatabase) GetMostLostShipSorted(faction int) []*Item {

	shipSlice := make([]*Item, 0)

	for _, system := range d.SolarSystems {

		//dit kan vast wel mooier
		switch faction {
		case Amarr:
			for _, shipStruct := range system.Amarr {
				addShipLossToSlice(&shipSlice, *shipStruct)
			}
		case Minmatar:
			for _, shipStruct := range system.Minmatar {
				addShipLossToSlice(&shipSlice, *shipStruct)
			}
		default:
			for _, shipStruct := range system.Amarr {
				addShipLossToSlice(&shipSlice, *shipStruct)
			}
			for _, shipStruct := range system.Minmatar {
				addShipLossToSlice(&shipSlice, *shipStruct)
			}
			for _, shipStruct := range system.Neutral {
				addShipLossToSlice(&shipSlice, *shipStruct)
			}
		}
	}

	//sort de slice
	sort.Slice(shipSlice, func(i, j int) bool {
		return shipSlice[i].Count > shipSlice[j].Count
	})

	return shipSlice
}

func (d *ZDatabase) GetMostKillerShipSorted(faction int) []*Item {

	shipSlice := make([]*Item, 0)

	for _, system := range d.SolarSystems {

		//dit kan vast wel mooier
		switch faction {
		case Amarr:
			for _, shipStruct := range system.Amarr {
				addShipKillerToSlice(&shipSlice, *shipStruct)
			}
		case Minmatar:
			for _, shipStruct := range system.Minmatar {
				addShipKillerToSlice(&shipSlice, *shipStruct)
			}
		default:
			for _, shipStruct := range system.Amarr {
				addShipKillerToSlice(&shipSlice, *shipStruct)
			}
			for _, shipStruct := range system.Minmatar {
				addShipKillerToSlice(&shipSlice, *shipStruct)
			}
			for _, shipStruct := range system.Neutral {
				addShipKillerToSlice(&shipSlice, *shipStruct)
			}
		}
	}

	//sort de slice
	sort.Slice(shipSlice, func(i, j int) bool {
		return shipSlice[i].Count > shipSlice[j].Count
	})

	return shipSlice
}

func addShipLossToSlice(shipSlice *[]*Item, ship Ship) {
	for _, value := range *shipSlice {
		if value.Id == ship.Id {
			value.Count += ship.Loss
			return
		}
	}
	*shipSlice = append(*shipSlice, &Item{Id: ship.Id, Quantity: 0, Count: ship.Loss})
}

func addShipKillerToSlice(shipSlice *[]*Item, ship Ship) {
	for _, value := range *shipSlice {
		if value.Id == ship.Id {
			value.Count += ship.Killer
			return
		}
	}
	*shipSlice = append(*shipSlice, &Item{Id: ship.Id, Quantity: 0, Count: ship.Killer})
}

func (d *ZDatabase) GetMostFittedItemSorted(faction int) []*Item {

	itemSlice := make([]*Item, 0)

	for _, system := range d.SolarSystems {

		//dit kan vast wel mooier
		switch faction {
		case Amarr:
			for _, shipStruct := range system.Amarr {
				addFittedItemsToSlice(&itemSlice, shipStruct.Fitted)
			}
		case Minmatar:
			for _, shipStruct := range system.Minmatar {
				addFittedItemsToSlice(&itemSlice, shipStruct.Fitted)
			}
		default:
			for _, shipStruct := range system.Amarr {
				addFittedItemsToSlice(&itemSlice, shipStruct.Fitted)
			}
			for _, shipStruct := range system.Minmatar {
				addFittedItemsToSlice(&itemSlice, shipStruct.Fitted)
			}
			for _, shipStruct := range system.Neutral {
				addFittedItemsToSlice(&itemSlice, shipStruct.Fitted)
			}
		}
	}

	//sort de slice
	sort.Slice(itemSlice, func(i, j int) bool {
		return itemSlice[i].Count > itemSlice[j].Count
	})

	return itemSlice
}

func (d *ZDatabase) GetMostCargoItemSorted(faction int) []*Item {

	itemSlice := make([]*Item, 0)

	for _, system := range d.SolarSystems {

		//dit kan vast wel mooier
		switch faction {
		case Amarr:
			for _, shipStruct := range system.Amarr {
				addFittedItemsToSlice(&itemSlice, shipStruct.Cargo)
			}
		case Minmatar:
			for _, shipStruct := range system.Minmatar {
				addFittedItemsToSlice(&itemSlice, shipStruct.Cargo)
			}
		default:
			for _, shipStruct := range system.Amarr {
				addFittedItemsToSlice(&itemSlice, shipStruct.Cargo)
			}
			for _, shipStruct := range system.Minmatar {
				addFittedItemsToSlice(&itemSlice, shipStruct.Cargo)
			}
			for _, shipStruct := range system.Neutral {
				addFittedItemsToSlice(&itemSlice, shipStruct.Cargo)
			}
		}
	}

	//sort de slice
	sort.Slice(itemSlice, func(i, j int) bool {
		return itemSlice[i].Count > itemSlice[j].Count
	})

	return itemSlice
}

func addFittedItemsToSlice(s *[]*Item, m map[int]*Item) {
	for _, value := range m {
		addItemCountToSlice(s, *value)
	}
}

func addItemCountToSlice(itemSlice *[]*Item, item Item) {
	for _, value := range *itemSlice {
		if value.Id == item.Id {
			value.Count += item.Count
			value.Quantity += item.Quantity
			return
		}
	}
	*itemSlice = append(*itemSlice, &item)
}

func New() *ZDatabase {
	return &ZDatabase{SolarSystems: make(map[int]*SolarSystem)}
}
