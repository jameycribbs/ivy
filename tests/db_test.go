package ivy

import (
	"fmt"
	"github.com/jameycribbs/ivy"
	"io/ioutil"
	"os"
	"testing"
)

var db *ivy.DB

func init() {
	var err error

	files, _ := ioutil.ReadDir("data/foos")
	for _, file := range files {
		if !file.IsDir() {
			err = os.Remove("data/foos/" + file.Name())
			if err != nil {
				fmt.Println("Failed to remove file", err)
			}
		}
	}

	fieldsToIndex := make(map[string][]string)
	fieldsToIndex["foos"] = []string{"tags", "bar"}

	db, err = ivy.OpenDB("data", fieldsToIndex)
	if err != nil {
		fmt.Println("Failed to open database:", err)
		os.Exit(1)
	}
}

func TestCreate(t *testing.T) {
	foo := Foo{Bar: "test", Tags: []string{"test"}}
	id, err := db.Create("foos", foo)
	if err != nil {
		t.Error("Find failed:", err)
	}

	foo = Foo{}

	err = db.Find("foos", &foo, id)
	if err != nil {
		t.Error("Find failed:", err)
	}

	if foo.Bar != "test" {
		t.Error("Expected 'test', got ", foo.Bar)
	}

	if foo.Tags[0] != "test" {
		t.Error("Expected first tag to be 'test', got ", foo.Tags[0])
	}
}

func TestUpdate(t *testing.T) {
	foo := Foo{Bar: "test", Tags: []string{"test"}}
	id, err := db.Create("foos", foo)
	if err != nil {
		t.Error("Create failed:", err)
	}

	foo.Bar = "test2"
	foo.Tags = []string{"test2", "test3"}

	err = db.Update("foos", foo, id)
	if err != nil {
		t.Error("Update failed:", err)
	}

	foo = Foo{}

	err = db.Find("foos", &foo, id)
	if err != nil {
		t.Error("Find failed:", err)
	}

	if foo.Bar != "test2" {
		t.Error("Expected 'test2', got ", foo.Bar)
	}

	if foo.Tags[0] != "test2" {
		t.Error("Expected first tag to be 'test2', got ", foo.Tags[0])
	}

	if foo.Tags[1] != "test3" {
		t.Error("Expected second tag to be 'test3', got ", foo.Tags[1])
	}
}

func TestDelete(t *testing.T) {
	foo := Foo{Bar: "test", Tags: []string{"test"}}
	id, err := db.Create("foos", foo)
	if err != nil {
		t.Error("Find failed:", err)
	}

	err = db.Delete("foos", id)
	if err != nil {
		t.Error("Delete failed:", err)
	}

	foo = Foo{}

	err = db.Find("foos", &foo, id)
	if err != nil {
		if !os.IsNotExist(err) {
			t.Error("Expected Find error to be 'file does not exist', got ", err)
		}
	} else {
		t.Error("Expected Find error, got no error.")
	}

}

//=============================================================================
// Setup Stuff
//=============================================================================
type Foo struct {
	FileId string   `json:"-"`
	Bar    string   `json:"bar"`
	Tags   []string `json:"tags"`
}

func (foo *Foo) AfterFind(db *ivy.DB, fileId string) {
	*foo = Foo(*foo)

	foo.FileId = fileId
}
