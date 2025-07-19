package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// encodeJson()
	decodeJson()
}

func encodeJson() {

	courses := []course{
		{"React js", 1000, "123", []string{"js", "client"}},
		{"Next js", 1500, "abc", []string{"js", "server"}},
		{"Express js", 1200, "abc123", nil},
	}

	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {
	jsonDatafromWeb := []byte(`
		{
                "coursename": "React js",
                "price": 1000,
                "tags": ["js","client"]
        }
	`)

	var newCourse course

	check := json.Valid(jsonDatafromWeb) // check json is valid or not

	if check {
		fmt.Println("JSON is Valid.")
		json.Unmarshal(jsonDatafromWeb, &newCourse) // string json to json format
		fmt.Printf("%#v\n", newCourse)
	} else {
		fmt.Println("JSON is not valid.")
	}

	// some times we just want to add data as key value pair
	fmt.Println("JSON as map.")
	var mapNewCourse map[string]interface{}
	json.Unmarshal(jsonDatafromWeb, &mapNewCourse)
	fmt.Printf("%#v\n", mapNewCourse)

}
