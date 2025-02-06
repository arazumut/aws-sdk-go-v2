package auth

import (
	"testing"

	smithy "github.com/aws/smithy-go"
)

func TestV4(t *testing.T) {
	// V4 özelliklerini tanımla
	propsV4 := smithy.Properties{}

	propsV4.Set("authSchemes", interface{}([]interface{}{
		map[string]interface{}{
			"disableDoubleEncoding": true,
			"name":                  "sigv4",
			"signingName":           "s3",
			"signingRegion":         "us-west-2",
		},
	}))

	// Kimlik doğrulama şemalarını al
	result, err := GetAuthenticationSchemes(&propsV4)
	if err != nil {
		t.Fatalf("Hata beklenmiyordu, aldık %v", err)
	}

	// Sonucun beklenen AuthenticationScheme türünde olup olmadığını kontrol et
	_, ok := result[0].(AuthenticationScheme)
	if !ok {
		t.Fatalf("Beklenen AuthenticationScheme alınamadı. %v", result[0])
	}

	// Sonucun beklenen AuthenticationSchemeV4 türünde olup olmadığını kontrol et
	v4Scheme, ok := result[0].(*AuthenticationSchemeV4)
	if !ok {
		t.Fatalf("Beklenen AuthenticationSchemeV4 alınamadı. %v", result[0])
	}

	// İmza sürüm adının doğru olup olmadığını kontrol et
	if v4Scheme.Name != "sigv4" {
		t.Fatalf("Beklenen AuthenticationSchemeV4 imza sürüm adı alınamadı")
	}
}

func TestV4A(t *testing.T) {
	// V4A özelliklerini tanımla
	propsV4A := smithy.Properties{}

	propsV4A.Set("authSchemes", []interface{}{
		map[string]interface{}{
			"disableDoubleEncoding": true,
			"name":                  "sigv4a",
			"signingName":           "s3",
			"signingRegionSet":      []string{"*"},
		},
	})

	// Kimlik doğrulama şemalarını al
	result, err := GetAuthenticationSchemes(&propsV4A)
	if err != nil {
		t.Fatalf("Hata beklenmiyordu, aldık %v", err)
	}

	// Sonucun beklenen AuthenticationScheme türünde olup olmadığını kontrol et
	_, ok := result[0].(AuthenticationScheme)
	if !ok {
		t.Fatalf("Beklenen AuthenticationScheme alınamadı. %v", result[0])
	}

	// Sonucun beklenen AuthenticationSchemeV4A türünde olup olmadığını kontrol et
	v4AScheme, ok := result[0].(*AuthenticationSchemeV4A)
	if !ok {
		t.Fatalf("Beklenen AuthenticationSchemeV4A alınamadı. %v", result[0])
	}

	// İmza sürüm adının doğru olup olmadığını kontrol et
	if v4AScheme.Name != "sigv4a" {
		t.Fatalf("Beklenen AuthenticationSchemeV4A imza sürüm adı alınamadı")
	}
}

func TestV4S3Express(t *testing.T) {
	// V4S3Express özelliklerini tanımla
	props := smithy.Properties{}
	props.Set("authSchemes", []interface{}{
		map[string]interface{}{
			"name":                  SigV4S3Express,
			"signingName":           "s3",
			"signingRegion":         "us-east-1",
			"disableDoubleEncoding": true,
		},
	})

	// Kimlik doğrulama şemalarını al
	result, err := GetAuthenticationSchemes(&props)
	if err != nil {
		t.Fatalf("Hata beklenmiyordu, aldık %v", err)
	}

	// Sonucun beklenen AuthenticationSchemeV4 türünde olup olmadığını kontrol et
	scheme, ok := result[0].(*AuthenticationSchemeV4)
	if !ok {
		t.Fatalf("Beklenen AuthenticationSchemeV4 alınamadı. %v", result[0])
	}

	// İmza sürüm adının doğru olup olmadığını kontrol et
	if scheme.Name != SigV4S3Express {
		t.Fatalf("beklenen %s, aldık %s", SigV4S3Express, scheme.Name)
	}
}
