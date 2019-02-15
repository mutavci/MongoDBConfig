package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/mholt/archiver"
)

//Download mongodb dosyası var mı ? yok mu ? kontrolu yapar
func Download(dbversion string) {
	splits := strings.Trim(dbversion, ".zip")
	if _, err := os.Stat(splits); os.IsNotExist(err) {
		fmt.Println(err)
		dowloandInPath(dbversion)
	} else {
		var question string
		fmt.Print("Daha Önce indirilmiştir. Üzerine yazalım mı ? (Y/N) :")
		fmt.Scanf("%s", &question)
		if question == "Y" {
			dowloandInPath(dbversion)

		}

	}

}

//downloadInPath İndirme işlemi gerçekleştirir.
func dowloandInPath(dbversion string) {
	client := grab.NewClient()
	req, _ := grab.NewRequest(".", "https://fastdl.mongodb.org/win32/"+dbversion)

	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

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
			break Loop
		}
	}
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)
	err := archiver.Unarchive(dbversion, "./")
	if err != nil {
		fmt.Println(err)
	}
}
