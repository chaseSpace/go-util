package ucrypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"

	"github.com/spaolacci/murmur3"
)

// MD5 计算字符串的MD5哈希值，返回16进制字符串
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// SHA1 计算字符串的SHA1哈希值，返回16进制字符串
func SHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256 计算字符串的SHA256哈希值，返回16进制字符串
func SHA256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// SHA512 计算字符串的SHA512哈希值，返回16进制字符串
func SHA512(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Hash 计算指定哈希算法的字符串哈希值，返回16进制字符串
func Hash(s string, h hash.Hash) string {
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Murmur32 计算字符串的MurmurHash32哈希值，返回16进制字符串
func Murmur32(s string, seed ...uint32) string {
	seed_ := uint32(0)
	if len(seed) > 0 {
		seed_ = seed[0]
	}
	h := murmur3.New32WithSeed(seed_)
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Murmur64 计算字符串的MurmurHash64哈希值，返回16进制字符串
func Murmur64(s string, seed ...uint32) string {
	seed_ := uint32(0)
	if len(seed) > 0 {
		seed_ = seed[0]
	}
	h := murmur3.New64WithSeed(seed_)
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Murmur128 计算字符串的MurmurHash128哈希值，返回16进制字符串
func Murmur128(s string, seed ...uint32) string {
	seed_ := uint32(0)
	if len(seed) > 0 {
		seed_ = seed[0]
	}
	h := murmur3.New128WithSeed(seed_)
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
