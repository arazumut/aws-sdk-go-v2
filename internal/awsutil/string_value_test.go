package awsutil_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/smithy-go/ptr"
)

// testStruct adında bir yapı tanımlıyoruz
type testStruct struct {
	Field1 string
	Field2 *string
	Field3 []byte `sensitive:"true"` // Field3 alanı hassas veri olarak işaretlenmiş
	Value  []string
}

// TestStringValue fonksiyonu, awsutil.StringValue fonksiyonunu test ediyor
func TestStringValue(t *testing.T) {
	// Test vakalarını tanımlıyoruz
	cases := map[string]struct {
		Value  interface{}
		Expect string
	}{
		"general": {
			Value: testStruct{
				Field1: "abc123",
				Field2: ptr.String("abc123"),
				Field3: []byte("don't show me"), // Bu alan hassas veri içeriyor
				Value: []string{
					"first",
					"second",
				},
			},
			Expect: `{
  Field1: "abc123",
  Field2: "abc123",
  Field3: <sensitive>, // Hassas veri gizlenmiş
  Value: ["first","second"],
}`,
		},
	}

	// Her bir test vakasını çalıştırıyoruz
	for d, c := range cases {
		t.Run(d, func(t *testing.T) {
			actual := awsutil.StringValue(c.Value)
			if e, a := c.Expect, actual; e != a {
				t.Errorf("beklenen:\n%v\ngerçekleşen:\n%v\n", e, a)
			}
		})
	}
}
