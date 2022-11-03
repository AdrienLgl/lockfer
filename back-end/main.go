package main

import (
	"fmt"
	"io/ioutil"
	"lockfer/upload"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go cleanDirectory(&wg)
	go upload.SetupRoutes(&wg)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}

func cleanDirectory(wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			yesterday := time.Now().AddDate(0, 0, -1)
			files, err := ioutil.ReadDir("files/encrypted")
			if err != nil {
				fmt.Println(err.Error())
			}
			for _, f := range files {
				if f.ModTime().Before(yesterday) {
					upload.AddLog("removed old directory", "server")
					fmt.Println("Removed old directory: " + f.Name())
					os.RemoveAll("files/encrypted/" + f.Name())
				}
			}
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
