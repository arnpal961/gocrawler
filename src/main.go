package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	/*
		// Example of readTextFile
		fmt.Println("Started project to crawl web concurrently")
		fc := readTextFile("../test/test.txt")
		fmt.Println(fc)

		// Example of writeTextFile
		writeString := "Hello How are you ?"
		writeTextFile("../test/writeTest.txt", writeString)


		// jsonMovieExample
		jsonMovieExample(true)
		jsonMovieExample(false)

		decodeJSONFile("../test/fortune_email_pat.json")

		var urls = []string{"https://www.youtube.com/", "https://www.pdfdrive.com/"}
		fetchURLs(urls)
	*/

}
func checkError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func readTextFile(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(b)
	return fileContent
}

func writeTextFile(path string, content string) {
	b := []byte(content)
	err := ioutil.WriteFile(path, b, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func jsonMovieExample(isIndent bool) {
	type Movie struct {
		Title  string
		Year   int  `json:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}}}
	var (
		data []byte
		err  error
	)

	if isIndent {
		data, err = json.MarshalIndent(movies, "", "	")
	} else {
		data, err = json.Marshal(movies)
	}
	checkError(err, "JSON marshalling failed")
	fmt.Printf("%s\n", data)

}

// Reading a json file
func decodeJSONFile(filePath string) {
	rb, err := ioutil.ReadFile(filePath)
	checkError(err, "File opening failed")
	var result map[string]interface{}
	json.Unmarshal(rb, &result)
	for k, v := range result {
		fmt.Println(k, "---->", v.(string))
	}
}

func encodeJSONData(filepath string, wb []byte) {

}

func fetchURL(url string, ch chan<- []byte) {
	resp, err := http.Get(url)
	checkError(err, "Failed during fetching")
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("%T\n", resp.Body)
		body, err := ioutil.ReadAll(resp.Body)
		checkError(err, "Failed")
		resp.Body.Close()
		ch <- body
	}
}

func fetchURLs(urls []string) {
	ch := make(chan []byte)
	for i, url := range urls {
		fmt.Println("Launching routine no -> ", i)
		go fetchURL(url, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(len(<-ch))
	}
}

// visit() appends to links each link found in
// n and returns the result
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
