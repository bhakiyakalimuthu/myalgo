package stringsalgo

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func Difference(a, b []string) []string {
	mb := make(map[string]string, len(b))
	for _, x := range b {
		mb[x] = x
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func Reader(path string, col int) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("failed to read file data")
	}
	recordTOReturn := make([]string, 0, len(records))
	for _, item := range records {
		recordTOReturn = append(recordTOReturn, item[col])
	}
	return recordTOReturn, nil

}

func JSONReader(path string) (*zone, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	defer file.Close()
	reader, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("failed to read file data")
	}
	var z zone
	if err := json.Unmarshal(reader, &z); err != nil {
		return nil, errors.New("failed to unmarshal")
	}
	return &z, nil
}

type zone struct {
	Zones []struct {
		ZoneID              string      `json:"zone_id"`
		Name                string      `json:"name"`
		Country             string      `json:"country"`
		CompanyName         string      `json:"company_name"`
		CompanyID           string      `json:"company_id"`
		Currency            string      `json:"currency"`
		Vat                 float64     `json:"vat"`
		StartCost           int         `json:"start_cost"`
		MinuteCost          int         `json:"minute_cost"`
		FeeOutOfZone        int         `json:"fee_out_of_zone"`
		CreditsExchangeRate interface{} `json:"credits_exchange_rate"`
		Language            string      `json:"language"`
		DefaultLocale       string      `json:"default_locale"`
		Locales             string      `json:"locales"`
		TimeZone            string      `json:"time_zone"`
		MaxSpeed            int         `json:"max_speed"`
		Boundaries          struct {
			Lo struct {
				Lat float64 `json:"Lat"`
				Lng float64 `json:"Lng"`
			} `json:"Lo"`
			Hi struct {
				Lat float64 `json:"Lat"`
				Lng float64 `json:"Lng"`
			} `json:"Hi"`
		} `json:"boundaries"`
		BountyStateExpr  interface{} `json:"bounty_state_expr"`
		BountyAmountExpr string      `json:"bounty_amount_expr"`
		MinutePrice      []struct {
			Day    string `json:"day"`
			Period string `json:"period"`
			Amount int    `json:"amount"`
		} `json:"minute_price"`
		SleepHours          interface{} `json:"sleep_hours"`
		IsActive            bool        `json:"is_active"`
		IsSleeping          bool        `json:"is_sleeping"`
		ParkingMode         string      `json:"parking_mode"`
		ParkingSnapshotMode string      `json:"parking_snapshot_mode"`
	} `json:"zones"`
}
