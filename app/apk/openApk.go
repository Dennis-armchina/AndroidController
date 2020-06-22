package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"time"
)

//OpenApk for openeing the pacakge
func OpenApk(w http.ResponseWriter, r *http.Request) {
	head := "apk_name"
	r.ParseForm()
	apkName := parseOpen(r.Form, head)
	cmd := exec.Command("am", "start", apkName)
	val, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error : ", err)
		response := ApkResponse{
			Code:   BadRequest,
			Msg:    "fail",
			Pkg:    apkName,
			Method: "open-package",
			Time:   time.Now().String(),
		}
		resJSON, _ := json.Marshal(response)
		io.WriteString(w, string(resJSON))
	} else {
		//check if the process has already been brought to the front or with errors
		if checkWarning(string(val)) && checkError(string(val)) {
			response := ApkResponse{
				Code:   Success,
				Msg:    "success",
				Pkg:    apkName,
				Method: "open-package",
				Time:   time.Now().String(),
			}
			resJSON, _ := json.Marshal(response)
			io.WriteString(w, string(resJSON))
		} else {
			if !checkWarning(string(val)) {
				response := ApkResponse{
					Code:   AlreadyStart,
					Msg:    "warning: process already started",
					Pkg:    apkName,
					Method: "open-package",
					Time:   time.Now().String(),
				}
				resJSON, _ := json.Marshal(response)
				io.WriteString(w, string(resJSON))
			} else {
				response := ApkResponse{
					Code:   BadRequest,
					Msg:    "error： package not found",
					Pkg:    apkName,
					Method: "open-package",
					Time:   time.Now().String(),
				}
				resJSON, _ := json.Marshal(response)
				io.WriteString(w, string(resJSON))
			}
		}
	}
}

//GetDir get working directory
func GetDir() string {
	_, filename, _, ok := runtime.Caller(1)
	var cwdPath string
	if ok {
		cwdPath = path.Join(path.Dir(filename), "") // the the main function file directory
	} else {
		cwdPath = "./"
	}
	return cwdPath
}

func parseOpen(form url.Values, head string) (item string) {
	for index, apk := range form {
		if index == head {
			for _, item := range apk {
				return item
			}
		}
	}
	return
}

//check if warning msg exists
func checkWarning(input string) bool {
	head := "Warning"
	if strings.Index(input, head) == -1 {
		return true
	}
	return false
}

func checkError(input string) bool {
	head := "Error"
	if strings.Index(input, head) == -1 {
		return true
	}
	return false
}
