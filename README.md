# go-util

一个轻量级、高性能的 Go 实用工具库，提供常用的字符串处理、随机数生成、时间操作、加密解密、文件操作等实用函数。

## 特性

- 🚀 高性能：优先使用高性能库（如 Sonic JSON）
- 🔒 安全：内置完善的加密解密功能
- 🧪 测试完备：核心功能均包含详尽的测试用例
- 🌍 国际化：支持国际手机号解析等功能

## 安装

```bash
go get github.com/chasespace/go-util
```

## 模块列表

| 模块          | 功能描述               | 包路径                                     |
|-------------|--------------------|-----------------------------------------|
| `ustr`      | 字符串处理              | `github.com/chasespace/go-util/ustr`    |
| `rand`      | 随机字符/数字/emoji生成等   | `github.com/chasespace/go-util/rand`    |
| `utime`     | 时间日期操作             | `github.com/chasespace/go-util/utime`   |
| `uregex`    | 正则表达式验证            | `github.com/chasespace/go-util/uregex`  |
| `ucrypto`   | 加密解密与哈希            | `github.com/chasespace/go-util/ucrypto` |
| `ufile`     | 文件操作               | `github.com/chasespace/go-util/ufile`   |
| `ugo`       | Go 语言工具函数          | `github.com/chasespace/go-util/ugo`     |
| `ujson`     | JSON 处理（基于 Sonic）  | `github.com/chasespace/go-util/ujson`   |
| `uhttp`     | HTTP 客户端（基于 Resty） | `github.com/chasespace/go-util/uhttp`   |
| `ui18`      | 国际化工具              | `github.com/chasespace/go-util/ui18`    |
| `uerr`      | 统一错误包装与码管理         | `github.com/chasespace/go-util/uerr`    |
| `uip`       | IP 归属与地理位置查询       | `github.com/chasespace/go-util/uip`     |
| `uid`       | UUID 拓展 UID 生成     | `github.com/chasespace/go-util/uid`     |
| `lifecycle` | 优雅生命周期管理           | `github.com/chasespace/go-util`         |

## 快速开始

### 字符串处理 (ustr)

```go
package main

import (
	"fmt"

	"github.com/chasespace/go-util/ustr"
)

func main() {
	truncated := ustr.TruncateUTF8("你好世界", 2, "...")
	snake := ustr.ToSnakeCase("CamelCase")
	camel := ustr.ToCamelCase("under_score")
	kebab := ustr.ToKebabCase("CamelCase")

	fmt.Println("truncated:", truncated)
	fmt.Println("snake:", snake)
	fmt.Println("camel:", camel)
	fmt.Println("kebab:", kebab)
}
```

### 随机数生成 (rand)

```go
package main

import (
	"fmt"

	"github.com/chasespace/go-util"
)

func main() {
	str := go_util.RandStr(16)
	digits := go_util.RandStrDigit(6)
	letters := go_util.RandStrLetter(8)
	num := go_util.RandInt(1, 100)
	f := go_util.RandFloat(0.0, 1.0)

	fmt.Println("random string:", str)
	fmt.Println("random digits:", digits)
	fmt.Println("random letters:", letters)
	fmt.Println("random int:", num)
	fmt.Printf("random float: %.4f\n", f)
}
```

### 时间日期操作 (utime)

```go
package main

import (
	"fmt"
	"time"

	"github.com/chasespace/go-util"
)

func main() {
	now := time.Now()
	dateNum := go_util.DateNumber()
	start := go_util.DayStartTime()
	end := go_util.DayEndTime()
	yesterdayStart := go_util.YesterdayStartTime()

	sameDay := go_util.IsSameDay(now, now)
	days := go_util.DaysBetween(now.AddDate(0, 0, -3), now)

	weekStart := go_util.BeginningOfWeek(now)
	weekEnd := go_util.EndOfWeek(now)
	monthStart, monthEnd, err := go_util.GetMonthStartEnd(now)
	if err != nil {
		fmt.Println("GetMonthStartEnd error:", err)
		return
	}

	fmt.Println("date number:", dateNum)
	fmt.Println("today range:", start, "->", end)
	fmt.Println("yesterday start:", yesterdayStart)
	fmt.Println("same day now:", sameDay)
	fmt.Println("days between:", days)
	fmt.Println("week range:", weekStart, "->", weekEnd)
	fmt.Println("month range:", monthStart, "->", monthEnd)
}
```

### 正则验证 (uregex)

```go
package main

import (
	"fmt"

	"github.com/chasespace/go-util"
)

func main() {
	fmt.Println("IsIP:", go_util.IsIP("192.168.1.1"))
	fmt.Println("IsIPv4:", go_util.IsIPv4("192.168.1.1"))
	fmt.Println("IsIPv6:", go_util.IsIPv6("::1"))
	fmt.Println("IsChinaPhone:", go_util.IsChinaPhone("13812345678"))
	fmt.Println("IsEmail:", go_util.IsEmail("test@example.com"))
	fmt.Println("IsURL:", go_util.IsURL("https://example.com"))
	fmt.Println("IsChinaIDCard:", go_util.IsChinaIDCard("110101199001011234"))
	fmt.Println("Masked phone:", go_util.MaskPhone("13812345678"))
	fmt.Println("Masked email:", go_util.MaskEmail("test@example.com"))
}
```

### 加密解密 (ucrypto)

```go
package main

import (
	"fmt"

	"github.com/chasespace/go-util/ucrypto"
)

func main() {
	key := []byte("16-byte-key-1234")
	plaintext := []byte("Hello, World!")
	ciphertext, err := ucrypto.AESEncrypt(plaintext, key)
	if err != nil {
		fmt.Println("AES encrypt error:", err)
		return
	}
	decrypted, err := ucrypto.AESDecrypt(ciphertext, key)
	if err != nil {
		fmt.Println("AES decrypt error:", err)
		return
	}
	fmt.Println("AES decrypted:", string(decrypted))

	priv, pub, err := ucrypto.GenerateRSAKeyPair(2048)
	if err != nil {
		fmt.Println("RSA key gen error:", err)
		return
	}
	encrypted, err := ucrypto.RSAEncrypt([]byte("secret"), pub)
	if err != nil {
		fmt.Println("RSA encrypt error:", err)
		return
	}
	decryptedRSA, err := ucrypto.RSADecrypt(encrypted, priv)
	if err != nil {
		fmt.Println("RSA decrypt error:", err)
		return
	}
	fmt.Println("RSA decrypted:", string(decryptedRSA))

	signature, err := ucrypto.RSASign([]byte("data"), priv)
	if err != nil {
		fmt.Println("RSA sign error:", err)
		return
	}
	if err := ucrypto.RSAVerify([]byte("data"), signature, pub); err != nil {
		fmt.Println("RSA verify failed:", err)
		return
	}
	fmt.Println("RSA signature verified")

	fmt.Println("MD5:", ucrypto.MD5("hello"))
	fmt.Println("SHA1:", ucrypto.SHA1("hello"))
	fmt.Println("SHA256:", ucrypto.SHA256("hello"))
	fmt.Println("SHA512:", ucrypto.SHA512("hello"))
	fmt.Println("Murmur32:", ucrypto.Murmur32("hello"))

	pemPriv := ucrypto.PrivateKeyToPEM(priv)
	pemPub := ucrypto.PublicKeyToPEM(pub)
	restoredPriv, err := ucrypto.PEMToPrivateKey(pemPriv)
	if err != nil {
		fmt.Println("PEM to private key error:", err)
		return
	}
	restoredPub, err := ucrypto.PEMToPublicKey(pemPub)
	if err != nil {
		fmt.Println("PEM to public key error:", err)
		return
	}
	fmt.Println("PEM roundtrip successful:", restoredPriv != nil && restoredPub != nil)
}
```

### 文件操作 (ufile)

```go
package main

import (
	"fmt"

	"github.com/chasespace/go-util/ufile"
)

func main() {
	md5Hash, err := ufile.GetFileMD5("README.md")
	if err != nil {
		fmt.Println("GetFileMD5 error:", err)
		return
	}
	fmt.Println("README.md MD5:", md5Hash)
}
```

### Go 语言工具 (ugo)

```go
package main

import (
	"fmt"
	"time"

	"github.com/chasespace/go-util/ugo"
)

func main() {
	ugo.Protect(func() {
		panic("simulated panic")
	}, func(err interface{}) {
		fmt.Println("caught panic:", err)
	})

	err := ugo.Retry(func() error {
		return fmt.Errorf("simulated failure")
	}, 3, time.Second)
	if err != nil {
		fmt.Println("retry failed:", err)
		return
	}
	fmt.Println("retry succeeded")
}
```

### JSON 处理 (ujson)

```go
package main

import (
	"fmt"

	"github.com/chasespace/go-util/ujson"
)

func main() {
	payload := map[string]interface{}{
		"name": "John",
		"age":  30,
	}
	jsonStr := ujson.MustJSON(payload)
	fmt.Println("serialized:", jsonStr)

	var result map[string]interface{}
	ujson.MustUnmarshal(jsonStr, &result)
	fmt.Println("deserialized:", result)

	ujson.MustUnmarshalBytes([]byte(jsonStr), &result)
	fmt.Println("deserialized from bytes:", result)
}
```

### 国际化工具 (ui18)

```go
package main

import (
	"fmt"

	"github.com/chasespace/go-util/ui18"
)

func main() {
	countryCode, mobile, err := ui18.ParsePhoneNum("8613812345678")
	if err != nil {
		fmt.Println("parse error:", err)
		return
	}
	fmt.Println("countryCode:", countryCode)
	fmt.Println("mobile:", mobile)
}
```

### 统一错误管理 (uerr)

```go
package main

import (
	"fmt"

	"github.com/chasespace/go-util/uerr"
)

func main() {
	sampleErr := fmt.Errorf("connection refused")
	handled := handle(sampleErr)
	fmt.Println("handled code:", handled.Code())

	mysqlErr := fmt.Errorf("timeout")
	wrapped := wrapMySQLError(mysqlErr)
	fmt.Println("wrapped detail:", wrapped.Detail())
}

func handle(err error) *uerr.UErr {
	// 任何 error 都可以转成 UErr，链路上统一携带码、明细、堆栈
	x := uerr.ToUErr(err)
	fmt.Println("detail:", x.Detail())
	fmt.Println("stack:\n", x.FormatStack())
	return x
}

func wrapMySQLError(mysqlErr error) *uerr.UErr {
	return uerr.NewWithError(uerr.CodeMySQLError, mysqlErr, "query failed")
}
```

### IP 归属 (uip)

```go
package main

import (
	"fmt"
	"log"

	"github.com/chasespace/go-util/uip"
)

func main() {
	uip.MustInitIp2Region("testdata/ip2region_v4.xdb")
	defer uip.CloseIp2Region()

	detail, err := uip.SearchIPv4("114.114.114.114")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ip detail: %+v\n", detail)
}
```

### UID 生成 (uid)

```go
package main

import (
	"fmt"
	"log"

	uidsdk "github.com/chasespace/go-util/uid"
)

func main() {
	uid, err := uidsdk.GetUIDWithUUID(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("uid:", uid)
}
```

### 生命周期管理 (lifecycle)

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chasespace/go-util"
)

func main() {
	fmt.Println("service running; send SIGINT/SIGTERM to trigger shutdown")
	go_util.Listening(5*time.Second, func(ctx context.Context) {
		fmt.Println("graceful shutdown starting")
	})
}
```

## 其他常用库推荐

- github.com/bytedance/sonic - 高性能JSON库
- github.com/casbin/casbin/v2 - 权限控制框架
- github.com/chaseSpace/lumberjack/v2 - 日志轮转
- github.com/google/uuid - UUID生成
- github.com/gorilla/websocket - WebSocket支持
- github.com/minio/minio-go/v7 - 对象存储客户端
- github.com/nyaruka/phonenumbers - 电话号码处理
- github.com/patrickmn/go-cache - 内存缓存
- github.com/redis/go-redis/v9 - Redis客户端
- github.com/robfig/cron/v3 - 定时任务
- github.com/samber/lo - 函数式编程工具
- github.com/segmentio/ksuid - KSUID生成
- github.com/spf13/cast - 类型转换
- github.com/stretchr/testify - 测试工具
- github.com/tidwall/gjson - JSON解析
- github.com/xuri/excelize/v2 - Excel处理
- gopkg.in/gomail.v2 - 邮件发送
- gorm.io/driver/mysql - MySQL驱动
- gorm.io/gorm - ORM框架

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
