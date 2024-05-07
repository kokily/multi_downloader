package main

import (
	"fmt"
	scrapping "goproject/multi_downloader/go_pkg"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/cavaliergopher/grab/v3"
)

func main() {
	data := scrapping.WebScrapping()

	target := strings.Split(data, ",")
	count := strings.Count(data, ",")

	request := make(chan *grab.Request)
	response := make(chan *grab.Response)

	// Start 4 Workers
	client := grab.NewClient()

	wg := sync.WaitGroup{}

	for i := 0; i < 4; i++ {
		wg.Add(1)

		go func() {
			client.DoChannel(request, response)
			wg.Done()
		}()
	}

	go func() {
		for i := 0; i < count; i++ {
			url := target[i]
			req, err := grab.NewRequest("./data", url)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(strconv.Itoa(count) + " 파일 중 " + strconv.Itoa(i+1) + "번째 파일 다운로드 중")

			request <- req
		}

		close(request)

		wg.Wait()
		close(response)
	}()

	for resp := range response {
		if err := resp.Err(); err != nil {
			log.Fatal()
		}
	}

	fmt.Printf("%d 개 파일 다운로드 완료\n", count)
}
