// Copyright 2013 Unknown
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	"github.com/astaxie/beego"
	"github.com/unknwon/gowalker/doc"
	"github.com/unknwon/gowalker/utils"
)

type HomeController struct {
	beego.Controller
}

// Get implemented Get method for HomeController.
// It serves home page of Go Walker.
func (this *HomeController) Get() {
	// Get query field
	q := this.Input().Get("q")
	// Set properties
	this.TplNames = "home.html"
	this.Layout = "layout.html"

	// Empty query string shows home page
	if len(q) == 0 {
		beego.Info("Show home page")
		this.Render()
		return
	}

	if path, ok := utils.IsBrowseURL(q); ok {
		q = path
	}

	// Check remote path
	if doc.IsValidRemotePath(q) {
		// Check documentation in database
		beego.Info("Remote path")
		this.Render()
		return
	}

	// Show search page
	this.Redirect("/search?q="+q, 302)
}