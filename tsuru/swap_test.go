// Copyright 2016 tsuru-client authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/tsuru/tsuru/cmd"
	"github.com/tsuru/tsuru/cmd/cmdtest"
	"gopkg.in/check.v1"
)

func (s *S) TestSwapInfo(c *check.C) {
	command := appSwap{}
	c.Assert(command.Info(), check.NotNil)
}

func (s *S) TestSwap(c *check.C) {
	var buf bytes.Buffer
	var called bool
	transport := cmdtest.ConditionalTransport{
		Transport: cmdtest.Transport{Status: http.StatusOK, Message: ""},
		CondFunc: func(r *http.Request) bool {
			called = true
			return r.Method == "PUT" && strings.HasSuffix(r.URL.Path, "/swap")
		},
	}
	context := cmd.Context{
		Args:   []string{"app1", "app2"},
		Stdout: &buf,
	}
	client := cmd.NewClient(&http.Client{Transport: &transport}, nil, manager)
	command := appSwap{}
	err := command.Run(&context, client)
	c.Assert(err, check.IsNil)
	c.Assert(called, check.Equals, true)
	expected := "Apps successfully swapped!\n"
	c.Assert(buf.String(), check.Equals, expected)
}

func (s *S) TestSwapFlags(c *check.C) {
	command := appSwap{}
	flagset := command.Flags()
	c.Assert(flagset, check.NotNil)
	flagset.Parse(true, []string{"app1", "app2", "-f", "-c"})
	forces := flagset.Lookup("force")
	c.Check(forces, check.NotNil)
	c.Check(forces.Name, check.Equals, "force")
	c.Check(forces.Usage, check.Equals, "Force Swap among apps with different number of units or different platform.")
	c.Check(forces.Value.String(), check.Equals, "true")
	c.Check(forces.DefValue, check.Equals, "false")
	cname := flagset.Lookup("cname-only")
	c.Check(cname, check.NotNil)
	c.Check(cname.Name, check.Equals, "cname-only")
	c.Check(cname.Usage, check.Equals, "Swap all cnames except the default cname.")
	c.Check(cname.Value.String(), check.Equals, "true")
	c.Check(cname.DefValue, check.Equals, "false")
}

func (s *S) TestSwapCnameOnlyFlag(c *check.C) {
	var buf bytes.Buffer
	var called bool
	transport := cmdtest.ConditionalTransport{
		Transport: cmdtest.Transport{Status: http.StatusOK, Message: ""},
		CondFunc: func(r *http.Request) bool {
			called = true
			return r.Method == "PUT" && strings.HasSuffix(r.URL.Path, "/swap") && r.URL.RawQuery == "app1=app1&app2=app2&force=false&cnameOnly=true"
		},
	}
	context := cmd.Context{
		Args:   []string{"app1", "app2"},
		Stdout: &buf,
	}
	client := cmd.NewClient(&http.Client{Transport: &transport}, nil, manager)
	command := appSwap{}
	command.Flags().Parse(true, []string{"-c"})
	err := command.Run(&context, client)
	c.Assert(err, check.IsNil)
	c.Assert(called, check.Equals, true)
	expected := "Apps successfully swapped!\n"
	c.Assert(buf.String(), check.Equals, expected)
}

func (s *S) TestSwapWhenAppsAreNotEqual(c *check.C) {
	var buf bytes.Buffer
	var called int
	stdin := bytes.NewBufferString("yes")
	transportError := cmdtest.ConditionalTransport{
		Transport: cmdtest.Transport{Status: http.StatusPreconditionFailed, Message: "Apps are not equal."},
		CondFunc: func(r *http.Request) bool {
			called += 1
			return r.URL.RawQuery == "app1=app1&app2=app2&force=false&cnameOnly=false"
		},
	}
	transportOk := cmdtest.ConditionalTransport{
		Transport: cmdtest.Transport{Status: http.StatusOK, Message: ""},
		CondFunc: func(r *http.Request) bool {
			called += 1
			return r.URL.RawQuery == "app1=app1&app2=app2&force=true&cnameOnly=false"
		},
	}
	multiTransport := cmdtest.MultiConditionalTransport{
		ConditionalTransports: []cmdtest.ConditionalTransport{transportError, transportOk},
	}
	context := cmd.Context{
		Args:   []string{"app1", "app2"},
		Stdout: &buf,
		Stdin:  stdin,
	}
	client := cmd.NewClient(&http.Client{Transport: &multiTransport}, nil, manager)
	command := appSwap{}
	err := command.Run(&context, client)
	c.Assert(err, check.IsNil)
	c.Assert(called, check.Equals, 2)
}

func (s *S) TestSwapIsACommand(c *check.C) {
	var _ cmd.Command = &appSwap{}
}
