package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"time"
)

//ShowSystem return info of the android systme
func ShowSystem(w http.ResponseWriter, r *http.Request) {
	cmdVersion := exec.Command("getprop", "ro.build.version.release")
	outVersion, err1 := cmdVersion.CombinedOutput()
	Version := Trim(string(outVersion))
	cmdCPU := exec.Command("cat", "/proc/cpuinfo")
	outCPU, err2 := cmdCPU.CombinedOutput()
	//replace '\t'
	CPU := Trim(string(outCPU))
	//replace '\n'
	if err1 != nil || err2 != nil {
		response := DeviceResponse{
			Code: 400,
			Msg:  "error: " + err1.Error(),
			// Version: string(Version),
			// CPU:     []string{"not available"},
			Method: "list-system",
			Time:   time.Now().String(),
			Device: GetDevice(),
		}
		resJSON, _ := json.Marshal(response)
		fmt.Printf("%s\n", resJSON)
		io.WriteString(w, string(resJSON))
	} else {
		response := DeviceResponse{
			Code: 200,
			Msg:  "success",
			Data: []string{"Version:" + Version, "CPU:" + CPU},
			// Version: string(Version),
			// CPU:     info[:len(info)-1],
			Method: "list-system",
			Time:   time.Now().String(),
			Device: GetDevice(),
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
