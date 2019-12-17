/*
 * Copyright (c) 2019, salesforce.com, inc.
 * All rights reserved.
 * SPDX-License-Identifier: BSD-3-Clause
 * For full license text, see LICENSE.txt file in the repo root or https://opensource.org/licenses/BSD-3-Clause
 */

package main

import (
	"github.com/golang/glog"
	"github.com/salesforce/sloop/pkg/sloop/server"
	_ "net/http/pprof"
	"os"
)

func main() {
	err := server.RealMain()
	if err != nil {
		glog.Errorf("Main exited with error: %v\n", err)
		os.Exit(1)
	} else {
		glog.Infof("Shutting down gracefully")
	}

}
