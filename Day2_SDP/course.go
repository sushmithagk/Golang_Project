package main

import (
	// "fmt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Course struct {
	Id          string
	Name        string
	Duration    int
	Trainername string
	Level       string
	Description string
	Price       float32
}

func readAllCourses(c *gin.Context) {
	courses := []Course{
		{Id: "1", Name: "Machine Learning",
			Duration:    20,
			Trainername: "Nithin", Level: " Beginner", Description: "Tech stack related to machine learning",
			Price: 4000.00},
		{Id: "2", Name: "Generative AI",
			Duration:    20,
			Trainername: "Mahesh", Level: " Intermediate", Description: "Learn about AI",
			Price: 5000.00},
	}
	fmt.Println(courses)

	c.JSON(http.StatusOK, courses)
}

func readCourseById(c *gin.Context) {
	id := c.Param("id")
	course := Course{Id: id, Name: "Generative AI",
		Duration:    20,
		Trainername: "Mahesh", Level: " Intermediate", Description: "Learn about AI ",
		Price: 5000.00}
	c.JSON(http.StatusOK, course)
}

func createCourse(c *gin.Context) {
	var jbodyCourse Course
	err := c.BindJSON(&jbodyCourse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error." + err.Error()})
		return
	}
	createdCouse := Course{Id: "10", Name: "Data Science",
		Duration:    20,
		Trainername: "Nithin", Level: " Beginner", Description: "Tech stack related to data science",
		Price: 4500.00}
	c.JSON(http.StatusCreated, gin.H{"message": "Course created succeessfully", "course": createdCouse})
}

func updateCourse(c *gin.Context) {
	id := c.Param("id")
	var jbodyCourse Course
	err := c.BindJSON(&jbodyCourse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error." + err.Error()})
		return
	}
	updatedCourse := Course{Id: id, Name: "Machine Learning",
		Duration:    20,
		Trainername: "Nithin", Level: " Beginner", Description: "Tech stack related to machine learning",
		Price: 2000.00}
	c.JSON(http.StatusOK, gin.H{"message": "Course updated succeessfully", "course": updatedCourse})
}

func deleteCourse(c *gin.Context) {
	// id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted succeessfully"})
}
func main() {
	//router
	r := gin.Default()
	//routes | API endpoints
	r.GET("/Courses", readAllCourses)
	r.GET("/Courses/:id", readCourseById)
	r.POST("/Courses", createCourse)
	r.PUT("/Courses/:id", updateCourse)
	r.DELETE("/Courses/:id", deleteCourse)

	//server (default portal 8080)
	r.Run(":8080")
}
