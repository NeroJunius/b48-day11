package main 

import (
	"fmt"
	"html/template"
	"net/http"
	

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/assets", "assets")

	e.GET("/", Home)
	e.GET("/contactMe", contactMe)
	e.GET("/testimonial", testimonials)
	e.GET("/createProject", createProject)
	e.GET("/projectDetail/:id", projectDetail)

	e.POST("/add-project", addProject)

	fmt.Println("server started on port 5900")
	e.Logger.Fatal(e.Start("localhost:5900"))
}

func Home(c echo.Context) error {
	 tmpl, err := template.ParseFiles("tabs/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return tmpl.Execute(c.Response(), nil)
}

func contactMe(c echo.Context) error {
	var tmpl, err = template.ParseFiles("tabs/contact.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func createProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("tabs/project.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func testimonials(c echo.Context) error {
	var tmpl, err = template.ParseFiles("tabs/testimonial.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id := c.Param("id")

	var tmpl, err = template.ParseFiles("tabs/project-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data := map[string]interface{}{
		"Id":           id,
		"Title":        "Ameno ameno latire Latiremo Dori me",
		"detailDate":   "durasi : 4 bulan",
		"projectDescription": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin quis risus ut mi euismod sodales. Mauris id quam ut massa sodales faucibus consectetur sit amet dolor. ",
		"nodeJS":       "nodejs",
		"reactJS":      "reactjs",
		"nextJS":       "nextjs",
		"typeScript":   "typescript",
		"img": "quotes",

	}

	return tmpl.Execute(c.Response(), data)
}

func addProject(c echo.Context) error {
	title := c.FormValue("projectName")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	DescProjects := c.FormValue("projectDescription")
	img := c.FormValue("imageProject")
	
	fmt.Println("projectName:", title)
	fmt.Println("startDate:", startDate)
	fmt.Println("endDate:", endDate)
	fmt.Println("projectDescription:", DescProjects)
	fmt.Println("imageProject:", img)
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}


