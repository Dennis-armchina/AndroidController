package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"time"
)

//ApkResponse base format
type ApkResponse struct {
	Code   int
	Msg    string
	Data   []string
	Method string
	Time   string
	Pkg    string
}

//status code
const (
	Success      = 200
	AlreadyStart = 201
	BadRequest   = 400
)

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
		response := ApkResponse{
			Code:   BadRequest,
			Msg:    "error: " + err.Error(),
			Method: "install-package",
			Time:   time.Now().String(),
			Pkg:    apkURL,
		}
		resJSON, _ := json.Marshal(response)
		io.WriteString(w, string(resJSON))
	} else {
		response := ApkResponse{
			Code:   Success,
			Msg:    "success",
			Pkg:    apkURL,
			Method: "install-package",
			Time:   time.Now().String(),
		}
		resJSON, _ := json.Marshal(response)
		io.WriteString(w, string(resJSON))
	}
}
