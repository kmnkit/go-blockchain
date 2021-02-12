package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

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

func outputFinally(c chan []string, j extractedJob) {
	jobSlice := []string{
		"https://kr.indeed.com/viewjob?jk=" + j.id,
		j.title,
		j.location,
		j.salary,
		j.summary,
	}
	c <- jobSlice
}

// 구인 내용들을 csv 파일에 쓴다.
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	utils.CheckErr(err)
	w := csv.NewWriter(file)

	defer w.Flush()

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}
	wErr := w.Write(headers)
	utils.CheckErr(wErr)
	c := make(chan []string)
	// job의 type은 extractedJob
	for _, job := range jobs {
		go outputFinally(c, job)
	}
	for i := 0; i < len(jobs); i++ {
		jwErr := w.Write(<-c)
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
	start := time.Now()
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
	elapsedTime := time.Since(start)
	fmt.Println("실행시간:", elapsedTime.Seconds())
	fmt.Println("Done, Extracted!", len(jobs))
}
