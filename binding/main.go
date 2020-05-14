package main

//#include <stdint.h>
//typedef struct { char *publicKey; char *privateKey; } KeyPair;
//typedef struct { char *hash; char *cipher; char *compression; char *compressionLevel; char *rsaBits; } KeyOptions;
//typedef struct { char *name; char *comment; char *email; char *passphrase; KeyOptions keyOptions; } Options;
import "C"
import (
	"fmt"
	"github.com/jerson/openpgp-mobile/openpgp"
	"strconv"
	"unsafe"
)

var instance = openpgp.NewFastOpenPGP()

func errorThrow(err error) {
	fmt.Println(err.Error())

	//openpgp_bridge.ErrorGenerateThrow(err.Error())
}

//export Encrypt
func Encrypt(message, publicKey *C.char) *C.char {
	result, err := instance.Encrypt(C.GoString(message), C.GoString(publicKey))
	if err != nil {
		errorThrow(err)
		return nil
	}
	return C.CString(result)
}

//export EncryptBytes
func EncryptBytes(message unsafe.Pointer, messageSize C.int, publicKey *C.char) (unsafe.Pointer, C.int) {
	result, err := instance.EncryptBytes(C.GoBytes(message, messageSize), C.GoString(publicKey))
	if err != nil {
		errorThrow(err)
		return nil, C.int(0)
	}
	return C.CBytes(result), C.int(len(result))
}

//export Decrypt
func Decrypt(message, privateKey, passphrase *C.char) *C.char {
	result, err := instance.Decrypt(C.GoString(message), C.GoString(privateKey), C.GoString(passphrase))
	if err != nil {
		errorThrow(err)
		return nil
	}
	return C.CString(result)
}

//export DecryptBytes
func DecryptBytes(message unsafe.Pointer, messageSize C.int, privateKey, passphrase *C.char) (unsafe.Pointer, C.int) {
	result, err := instance.DecryptBytes(C.GoBytes(message, messageSize), C.GoString(privateKey), C.GoString(passphrase))
	if err != nil {
		errorThrow(err)
		return nil, C.int(0)
	}
	return C.CBytes(result), C.int(len(result))
}

//export Sign
func Sign(message, publicKey, privateKey, passphrase *C.char) *C.char {
	result, err := instance.Sign(C.GoString(message), C.GoString(publicKey), C.GoString(privateKey), C.GoString(passphrase))
	if err != nil {
		errorThrow(err)
		return nil
	}
	return C.CString(result)
}

//export SignBytes
func SignBytes(message unsafe.Pointer, messageSize C.int, publicKey, privateKey, passphrase *C.char) (unsafe.Pointer, C.int) {
	result, err := instance.SignBytes(C.GoBytes(message, messageSize), C.GoString(publicKey), C.GoString(privateKey), C.GoString(passphrase))
	if err != nil {
		errorThrow(err)
		return nil, C.int(0)
	}
	return C.CBytes(result), C.int(len(result))
}

//export SignBytesToString
func SignBytesToString(message unsafe.Pointer, messageSize C.int, publicKey, privateKey, passphrase *C.char) *C.char {
	result, err := instance.SignBytesToString(C.GoBytes(message, messageSize), C.GoString(publicKey), C.GoString(privateKey), C.GoString(passphrase))
	if err != nil {
		errorThrow(err)
		return nil
	}
	return C.CString(result)
}

//export Verify
func Verify(signature, message, publicKey *C.char) *C.char {
	result, err := instance.Verify(C.GoString(signature), C.GoString(message), C.GoString(publicKey))
	if err != nil {
		errorThrow(err)
		return nil
	}
	if result {
		return C.CString("1")
	}
	return C.CString("")
}

//export VerifyBytes
func VerifyBytes(signature *C.char, message unsafe.Pointer, messageSize C.int, publicKey *C.char) *C.char {
	result, err := instance.VerifyBytes(C.GoString(signature), C.GoBytes(message, messageSize), C.GoString(publicKey))
	if err != nil {
		errorThrow(err)
		return nil
	}
	if result {
		return C.CString("1")
	}
	return C.CString("")
}

//export EncryptSymmetric
func EncryptSymmetric(message, passphrase *C.char, options C.KeyOptions) *C.char {

	result, err := instance.EncryptSymmetric(C.GoString(message), C.GoString(passphrase), getKeyOptions(options))
	if err != nil {
		errorThrow(err)
		return nil
	}
	return C.CString(result)
}

//export EncryptSymmetricBytes
func EncryptSymmetricBytes(message unsafe.Pointer, messageSize C.int, passphrase *C.char, options C.KeyOptions) (unsafe.Pointer, C.int) {

	result, err := instance.EncryptSymmetricBytes(C.GoBytes(message, messageSize), C.GoString(passphrase), getKeyOptions(options))
	if err != nil {
		errorThrow(err)
		return nil, C.int(0)
	}
	return C.CBytes(result), C.int(len(result))
}

//export DecryptSymmetric
func DecryptSymmetric(message, passphrase *C.char, options C.KeyOptions) *C.char {
	result, err := instance.DecryptSymmetric(C.GoString(message), C.GoString(passphrase), getKeyOptions(options))
	if err != nil {
		errorThrow(err)
		return nil
	}
	return C.CString(result)
}

//export DecryptSymmetricBytes
func DecryptSymmetricBytes(message unsafe.Pointer, messageSize C.int, passphrase *C.char, options C.KeyOptions) (unsafe.Pointer, C.int) {
	result, err := instance.DecryptSymmetricBytes(C.GoBytes(message, messageSize), C.GoString(passphrase), getKeyOptions(options))
	if err != nil {
		errorThrow(err)
		return nil, C.int(0)
	}
	return C.CBytes(result), C.int(len(result))
}

//export Generate
func Generate(options C.Options) C.KeyPair {
	result, err := instance.Generate(getOptions(options))
	if err != nil {
		errorThrow(err)
		return C.KeyPair{C.CString(""), C.CString("")}

	}
	return C.KeyPair{C.CString(result.PublicKey), C.CString(result.PrivateKey)}

}

func getKeyOptions(options C.KeyOptions) *openpgp.KeyOptions {

	compressionLevel, _ := strconv.Atoi(C.GoString(options.compressionLevel))
	rsaBits, _ := strconv.Atoi(C.GoString(options.rsaBits))

	return &openpgp.KeyOptions{
		Hash:             C.GoString(options.hash),
		Cipher:           C.GoString(options.cipher),
		Compression:      C.GoString(options.compression),
		CompressionLevel: compressionLevel,
		RSABits:          rsaBits,
	}
}

func getOptions(options C.Options) *openpgp.Options {

	return &openpgp.Options{
		KeyOptions: getKeyOptions(options.keyOptions),
		Name:       C.GoString(options.name),
		Comment:    C.GoString(options.comment),
		Email:      C.GoString(options.email),
		Passphrase: C.GoString(options.passphrase),
	}
}

func main() {}
