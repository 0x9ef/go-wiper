// Copyright (c) 2021 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	wipe "github.com/0x9ef/go-wiper/wipe"
)

var mappedRules = map[int]*wipe.Rule{
	1: wipe.RuleFast,
	2: wipe.RuleVSITR,
	3: wipe.RuleUsDod5220_22_M,
	4: wipe.RuleGutmann,
}

var mappedPolicy = map[int]*wipe.Policy{
	1: &wipe.Policy{"Fast", "Data will be overwrited with zeroes (1 passes)", wipe.RuleFast},
	2: &wipe.Policy{"VSITR", "German VSITR (7 passes)", wipe.RuleVSITR},
	3: &wipe.Policy{"UsDod5220_22_M", "US Department of Defense DoD 5220.22-M (3 passes)", wipe.RuleUsDod5220_22_M},
	4: &wipe.Policy{"Gutmann", "Peter Gutmann Secure Method (35 passes)", wipe.RuleGutmann},
}

func report(a ...interface{}) {
	if *flagReport {
		fmt.Fprintln(os.Stdout, a...)
	}
}

var (
	flagPath   = flag.String("path", "", "Path of file or folder to permanently delete")
	flagRule   = flag.Int("rule", 1, "Name of rule which will be used for data wiping")
	flagKeep   = flag.Bool("keep", false, "Do not delete a file after wiping")
	flagReport = flag.Bool("report", false, "Reporting about all operations?")
)

func main() {
	print(`
     ________       __      __.__                     
    /  _____/  ____/  \    /  \__|_____   ___________ 
   /   \  ___ /  _ \   \/\/   /  \____ \_/ __ \_  __ \
   \    \_\  (  <_> )        /|  |  |_> >  ___/|  | \/
    \______  /\____/ \__/\  / |__|   __/ \___  >__|   
           \/  by 0x9ef   \/     |__|        \/       
   

Currently supported 4 algorithms, see list below:
 #1 Fast = data will be overwrited with zeroes (1 passes)
 #2 Vsitr = German VSITR (7 passes)
 #3 DoD 5220.22-M = US Department of Defense DoD 5220.22-M (3 passes)
 #4 Gutmann = Peter Gutmann Secure Method (35 passes)` + "\n\n")

	flag.Parse()
	flag.Usage = func() { flag.PrintDefaults() }

	policy, ok := mappedPolicy[*flagRule]
	if !ok {
		panic("wiper: provided unknown wipe rule")
	}

	rule := policy.Rule
	fmt.Printf("Selected rule:\n%s\n", policy.String())

	path := *flagPath
	fstat, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	if fstat.IsDir() {
		// Recursively delete all files and folders
		report("Recursively traversing files and directories ...")
		err := filepath.Walk(path,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if !info.IsDir() {
					report("Currently processing file:", path)
					err = wipe.Wipe(path, rule)
					if err != nil {
						panic(err)
					}
					if !*flagKeep {
						err = os.Remove(path)
						if err != nil {
							return err
						}
					}
				}
				return err
			})

		if err != nil {
			panic(err)
		}
	} else {
		if fstat.Size() == 0 {
			panic("wiper: provided empty file ")
		}

		report("Currently processing file:", path)
		err = wipe.Wipe(path, rule)
		if err != nil {
			panic(err)
		}

		if !*flagKeep {
			err = os.Remove(path)
			if err != nil {
				panic(err)
			}
		}
	}
}
