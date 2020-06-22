package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

type showResponse struct {
	Code   int
	Msg    string
	Data   []string
	Method string
	Time   string
}

//ShowApk method returns installed Apks
func ShowApk(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sh", "/system/bin/pm", "list", "packages", "-3")
	out, err := cmd.CombinedOutput()
	fmt.Printf("%s", out)
	if err != nil {
		fmt.Println("error : ", err)
		fmt.Println("cmd: ", cmd)
		response := showResponse{
			Code:   400,
			Msg:    "fail",
			Data:   []string{"not available"},
			Method: "list-package",
			Time:   time.Now().String(),
		}
		resJSON, _ := json.Marshal(response)
		io.WriteString(w, string(resJSON))
	} else {
		info := strings.Split(string(out), "\n")
		response := showResponse{
			Code:   200,
			Msg:    "success",
			Data:   info[:len(info)-1],
			Method: "list-package",
			Time:   time.Now().String(),
		}
		resJSON, _ := json.Marshal(response)
		fmt.Printf("%s\n", resJSON)
		io.WriteString(w, string(resJSON))
	}
}

//GetDevice returns device information
func GetDevice() string {
	cmd := exec.Command("getprop", "ro.product.model")
	info, _ := cmd.CombinedOutput()
	return string(info[:len(info)-1])
}