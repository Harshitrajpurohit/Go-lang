package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// course model
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake data
var courses []Course

// middleware or helper
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

// main function
func main() {

	r := mux.NewRouter()

	//seeding data
	courses = append(courses, Course{CourseId: "1", CourseName: "React", CoursePrice: 299, Author: &Author{Fullname: "Harshit", Website: "google.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Next", CoursePrice: 499, Author: &Author{Fullname: "Harshit Rajpurohit", Website: "mygoogle.com"}})

	r.HandleFunc("/", serveHomeRoute).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen on port 8080
	fmt.Println("server is listning on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}

// controllers - files

// serve home page
func serveHomeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is home route.</h1>"))
}

// get courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

// get one course
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found.")
	return
}

// create one course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// check is req.body is null
	if r.Body == nil {
		json.NewEncoder(w).Encode("send some data")
		return
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("json data is enpty")
		return
	}

	// check course is unique or not
	for _, co := range courses {
		if co.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("course is already available.")
			return
		}
	}

	// generate new ID
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	// append in courses
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

// update course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// check is req.body is null
	if r.Body == nil {
		json.NewEncoder(w).Encode("send some data")
		return
	}
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("json data is enpty")
		return
	}

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// if id not found
	json.NewEncoder(w).Encode("data not found in database.")

}

// delete course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("course not found. ")
}
