// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codegen

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"istio.io/istio/pkg/config/schema/ast"
	"istio.io/istio/pkg/test/env"
)

func Run() error {
	inp, err := buildInputs()
	if err != nil {
		return err
	}

	// Include synthetic types used for XDS pushes
	kindEntries := append([]colEntry{{
		Resource: &ast.Resource{Identifier: "Address", Kind: "Address", Version: "internal", Group: "internal"},
	}}, inp.Entries...)

	// filter to only types agent needs (to keep binary small)
	agentEntries := []colEntry{}
	for _, e := range inp.Entries {
		if strings.Contains(e.Resource.ProtoPackage, "istio.io") {
			agentEntries = append(agentEntries, e)
		}
	}
	return errors.Join(
		writeTemplate("pkg/config/schema/gvk/resources.gen.go", gvkTemplate, map[string]any{
			"Entries":     inp.Entries,
			"PackageName": "gvk",
		}),
		writeTemplate("pkg/config/schema/gvr/resources.gen.go", gvrTemplate, map[string]any{
			"Entries":     inp.Entries,
			"PackageName": "gvr",
		}),
		writeTemplate("pkg/config/schema/kind/resources.gen.go", kindTemplate, map[string]any{
			"Entries":     kindEntries,
			"PackageName": "kind",
		}),
		writeTemplate("pkg/config/schema/collections/collections.gen.go", collectionsTemplate, map[string]any{
			"Entries":     inp.Entries,
			"Packages":    inp.Packages,
			"PackageName": "collections",
			"FilePrefix":  "// +build !agent",
		}),
		writeTemplate("pkg/config/schema/collections/collections.agent.gen.go", collectionsTemplate, map[string]any{
			"Entries":     agentEntries,
			"Packages":    inp.Packages,
			"PackageName": "collections",
			"FilePrefix":  "// +build agent",
		}),
	)
}

func writeTemplate(path, tmpl string, i any) error {
	t, err := applyTemplate(tmpl, i)
	if err != nil {
		return fmt.Errorf("apply template %v: %v", path, err)
	}
	dst := filepath.Join(env.IstioSrc, path)
	if err = os.WriteFile(dst, []byte(t), os.ModePerm); err != nil {
		return fmt.Errorf("write template %v: %v", path, err)
	}
	c := exec.Command("goimports", "-w", "-local", "istio.io", dst)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

func applyTemplate(tmpl string, i any) (string, error) {
	t := template.New("tmpl").Funcs(template.FuncMap{
		"contains": strings.Contains,
	})

	t2 := template.Must(t.Parse(tmpl))

	var b bytes.Buffer
	if err := t2.Execute(&b, i); err != nil {
		return "", err
	}

	return b.String(), nil
}
