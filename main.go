package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}
type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

func main() {
	var r RSS

	data:= readGoogleTrends()
	//fmt.Println(string(data)) data are in XML format

	err:=xml.Unmarshal(data,&r)

	if err!=nil{
		fmt.Println("error:",err)
		os.Exit(1)
	}


}

func getGoogleTrends() *http.Response {
	resp,err:=http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")

	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}

	return resp
}

func readGoogleTrends() []byte {
	resp:= getGoogleTrends()
	data, err:=ioutil.ReadAll(resp.Body)

	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}

	return data

}