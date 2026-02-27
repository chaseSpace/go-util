# go-util

一个轻量级、高性能的 Go 实用工具库，提供常用的字符串处理、随机数生成、时间操作、加密解密、文件操作等实用函数。

## 特性

- 🚀 高性能：优先使用高性能库（如 Sonic JSON）
- 🔒 安全：内置完善的加密解密功能
- 📦 零依赖：核心功能零外部依赖
- ✅ 完整测试：所有函数都有单元测试覆盖
- 🌍 国际化：支持国际手机号解析等功能

## 安装

```bash
go get github.com/chasespace/go-util
```

## 模块列表

| 模块        | 功能描述               | 包路径                                     |
|-----------|--------------------|-----------------------------------------|
| `ustr`    | 字符串处理              | `github.com/chasespace/go-util/ustr`    |
| `rand`    | 随机数生成              | `github.com/chasespace/go-util/rand`    |
| `utime`   | 时间日期操作             | `github.com/chasespace/go-util/utime`   |
| `uregex`  | 正则表达式验证            | `github.com/chasespace/go-util/uregex`  |
| `ucrypto` | 加密解密与哈希            | `github.com/chasespace/go-util/ucrypto` |
| `ufile`   | 文件操作               | `github.com/chasespace/go-util/ufile`   |
| `ugo`     | Go 语言工具函数          | `github.com/chasespace/go-util/ugo`     |
| `ujson`   | JSON 处理（基于 Sonic）  | `github.com/chasespace/go-util/ujson`   |
| `uhttp`   | HTTP 客户端（基于 Resty） | `github.com/chasespace/go-util/uhttp`   |
| `ui18`    | 国际化工具              | `github.com/chasespace/go-util/ui18`    |

## 快速开始

### 字符串处理 (ustr)

```go
import "github.com/chasespace/go-util/ustr"

// UTF-8 截断
s := ustr.TruncateUTF8("你好世界", 2, "...") // "你好..."

// 命名转换
str := ustr.ToSnakeCase("CamelCase") // "camel_case"
str = ustr.ToCamelCase("under_score") // "UnderScore"
str = ustr.ToKebabCase("CamelCase") // "camel-case"
```

### 随机数生成 (rand)

```go
import "github.com/chasespace/go-util/rand"

// 生成随机字符串
str := rand.RandStr(16)        // "aB3dE5fG7hI9jK1l"
str = rand.RandStrDigit(6)     // "123456"
str = rand.RandStrLetter(8) // "aBcDeFgH"

// 随机数
num := rand.RandInt(1, 100) // 42
f := rand.RandFloat(0.0, 1.0) // 0.7234
```

### 时间日期操作 (utime)

```go
import "github.com/chasespace/go-util/utime"

// 日期数字
dateNum := utime.DateNumber() // 20250227

// 时间范围
start := utime.DayStartTime() // 今日 00:00:00
end := utime.DayEndTime() // 今日 23:59:59
yesterdayStart := utime.YesterdayStartTime()

// 日期判断
sameDay := utime.IsSameDay(time.Now(), time.Now()) // true
days := utime.DaysBetween(t1, t2) // 3

// 周与月
weekStart := utime.BeginningOfWeek(time.Now())
weekEnd := utime.EndOfWeek(time.Now())
monthStart, monthEnd, _ := utime.GetMonthStartEnd(time.Now())
```

### 正则验证 (uregex)

```go
import "github.com/chasespace/go-util/uregex"

// 验证
valid := uregex.IsIP("192.168.1.1")           // true
valid = uregex.IsIPv4("192.168.1.1")          // true
valid = uregex.IsIPv6("::1") // true
valid = uregex.IsChinaPhone("13812345678") // true
valid = uregex.IsEmail("test@example.com") // true
valid = uregex.IsURL("https://example.com") // true
valid = uregex.IsChinaIDCard("110101199001011234") // true

// 数据脱敏
masked := uregex.MaskPhone("13812345678") // "138****5678"
masked = uregex.MaskEmail("test@example.com") // "t***t@example.com"
```

### 加密解密 (ucrypto)

```go
import "github.com/chasespace/go-util/ucrypto"

// AES 加密解密
key := []byte("16-byte-key-1234")
plaintext := []byte("Hello, World!")
ciphertext, _ := ucrypto.AESEncrypt(plaintext, key)
decrypted, _ := ucrypto.AESDecrypt(ciphertext, key)

// RSA 加密解密
privateKey, publicKey, _ := ucrypto.GenerateRSAKeyPair(2048)
encrypted, _ := ucrypto.RSAEncrypt([]byte("secret"), publicKey)
decrypted, _ := ucrypto.RSADecrypt(encrypted, privateKey)

// RSA 签名验签
signature, _ := ucrypto.RSASign([]byte("data"), privateKey)
err := ucrypto.RSAVerify([]byte("data"), signature, publicKey)

// 哈希计算
hash := ucrypto.MD5("hello") // "5d41402abc4b2a76b9719d911017c592"
hash = ucrypto.SHA1("hello") // "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
hash = ucrypto.SHA256("hello") // "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
hash = ucrypto.SHA512("hello") // "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"
hash = ucrypto.Murmur32("hello")      // "3a6c7b2a"

// PEM 格式转换
pemPriv := ucrypto.PrivateKeyToPEM(privateKey)
pemPub := ucrypto.PublicKeyToPEM(publicKey)
privKey, _ := ucrypto.PEMToPrivateKey(pemPriv)
pubKey, _ := ucrypto.PEMToPublicKey(pemPub)
```

### 文件操作 (ufile)

```go
import "github.com/chasespace/go-util/ufile"

// 计算文件 MD5
md5Hash, err := ufile.GetFileMD5("/path/to/file.txt")
```

### Go 语言工具 (ugo)

```go
import "github.com/chasespace/go-util/ugo"

// 异常捕获
ugo.Protect(func () {
// 可能 panic 的代码
}, func (err interface{}) {
// 异常处理
})

// 重试机制
err := ugo.Retry(func () error {
// 可能失败的操作
return nil
}, 3, time.Second) // 最多重试3次，间隔1秒
```

### 国际化工具 (ui18)

```go
import "github.com/chasespace/go-util/ui18"

// 解析国际手机号
countryCode, mobile, err := ui18.ParsePhoneNum("8613812345678")
// countryCode: "86", mobile: "13812345678"
```

## 测试

运行所有测试：

```bash
go test ./...
```

运行特定模块测试：

```bash
go test ./ustr
go test ./ucrypto
```

## 依赖

本项目部分模块依赖以下第三方库：

- [Sonic](https://github.com/bytedance/sonic) - 高性能 JSON 库
- [Resty](https://github.com/go-resty/resty) - HTTP 客户端
- [phonenumbers](https://github.com/nyaruka/phonenumbers) - 国际手机号解析
- [murmur3](https://github.com/spaolacci/murmur3) - MurmurHash 算法

## 许可证

本项目采用 MIT 许可证。详见 [LICENSE](LICENSE) 文件。

## 贡献

欢迎提交 Issue 和 Pull Request！
