package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"time"
)

type installResponse struct {
	Code   int
	Msg    string
	Pkg    string
	Method string
	Time   string
}

//InstallApk method parse the path of Apk and install
func InstallApk(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	head := "apk_url"
	apkURL := parseOpen(r.Form, head)
	// fmt.Fprintln(w, "initializing...")
	cmd := exec.Command("sh", "/system/bin/pm", "install", "-d", "-r", apkURL)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error to install:", err)
		response := installResponse{
			Code:   400,
			Msg:    "error: " + err.Error(),
			Pkg:    apkURL,
			Method: "install-package",
			Time:   time.Now().String(),
		}
		resJSON, _ := json.Marshal(response)
		io.WriteString(w, string(resJSON))
	} else {
		// fmt.Fprintln(w, "install successfully")
		response := installResponse{
			Code:   200,
			Msg:    "success",
			Pkg:    apkURL,
			Method: "install-package",
			Time:   time.Now().String(),
		}
		resJSON, _ := json.Marshal(response)
		io.WriteString(w, string(resJSON))
	}
}
