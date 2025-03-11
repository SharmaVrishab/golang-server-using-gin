package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// model for course
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json :"author"`
}
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
// var courses []Course
// seeding value in the main method
var courses = []Course{
	{
		CourseId:    "C101",
		CourseName:  "Go Programming",
		CoursePrice: 999,
		Author:      &Author{FullName: "John Doe", Website: "https://johndoe.com"},
	},
	{
		CourseId:    "C102",
		CourseName:  "DevOps Bootcamp",
		CoursePrice: 1499,
		Author:      &Author{FullName: "Jane Smith", Website: "https://janesmith.com"},
	},
	{
		CourseId:    "C103",
		CourseName:  "Cloud Fundamentals",
		CoursePrice: 1999,
		Author:      &Author{FullName: "Alice Brown", Website: "https://alicebrown.com"},
	},
}

// middleware , helper
func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}
func main() {

	r := gin.Default()
	r.GET("/", serverHome)
	r.GET("/courses", getAllCourse)
	r.GET("/course/:id", getOneCourse)
	r.POST("/createcourse", createOneCourse)
	r.POST("/updatecourse", updateOneCourse)
	r.POST("/deletecourse", deleteOneCourse)
	r.Run(":4000")

}

// controller file
// server home route
func serverHome(c *gin.Context) {
	c.Data(200, "text/html", []byte("<h1>WELCOME TO THE HOME PAGE</h1>"))
}

func getAllCourse(c *gin.Context) {
	fmt.Println("get all courses")
	c.Header("Content-Type", "application/json")
	// if you want to retunrn vlaue directly do this or
	// c.JSON(200, courses)
	// c.JSON(200, gin.H{"course": courses})
	// beautify json
	c.IndentedJSON(200, gin.H{"courses": courses})
}

func getOneCourse(c *gin.Context) {
	fmt.Println("get one couse based on id")
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	for _, course := range courses {
		if course.CourseId == id {
			c.IndentedJSON(200, gin.H{"data": course})
			return
		}
	}
	c.IndentedJSON(404, gin.H{
		"message": "no course was found for the given id",
	})
}

func createOneCourse(c *gin.Context) {
	fmt.Println("create one course")
	c.Header("Content-Type", "application/json")
	var course Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(404, gin.H{"message": "Please send some data"})
		return
	}
	//  generate UNIQUE ID
	// CONVERT TO STRING
	// APPEND

	// for getting random value
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	course.CourseId = strconv.Itoa(r.Intn(100))
	courses = append(courses, course)
	c.IndentedJSON(200, course)
}

func updateOneCourse(c *gin.Context) {
	fmt.Println("update one course")
	c.Header("Content-Type", "application/json")
	// grab id from request
	id := c.Param("id")
	// loop and remove the id and than the add the value from body
	var new_course Course
	for index, course := range courses {
		if course.CourseId == id {
			courses = append(courses[:index], courses[index+1:]...)
			if err := c.ShouldBindJSON(&new_course); err != nil {
				c.JSON(404, gin.H{"message": "Please send some data"})
				return
			}
			course.CourseId = id
			courses = append(courses, new_course)
			c.IndentedJSON(200, new_course)
		}
	}
}

func deleteOneCourse(c *gin.Context) {
	fmt.Println("delete one course")
	c.Header("Content-Type", "applciation/json")
	id := c.Param("id")

	for index, course := range courses {
		if course.CourseId == id {
			//  you have found the id
			courses = append(courses[:index], courses[index+1:]...)
			c.JSON(200, gin.H{"message": "data was deleted"})
			break
		}
	}
	c.JSON(404, gin.H{"message": "not found "})
}
