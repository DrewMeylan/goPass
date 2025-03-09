package crypto

import (
  "bytes"
  "fmt"
  "os"
  "runtime"
  "path/filepath"
  "github.com/ProtonMail/gopenpgp/v3"
  "github.com/ProtonMail/gopenpgp/v3/crypto
  "time"
)

func generateKey(user_name string, user_email string, password string) (*openpgp.Entity, error) {
	// Define key properties: name, email, passphrase.
	name := user_name
	email := user_email
	passphrase := []byte(password)
	// Create a key signing entity (identity) for generating keys
	identity := &packet.Identity{
		Name:  name,
		Email: email,
	}
	// Define the key generation options.
	keyParams := packet.NewKeyGenParams(packet.PubKeyAlgoRSA, packet.SigKeyAlgoRSA, 4096, time.Now().Add(365*24*time.Hour)) // RSA 2048-bit, valid for 1 year

	// Generate the entity (public + private key)
	entity, err := openpgp.NewEntityWithConfig(identity, nil, nil, keyParams)
	if err != nil {
		return nil, fmt.Errorf("error generating GPG key: %v", err)
	}
	// Encrypt the private key with the passphrase
	entity.PrivateKey.Encrypt(passphrase)
	return entity, nil

  // Logic for Exporting Keys to local system
}

// ExportGPGKeys exports the generated OpenPGP entity's public and private keys to local files.
func ExportGPGKeys(entity *openpgp.Entity) error {
	// Ensure the directory exists
  // We should be writing these keys to the default directories based on the operating sytem..
  outputDir, err := getDefaultGPGDir()
	if err != nil {
    fmt.Println("Error encountered: ", err)
    return err
  }
  if outputDir == nil {
    fmt.Println("GPG directory not found!")
    return nil
  }

	// Define file paths
	pubKeyPath := fmt.Sprintf("%s/public-key.asc", outputDir)
	privKeyPath := fmt.Sprintf("%s/private-key.asc", outputDir)

	// Export public key
	pubFile, err := os.Create(pubKeyPath)
	if err != nil {
		return fmt.Errorf("failed to create public key file: %v", err)
	}
	defer pubFile.Close()

	err = entity.Serialize(pubFile)
	if err != nil {
		return fmt.Errorf("failed to write public key: %v", err)
	}

	// Export private key
	privFile, err := os.Create(privKeyPath)
	if err != nil {
		return fmt.Errorf("failed to create private key file: %v", err)
	}
	defer privFile.Close()

	err = entity.SerializePrivate(privFile, nil)
	if err != nil {
		return fmt.Errorf("failed to write private key: %v", err)
	}

	fmt.Println("Keys exported successfully:")
	fmt.Println("Public Key:", pubKeyPath)
	fmt.Println("Private Key:", privKeyPath)

	return nil
}

func getDefaultGPGDir() (string, erro) {
  var gpgDir string

  switch runtime.GOOS {
  case "linux", "darwin":
    gpgDir = filepath.Join(os.Getenv("HOME"), ".gnupg")
  case "windows":
    appData := os.Getenv("APPDATA")
    if appData == "" {
      return "", fmt.Errorf("APPDATA environment variable not set!")
    }
    gpgDir = filepath.Join(appData, "gnupg")
  defualt:
    return "", fmt.Errorf("Unsupported operating system type")
  }

  return gpgDir, nil
}

