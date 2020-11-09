package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// DecodeWeAppUserInfo 解密微信小程序用户信息
func DecodeWeAppUserInfo(encryptedData string, sessionKey string, iv string) (result []byte, e error) {
	bytes, e := base64.StdEncoding.DecodeString(encryptedData)
	if e != nil {
		fmt.Println("encryptedData: ", encryptedData, "\n", e.Error())
		return nil, e
	}

	key, keyErr := base64.StdEncoding.DecodeString(sessionKey)
	if keyErr != nil {
		fmt.Println("sessionKey: ", sessionKey, "\n", keyErr.Error())
		return nil, keyErr
	}

	theIV, ivErr := base64.StdEncoding.DecodeString(iv)
	if ivErr != nil {
		fmt.Println("iv: ", iv, "\n", ivErr.Error())
		return nil, ivErr
	}

	result, e = AESDecrypt(bytes, key, theIV)
	return
}

func AESDecrypt(cipherText, key, iv []byte) ([]byte, error) {
	block, e := aes.NewCipher(key) //选择加密算法
	if e != nil {
		return nil, e
	}
	blockModel := cipher.NewCBCDecrypter(block, iv)
	plantText := make([]byte, len(cipherText))
	blockModel.CryptBlocks(plantText, cipherText)
	//if len(plantText)==0{
	//	return nil,errors.New("invalid plantText")
	//}
	for i, ch := range plantText {
		if ch == '\x0e' {
			plantText[i] = ' '
		}
	}
	//plantText = PKCS7UnPadding(plantText, block.BlockSize())
	return plantText, nil
}

func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}
