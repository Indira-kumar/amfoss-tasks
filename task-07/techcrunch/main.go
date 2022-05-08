package main

import(
  "net/http"
  "fmt"
  "os"
  "strings"
  "encoding/csv"
  "github.com/PuerkitoBio/goquery"
)
func writeFile(data, filename string) {
  file , error := os.Create(filename)
  defer file.Close()
  checkError(error)
  file.WriteString(data)
}
func checkError(err error){
  if err != nil {
    fmt.Println(err)
  }
}
func main(){
  url := "https://techcrunch.com/"

  response, error := http.Get(url)
  defer response.Body.Close()
  checkError(error)

  if response.StatusCode > 400 {
    fmt.Println("Status Code: ",response.StatusCode)
  }
  doc, error := goquery.NewDocumentFromReader(response.Body)
  checkError(error)
  file, error := os.Create("posts.csv")
  checkError(error)
  writer := csv.NewWriter(file)

  doc.Find("div.river").Find("div.post-block").Each(func(index int, item *goquery.Selection){
    h2 := item.Find("h2")
    heading := strings.TrimSpace(h2.Text())
    url, _ := h2.Find("a").Attr("href")
    fmt.Println(heading, url)
    content := []string{heading, url}
    writer.Write(content)
  })
  //checkError(error)
  //writeFile(river, "index.html")
writer.Flush()
}
