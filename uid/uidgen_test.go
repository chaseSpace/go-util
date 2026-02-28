package uid

import (
	"fmt"
	"testing"
)

func TestImplUIDGenAPIWithUUID_GetOne(t *testing.T) {
	gen := &implUIDGenAPIWithUUID{}

	tests := []struct {
		name    string
		length  int
		wantErr bool
	}{
		{
			name:    "正常长度1",
			length:  1,
			wantErr: false,
		},
		{
			name:    "正常长度6",
			length:  6,
			wantErr: false,
		},
		{
			name:    "正常长度10",
			length:  10,
			wantErr: false,
		},
		{
			name:    "长度为0",
			length:  0,
			wantErr: true,
		},
		{
			name:    "负数长度",
			length:  -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gen.GetOne(tt.length)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// 验证生成的UID长度是否符合要求
				uidStr := toString(got)
				if len(uidStr) > tt.length {
					t.Errorf("GetOne() generated UID length = %d, want <= %d", len(uidStr), tt.length)
				}

				// 验证UID为非负数
				if got < 0 {
					t.Errorf("GetOne() generated negative UID = %d", got)
				}
			}
		})
	}
}

func TestGetOneMultipleTimes(t *testing.T) {
	gen := &implUIDGenAPIWithUUID{}
	length := 8

	// 生成多个UID，验证唯一性
	uids := make(map[int64]bool)
	count := 100

	for i := 0; i < count; i++ {
		uid, err := gen.GetOne(length)
		if err != nil {
			t.Fatalf("GetOne() failed: %v", err)
		}

		// 验证UID长度
		uidStr := toString(uid)
		if len(uidStr) > length {
			t.Errorf("Generated UID length = %d, want <= %d", len(uidStr), length)
		}

		// 检查重复
		if uids[uid] {
			t.Errorf("Duplicate UID generated: %d", uid)
		}
		uids[uid] = true
	}

	// 验证生成了期望数量的唯一UID
	if len(uids) != count {
		t.Errorf("Expected %d unique UIDs, got %d", count, len(uids))
	}
}

func TestDifferentLengths(t *testing.T) {
	gen := &implUIDGenAPIWithUUID{}

	testCases := []int{1, 3, 5, 8, 10}

	for _, length := range testCases {
		t.Run("length_"+toString(int64(length)), func(t *testing.T) {
			uid, err := gen.GetOne(length)
			if err != nil {
				t.Fatalf("GetOne() failed for length %d: %v", length, err)
			}

			uidStr := toString(uid)
			if len(uidStr) > length {
				t.Errorf("UID length mismatch for length %d: got %d digits", length, len(uidStr))
			}

			t.Logf("Generated UID for length %d: %d (%s)", length, uid, uidStr)
		})
	}
}

// 辅助函数：将int64转换为字符串
func toString(i int64) string {
	if i == 0 {
		return "0"
	}

	result := ""
	num := i
	if num < 0 {
		num = -num
	}

	for num > 0 {
		digit := num % 10
		result = fmt.Sprintf("%d%s", digit, result)
		num /= 10
	}

	if i < 0 {
		result = "-" + result
	}

	return result
}
