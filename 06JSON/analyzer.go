package main

//2.1 Gom tất cả những người trong cùng một thành phố lại
func GroupPeopleByCity(p []Person) (result map[string][]Person) {
	result = make(map[string][]Person)
	for _, person := range p {
		result[person.City] = append(result[person.City], person)
	}
	return result
}

//2.2 Nhóm các nghề nghiệp và đếm số người làm
func GroupPeopleByJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.Job]++
	}
	return result
}

//2.3 Tìm 5 nghề có nhiều người làm nhất, đếm từ cao xuống thấp
//https://code-maven.com/slides/golang/sort-map-by-value
func Top5JobsByNumer(p []Person) []KeyValue {
	jobCount := GroupPeopleByJob(p)
	sortedJobByNumberDESC := SortMapByValue(jobCount, false)
	return sortedJobByNumberDESC[0:5]
}

//2.4 Tìm 5 thành phố có nhiều người trong danh sách ở nhất, đếm từ cao xuống thấp
func Top5CitiesByNumber(p []Person) []KeyValue {
	//Đếm số dân trong một thành phố
	cityPopulation := make(map[string]int)
	for _, person := range p {
		cityPopulation[person.City]++
	}

	sortedPopulationByNumberDESC := SortMapByValue(cityPopulation, false)
	return sortedPopulationByNumberDESC[0:5]
}

//2.5 Trong mỗi thành phố, hãy tìm ra nghề nào được làm nhiều nhất
func TopJobByNumerInEachCity(p []Person) (result []CityJobCount) {
	//Bước 1: GroupPeopleByCity
	groupPeopleByCity := GroupPeopleByCity(p)

	//Bước 2: tìm nghề nhiều nhất trong một thành phố
	for city, peopleInCity := range groupPeopleByCity {
		jobHasMaxCount, maxJobCount := TopJobInOneCity(peopleInCity)
		result = append(result, CityJobCount{City: city, Job: jobHasMaxCount, Count: maxJobCount})
	}
	return
}

//Truyền vào danh sách người trong một thành phố, trả về nghề có số lượng cao nhất
func TopJobInOneCity(peopleInACity []Person) (string, int) {
	jobCountInOneCity := make(map[string]int)
	jobHasMaxCount := ""
	maxJobCount := 0

	for _, person := range peopleInACity {
		jobCountInOneCity[person.Job]++
		if jobCountInOneCity[person.Job] > maxJobCount {
			jobHasMaxCount = person.Job
			maxJobCount = jobCountInOneCity[person.Job]
		}
	}
	return jobHasMaxCount, maxJobCount
}
