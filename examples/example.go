package main

import (
	"fmt"
	"github.com/jameycribbs/ivy"
)

// Need a struct for every table
type Plane struct {
	Name  string   `json:"name"`
	Speed int      `json:"speed"`
	Range int      `json:"range"`
	Tags  []string `json:"tags"`
}

// Need to add this method to every table struct
func (plane *Plane) Transform() {
	*plane = Plane(*plane)
}

func main() {
	//
	// Open DB
	//
	db, err := ivy.OpenDB("data")
	if err != nil {
		fmt.Println("Failed to open database:", err)
	}

	defer db.Close()

	plane := Plane{}

	//
	// Find
	//
	err = db.Find("planes", &plane, "1")
	if err != nil {
		fmt.Println("Find failed:", err)
	}

	fmt.Println("\n======================= Plane with id '1' =======================================================================\n")
	fmt.Printf("%#v\n", plane.Name)

	//
	// FindAllIdsForTags
	//
	ids, err := db.FindAllIdsForTags("planes", []string{"german"})
	if err != nil {
		fmt.Println("FindAllIdsForTags failed:", err)
	}

	fmt.Println("\n======================= Planes with tag 'german' ==================================================================\n")
	for _, id := range ids {
		err = db.Find("planes", &plane, id)
		if err != nil {
			fmt.Println("Find failed:", err)
		}

		fmt.Printf("%#v\n", plane.Name)
	}

	//
	// FindFirstIdForField
	//
	id, err := db.FindFirstIdForField("planes", "name", "P-51D")
	if err != nil {
		fmt.Println("FindFirstIdForField failed:", err)
	}

	fmt.Println("\n======================= Plane with name 'P-51D' ==================================================================\n")
	err = db.Find("planes", &plane, id)
	if err != nil {
		fmt.Println("Find failed:", err)
	}

	fmt.Printf("%#v\n", plane.Name)

	//
	// Create
	//
	plane = Plane{Name: "Test", Speed: 111, Range: 111, Tags: []string{"test"}}
	id, err = db.Create("planes", plane)
	if err != nil {
		fmt.Println("Create failed:", err)
	}

	err = db.Find("planes", &plane, id)
	if err != nil {
		fmt.Println("Find failed:", err)
	}

	fmt.Printf("\n======================= New Plane with id '%v' ================================================================\n", id)
	fmt.Printf("%#v\n", plane.Name)

	//
	// Update
	//
	plane.Speed = 999
	err = db.Update("planes", plane, id)
	if err != nil {
		fmt.Println("Update failed:", err)
	}

	err = db.Find("planes", &plane, id)
	if err != nil {
		fmt.Println("Find failed:", err)
	}

	fmt.Printf("\n======================= Plane with id '%v' ===================================================================\n", id)
	fmt.Printf("%#v\n", plane.Speed)

	//
	// Delete
	//
	err = db.Delete("planes", id)
	if err != nil {
		fmt.Println("Delete failed:", err)
	}
}
