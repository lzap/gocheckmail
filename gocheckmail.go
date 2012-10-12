package main

import (
	"code.google.com/p/goconf/conf"
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"os/exec"
)

func main() {
	var path = os.ExpandEnv("$HOME/mail")
	var command = os.ExpandEnv("/usr/bin/notify-send 'New mail'")
	c, err := conf.ReadConfigFile(os.ExpandEnv("$HOME/.gocheckmail.conf"))
	if err == nil {
		path, _ = c.GetString("default", "path")
		command, _ = c.GetString("default", "command")
		path = os.ExpandEnv(path)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("error creating inotify watcher: %s", err)
	}

	err = watcher.Watch(path)
	if err != nil {
		log.Fatalf("error while opening %s: %s", path, err)
	}
	defer func() {
		watcher.Close()
	}()

	for {
		select {
		case ev := <-watcher.Event:
			//log.Print(ev)
			if ev.IsCreate() {
				cmd := exec.Command("/bin/bash", "-c", command)
				err := cmd.Start()
				if err != nil {
					log.Fatal(err)
				}
			}
		case err := <-watcher.Error:
			log.Printf("error event from inotify: %s", err)
		}
	}
}
