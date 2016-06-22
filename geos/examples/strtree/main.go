package main

import (
	"encoding/json"
	"fmt"

	"github.com/robertogyn19/gogeos/geos"
)

// Callback used to get items from query/iterate functions
func handleResult(item []byte) {
	i := Item{}
	json.Unmarshal(item, &i)

	fmt.Printf("Item: %v\n", i)
}

type Item struct {
	Index int
	Name  string
	Price float64
}

/**
Implement geos.STRTreeItem interface to be able to insert
*/
func (i Item) Parse() []byte {
	j, _ := json.Marshal(i)
	return j
}

func main() {
	strTreeCapacity := 10
	strtree := geos.NewSTRTree(strTreeCapacity)

	points := []geos.Coord{geos.NewCoord(-46, -23)}

	for i := 1; i < strTreeCapacity; i++ {
		points = append(points, geos.NewCoord(0, 0))
	}

	array := []geos.STRTreeItem{}
	for i := 0; i < strTreeCapacity; i++ {
		name := fmt.Sprintf("Product-%d", i+1)
		p := float64((i + 1) * 10)
		array = append(array, Item{i + 1, name, p})
	}

	for i, item := range array {
		g, _ := geos.NewPoint(points[i])
		strtree.Insert(g, item)
	}

	wktPolygon := "POLYGON ((-46.19064331 -23.15930069, -46.19064331 -22.79580631, -45.81367493 -22.79580631, -45.81367493 -23.15930069, -46.19064331 -23.15930069))"
	polygon, _ := geos.FromWKT(wktPolygon)

	fmt.Println("Query with polygon (should print only one item):")
	strtree.Query(polygon, handleResult)

	fmt.Println("\nIterate all items (should iterate all items):")
	strtree.Iterate(handleResult)
}
