// This file is part of arduino-lint.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-lint.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

// Package checkresult defines the possible result values returned by a check.
package checkresult

//go:generate stringer -type=Type -linecomment
type Type int

const (
	Pass Type = iota // pass
	Fail             // fail
	// The check is configured to be skipped in the current tool configuration mode
	Skip // skipped
	// An unrelated error prevented the check from running
	NotRun // unable to run
)
