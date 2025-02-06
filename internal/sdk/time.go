package sdk

import (
	"context"
	"time"
)

func init() {
	NowTime = time.Now
	Sleep = time.Sleep
	SleepWithContext = sleepWithContext
}

// NowTime, mevcut zamanı almak için kullanılan bir değerdir. Bu değer, testlerde mevcut zamanı taklit etmek için değiştirilebilir.
var NowTime func() time.Time

// Sleep, belirli bir süre uyumak için kullanılan bir değerdir. Bu değer, testlerde uyuma süresini taklit etmek için değiştirilebilir.
var Sleep func(time.Duration)

// SleepWithContext, zamanlayıcı süresi dolana kadar veya context iptal edilene kadar bekler. Hangisi önce gerçekleşirse. Eğer context iptal edilirse, Context'in hatası döndürülür.
// Bu değer, testlerde uyuma süresini taklit etmek için değiştirilebilir.
var SleepWithContext func(context.Context, time.Duration) error

// sleepWithContext, zamanlayıcı süresi dolana kadar veya context iptal edilene kadar bekler. Hangisi önce gerçekleşirse. Eğer context iptal edilirse, Context'in hatası döndürülür.
func sleepWithContext(ctx context.Context, dur time.Duration) error {
	t := time.NewTimer(dur)
	defer t.Stop()

	select {
	case <-t.C:
		break
	case <-ctx.Done():
		return ctx.Err()
	}

	return nil
}

// noOpSleepWithContext hiçbir şey yapmaz, hemen döner.
func noOpSleepWithContext(context.Context, time.Duration) error {
	return nil
}

// noOpSleep hiçbir şey yapmaz, hemen döner.
func noOpSleep(time.Duration) {}

// TestingUseNopSleep, SDK genelinde uyumayı devre dışı bırakmak için kullanılan bir yardımcı işlevdir.
func TestingUseNopSleep() func() {
	SleepWithContext = noOpSleepWithContext
	Sleep = noOpSleep

	return func() {
		SleepWithContext = sleepWithContext
		Sleep = time.Sleep
	}
}

// TestingUseReferenceTime, test amaçları için SDK genelinde belirli bir referans zamanı döndüren zaman işlevini değiştirmek için kullanılan bir yardımcı işlevdir.
func TestingUseReferenceTime(referenceTime time.Time) func() {
	NowTime = func() time.Time {
		return referenceTime
	}
	return func() {
		NowTime = time.Now
	}
}
