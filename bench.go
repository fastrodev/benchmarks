package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func execCmd(command string) string {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return err.Error()
	}
	return string(output)
}

func runServerAndWrk(cmd string, port string) string {
	serverCmd := cmd + " &"
	wrkCmd := "wrk http://localhost:" + port
	killCmd := "lsof -nti:" + port + " | xargs kill -9"
	defer execCmd(killCmd)
	go func() {
		execCmd(serverCmd)
	}()
	time.Sleep(1 * time.Second)
	data := execCmd(wrkCmd)
	fmt.Println(data)
	return data
}

func writeFile(fastrex, node, std, express string) {
	ctnByte, err := ioutil.ReadFile("tmpl.md")

	if err != nil {
		log.Fatal(err)
	}

	content := string(ctnByte)
	content = strings.Replace(content, "{{fastrex}}", fastrex, -1)
	content = strings.Replace(content, "{{node}}", node, -1)
	content = strings.Replace(content, "{{standard}}", std, -1)
	content = strings.Replace(content, "{{express}}", express, -1)

	f, err := os.Create("readme.md")
	if err != nil {
		log.Fatal(err)
	}

	_, err2 := f.WriteString(content)
	if err2 != nil {
		log.Fatal(err2)
	}

	defer f.Close()
	fmt.Println("done")
}

func main() {
	express := runServerAndWrk("node express/main.js", "3000")
	node := runServerAndWrk("node node/main.js", "3001")
	std := runServerAndWrk("go run standard/main.go", "9001")
	fastrex := runServerAndWrk("go run fastrex/main.go", "9000")
	writeFile(fastrex, node, std, express)
}
