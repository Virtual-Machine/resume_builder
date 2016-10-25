package main

import (
	"encoding/json"
	"io/ioutil"
    "os"
    "html/template"
    "log"
)

type TechnicalSkills struct {
	Skill string	`json:"skill"`
	Level int 	`json:"level"`
	Image string `json:"image"`
}

type ApplicationSkills struct {
	Skill string	`json:"skill"`
	Level int 	`json:"level"`
	Image string `json:"image"`
}

type Certifications struct {
	Name string			`json:"name"`
}

type Experience struct {
	Company		string `json:"company"`
	Position	string `json:"position"`
	Location	string `json:"location"`
	Dates		string `json:"dates"`
	Tasks		[]string `json:"tasks"`
}

type Education struct {
	School		string 	`json:"school"`
	Location 	string  `json:"location"`
	Program 	string  `json:"program"`
	Years 		string  `json:"years"`
	Notes 		[]string `json:"notes"`
}

type Resume struct {
	Name string			`json:"name"`
	Title string		`json:"title"`
	Objective string	`json:"objective"`
	Website string		`json:"website"`
	Email string		`json:"email"`
	Phone string		`json:"phone"`
	TechnicalSkills []TechnicalSkills	`json:"technical_skills"`
	ApplicationSkills []ApplicationSkills	`json:"application_skills"`
	Certifications	[]Certifications	`json:"certifications"`
	Experience	[]Experience		`json:"experience"`
	Education			`json:"education"`
}

func main(){
	byt, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}
	fo, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	var resume Resume
	
	if err := json.Unmarshal(byt, &resume); err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.ParseFiles("index.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	
	err = tmpl.ExecuteTemplate(fo, "index.tmpl", resume)
	if err != nil {
		log.Fatal(err)
	}
}