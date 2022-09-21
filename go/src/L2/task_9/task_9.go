package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/cavaliergopher/grab/v3"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var mainDir = "files"

func getFilesName(strBody, sep string) []string {
	strSrc := strings.Split(strBody, sep)

	srcFilesName := make([]string, 0)

	for i := 1; i < len(strSrc); i++ {
		for j := 0; j < len(strSrc[i]); j++ {
			if strSrc[i][j] == '"' {
				fmt.Println("src = ", strSrc[i][0:j])
				srcFilesName = append(srcFilesName, strSrc[i][0:j])
				break
			}
		}
	}
	return srcFilesName
}

func getUrls(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err = ", err)
	}

	// Defer close the response Body.
	defer resp.Body.Close()

	// Read everything from Body.

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	urls := make([]string, 0)

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		atr, err1 := s.Attr("src")
		if !err1 {
			fmt.Println("not exist")
		} else {
			if atr[:8] != "https://" && atr[:7] != "http://" {
				urls = append(urls, atr)
				fmt.Printf("img src  %d: %s\n", i, atr)
			}
		}
	})

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		atr, err1 := s.Attr("src")
		if !err1 {
			fmt.Println("not exist")
		} else {
			if atr[:8] != "https://" && atr[:7] != "http://" {
				urls = append(urls, atr)
				fmt.Printf("js src  %d: %s\n", i, atr)
			}
		}
	})

	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		if atr1, _ := s.Attr("type"); atr1 == "image/png" || atr1 == "text/css" {
			atr, err1 := s.Attr("href")
			if !err1 {
				fmt.Println("not exist")
			} else {
				if atr[:8] != "https://" && atr[:7] != "http://" {
					urls = append(urls, atr)
					fmt.Printf("link src %d: %s\n", i, atr)
				}
			}
		}
	})

	fmt.Println("urls = ", urls)
	return urls

	//body, _ := ioutil.ReadAll(resp.Body)
	//
	//// Convert body bytes to string.
	//bodyText := string(body)
	//
	//fmt.Println("DONE")
	//fmt.Println("Response length: ", len(bodyText))
	//
	//// Convert to rune slice to take substring.
	//
	//return bodyText
}

func makeFolderFromFile(file string) (string, string) {
	sepFile := strings.Split(file, "/")
	//fmt.Println("sepFIle = ", sepFile)

	res := ""
	for i := 0; i < len(sepFile)-1; i++ {
		if sepFile[i] != "" {
			res += sepFile[i] + "/"
		}
	}

	return res, sepFile[len(sepFile)-1]
}

func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func MakeDir(dir string) error {
	if !FileIsExisted(dir) {
		if err := os.MkdirAll(dir, 0777); err != nil { //os.ModePerm
			fmt.Println("MakeDir failed:", err)
			return err
		}
	}
	return nil
}

func downloadFile(url, file string) {
	folder, fileName := makeFolderFromFile(file)

	client := grab.NewClient()

	//folder1 := folder
	//if len(folder1) > 0 {
	//	if folder1[len(folder1)-1] == '/' {
	//		folder1 = folder1[:len(folder1)-1]
	//	}
	//}

	err := MakeDir(mainDir + "/" + folder)
	if err != nil {
		fmt.Println("err with creat dir, err = ", err.Error())
		return
	}

	fmt.Println("download file = ", url+folder+fileName)
	fmt.Println("file folder = ", mainDir+"/"+folder)

	req, _ := grab.NewRequest(mainDir+"/"+folder, url+folder+fileName)

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Download saved to ./%v \n\n\n", resp.Filename)

	// Output:
	// Downloading http://www.golang-book.com/public/pdf/gobook.pdf...
	//   200 OK
	//   transferred 42970 / 2893557 bytes (1.49%)
	//   transferred 1207474 / 2893557 bytes (41.73%)
	//   transferred 2758210 / 2893557 bytes (95.32%)
	// Download saved to ./gobook.pdf

}

func downloadPlaneFile(url string) {

	client := grab.NewClient()

	req, _ := grab.NewRequest(mainDir+"/", url)

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Download saved to ./%v \n\n\n", resp.Filename)

}

func main() {
	//create client
	//Get Wikipedia home page.
	argsWithProg := os.Args

	fmt.Println("args ='", argsWithProg[1], "'")
	if len(argsWithProg) < 3 {
		fmt.Println("should be some download")
		downloadPlaneFile(argsWithProg[1])
		return
	}

	//url := "https://www.google.ru/"

	urls := getUrls(argsWithProg[1])

	//srcFilesName := getFilesName(strBody, "src=\"")

	//fmt.Println("files = ", srcFilesName)

	downloadFile(argsWithProg[1], "index.html")

	for _, file := range urls {
		downloadFile(argsWithProg[1], file)

	}

	// Задержка операции после defer, обычно используется для освобождения связанных переменных

}
