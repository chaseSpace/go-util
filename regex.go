package go_util

import (
	"net"
	"regexp"
)

var (
	// IP地址正则表达式
	ipRegex = regexp.MustCompile(`^((2((5[0-5])|([0-4]\d)))|([0-1]?\d{1,2}))(\.((2((5[0-5])|([0-4]\d)))|([0-1]?\d{1,2}))){3}$`)

	// 手机号正则表达式（中国手机号）
	chinaPhoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)

	// 邮箱正则表达式
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// IPv6地址正则表达式
	ipv6Regex = regexp.MustCompile(`^([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}$|^::1$|^::$`)

	// URL正则表达式
	urlRegex = regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].\S*$`)

	// 身份证号正则表达式（18位）
	chinaIDCardRegex = regexp.MustCompile(`^[1-9]\d{5}(18|19|20)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
)

// IsIP 判断是否为有效的IPv4地址
func IsIP(ip string) bool {
	return ipRegex.MatchString(ip)
}

// IsIPv4 判断是否为有效的IPv4地址（使用标准库验证）
func IsIPv4(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil && parsedIP.To4() != nil
}

// IsIPv6 判断是否为有效的IPv6地址
func IsIPv6(ip string) bool {
	return ipv6Regex.MatchString(ip) || (net.ParseIP(ip) != nil && net.ParseIP(ip).To4() == nil)
}

// IsChinaPhone 判断是否为有效的中国手机号
func IsChinaPhone(phone string) bool {
	return chinaPhoneRegex.MatchString(phone)
}

// IsEmail 判断是否为有效的邮箱地址
func IsEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// IsURL 判断是否为有效的URL
func IsURL(url string) bool {
	return urlRegex.MatchString(url)
}

// IsChinaIDCard 判断是否为有效的18位身份证号
func IsChinaIDCard(idCard string) bool {
	return chinaIDCardRegex.MatchString(idCard)
}

// MaskPhone 隐藏手机号中间4位
func MaskPhone(phone string) string {
	if !IsChinaPhone(phone) {
		return phone
	}
	return phone[:3] + "****" + phone[7:]
}

// MaskEmail 隐藏邮箱用户名部分
func MaskEmail(email string) string {
	if !IsEmail(email) {
		return email
	}

	parts := regexp.MustCompile(`@`).Split(email, 2)
	if len(parts) != 2 {
		return email
	}

	username := parts[0]
	domain := parts[1]

	if len(username) <= 2 {
		return email
	}

	// 保留首尾字符，中间用*代替
	maskedUsername := string(username[0]) + "***" + string(username[len(username)-1])
	return maskedUsername + "@" + domain
}
