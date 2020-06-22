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

type screenResponse struct {
	Code    int
	Msg     string
	Size    string
	Density string
	Method  string
	Time    string
	Device  string
}

//ShowScreen returns info of the screen
func ShowScreen(w http.ResponseWriter, r *http.Request) {
	cmdSize := exec.Command("sh", "/system/bin/wm", "size")
	size, err1 := cmdSize.CombinedOutput()
	cmdDensity := exec.Command("sh", "/system/bin/wm", "density")
	density, err2 := cmdDensity.CombinedOutput()
	//Trim
	Size := strings.Replace(string(size), "\n", "", -1)
	Density := strings.Replace(string(density), "\n", "", -1)
	if err1 != nil || err2 != nil {
		response := screenResponse{
			Code:    400,
			Msg:     "error: " + err1.Error(),
			Size:    "not available",
			Density: "not available",
			Method:  "list-screen",
			Time:    time.Now().String(),
			Device:  GetDevice(),
		}
		resJSON, _ := json.Marshal(response)
		fmt.Printf("%s\n", resJSON)
		io.WriteString(w, string(resJSON))
	} else {
		response := screenResponse{
			Code:    200,
			Msg:     "success",
			Size:    string(Size),
			Density: string(Density),
			Method:  "list-screen",
			Time:    time.Now().String(),
			Device:  GetDevice(),
		}
		resJSON, _ := json.Marshal(response)
		fmt.Printf("%s\n", resJSON)
		io.WriteString(w, string(resJSON))
	}
}
