package swagger

import (
	"fmt"
	"net/http"
	"os/exec"
)

//Install method parse the path of Apk and install
func Install(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	head := "apk_url"
	apkURL := parseOpen(r.Form, head)
	fmt.Fprintln(w, "initializing...")
	cmd := exec.Command("sh", "/system/bin/pm", "install", "-d", "-r", apkURL)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error to install:", err)
	} else {
		fmt.Fprintln(w, "install successfully")
		fmt.Fprintln(w, "install url:", apkURL)
	}
}
