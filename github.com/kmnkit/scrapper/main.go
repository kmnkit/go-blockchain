package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/kmnkit/scrapper/utils"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&l=seoul"

// 구인 내용들을 csv 파일에 쓴다.
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	utils.CheckErr(err)
	w := csv.NewWriter(file)

	defer w.Flush()

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}
	wErr := w.Write(headers)
	utils.CheckErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{
			"https://kr.indeed.com/viewjob?jk=" + job.id,
			job.title,
			job.location,
			job.salary,
			job.summary,
		}
		jwErr := w.Write(jobSlice)
		utils.CheckErr(jwErr)
	}
}

// 구인 내용들을 모두 얻어온다.
func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := utils.CleanString(card.Find(".title>a").Text())
	location := utils.CleanString(card.Find(".sjcl>.location").Text())
	salary := utils.CleanString(card.Find(".salaryText").Text())
	summary := utils.CleanString(card.Find(".summary").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		summary:  summary,
		salary:   salary,
	}
}

// 한 페이지 내의 구인 내용들을 모두 얻는다.
func getPage(page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + "&start=" + strconv.Itoa(page*10)

	res, err := http.Get(pageURL)
	utils.CheckErr(err)
	utils.CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

// 페이지 내의 모든 페이지 수를 얻는다.
func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	utils.CheckErr(err)
	utils.CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckErr(err)

	doc.Find(".pagination-list").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("li").Length()
	})

	return pages
}

func main() {
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages() - 1 // 여기선 페이지 수만 얻음. 화살표 버튼이 있어서 하나 빼야 함.

	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}
	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}
	writeJobs(jobs)
	fmt.Println("Done, Extracted!", len(jobs))
}
