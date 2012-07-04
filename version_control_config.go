// Copyright 2012, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

/*  Filename:    version_control_spec.go
 *  Author:      Bryan Matsuo <bryan.matsuo [at] gmail.com>
 *  Created:     2012-07-03 18:02:20.295289 -0700 PDT
 *  Description: 
 */

import ()

type VersionControlConfig struct {
	Type   string
	Remote map[string]interface{}
}