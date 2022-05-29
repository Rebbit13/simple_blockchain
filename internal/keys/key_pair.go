package keys

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"log"
	"os"
)

const (
	keysPath       = ".keys"
	privateKeyFile = "key.priv"
)

func createKeyPair() (*btcec.PrivateKey, *btcec.PublicKey) {
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyString := fmt.Sprintf("%x", privateKey.Serialize())
	err = os.WriteFile(fmt.Sprintf("%s/%s", keysPath, privateKeyFile), []byte(privateKeyString), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	return privateKey, privateKey.PubKey()
}

func parseKeysFromFile() (*btcec.PrivateKey, *btcec.PublicKey, error) {
	data, err := os.ReadFile(fmt.Sprintf("%s/%s", keysPath, privateKeyFile))
	if err != nil {
		return nil, nil, err
	}
	s, err := hex.DecodeString(string(data))
	if err != nil {
		panic(err)
	}
	privateKey, pubKey := btcec.PrivKeyFromBytes(s)
	return privateKey, pubKey, nil
}

func GetOrCreateKeyPair() (*btcec.PrivateKey, *btcec.PublicKey) {
	privateKey, publicKey, err := parseKeysFromFile()
	if errors.As(err, &os.ErrNotExist) {
		privateKey, publicKey = createKeyPair()
	} else if err != nil {
		log.Fatal(err)
	}
	return privateKey, publicKey
}
