package swagger

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"path"
	"runtime"
)

//OpenPackage for openeing the pacakge
func OpenPackage(w http.ResponseWriter, r *http.Request) {
	head := "apk_name"
	r.ParseForm()
	apkName := parseOpen(r.Form, head)
	cmd := exec.Command("am", "start", apkName)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error : ", err)
		fmt.Fprintln(w, "fail to open the specific package")
	}
	fmt.Println("command to be executed:", cmd)
	fmt.Fprintln(w, "command to be executed:", cmd)
	fmt.Fprintln(w, "the package will open:", apkName)
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
