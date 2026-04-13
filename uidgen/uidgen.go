package uidgen

import (
	"fmt"
	"math/big"

	"github.com/google/uuid"
)

// UIDGenAPI UID生成器接口
type UIDGenAPI interface {
	// GetOne 生成指定长度的UID
	// length: UID的位数限制（1-18位）
	GetOne(length int) (int64, error)
}

// 初始化函数，启用UUID随机池优化
func init() {
	uuid.EnableRandPool()
}

// GetUIDWithUUID 获取基于UUID的UID
func GetUIDWithUUID(length int) (int64, error) {
	return new(implUIDGenAPIWithUUID).GetOne(length)
}

// implUIDGenAPIWithUUID 基于UUID的UID生成器实现
/*
算法原理和步骤：
1. 使用UUID v7生成128位的时间戳-based随机数
2. 将UUID字节转换为大整数
3. 通过对10^length取模来限制UID位数
4. 返回int64类型的UID

优点：
- 基于时间戳，具有一定的有序性
- UUID v7提供了良好的随机性和分布
- 算法简单高效，性能较好
- 生成的UID在指定长度内基本唯一

缺点：
- 当length较小时，碰撞概率会增加
- 只支持最大18位数字（int64范围限制）
- 不能保证全局唯一，只能保证在限定长度内的相对唯一
- 对于高并发场景可能需要额外的去重（重试）机制
*/
type implUIDGenAPIWithUUID struct {
}

// GetOne 生成指定长度的UID
func (*implUIDGenAPIWithUUID) GetOne(length int) (int64, error) {
	if length < 1 {
		return 0, fmt.Errorf("length too small")
	}
	if length > 18 {
		return 0, fmt.Errorf("length too large")
	}

	// 步骤3: 计算模数 10^length，用于限制UID位数
	// 使用Exp方法进行正确的幂运算（避免位运算XOR的错误）
	modulus := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length)), nil)

	var min = new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length-1)), nil)

	for {
		// 步骤1: 生成UUID v7（基于时间戳的UUID）
		// UUID v7格式: timestamp(48bit) + random(74bit) + version(4bit) + variant(2bit)
		u, err := uuid.NewV7() // v7根据timestamp生成
		if err != nil {
			return 0, err
		}

		// 步骤2: 将UUID字节数组转换为大整数
		// UUID是16字节的数组，转换为一个很大的整数
		num := new(big.Int).SetBytes(u[:])

		// 步骤4: 取模运算得到指定位数内的UID
		result := new(big.Int).Mod(num, modulus)

		if result.Cmp(min) >= 0 {
			// 步骤5: 转换为int64返回
			return result.Int64(), nil
		}
	}
}
