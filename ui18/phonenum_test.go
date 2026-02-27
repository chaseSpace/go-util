package ui18

import (
	"testing"
)

func TestParsePhoneNum(t *testing.T) {
	tests := []struct {
		name        string
		phone       string
		wantCountry string
		wantMobile  string
		expectError bool
	}{
		{
			name:        "中国手机号",
			phone:       "8613812345678",
			wantCountry: "86",
			wantMobile:  "13812345678",
			expectError: false,
		},
		{
			name:        "美国手机号",
			phone:       "12345678901",
			wantCountry: "1",
			wantMobile:  "2345678901",
			expectError: false,
		},
		{
			name:        "英国手机号",
			phone:       "447911123456",
			wantCountry: "44",
			wantMobile:  "7911123456",
			expectError: false,
		},
		{
			name:        "日本手机号",
			phone:       "819012345678",
			wantCountry: "81",
			wantMobile:  "9012345678",
			expectError: false,
		},
		{
			name:        "韩国手机号",
			phone:       "821012345678",
			wantCountry: "82",
			wantMobile:  "1012345678",
			expectError: false,
		},
		{
			name:        "德国手机号",
			phone:       "4915123456789",
			wantCountry: "49",
			wantMobile:  "15123456789",
			expectError: false,
		},
		{
			name:        "法国手机号",
			phone:       "33612345678",
			wantCountry: "33",
			wantMobile:  "612345678",
			expectError: false,
		},
		{
			name:        "加拿大手机号",
			phone:       "14161234567",
			wantCountry: "1",
			wantMobile:  "4161234567",
			expectError: false,
		},
		{
			name:        "澳大利亚手机号",
			phone:       "61412345678",
			wantCountry: "61",
			wantMobile:  "412345678",
			expectError: false,
		},
		{
			name:        "印度手机号",
			phone:       "919876543210",
			wantCountry: "91",
			wantMobile:  "9876543210",
			expectError: false,
		},
		{
			name:        "巴西手机号",
			phone:       "5511987654321",
			wantCountry: "55",
			wantMobile:  "11987654321",
			expectError: false,
		},
		{
			name:        "俄罗斯手机号",
			phone:       "79123456789",
			wantCountry: "7",
			wantMobile:  "9123456789",
			expectError: false,
		},
		{
			name:        "无效号码",
			phone:       "invalid",
			wantCountry: "",
			wantMobile:  "",
			expectError: true,
		},
		{
			name:        "空号码",
			phone:       "",
			wantCountry: "",
			wantMobile:  "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCountry, gotMobile, err := ParsePhoneNum(tt.phone)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if gotCountry != tt.wantCountry {
				t.Errorf("country code = %v, want %v", gotCountry, tt.wantCountry)
			}

			if gotMobile != tt.wantMobile {
				t.Errorf("mobile = %v, want %v", gotMobile, tt.wantMobile)
			}
		})
	}
}
