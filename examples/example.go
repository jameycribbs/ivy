package main

import (
	"fmt"
	"github.com/jameycribbs/ivy"
)

// Need a struct for every table
type Plane struct {
	FileId     string   `json:"-"`
	Name       string   `json:"name"`
	Speed      int      `json:"speed"`
	Range      int      `json:"range"`
	EngineType string   `json:"enginetype"`
	Tags       []string `json:"tags"`
}

// AfterFind is a callback that runs inside the Find method after the record has been found and populated.
// You need to define an AfterFind for each one of your tables.
func (plane *Plane) AfterFind(db *ivy.DB, fileId string) {
	// IMPORTANT!  Make sure you do something like this in your AfterFind methods for your tables.
	// This turns the interface record inside the Find method back into your particular struct type
	// before returning the record to you.
	*plane = Plane(*plane)

	plane.FileId = fileId
}

func main() {
	// Specify which tables Ivy should build indexes for.
	fieldsToIndex := make(map[string][]string)
	fieldsToIndex["planes"] = []string{"tags", "name", "enginetype"}

	//
	// Open DB
	//
	db, err := ivy.OpenDB("data", fieldsToIndex)
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
		plane = Plane{}

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

	plane = Plane{}

	err = db.Find("planes", &plane, id)
	if err != nil {
		fmt.Println("Find failed:", err)
	}

	fmt.Printf("%#v\n", plane.Name)

	//
	// FindAllIdsForField
	//
	ids, err = db.FindAllIdsForField("planes", "enginetype", "radial")
	if err != nil {
		fmt.Println("FindAllIdsForField failed:", err)
	}

	fmt.Println("\n======================= Planes with enginetype 'radial' ===========================================================\n")
	for _, id := range ids {
		plane = Plane{}

		err = db.Find("planes", &plane, id)
		if err != nil {
			fmt.Println("Find failed:", err)
		}

		fmt.Printf("%#v\n", plane.Name)
	}

	//
	// Create
	//
	plane = Plane{Name: "Test", Speed: 111, Range: 111, EngineType: "radial", Tags: []string{"test"}}
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

	plane = Plane{}

	err = db.Find("planes", &plane, id)
	if err != nil {
		fmt.Println("Find failed:", err)
	}

	fmt.Printf("\n======================= Plane with updated speed ==============================================================\n")
	fmt.Printf("%#v\n", plane.Speed)

	//
	// Delete
	//
	err = db.Delete("planes", id)
	if err != nil {
		fmt.Println("Delete failed:", err)
	}
}
