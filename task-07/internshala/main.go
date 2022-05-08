package main

import(
  "os"
  "encoding/csv"
  "github.com/gocolly/colly"
  "fmt"
  "strings"
)
func main(){
  fName :="data.csv"
  file, err := os.Create(fName)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()
  writer := csv.NewWriter(file)
  defer writer.Flush()
  c := colly.NewCollector(
    //colly.AllowedDomains("www.forbes.com/real-time-billionaires/#3779f96e3d7","forbes.com/real-time-billionaires/",
      //"www.forbes.com/real-time-billionaires/","forbes.com/","www.forbes.com/"),
      colly.AllowedDomains("internshala.com"),
  )
  c.OnHTML(".internship_meta",func(e *colly.HTMLElement){
    writer.Write([]string{
      strings.TrimSpace(e.ChildText("a")),
      strings.TrimSpace(e.ChildText("span")),
    })
  })
  c.Visit("https://internshala.com/internships")

}
