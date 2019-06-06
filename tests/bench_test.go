package test

import (
	tjsm2 "gmsm/sm2"
	"gmsm/sm3"
	opensm2 "transaction_service/openssl/sm2"

	"testing"
)

func BenchmarkOpenSSLVerify(b *testing.B) {
	p, err := tjsm2.GenerateKey()
	if err != nil {
		b.Error(err)
	}
	raw, err := tjsm2.MarshalSm2UnecryptedPrivateKey(p)
	if err != nil {
		b.Error(err)
	}
	priv, err := opensm2.ParsePKCS8UnecryptedPrivateKey(raw)
	if err != nil {
		b.Error(err)
	}
	hash := sm3.Sm3Sum([]byte("Hello"))
	r, s, err := opensm2.Sign(priv, hash)
	if err != nil {
		b.Error(err)
	}
	passed := tjsm2.Verify(&p.PublicKey, hash, r, s)
	if !passed {
		b.Error("Verify Failed")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		passed := opensm2.Verify(priv.PublicKey, hash, r, s)
		if !passed {
			b.Error("Verify Failed")
		}
	}
}

func BenchmarkOpenSSLVerifyParallel(b *testing.B) {
	p, err := tjsm2.GenerateKey()
	if err != nil {
		b.Error(err)
	}
	raw, err := tjsm2.MarshalSm2UnecryptedPrivateKey(p)
	if err != nil {
		b.Error(err)
	}
	priv, err := opensm2.ParsePKCS8UnecryptedPrivateKey(raw)
	if err != nil {
		b.Error(err)
	}
	hash := sm3.Sm3Sum([]byte("Hello"))
	r, s, err := opensm2.Sign(priv, hash)
	if err != nil {
		b.Error(err)
	}
	passed := tjsm2.Verify(&p.PublicKey, hash, r, s)
	if !passed {
		b.Error("Verify Failed")
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if !opensm2.Verify(priv.PublicKey, hash, r, s) {
				b.Error("Verify Failed")
			}
		}
	})
}

func BenchmarkTJVerify(b *testing.B) {
	p, err := opensm2.GenerateKey()
	if err != nil {
		b.Error(err)
	}
	raw, err := opensm2.MarshalSm2UnecryptedPrivateKey(p)
	if err != nil {
		b.Error(err)
	}
	priv, err := tjsm2.ParsePKCS8UnecryptedPrivateKey(raw)
	if err != nil {
		b.Error(err)
	}
	hash := sm3.Sm3Sum([]byte("Hello"))
	r, s, err := tjsm2.Sign(priv, hash)
	if err != nil {
		b.Error(err)
	}
	passed := opensm2.Verify(p.PublicKey, hash, r, s)
	if !passed {
		b.Error("Verify Failed")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		passed := tjsm2.Verify(&priv.PublicKey, hash, r, s)
		if !passed {
			b.Error("Verify Failed")
		}
	}

}

func BenchmarkTJVerifyParallel(b *testing.B) {
	p, err := opensm2.GenerateKey()
	if err != nil {
		b.Error(err)
	}
	raw, err := opensm2.MarshalSm2UnecryptedPrivateKey(p)
	if err != nil {
		b.Error(err)
	}
	priv, err := tjsm2.ParsePKCS8UnecryptedPrivateKey(raw)
	if err != nil {
		b.Error(err)
	}
	hash := sm3.Sm3Sum([]byte("Hello"))
	r, s, err := tjsm2.Sign(priv, hash)
	if err != nil {
		b.Error(err)
	}
	passed := opensm2.Verify(p.PublicKey, hash, r, s)
	if !passed {
		b.Error("Verify Failed")
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if !tjsm2.Verify(&priv.PublicKey, hash, r, s) {
				b.Error("Verify Failed")
			}
		}
	})
}

func BenchmarkOpenSSLSign(b *testing.B) {
	p, err := tjsm2.GenerateKey()
	if err != nil {
		b.Error(err)
	}
	raw, err := tjsm2.MarshalSm2UnecryptedPrivateKey(p)
	if err != nil {
		b.Error(err)
	}
	priv, err := opensm2.ParsePKCS8UnecryptedPrivateKey(raw)
	if err != nil {
		b.Error(err)
	}
	hash := sm3.Sm3Sum([]byte("Hello"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, s, err := opensm2.Sign(priv, hash)
		if err != nil || r == nil || s == nil {
			b.Error(err)
		}
	}
}

func BenchmarkOpenSSLSignParallel(b *testing.B) {
	p, err := tjsm2.GenerateKey()
	if err != nil {
		b.Error(err)
	}
	raw, err := tjsm2.MarshalSm2UnecryptedPrivateKey(p)
	if err != nil {
		b.Error(err)
	}
	priv, err := opensm2.ParsePKCS8UnecryptedPrivateKey(raw)
	if err != nil {
		b.Error(err)
	}
	hash := sm3.Sm3Sum([]byte("Hello"))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r, s, err := opensm2.Sign(priv, hash)
			if err != nil || r == nil || s == nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkTJSign(b *testing.B) {
	p, err := opensm2.GenerateKey()
	if err != nil {
		b.Error(err)
	}
	raw, err := opensm2.MarshalSm2UnecryptedPrivateKey(p)
	if err != nil {
		b.Error(err)
	}
	priv, err := tjsm2.ParsePKCS8UnecryptedPrivateKey(raw)
	if err != nil {
		b.Error(err)
	}
	hash := sm3.Sm3Sum([]byte("Hello"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, s, err := tjsm2.Sign(priv, hash)
		if err != nil || r == nil || s == nil {
			b.Error(err)
		}
	}
}

func BenchmarkTJSignParallel(b *testing.B) {
	p, err := opensm2.GenerateKey()
	if err != nil {
		b.Error(err)
	}
	raw, err := opensm2.MarshalSm2UnecryptedPrivateKey(p)
	if err != nil {
		b.Error(err)
	}
	priv, err := tjsm2.ParsePKCS8UnecryptedPrivateKey(raw)
	if err != nil {
		b.Error(err)
	}
	hash := sm3.Sm3Sum([]byte("Hello"))

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r, s, err := tjsm2.Sign(priv, hash)
			if err != nil || r == nil || s == nil {
				b.Error(err)
			}
		}
	})
}
