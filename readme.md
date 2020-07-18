# Nomadcoders.co의 "쉽고 빠른 Go 시작하기" 강의와 관련된 코드입니다

- Python으로 Web Scraper 만들기 강의와 같이 Go를 이용해 Web Scraper와 웹서버를 만들어보는 강의 입니다. 
- Scraper를 만들기 앞서서 간단한 프로젝트를 통해 go의 syntax를 살펴봤습니다.
- Python으로 Web Scraper를 Go로 똑같이 만들어 보면서 Go를 학습하기 위해 이 프로젝트를 진행하였습니다. 
- Indeed에서 python 키워드로 job을 검색하여, job list를 가져오는 scraper로, apply link, title, location, summary, salary 등을 가져옵니다. 그리고 가져온 정보를 csv 파일로 저장합니다. (Directory : jobScraper)
- 각 페이지 별로, 그리고 각 jobs card 별로 순차적으로 scraping 하는 기존 python scraper와 달리 goroutines와 channel을 이용해 scraper의 속도를 비약적으로 향상시킬 수 있었습니다.
- 추후 파일 쓰기 부분도 goroutines와 channel을 이용한다면 더욱 속도를 향상시킬 수 있을 것입니다.
- echo를 이용한 간단한 웹서버를 만들면서 python 키워드로 job scraping 하는 것에서 원하는 키워드를 넣어 scraping 하는 것으로 기능을 변경하였습니다. 


## 학습한 내용 

- golang syntax 
- goroutines, channel
- URL checkers 
- Job scraper(goquery)
- strings
- go echo (웹서버 만들기)

## 완성된 프로젝트 

검색어 입력 <br>
----------------
![main](/jobScraper/main.png)

csv에 저장 <br>
----------------
![csv](/jobScraper/csv.png)