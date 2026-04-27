package tasks

import (
	"errors"
	"strings"
)

// Codec ID编解码器
type Codec struct {
	alphabet string
	base     int
	minLen   int
	salt     []int // 混淆用的盐值数组
}

// NewCodec 创建编解码器
// alphabet: 使用的字符集（建议唯一且不重复）
// minLen: 最小编码长度
// salt: 混淆盐值（可选）
func NewCodec(alphabet string, minLen int, salt ...int) (*Codec, error) {
	if len(alphabet) < 2 {
		return nil, errors.New("alphabet length must be at least 2")
	}

	// 检查字符是否重复
	seen := make(map[rune]bool)
	for _, ch := range alphabet {
		if seen[ch] {
			return nil, errors.New("alphabet contains duplicate characters")
		}
		seen[ch] = true
	}

	codec := &Codec{
		alphabet: alphabet,
		base:     len(alphabet),
		minLen:   minLen,
	}

	// 初始化盐值
	if len(salt) > 0 && salt[0] != 0 {
		codec.salt = make([]int, minLen)
		for i := 0; i < minLen; i++ {
			codec.salt[i] = (salt[0] >> uint(i)) & 0xFF
		}
	}

	return codec, nil
}

// Encode 将数字ID编码为字符串
func (c *Codec) Encode(id int64) string {
	if id < 0 {
		id = -id
	}

	// 添加混淆（如果有盐值）
	encodedId := c.mix(id)

	// 转换为指定进制
	result := c.toBase(encodedId)

	// 确保最小长度
	if len(result) < c.minLen {
		padding := c.minLen - len(result)
		// 使用第一个字符作为填充
		result = strings.Repeat(string(c.alphabet[0]), padding) + result
	}

	return result
}

// Decode 将字符串解码为数字ID
func (c *Codec) Decode(code string) (int64, error) {
	if code == "" {
		return 0, errors.New("empty code")
	}

	// 去除前导填充字符
	trimmed := strings.TrimLeft(code, string(c.alphabet[0]))
	if trimmed == "" {
		return 0, nil
	}

	// 从指定进制转换回来
	id := c.fromBase(trimmed)
	if id == -1 {
		return 0, errors.New("invalid code")
	}

	// 去混淆
	decodedID := c.unmix(id)

	return decodedID, nil
}

// toBase 转换为指定进制字符串
func (c *Codec) toBase(num int64) string {
	if num == 0 {
		return string(c.alphabet[0])
	}

	var result strings.Builder
	n := num

	for n > 0 {
		remainder := n % int64(c.base)
		result.WriteByte(c.alphabet[remainder])
		n = n / int64(c.base)
	}

	// 反转字符串（因为是从低位到高位构建的）
	return c.reverseString(result.String())
}

// fromBase 从进制字符串转换为数字
func (c *Codec) fromBase(str string) int64 {
	var result int64 = 0

	for i := 0; i < len(str); i++ {
		char := str[i]
		pos := strings.IndexByte(c.alphabet, char)
		if pos == -1 {
			return -1
		}
		result = result*int64(c.base) + int64(pos)
	}

	return result
}

// mix 混淆ID（简单异或和位移）
func (c *Codec) mix(id int64) int64 {
	if len(c.salt) == 0 {
		return id
	}

	// 使用salt进行简单混淆
	mixed := id
	for i, s := range c.salt {
		mixed ^= int64(s) << uint(i*4)
	}
	// 添加一些额外的混淆
	mixed = (mixed ^ 0x9E3779B97F4A7C15) ^ (mixed >> 31)

	return mixed
}

// unmix 解混淆
func (c *Codec) unmix(mixed int64) int64 {
	if len(c.salt) == 0 {
		return mixed
	}

	// 逆向混淆
	unmixed := mixed ^ (mixed >> 31) ^ 0x9E3779B97F4A7C15

	for i, s := range c.salt {
		unmixed ^= int64(s) << uint(i*4)
	}

	return unmixed
}

// reverseString 反转字符串
func (c *Codec) reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
