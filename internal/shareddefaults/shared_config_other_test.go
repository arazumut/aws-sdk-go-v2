//go:build !windows
// +build !windows

package shareddefaults_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go-v2/internal/awstesting"

	"github.com/aws/aws-sdk-go-v2/internal/shareddefaults"
)

func TestSharedCredsFilename(t *testing.T) {
	restoreEnv := awstesting.StashEnv()
	defer awstesting.PopEnv(restoreEnv)

	os.Setenv("HOME", "home_dir")
	os.Setenv("USERPROFILE", "profile_dir")

	expect := filepath.Join("home_dir", ".aws", "credentials")

	name := shareddefaults.SharedCredentialsFilename()
	if e, a := expect, name; e != a {
		t.Errorf("expect %q shared creds filename, got %q", e, a)
	}
}

func TestSharedConfigFilename(t *testing.T) {
	restoreEnv := awstesting.StashEnv()
	defer awstesting.PopEnv(restoreEnv)

	os.Setenv("HOME", "home_dir")
	os.Setenv("USERPROFILE", "profile_dir")

	expect := filepath.Join("home_dir", ".aws", "config")

	name := shareddefaults.SharedConfigFilename()
	if e, a := expect, name; e != a {
		t.Errorf("expect %q shared config filename, got %q", e, a)
	}
}
func TestCustomSharedCredsFilename(t *testing.T) {
	restoreEnv := awstesting.StashEnv()
	defer awstesting.PopEnv(restoreEnv)

	customPath := filepath.Join("custom_dir", ".aws", "credentials")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", customPath)

	name := shareddefaults.SharedCredentialsFilename()
	if e, a := customPath, name; e != a {
		t.Errorf("expect %q shared creds filename, got %q", e, a)
	}
}

func TestCustomSharedConfigFilename(t *testing.T) {
	restoreEnv := awstesting.StashEnv()
	defer awstesting.PopEnv(restoreEnv)

	customPath := filepath.Join("custom_dir", ".aws", "config")
	os.Setenv("AWS_CONFIG_FILE", customPath)

	name := shareddefaults.SharedConfigFilename()
	if e, a := customPath, name; e != a {
		t.Errorf("expect %q shared config filename, got %q", e, a)
	}
}
