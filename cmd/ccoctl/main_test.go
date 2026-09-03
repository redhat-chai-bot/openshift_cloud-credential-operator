package main

import (
	"bytes"
	"testing"

	"github.com/openshift/cloud-credential-operator/pkg/version"
)

func TestVersionFlag(t *testing.T) {
	cmd := newRootCmd()

	// The Version field must be set on the root command.
	if cmd.Version == "" {
		t.Fatal("expected Version to be set on the root command")
	}

	// When built without ldflags, version.Get().String() returns "" so
	// the command should fall back to "<unknown>".
	v := version.Get().String()
	if len(v) == 0 {
		if cmd.Version != "<unknown>" {
			t.Fatalf("expected Version to be %q when built without ldflags, got %q", "<unknown>", cmd.Version)
		}
	} else {
		if cmd.Version != v {
			t.Fatalf("expected Version to be %q, got %q", v, cmd.Version)
		}
	}

	// Executing --version should succeed and print the version.
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"--version"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error executing --version: %v", err)
	}

	output := buf.String()
	if len(output) == 0 {
		t.Fatal("expected --version to produce output")
	}

	// The output should contain the version string.
	expected := cmd.Version
	if !bytes.Contains(buf.Bytes(), []byte(expected)) {
		t.Fatalf("expected --version output to contain %q, got %q", expected, output)
	}
}

func TestRootCmdSubcommands(t *testing.T) {
	cmd := newRootCmd()

	// Verify cloud provider subcommands are registered.
	expectedSubcmds := []string{"aws", "gcp", "ibmcloud", "nutanix", "azure"}
	for _, name := range expectedSubcmds {
		found := false
		for _, sub := range cmd.Commands() {
			if sub.Name() == name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected subcommand %q to be registered", name)
		}
	}
}
