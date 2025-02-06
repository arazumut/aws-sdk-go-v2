//go:build windows
// +build windows

package shareddefaults_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go-v2/internal/awstesting"
	"github.com/aws/aws-sdk-go-v2/internal/shareddefaults"
)

// TestSharedCredsFilename fonksiyonu, paylaşılan kimlik bilgileri dosyasının adını test eder.
func TestSharedCredsFilename(t *testing.T) {
	restoreEnv := awstesting.StashEnv() // Mevcut ortam değişkenlerini sakla
	defer awstesting.PopEnv(restoreEnv) // Testten sonra ortam değişkenlerini geri yükle

	os.Setenv("HOME", "home_dir")           // HOME ortam değişkenini ayarla
	os.Setenv("USERPROFILE", "profile_dir") // USERPROFILE ortam değişkenini ayarla

	expect := filepath.Join("profile_dir", ".aws", "credentials") // Beklenen dosya yolu

	name := shareddefaults.SharedCredentialsFilename() // Gerçek dosya yolunu al
	if e, a := expect, name; e != a {
		t.Errorf("beklenen %q paylaşılan kimlik bilgileri dosya adı, alınan %q", e, a)
	}
}

// TestSharedConfigFilename fonksiyonu, paylaşılan yapılandırma dosyasının adını test eder.
func TestSharedConfigFilename(t *testing.T) {
	restoreEnv := awstesting.StashEnv() // Mevcut ortam değişkenlerini sakla
	defer awstesting.PopEnv(restoreEnv) // Testten sonra ortam değişkenlerini geri yükle

	os.Setenv("HOME", "home_dir")           // HOME ortam değişkenini ayarla
	os.Setenv("USERPROFILE", "profile_dir") // USERPROFILE ortam değişkenini ayarla

	expect := filepath.Join("profile_dir", ".aws", "config") // Beklenen dosya yolu

	name := shareddefaults.SharedConfigFilename() // Gerçek dosya yolunu al
	if e, a := expect, name; e != a {
		t.Errorf("beklenen %q paylaşılan yapılandırma dosya adı, alınan %q", e, a)
	}
}
