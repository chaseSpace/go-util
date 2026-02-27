package ucrypto

import (
	"bytes"
	"testing"
)

func TestAESEncryptDecrypt(t *testing.T) {
	// 测试数据
	plaintext := []byte("hello world")
	key := []byte("1234567890123456") // 16字节密钥

	// 加密
	ciphertext, err := AESEncrypt(plaintext, key)
	if err != nil {
		t.Fatalf("AES加密失败: %v", err)
	}

	// 解密
	decrypted, err := AESDecrypt(ciphertext, key)
	if err != nil {
		t.Fatalf("AES解密失败: %v", err)
	}

	// 验证
	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("解密结果不匹配. 期望: %s, 实际: %s", plaintext, decrypted)
	}
}

func TestAESInvalidKey(t *testing.T) {
	plaintext := []byte("test")
	invalidKey := []byte("short") // 无效密钥长度

	_, err := AESEncrypt(plaintext, invalidKey)
	if err == nil {
		t.Error("应该返回密钥长度错误")
	}
}

func TestAESDecryptTooShort(t *testing.T) {
	key := []byte("1234567890123456")
	shortCiphertext := []byte("short")

	_, err := AESDecrypt(shortCiphertext, key)
	if err == nil {
		t.Error("应该返回密文太短的错误")
	}
}

func TestGenerateRSAKeyPair(t *testing.T) {
	privateKey, publicKey, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("生成RSA密钥对失败: %v", err)
	}

	if privateKey == nil {
		t.Error("私钥不能为空")
	}

	if publicKey == nil {
		t.Error("公钥不能为空")
	}

	// 验证密钥长度
	if privateKey.N.BitLen() != 2048 {
		t.Errorf("私钥长度不正确，期望2048位，实际%d位", privateKey.N.BitLen())
	}
}

func TestRSAEncryptDecrypt(t *testing.T) {
	// 生成密钥对
	privateKey, publicKey, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	plaintext := []byte("hello RSA encryption")

	// 加密
	ciphertext, err := RSAEncrypt(plaintext, publicKey)
	if err != nil {
		t.Fatalf("RSA加密失败: %v", err)
	}

	// 解密
	decrypted, err := RSADecrypt(ciphertext, privateKey)
	if err != nil {
		t.Fatalf("RSA解密失败: %v", err)
	}

	// 验证
	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("RSA解密结果不匹配. 期望: %s, 实际: %s", plaintext, decrypted)
	}
}

func TestRSASignVerify(t *testing.T) {
	// 生成密钥对
	privateKey, publicKey, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	data := []byte("important data to sign")

	// 签名
	signature, err := RSASign(data, privateKey)
	if err != nil {
		t.Fatalf("RSA签名失败: %v", err)
	}

	// 验证签名
	err = RSAVerify(data, signature, publicKey)
	if err != nil {
		t.Errorf("RSA验签失败: %v", err)
	}

	// 验证错误签名
	wrongData := []byte("wrong data")
	err = RSAVerify(wrongData, signature, publicKey)
	if err == nil {
		t.Error("应该验签失败")
	}
}

func TestPEMConversion(t *testing.T) {
	// 生成密钥对
	privateKey, publicKey, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 私钥PEM转换
	privatePEM := PrivateKeyToPEM(privateKey)
	if len(privatePEM) == 0 {
		t.Error("私钥PEM转换失败")
	}

	// 公钥PEM转换
	publicPEM := PublicKeyToPEM(publicKey)
	if len(publicPEM) == 0 {
		t.Error("公钥PEM转换失败")
	}

	// PEM转回私钥
	restoredPrivate, err := PEMToPrivateKey(privatePEM)
	if err != nil {
		t.Fatalf("PEM转私钥失败: %v", err)
	}

	// PEM转回公钥
	restoredPublic, err := PEMToPublicKey(publicPEM)
	if err != nil {
		t.Fatalf("PEM转公钥失败: %v", err)
	}

	// 验证转换后的密钥是否一致
	if restoredPrivate.N.Cmp(privateKey.N) != 0 {
		t.Error("转换后的私钥不匹配")
	}

	if restoredPublic.N.Cmp(publicKey.N) != 0 {
		t.Error("转换后的公钥不匹配")
	}
}
