/*
Copyright 2021 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package standalone

import (
	"bufio"
	"os/exec"
	"strings"
)

// Values for these are injected by the build.
var (
	gitcommit, gitversion string
)

// GetRuntimeVersion returns the version for the local Dapr runtime.
func GetRuntimeVersion() string {
	daprBinDir := defaultDaprBinPath()
	daprCMD := binaryFilePath(daprBinDir, "daprd")

	out, err := exec.Command(daprCMD, "--version").Output()
	if err != nil {
		return "n/a\n"
	}
	return string(out)
}

// GetDashboardVersion returns the version for the local Dapr dashboard.
func GetDashboardVersion() string {
	daprBinDir := defaultDaprBinPath()
	dashboardCMD := binaryFilePath(daprBinDir, "dashboard")

	out, err := exec.Command(dashboardCMD, "--version").Output()
	if err != nil {
		return "n/a\n"
	}
	return string(out)
}

// GetBuildInfo returns build info for the CLI and the local Dapr runtime.
func GetBuildInfo(version string) string {
	daprBinDir := defaultDaprBinPath()
	daprCMD := binaryFilePath(daprBinDir, "daprd")

	strs := []string{
		"CLI:",
		"\tVersion: " + version,
		"\tGit Commit: " + gitcommit,
		"\tGit Version: " + gitversion,
		"Runtime:",
	}

	out, err := exec.Command(daprCMD, "--build-info").Output()
	if err != nil {
		// try '--version' for older runtime version.
		out, err = exec.Command(daprCMD, "--version").Output()
	}
	if err != nil {
		strs = append(strs, "\tN/A")
	} else {
		scanner := bufio.NewScanner(strings.NewReader(string(out)))
		for scanner.Scan() {
			strs = append(strs, "\t"+scanner.Text())
		}
	}
	return strings.Join(strs, "\n")
}
