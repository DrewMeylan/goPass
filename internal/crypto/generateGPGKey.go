package crypto

import (
  "bytes"
  "fmt"
  "github.com/ProtonMail/go-crypto/openpgp"
  "github.com/ProtonMail/go-crypto/openpgp/packet"
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
}
