package main

import (
	"./registry"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var readFile = ioutil.ReadFile
var httpListenAndServe = http.ListenAndServe
var httpWriterSetContentType = func(w http.ResponseWriter, value string) {
	w.Header().Set("Content-Type", value)
}
var logPrintf = log.Printf

type Executable interface {
	Execute(args []string) error
}

var lookupHost = net.LookupHost
var registryInstance registry.Registrarable = registry.Consul{}
