package enc

import "testing"

func TestEncryption(t *testing.T) {
	plainText := "some text"
	key := "passphrasewhichneedstobe32bytes!"
	encrypted := Encrypt(plainText, key)

	if encrypted == plainText {
		t.Error("encryption failed")
	}

	decrypted := Decrypt(encrypted, key)
	if decrypted != plainText {
		t.Error("decryption failed")
	}

}
