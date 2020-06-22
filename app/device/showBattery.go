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

//DeviceResponse base format
type DeviceResponse struct {
	Code   int
	Msg    string
	Data   []string
	Method string
	Time   string
	Device string
}

//status code
const (
	Success      = 200
	AlreadyStart = 201
	BadRequest   = 400
)

//ShowBattery returns info of the battery
func ShowBattery(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("dumpsys", "battery")
	out, err := cmd.CombinedOutput()
	if err != nil {
		response := DeviceResponse{
			Code:   BadRequest,
			Msg:    "error: " + err.Error(),
			Method: "list-battery",
			Time:   time.Now().String(),
			Device: GetDevice(),
		}
		resJSON, _ := json.Marshal(response)
		fmt.Printf("%s\n", resJSON)
		io.WriteString(w, string(resJSON))
	} else {
		info := strings.Split(string(out), "\n")
		response := DeviceResponse{
			Code:   Success,
			Msg:    "success",
			Data:   info,
			Method: "list-battery",
			Time:   time.Now().String(),
			Device: GetDevice(),
		}
		resJSON, _ := json.Marshal(response)
		fmt.Printf("%s\n", resJSON)
		io.WriteString(w, string(resJSON))
	}
}

//Trim the string to get rid of "\t", "\n"
func Trim(input string) string {
	trim := strings.Replace(input, "\n", " ", -1)
	trim = strings.Replace(trim, "\t", " ", -1)
	return trim
}
