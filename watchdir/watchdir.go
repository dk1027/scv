package main

import (
	"github.com/dk1027/scv/shared"
	"github.com/go-fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
)

type watchdirTrigger struct {
	rule    shared.IRule
	param   string
	watcher *fsnotify.Watcher
}

func (this *watchdirTrigger) Close() {
	log.Printf("Closing watcher")
	this.watcher.Close()
}

func (this *watchdirTrigger) Run() {
	log.Println("watchdirTrigger running")
	watcher, err := fsnotify.NewWatcher()
	this.watcher = watcher

	if err != nil {
		log.Fatal(err)
	}

	if err := filepath.Walk(this.param, bindWatchDir(this.watcher)); err != nil {
		log.Printf("ERROR %v", err)
	}
	for {
		select {
		// watch for events
		case event := <-this.watcher.Events:
			log.Printf("Received event")
			if event.Op == fsnotify.Write {
				log.Printf("%s\n", event.Name)
				//this.rule.Fire(event.Name)
			}
		// watch for errors
		case err := <-this.watcher.Errors:
			log.Printf("ERROR %v", err)
			panic("Busted")
		}
	}
}

func bindWatchDir(watcher *fsnotify.Watcher) func(string, os.FileInfo, error) error {
	return func(path string, fi os.FileInfo, err error) error {
		return watchDir(watcher, path, fi, err)
	}
}

// watchDir gets run as a walk func, searching for directories to add watchers to
func watchDir(watcher *fsnotify.Watcher, path string, fi os.FileInfo, err error) error {
	// since fsnotify can watch all the files in a directory, watchers only need
	// to be added to each nested directory
	if fi.Mode().IsDir() {
		log.Printf("Added %v", path)
		return watcher.Add(path)
	}

	return nil
}

func CreateTriggerImpl(rule shared.IRule, param string) shared.ITrigger {
	log.Println("Loaded watchdir plugin")
	var res shared.ITrigger = &watchdirTrigger{rule: rule, param: param}
	return res
}
