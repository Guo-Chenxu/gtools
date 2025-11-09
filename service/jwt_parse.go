package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"gtools/biz/model/gtools"
	"gtools/consts"

	"github.com/golang-jwt/jwt/v5"
)

// ParseJWT 解析JWT token
func ParseJWT(ctx context.Context, req *gtools.ParseJWTReq) (*gtools.JWTParseData, *consts.BizCode) {
	result := &gtools.JWTParseData{
		Header:   make(map[string]string),
		Claims:   make(map[string]string),
		Valid:    false,
		ErrorMsg: "",
	}

	// 分割JWT字符串
	parts := strings.Split(req.JwtToken, ".")
	if len(parts) != 3 {
		result.ErrorMsg = "Invalid JWT format: must have 3 parts separated by dots"
		return result, nil
	}

	// 解析Header
	headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		result.ErrorMsg = fmt.Sprintf("Failed to decode header: %v", err)
		return result, nil
	}

	var headerMap map[string]interface{}
	decoder := json.NewDecoder(bytes.NewReader(headerBytes))
	decoder.UseNumber() // 使用json.Number保持数字精度
	if err := decoder.Decode(&headerMap); err != nil {
		result.ErrorMsg = fmt.Sprintf("Failed to parse header JSON: %v", err)
		return result, nil
	}

	// 将header转换为string map，处理int64精度问题
	for k, v := range headerMap {
		result.Header[k] = convertToString(v)
	}

	// 解析Claims (不验证签名)
	claimsBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		result.ErrorMsg = fmt.Sprintf("Failed to decode claims: %v", err)
		return result, nil
	}

	var claimsMap map[string]interface{}
	decoder = json.NewDecoder(bytes.NewReader(claimsBytes))
	decoder.UseNumber() // 使用json.Number保持数字精度
	if err := decoder.Decode(&claimsMap); err != nil {
		result.ErrorMsg = fmt.Sprintf("Failed to parse claims JSON: %v", err)
		return result, nil
	}

	// 将claims转换为string map，处理int64精度问题
	for k, v := range claimsMap {
		result.Claims[k] = convertToString(v)
	}

	// 如果提供了密钥，则验证签名
	if req.Secret != nil && *req.Secret != "" {
		token, err := jwt.Parse(req.JwtToken, func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				// 尝试RSA
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					// 尝试ECDSA
					if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
				}
			}
			// 对于HMAC算法，返回密钥字节
			return []byte(*req.Secret), nil
		})

		if err != nil {
			result.ErrorMsg = fmt.Sprintf("Token validation failed: %v", err)
			result.Valid = false
		} else if token.Valid {
			result.Valid = true
			result.ErrorMsg = ""
		} else {
			result.ErrorMsg = "Token is invalid"
			result.Valid = false
		}
	} else {
		// 没有提供密钥，只解析不验证
		result.ErrorMsg = "Token parsed without validation (no secret provided)"
	}

	return result, nil
}

// convertToString 将interface{}转换为字符串，保持精度
func convertToString(v interface{}) string {
	switch val := v.(type) {
	case json.Number:
		// json.Number保持了原始数字字符串，不会丢失精度
		return val.String()
	case string:
		return val
	case float64:
		// 检查是否为整数
		if val == float64(int64(val)) {
			return fmt.Sprintf("%d", int64(val))
		}
		return fmt.Sprintf("%f", val)
	case int64:
		return fmt.Sprintf("%d", val)
	case int:
		return fmt.Sprintf("%d", val)
	case bool:
		return fmt.Sprintf("%t", val)
	case nil:
		return "null"
	default:
		// 对于复杂类型，转换为JSON字符串
		bytes, _ := json.Marshal(val)
		return string(bytes)
	}
}
