package shareddefaults

import (
	"os"
	"os/user"
	"path/filepath"
)

// SharedCredentialsFilename, SDK'nın paylaşılan kimlik bilgileri dosyası için
// varsayılan dosya yolunu döndürür.
//
// İşletim sisteminin platformuna göre paylaşılan yapılandırma dosyası yolunu oluşturur.
//
//   - Linux/Unix: $HOME/.aws/credentials
//   - Windows: %USERPROFILE%\.aws\credentials
func SharedCredentialsFilename() string {
	return filepath.Join(UserHomeDir(), ".aws", "credentials")
}

// SharedConfigFilename, SDK'nın paylaşılan yapılandırma dosyası için
// varsayılan dosya yolunu döndürür.
//
// İşletim sisteminin platformuna göre paylaşılan yapılandırma dosyası yolunu oluşturur.
//
//   - Linux/Unix: $HOME/.aws/config
//   - Windows: %USERPROFILE%\.aws\config
func SharedConfigFilename() string {
	return filepath.Join(UserHomeDir(), ".aws", "config")
}

// UserHomeDir, işlemin çalıştığı kullanıcının ana dizinini döndürür.
func UserHomeDir() string {
	// Hataları yok say çünkü sadece Windows ve *nix ile ilgileniyoruz.
	home, _ := os.UserHomeDir()

	if len(home) > 0 {
		return home
	}

	currUser, _ := user.Current()
	if currUser != nil {
		home = currUser.HomeDir
	}

	return home
}
