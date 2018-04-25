package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("使用说明:\tTextKey.exe key file [mode]\nkey:\t任何可显示字符\nfile:\t要处理的文件\nmode:\t省略为加密,带有该项为解密(值任意)")
		os.Exit(0)
	}
	//第一个参数用于生成密钥
	arg1 := sha256.Sum224([]byte(os.Args[1]))
	key := arg1[:24]

	//第二个参数是目标文件
	fileName := os.Args[2]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("未找到待处理文件")
		os.Exit(0)
	}
	defer file.Close()
	//读取文件内容
	plain, _ := ioutil.ReadAll(file)
	//创建block
	block, _ := des.NewTripleDESCipher(key)

	//第三个参数表明是解密
	if len(os.Args) > 3 {
		DecryptMode := cipher.NewCBCDecrypter(block, key[:8])
		plain, _ = base64.StdEncoding.DecodeString(string(plain))
		DecryptMode.CryptBlocks(plain, plain)
		plain = PKCS5remove(plain)
		err := ioutil.WriteFile(fileName, plain, 0777)
		if err != nil {
			fmt.Println("保存解密后文件失败!")
		} else {
			fmt.Println("文件已解密!")
		}

	} else {
		EncryptMode := cipher.NewCBCEncrypter(block, key[:8])
		//明文补足PKCS5Padding
		plain = PKCS5append(plain)
		EncryptMode.CryptBlocks(plain, plain)
		err := ioutil.WriteFile(fileName, []byte(base64.StdEncoding.EncodeToString(plain)), 0777)
		if err != nil {
			fmt.Println("保存加密后文件失败!")
		} else {
			fmt.Println("文件已加密,务必记住加密key!")
		}
	}
}

func PKCS5append(plaintext []byte) []byte {
	num := 8 - len(plaintext)%8
	for i := 0; i < num; i++ {
		plaintext = append(plaintext, byte(num))
	}
	return plaintext
}

func PKCS5remove(plaintext []byte) []byte {
	length := len(plaintext)
	num := int(plaintext[length-1])
	return plaintext[:(length - num)]
}
