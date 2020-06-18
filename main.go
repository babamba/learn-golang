package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=frontend&l=seoul&limit=20"

func main() {
	var jobs []extractedJob
	c := make(chan []extractedJob)

	totalPages := getPages() // 페이지 숫자를 가져오고
	//fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ { // 각 페이지에서 일자리 정보를 수집
		//extractedJobs := getPage(i)           // 패아자애 았는 요소셀렉터로 extractedJob struct의 형태로 모델을 만들어주고 모델들의 배열로 만든다 .
		go getPage(i, c)
		//jobs = append(jobs, extractedJobs...) // 일자리 정보가 담긴 배열을 하나의 배열로 합침
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	//fmt.Println(jobs)
	writeJobs(jobs) // 파일을 생성하고 데이터를 집어넣고
	fmt.Println("Done, extracted ", len(jobs))
}

func getPage(page int, mainChan chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)

	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("requesting : ", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c) // 고루틴을 생성하고 채널을 통해 메시지를 전송
		// jobs = append(jobs, job)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c               // 메세지가 전달되기를 기다렸다가,
		jobs = append(jobs, job) // 메세지를 받으면 jobs 배열에 추가
	}

	mainChan <- jobs
	//return jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := cleanstring(card.Find(".title>a").Text())
	location := cleanstring(card.Find(".sjcl").Text())
	salary := cleanstring(card.Find(".salaryText").Text())
	summary := cleanstring(card.Find(".summary").Text())

	// 고루틴 채널에 extractedJob 타입으로 메시지로 전송
	c <- extractedJob{id: id, title: title, location: location, salary: salary, summary: summary}

	//fmt.Println(id, "/", title, "/", location, "/ ", salary, "/ ", summary)
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	fmt.Println(doc)

	return pages
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	utf8 := []byte{0xEF, 0xBB, 0xBF}
	w.Write([]string{string(utf8[:])})

	defer w.Flush()

	headers := []string{"Link", "Title", "Location",
		"Salary",
		"Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/jobs?q=frontend&l=seoul&vjk=" + job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status : ", res.StatusCode)
	}
}

func cleanstring(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
