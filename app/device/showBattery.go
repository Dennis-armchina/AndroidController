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

type batteryResponse struct {
	Code   int
	Msg    string
	Data   []string
	Method string
	Time   string
	Device string
}

//ShowBattery returns info of the battery
func ShowBattery(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("dumpsys", "battery")
	out, err := cmd.CombinedOutput()
	if err != nil {
		response := batteryResponse{
			Code:   400,
			Msg:    "error: " + err.Error(),
			Data:   []string{"not available"},
			Method: "list-battery",
			Time:   time.Now().String(),
			Device: GetDevice(),
		}
		resJSON, _ := json.Marshal(response)
		fmt.Printf("%s\n", resJSON)
		io.WriteString(w, string(resJSON))
	} else {
		info := strings.Split(string(out), "\n")
		response := batteryResponse{
			Code:   200,
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
