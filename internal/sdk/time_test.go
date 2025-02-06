package sdk

import (
	"context"
	"strings"
	"testing"
	"time"
)

// sleepWithContext fonksiyonunu test eder
func TestSleepWithContext(t *testing.T) {
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	err := sleepWithContext(ctx, 1*time.Millisecond)
	if err != nil {
		t.Errorf("beklenen: context iptal edilmemiş olmalı, alınan: %v", err)
	}
}

// sleepWithContext fonksiyonunun iptal edilmiş context ile test eder
func TestSleepWithContext_Canceled(t *testing.T) {
	ctx, cancelFn := context.WithCancel(context.Background())
	cancelFn()

	err := sleepWithContext(ctx, 10*time.Second)
	if err == nil {
		t.Fatalf("beklenen: hata, alınan: hata yok")
	}

	if e, a := "context canceled", err.Error(); !strings.Contains(a, e) {
		t.Errorf("beklenen: %v hatası, alınan: %v", e, a)
	}
}
