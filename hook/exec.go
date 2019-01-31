package exec

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/baidu/gls"
)

//arg 均非os/exec内定义
func Command(name string, arg ...string) *exec.Cmd {
	fmt.Printf("The command is: %s\n", name)
	var req = gls.Get("request").(*http.Request)
	fmt.Printf("%s %s %s %s\n", req.Method, req.Host, req.URL.Path, req.URL.RawQuery)
	return exec.Command(name, arg...)
}
