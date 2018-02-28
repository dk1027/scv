package scv

import (
	"bytes"
	"fmt"
	"github.com/dk1027/scv/shared"
	"log"
	"os/exec"
)

type Engine struct {
	Rules []TaskConfig
}

type Rule struct {
	trigger shared.ITrigger
	action  shared.IAction
}

func (this *Rule) Fire(runtimeParam string) {
	go this.action.Run(runtimeParam)
}

func (this *Rule) Start() {
	go this.trigger.Run() //TODO: Pass in a Close channel?
}

func (this *Rule) Close() {
	log.Println("Rule: Closing")
	this.trigger.Close()
}

//
// // DummyTrigger
// type watchdirTrigger struct {
// 	rule    *Rule
// 	param   string
// 	watcher *fsnotify.Watcher
// }
//
// func (this *watchdirTrigger) Close() {
// 	log.Printf("Closing watcher")
// 	this.watcher.Close()
// }
//
// func (this *watchdirTrigger) Run() {
// 	this.watcher, _ = fsnotify.NewWatcher()
//
// 	if err := filepath.Walk(this.param, bindWatchDir(this.watcher)); err != nil {
// 		log.Printf("ERROR %v", err)
// 	}
// 	for {
// 		select {
// 		// watch for events
// 		case event := <-this.watcher.Events:
// 			if event.Op == fsnotify.Write {
// 				log.Printf("%s\n", event.Name)
// 				this.rule.Fire(event.Name)
// 			}
// 		// watch for errors
// 		case err := <-this.watcher.Errors:
// 			log.Printf("ERROR %v", err)
// 			panic("Busted")
// 		}
// 	}
// 	//TODO: Wait for Close ... possibly using a channel
// }
//
// func bindWatchDir(watcher *fsnotify.Watcher) func(string, os.FileInfo, error) error {
// 	return func(path string, fi os.FileInfo, err error) error {
// 		return watchDir(watcher, path, fi, err)
// 	}
// }
//
// // watchDir gets run as a walk func, searching for directories to add watchers to
// func watchDir(watcher *fsnotify.Watcher, path string, fi os.FileInfo, err error) error {
// 	// since fsnotify can watch all the files in a directory, watchers only need
// 	// to be added to each nested directory
// 	if fi.Mode().IsDir() {
// 		return watcher.Add(path)
// 	}
//
// 	return nil
// }

// buildAction
type buildAction struct {
	staticParam string
}

func (this *buildAction) Run(runtimeParam string) {
	log.Printf("Runtime Param: %+v\n", runtimeParam)
	log.Printf("Running command: %+v\n", this.staticParam)
	cmd := exec.Command("sh", "-c", this.staticParam)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Printf(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	log.Printf(out.String())
}

func NewEngine(rules []TaskConfig) Engine {
	var ruleObjs []*Rule
	for _, rule := range rules {
		ruleObj := &Rule{}
		trigger := newTrigger(ruleObj, rule.Trigger.Type, rule.Trigger.Param)
		action := newAction(rule.Action.Type, rule.Action.Param)
		ruleObj.trigger = trigger
		ruleObj.action = action
		ruleObjs = append(ruleObjs, ruleObj)
	}

	// TODO: start the Rules. i.e. each trigger will do its thing waiting for its event
	for _, ruleObj := range ruleObjs {
		ruleObj.Start()
		defer ruleObj.Close()
	}
	engine := Engine{}
	done := make(chan bool)
	<-done
	return engine
}

func newTrigger(ruleObj shared.IRule, name string, param string) shared.ITrigger {
	switch name {
	case "watchdir":
		var result shared.ITrigger = shared.TriggerLoader("watchdir/watchdir.so")(ruleObj, param)
		//var result shared.ITrigger = &watchdirTrigger{rule: ruleObj, param: param}
		return result
	default:
		log.Printf("unknown trigger\n")
		return nil
	}
}

func newAction(name string, param string) shared.IAction {
	switch name {
	case "build":
		return &buildAction{staticParam: param}
	default:
		log.Printf("unknown action\n")
		return nil
	}
}
