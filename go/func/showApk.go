package swagger

import (
	"fmt"
	"net/http"
	"os/exec"
)

//ListPackage method returns installed Apks
func ListPackage(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sh", "/system/bin/pm", "list", "packages", "-3")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error : ", err)
		fmt.Println("cmd: ", cmd)
		fmt.Fprintln(w, "fail to get package info")
	}
	fmt.Fprintf(w, "---Installed APK list---\n %s", out)
}
