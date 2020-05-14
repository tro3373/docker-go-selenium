package main

import (
	"fmt"
	"os"

	"github.com/sclevine/agouti"
)

func main() {

	driver := agouti.ChromeDriver(
		agouti.Browser("chrome"),
		agouti.Debug,
		agouti.ChromeOptions("prefs", map[string]interface{}{
			"download.directory_upgrade": true,
			// "download.default_directory":              "/home/xxx/Downloads", // ダウンロードするディレクトリ
			"download.prompt_for_download":            false,
			"plugins.always_open_pdf_externally":      true,
			"plugins.plugins_disabled":                "Chrome PDF Viewer",
			"profile.default_content_settings.popups": 0,
		}),
		agouti.ChromeOptions("args", []string{
			// "--start-maximized", // ブラウザのサイズを指定する
			"--headless",
			"--disable-gpu",
			// "--debuggerAddress=http://localhost:4444/wd/hub",
			"--debuggerAddress=http://hub:4444/wd/hub",
			// "--kiosk-printing",
			// "--user-data-dir=/path/to/profile",
		}),
	)

	defer driver.Stop()
	// driver.Start()

	fmt.Println("> Starting ..")
	if err := driver.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "> Start Err: %s\n", err)
		return
	}

	fmt.Println("> NewPaging ..")
	page, err := driver.NewPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "> NewPage Err: %s\n", err)
		return
	}

	fmt.Println("> Navigating ..")
	page.Navigate("https://example.com")
	fmt.Println(page.Title())
}
