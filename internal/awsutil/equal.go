package awsutil

import (
	"reflect"
)

// DeepEqual, iki değerin reflect.DeepEqual gibi derinlemesine eşit olup olmadığını döndürür.
// Buna ek olarak, bu yöntem mümkünse giriş değerlerini de referansını çözerek
// bir parametre pointer iken diğerinin olmaması durumunda DeepEqual işleminin başarısız olmasını önler.
//
// DeepEqual, giriş parametrelerinin iç içe geçmiş değerlerinin dolaylılığını gerçekleştirmez.
func DeepEqual(a, b interface{}) bool {
	ra := reflect.Indirect(reflect.ValueOf(a))
	rb := reflect.Indirect(reflect.ValueOf(b))

	if raValid, rbValid := ra.IsValid(), rb.IsValid(); !raValid && !rbValid {
		// Eğer elemanlar her ikisi de nil ise ve aynı türde iseler eşittirler
		// Farklı türde iseler eşit değillerdir
		return reflect.TypeOf(a) == reflect.TypeOf(b)
	} else if raValid != rbValid {
		// Her iki değer de eşit olmak için geçerli olmalıdır
		return false
	}

	// Stringler için özel durum, çünkü tiplenmiş enumerasyonlar string alias'larıdır
	// fakat derinlemesine eşit değildirler.
	if ra.Kind() == reflect.String && rb.Kind() == reflect.String {
		return ra.String() == rb.String()
	}

	return reflect.DeepEqual(ra.Interface(), rb.Interface())
}
