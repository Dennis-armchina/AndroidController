package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"time"
)

//ShowScreen returns info of the screen
func ShowScreen(w http.ResponseWriter, r *http.Request) {
	cmdSize := exec.Command("sh", "/system/bin/wm", "size")
	size, err1 := cmdSize.CombinedOutput()
	cmdDensity := exec.Command("sh", "/system/bin/wm", "density")
	density, err2 := cmdDensity.CombinedOutput()
	//Trim
	Size := Trim(string(size))
	Density := Trim(string(density))
	if err1 != nil || err2 != nil {
		response := DeviceResponse{
			Code:   400,
			Msg:    "error: " + err1.Error(),
			Method: "list-screen",
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
			Data: []string{"Size:" + Size, "Density:" + Density},
			// Size:    string(Size),
			// Density: string(Density),
			Method: "list-screen",
			Time:   time.Now().String(),
			Device: GetDevice(),
		}
		resJSON, _ := json.Marshal(response)
		fmt.Printf("%s\n", resJSON)
		io.WriteString(w, string(resJSON))
	}
}
