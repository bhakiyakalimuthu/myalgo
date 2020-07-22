package main

import (
	"log"

	"github.com/bhakiyakalimuthu/myalgo/pkg/stringsalgo"
)

func main() {
	list1, _ := stringsalgo.Reader("/Users/bhakiyarajkalimuthu/Documents/voi/rides/vehicle_riding.csv", 0)
	list2, _ := stringsalgo.Reader("/Users/bhakiyarajkalimuthu/Documents/voi/rides/active_rides.csv", 2)
	diff := stringsalgo.Difference(list1, list2)
	log.Print(len(diff))
	log.Print(diff)
	// in := []string{"bbaa", "aaa", "cccccd", "dfdddy"}
	// op := stringsalgo.BabySorting(in, "y", nil)
	// fmt.Println(op)

}
