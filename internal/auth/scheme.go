package auth

import (
	"context"
	"fmt"

	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

// SigV4, SigV4A, SigV4S3Express ve None sabitleri AWS kimlik doğrulama şemalarını temsil eder.
const (
	SigV4          = "sigv4"
	SigV4A         = "sigv4a"
	SigV4S3Express = "sigv4-s3express"
	None           = "none"
)

// SupportedSchemes, desteklenen AWS kimlik doğrulama şemalarının listesini gösteren bir veri yapısıdır.
var SupportedSchemes = map[string]bool{
	SigV4:          true,
	SigV4A:         true,
	SigV4S3Express: true,
	None:           true,
}

// AuthenticationScheme, AWS kimlik doğrulama şemalarının bir temsilidir.
type AuthenticationScheme interface {
	isAuthenticationScheme()
}

// AuthenticationSchemeV4, AWS SigV4'ün bir temsilidir.
type AuthenticationSchemeV4 struct {
	Name                  string
	SigningName           *string
	SigningRegion         *string
	DisableDoubleEncoding *bool
}

func (a *AuthenticationSchemeV4) isAuthenticationScheme() {}

// AuthenticationSchemeV4A, AWS SigV4A'nın bir temsilidir.
type AuthenticationSchemeV4A struct {
	Name                  string
	SigningName           *string
	SigningRegionSet      []string
	DisableDoubleEncoding *bool
}

func (a *AuthenticationSchemeV4A) isAuthenticationScheme() {}

// AuthenticationSchemeNone, none kimlik doğrulama şemasının bir temsilidir.
type AuthenticationSchemeNone struct{}

func (a *AuthenticationSchemeNone) isAuthenticationScheme() {}

// NoAuthenticationSchemesFoundError, hiçbir kimlik doğrulama şeması belirtilmediğini işaret etmek için kullanılır.
type NoAuthenticationSchemesFoundError struct{}

func (e *NoAuthenticationSchemesFoundError) Error() string {
	return fmt.Sprint("Hiçbir kimlik doğrulama şeması belirtilmedi.")
}

// UnSupportedAuthenticationSchemeSpecifiedError, yalnızca desteklenmeyen kimlik doğrulama şemalarının belirtildiğini işaret etmek için kullanılır.
type UnSupportedAuthenticationSchemeSpecifiedError struct {
	UnsupportedSchemes []string
}

func (e *UnSupportedAuthenticationSchemeSpecifiedError) Error() string {
	return fmt.Sprint("Desteklenmeyen kimlik doğrulama şeması belirtildi.")
}

// GetAuthenticationSchemes, ilgili kimlik doğrulama şeması verilerini özel olarak güçlü bir şekilde yazılmış Go veri yapısına çıkarır.
func GetAuthenticationSchemes(p *smithy.Properties) ([]AuthenticationScheme, error) {
	var result []AuthenticationScheme
	if !p.Has("authSchemes") {
		return nil, &NoAuthenticationSchemesFoundError{}
	}

	authSchemes, _ := p.Get("authSchemes").([]interface{})

	var unsupportedSchemes []string
	for _, scheme := range authSchemes {
		authScheme, _ := scheme.(map[string]interface{})

		version := authScheme["name"].(string)
		switch version {
		case SigV4, SigV4S3Express:
			v4Scheme := AuthenticationSchemeV4{
				Name:                  version,
				SigningName:           getSigningName(authScheme),
				SigningRegion:         getSigningRegion(authScheme),
				DisableDoubleEncoding: getDisableDoubleEncoding(authScheme),
			}
			result = append(result, AuthenticationScheme(&v4Scheme))
		case SigV4A:
			v4aScheme := AuthenticationSchemeV4A{
				Name:                  SigV4A,
				SigningName:           getSigningName(authScheme),
				SigningRegionSet:      getSigningRegionSet(authScheme),
				DisableDoubleEncoding: getDisableDoubleEncoding(authScheme),
			}
			result = append(result, AuthenticationScheme(&v4aScheme))
		case None:
			noneScheme := AuthenticationSchemeNone{}
			result = append(result, AuthenticationScheme(&noneScheme))
		default:
			unsupportedSchemes = append(unsupportedSchemes, authScheme["name"].(string))
			continue
		}
	}

	if len(result) == 0 {
		return nil, &UnSupportedAuthenticationSchemeSpecifiedError{
			UnsupportedSchemes: unsupportedSchemes,
		}
	}

	return result, nil
}

type disableDoubleEncoding struct{}

// SetDisableDoubleEncoding, bağlam üzerinde çift kodlamayı devre dışı bırakma seçeneğini ayarlar veya değiştirir.
func SetDisableDoubleEncoding(ctx context.Context, value bool) context.Context {
	return middleware.WithStackValue(ctx, disableDoubleEncoding{}, value)
}

// GetDisableDoubleEncoding, bağlamdan çift kodlamayı devre dışı bırakma seçeneğini alır.
func GetDisableDoubleEncoding(ctx context.Context) (value bool, ok bool) {
	value, ok = middleware.GetStackValue(ctx, disableDoubleEncoding{}).(bool)
	return value, ok
}

func getSigningName(authScheme map[string]interface{}) *string {
	signingName, ok := authScheme["signingName"].(string)
	if !ok || signingName == "" {
		return nil
	}
	return &signingName
}

func getSigningRegionSet(authScheme map[string]interface{}) []string {
	untypedSigningRegionSet, ok := authScheme["signingRegionSet"].([]interface{})
	if !ok {
		return nil
	}
	signingRegionSet := []string{}
	for _, item := range untypedSigningRegionSet {
		signingRegionSet = append(signingRegionSet, item.(string))
	}
	return signingRegionSet
}

func getSigningRegion(authScheme map[string]interface{}) *string {
	signingRegion, ok := authScheme["signingRegion"].(string)
	if !ok || signingRegion == "" {
		return nil
	}
	return &signingRegion
}

func getDisableDoubleEncoding(authScheme map[string]interface{}) *bool {
	disableDoubleEncoding, ok := authScheme["disableDoubleEncoding"].(bool)
	if !ok {
		return nil
	}
	return &disableDoubleEncoding
}
