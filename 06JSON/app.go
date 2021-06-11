package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/* https://stackoverflow.com/questions/45303326/how-to-parse-non-standard-time-format-from-json
"name":"Dee Leng",
"email":"dleng0@cocolog-nifty.com",
"job":"developer",
"gender":"Female",
"city":"London",
"salary":9662,
"birthdate":"2007-09-30" */
type Person struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Job      string `json:"job"`
	City     string `json:"city"`
	Salary   int    `json:"salary"`
	Birthday string `json:"birthdate"`
}

func (p *Person) String() string {
	return fmt.Sprintf("name: %s, email: %s, job: %s, city: %s, salary: %d, birthday: %s",
		p.Name, p.Email, p.Job, p.City, p.Salary, p.Birthday)
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("person.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened person.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var people []Person

	json.Unmarshal(byteValue, &people)

	/*
		for i := 0; i < 10; i++ {
			fmt.Println(&people[i])
		}
	*/
	peopleByCity := GroupPeopleByCity(people)
	for key, value := range peopleByCity {
		fmt.Println(key)
		for _, person := range value {
			fmt.Println("  ", (&person).Name)
		}
	}

	fmt.Println("\n2.2 Nhóm các nghề nghiệp và đếm số người làm")
	groupPeopleByJob := GroupPeopleByJob(people)
	fmt.Println(groupPeopleByJob)

	fmt.Println("\n2.3 Tìm 5 nghề có nhiều người làm nhất, đếm từ cao xuống thấp")
	top5JobsByNumer := Top5JobsByNumer(people)
	PrintSliceKeyValue(top5JobsByNumer)

	fmt.Println("\n2.4 Tìm 5 thành phố có nhiều người trong danh sách ở nhất, đếm từ cao xuống thấp")
	top5CitiesByNumber := Top5CitiesByNumber(people)
	PrintSliceKeyValue(top5CitiesByNumber)

	fmt.Println("\n2.5 Top Job in Each City")
	topJobByNumerInEachCity := TopJobByNumerInEachCity(people)
	for _, topJob := range topJobByNumerInEachCity {
		fmt.Println(topJob)
	}
}
