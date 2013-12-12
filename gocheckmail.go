package main

import (
	"log"
	"os"
	"os/exec"
	"time"
	"bufio"
	"regexp"
	//"io"
	"path"
)

var subject_regexp = regexp.MustCompile(`^Subject:\s*(.*)$`)

func report(msg string) {
	cmd := exec.Command("/bin/bash", "-c", "/usr/bin/notify-send '" + msg + "'")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func read_subject(filename string) string {
		// check for new mail
		d, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
			return "?"
		}
		defer d.Close()
		//scanner := bufio.NewScanner(io.LimitReader(d, 5000))
		scanner := bufio.NewScanner(d)
		for scanner.Scan() {
			if subject_regexp.Match([]byte(scanner.Text())) {
				return subject_regexp.FindStringSubmatch(scanner.Text())[0]
			}
		}
		return "?"
}

func main() {
	var dirname = os.ExpandEnv("$HOME/Mail/INBOX")
	var latest = time.Now()

	for {
		time.Sleep(4 * time.Minute)

		// check for new mail
		d, err := os.Open(dirname)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer d.Close()
		fi, err := d.Readdir(-1)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		for _, fi := range fi {
			if fi.Mode().IsRegular() {
				if fi.ModTime().After(latest) {
					report("New mail: " + read_subject(path.Join(dirname, fi.Name())))
					latest = fi.ModTime()
					// no rush with opening bubbles
					time.Sleep(5 * time.Second)
				}
			}
		}

		// report hours every hour (break time!)
		if time.Now().Minute() >= 55 {
			report(time.Now().Format("15:04"))
			time.Sleep(64 * time.Second)
		}
	}
}
