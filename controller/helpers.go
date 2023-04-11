package controller

// 1
import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"CronJob/model"
)

// 2
// func hello(name string) {
// 	message := fmt.Sprintf("Hi, %v", name)
// 	fmt.Println(message)
// }

func downloadGitHubRepoAsZip(url string, dest string) error {
	// Create the destination directory if it doesn't exist
	err := os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		return err
	}

	// Send HTTP GET request to the GitHub URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Copy the contents of the response body to a buffer
	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return err
	}

	// Open a new zip archive for reading
	archive, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		return err
	}

	// Iterate over each file in the archive
	for _, file := range archive.File {
		// Open the file in the archive
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// Create the destination file
		path := filepath.Join(dest, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
			continue
		}

		writer, err := os.Create(path)
		if err != nil {
			return err
		}
		defer writer.Close()

		// Copy the contents of the file to the destination file
		_, err = io.Copy(writer, rc)
		if err != nil {
			return err
		}
	}

	return nil
}

func getzipfile(repoUrl string) (int64, *os.File, error) {
	// Get the repository URL from the user

	// Replace "github.com" with "codeload.github.com" to download a ZIP file
	zipUrl := strings.Replace(repoUrl, ".git", "/archive/refs/heads/main.zip", 1)

	// Download the ZIP file
	resp, err := http.Get(zipUrl)
	if err != nil {
		fmt.Println("Error downloading ZIP file:", err)
		return 0, nil, err
	}
	defer resp.Body.Close()

	component := path.Base(repoUrl)

	// Remove the ".git" extension
	component = strings.TrimSuffix(component, ".git")
	// Create a new ZIP file
	zipFile, err := os.Create(component + ".zip")
	if err != nil {
		fmt.Println("Error creating ZIP file:", err)
		return 0, nil, err
	}
	defer zipFile.Close()

	// Copy the downloaded ZIP file to the new ZIP file
	zsize, err := io.Copy(zipFile, resp.Body)
	if err != nil {
		fmt.Println("Error copying ZIP file:", err)
		return 0, nil, err
	}
	return zsize, zipFile, nil
}

func runCronJobs(request map[string]interface{}, ctx context.Context) {
	// 3
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task stopped")
			return
		default:
			// fmt.Println("Task running...")
			// time.Sleep(time.Second)

			_, zFile, zrr := getzipfile(request["repositaryLink"].(string))

			log.Println(zFile.Name())

			if zrr != nil {
				log.Fatal(zrr)
			}

			// err := downloadGitHubRepoAsZip(request["repositaryLink"].(string), request["backupLocation"].(string))
			// if err != nil {
			// 	panic(err)
			// }
			waitFreq := request["backupFrequency"].(float64)
			time.Sleep(time.Duration(waitFreq) * time.Second)
		}
	}

}

func stopHelper(cancelMap map[model.CronJob]context.CancelFunc, cronJob model.CronJob) {
	var cancel = cancelMap[cronJob]
	cancel()
	fmt.Println("Task stopping...")
}

// 6
// func cron() {
// 	runCronJobs()
// }
