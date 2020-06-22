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

type systemResponse struct {
	Code    int
	Msg     string
	Version string
	CPU     []string
	Method  string
	Time    string
	Device  string
}

//ShowSystem return info of the android systme
func ShowSystem(w http.ResponseWriter, r *http.Request) {
	cmdVersion := exec.Command("getprop", "ro.build.version.release")
	outVersion, err1 := cmdVersion.CombinedOutput()
	Version := strings.Replace(string(outVersion), "\n", "", -1)
	cmdCPU := exec.Command("cat", "/proc/cpuinfo")
	outCPU, err2 := cmdCPU.CombinedOutput()
	//replace '\t'
	CPU := strings.Replace(string(outCPU), "\t", "", -1)
	//replace '\n'
	if err1 != nil || err2 != nil {
		response := systemResponse{
			Code:    400,
			Msg:     "error: " + err1.Error(),
			Version: string(Version),
			CPU:     []string{"not available"},
			Method:  "list-system",
			Time:    time.Now().String(),
			Device:  GetDevice(),
		}
		resJSON, _ := json.Marshal(response)
		fmt.Printf("%s\n", resJSON)
		io.WriteString(w, string(resJSON))
	} else {
		info := strings.Split(CPU, "\n")
		response := systemResponse{
			Code:    200,
			Msg:     "success",
			Version: string(Version),
			CPU:     info[:len(info)-1],
			Method:  "list-system",
			Time:    time.Now().String(),
			Device:  GetDevice(),
		}
		resJSON, _ := json.Marshal(response)
		fmt.Printf("%s\n", resJSON)
		io.WriteString(w, string(resJSON))
	}
}

//GetDevice return system info
func GetDevice() string {
	cmd := exec.Command("getprop", "ro.product.model")
	info, _ := cmd.CombinedOutput()
	return string(info[:len(info)-1])
}
