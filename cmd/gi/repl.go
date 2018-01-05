package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-interpreter/gi/pkg/verb"
)

var p = verb.P
var pp = verb.PP

func main() {

	myflags := flag.NewFlagSet("gi", flag.ExitOnError)
	cfg := &GIConfig{}
	cfg.DefineFlags(myflags)

	err := myflags.Parse(os.Args[1:])
	err = cfg.ValidateConfig()
	if err != nil {
		log.Fatalf("%s command line flag error: '%s'", ProgramName, err)
	}

	verb.Verbose = cfg.Verbose || cfg.VerboseVerbose
	verb.VerboseVerbose = cfg.VerboseVerbose

	if !cfg.Quiet {
		fmt.Printf(
			`====================
gi: a go interpreter
====================
https://github.com/go-interpreter/gi
Copyright (c) 2018, Jason E. Aten, Ph.D.
License: 3-clause BSD. See the LICENSE file at
https://github.com/go-interpreter/gi/blob/master/LICENSE
====================
  [gi is an interactive Golang environment,
   also known as a REPL or Read-Eval-Print-Loop.]
  [type ctrl-d to exit]
  [type :help for help]
  [gi -h for flag help]
  [gi -q to start quietly]
====================
%s
==================
`, Version())
	}

	fmt.Printf("using this prelude directory: '%s'\n", cfg.PreludePath)
	fmt.Printf("using these files as prelude: '%#v'\n", cfg.preludeFiles)

	cfg.LuajitMain()
	//NodeChildMain()
	//OttoReplMain()
}
