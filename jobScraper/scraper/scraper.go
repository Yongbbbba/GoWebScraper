package scraper

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

// Scrape Indeed by a term
func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL) // 총 페이지 숫자 확인

	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, c) // 각 페이지 별로 getPage 함수 호출

	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, Extracted", len(jobs))
}

// 각 페이지에 있는 일자리를 모두 반환
func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)

	pageURL := url + "&start=:" + strconv.Itoa(page*50)
	fmt.Println("Resquesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err) //에러가 나오면 프로그램 종료

	searchCards := doc.Find(".jobsearch-SerpJobCard") //jobsearch-SerpJobCard 클래스 가져오기

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c) // 고루틴을 생성하고 채널을 통해 getPage 함수로 메세지(job card) 전송

	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

// extractedJob struct를 반환하여 이를 job 변수에 저장
func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title>a").Text())
	location := CleanString(card.Find(".sjcl").Text())
	salary := CleanString(card.Find(".salaryText").Text())
	summary := CleanString(card.Find(".summary").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
}

// 총 페이지 숫자 가져오기
func getPages(url string) int {
	pages := 0

	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err) //에러가 나오면 프로그램 종료

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()

	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}

}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}

}

// CleanString cleans a string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")

}

//  일자리를 csv 파일로 저장, go 패키지 사용
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)
	w := csv.NewWriter(file)
	// Flush 함수는 버퍼에 있는 데이터를 파일에 입력
	defer w.Flush() // 내가 공부한 바로는 defer는 python에서 finally와 비슷한 기능을 하는 것 같고, 함수가 return 되는 시점 전에 defef에 들어간 함수가 실행이 된다.

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)

	}

}
