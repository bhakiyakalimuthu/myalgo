package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/bhakiyakalimuthu/myalgo/pkg/stringsalgo"
)

// func main() {
// 	list1, _ := stringsalgo.Reader("/Users/bhakiyarajkalimuthu/Documents/voi/rides/vehicle_riding.csv", 0)
// 	list2, _ := stringsalgo.Reader("/Users/bhakiyarajkalimuthu/Documents/voi/rides/active_rides.csv", 2)
// 	// list3, _ := stringsalgo.Reader("/Users/bhakiyarajkalimuthu/Downloads/fm.csv", 0)
// 	diff := stringsalgo.Difference(list1, list2)
// 	log.Print(len(diff))
// 	log.Print(diff)
// // in := []string{"bbaa", "aaa", "cccccd", "dfdddy"}
// // op := stringsalgo.BabySorting(in, "y", nil)
// // fmt.Println(op)
// for _, i := range list3 {
// 	trim(i)
// }

// log.Print(list3)
// }

func main() {
	zone, _ := stringsalgo.JSONReader("/Users/bhakiyarajkalimuthu/Documents/voi/zones/zones.json")
	op := make([]zones, 0, len(zone.Zones))
	for _, z := range zone.Zones {
		op = append(op, zones{
			City: z.Name,
			ID:   z.ZoneID,
		})
	}
	println(len(zone.Zones))
	output, _ := json.Marshal(&op)
	log.Print(output)
}
func trim(str string) {
	s1 := strings.Replace(str, ";", "", -1)
	s2 := strings.Split(s1, "/")
	println(s2[4])
}

// func builder(city string, id string) {
// 	"zones":[{"city":"Berlin","id":"118"}]
// }

type zones struct {
	City string `json:"city"`
	ID   string `json:"id"`
}
