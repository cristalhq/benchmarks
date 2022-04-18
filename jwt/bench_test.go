package jwt_test

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"fmt"
	"testing"
	"time"

	cristalhq_jwt "github.com/cristalhq/jwt/v4"
	golang_jwt_jwt "github.com/golang-jwt/jwt"
	lestrrat_go_jwx "github.com/lestrrat-go/jwx"
	pascaldekloe_jwt "github.com/pascaldekloe/jwt"
)

var cristalhq_jwt_benchClaims = &struct {
	cristalhq_jwt.RegisteredClaims
}{
	RegisteredClaims: cristalhq_jwt.RegisteredClaims{
		Issuer:   "benchmark",
		IssuedAt: cristalhq_jwt.NewNumericDate(time.Now()),
	},
}

func Benchmark_cristalhq_jwt_EdDSA(b *testing.B) {
	signer, err := cristalhq_jwt.NewSignerEdDSA(testKeyEd25519Private)
	failIfErr(b, err)

	bui := cristalhq_jwt.NewBuilder(signer)
	b.Run("sign-"+cristalhq_jwt.EdDSA.String(), func(b *testing.B) {
		var tokenLen int
		for i := 0; i < b.N; i++ {
			token, err := bui.Build(cristalhq_jwt_benchClaims)
			failIfErr(b, err)
			tokenLen += len(token.Bytes())
		}
		b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
	})

	token, err := cristalhq_jwt.NewBuilder(signer).Build(cristalhq_jwt_benchClaims)
	failIfErr(b, err)

	verifier, err := cristalhq_jwt.NewVerifierEdDSA(testKeyEd25519Public)
	failIfErr(b, err)

	b.Run("check-"+cristalhq_jwt.EdDSA.String(), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := verifier.Verify(token)
			failIfErr(b, err)
		}
	})
}

func Benchmark_cristalhq_jwt_HMAC(b *testing.B) {
	algs := []cristalhq_jwt.Algorithm{cristalhq_jwt.HS256, cristalhq_jwt.HS384, cristalhq_jwt.HS512}

	for _, alg := range algs {
		signer, err := cristalhq_jwt.NewSignerHS(alg, keysHMAC)
		failIfErr(b, err)
		bui := cristalhq_jwt.NewBuilder(signer)
		b.Run("sign-"+alg.String(), func(b *testing.B) {
			var tokenLen int
			for i := 0; i < b.N; i++ {
				token, err := bui.Build(cristalhq_jwt_benchClaims)
				failIfErr(b, err)
				tokenLen += len(token.Bytes())
			}
			b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
		})
	}

	for _, alg := range algs {
		signer, err := cristalhq_jwt.NewSignerHS(alg, keysHMAC)
		failIfErr(b, err)
		token, err := cristalhq_jwt.NewBuilder(signer).Build(cristalhq_jwt_benchClaims)
		failIfErr(b, err)

		verifier, err := cristalhq_jwt.NewVerifierHS(alg, keysHMAC)
		failIfErr(b, err)
		b.Run("check-"+alg.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := verifier.Verify(token)
				failIfErr(b, err)
			}
		})
	}
}

func Benchmark_cristalhq_jwt_RSA(b *testing.B) {
	keys := []*rsa.PrivateKey{testKeyRSA1024, testKeyRSA2048, testKeyRSA4096}

	for _, key := range keys {
		signer, err := cristalhq_jwt.NewSignerRS(cristalhq_jwt.RS384, key)
		failIfErr(b, err)
		bui := cristalhq_jwt.NewBuilder(signer)
		b.Run(fmt.Sprintf("sign-%d-bit", key.Size()*8), func(b *testing.B) {
			var tokenLen int
			for i := 0; i < b.N; i++ {
				token, err := bui.Build(cristalhq_jwt_benchClaims)
				failIfErr(b, err)
				tokenLen += len(token.Bytes())
			}
			b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
		})
	}

	for _, key := range keys {
		signer, err := cristalhq_jwt.NewSignerRS(cristalhq_jwt.RS384, key)
		failIfErr(b, err)
		token, err := cristalhq_jwt.NewBuilder(signer).Build(cristalhq_jwt_benchClaims)
		failIfErr(b, err)

		verifier, err := cristalhq_jwt.NewVerifierRS(cristalhq_jwt.RS384, &key.PublicKey)
		failIfErr(b, err)
		b.Run(fmt.Sprintf("check-%d-bit", key.Size()*8), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := verifier.Verify(token)
				failIfErr(b, err)
			}
		})
	}
}

func Benchmark_cristalhq_jwt_ECDSA(b *testing.B) {
	tests := []struct {
		key *ecdsa.PrivateKey
		alg cristalhq_jwt.Algorithm
	}{
		{keysECDSA[0], cristalhq_jwt.ES256},
		{keysECDSA[1], cristalhq_jwt.ES384},
		{keysECDSA[2], cristalhq_jwt.ES512},
	}

	for _, test := range tests {
		signer, err := cristalhq_jwt.NewSignerES(test.alg, test.key)
		failIfErr(b, err)
		bui := cristalhq_jwt.NewBuilder(signer)
		b.Run("sign-"+test.alg.String(), func(b *testing.B) {
			var tokenLen int
			for i := 0; i < b.N; i++ {
				token, err := bui.Build(cristalhq_jwt_benchClaims)
				failIfErr(b, err)
				tokenLen += len(token.Bytes())
			}
			b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
		})
	}

	for _, test := range tests {
		signer, err := cristalhq_jwt.NewSignerES(test.alg, test.key)
		failIfErr(b, err)
		bui := cristalhq_jwt.NewBuilder(signer)
		token, err := bui.Build(cristalhq_jwt_benchClaims)
		failIfErr(b, err)

		verifier, err := cristalhq_jwt.NewVerifierES(test.alg, &test.key.PublicKey)
		failIfErr(b, err)
		b.Run("check-"+test.alg.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := verifier.Verify(token)
				failIfErr(b, err)
			}
		})
	}
}

var pascaldekloe_jwt_benchClaims = &pascaldekloe_jwt.Claims{
	Registered: pascaldekloe_jwt.Registered{
		Issuer: "benchmark",
		Issued: pascaldekloe_jwt.NewNumericTime(time.Now()),
	},
}

func Benchmark_pascaldekloe_jwt_EdDSA(b *testing.B) {
	b.Run("sign-"+pascaldekloe_jwt.EdDSA, func(b *testing.B) {
		var tokenLen int
		for i := 0; i < b.N; i++ {
			token, err := pascaldekloe_jwt_benchClaims.EdDSASign(testKeyEd25519Private)
			failIfErr(b, err)
			tokenLen += len(token)
		}
		b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
	})

	b.Run("check-"+pascaldekloe_jwt.EdDSA, func(b *testing.B) {
		token, err := pascaldekloe_jwt_benchClaims.EdDSASign(testKeyEd25519Private)
		failIfErr(b, err)

		for i := 0; i < b.N; i++ {
			_, err := pascaldekloe_jwt.EdDSACheck(token, testKeyEd25519Public)
			failIfErr(b, err)
		}
	})
}

func Benchmark_pascaldekloe_jwt_HMAC(b *testing.B) {
	algs := []string{pascaldekloe_jwt.HS256, pascaldekloe_jwt.HS384, pascaldekloe_jwt.HS512}

	for _, alg := range algs {
		b.Run("sign-"+alg, func(b *testing.B) {
			var tokenLen int
			for i := 0; i < b.N; i++ {
				token, err := pascaldekloe_jwt_benchClaims.HMACSign(alg, keysHMAC)
				failIfErr(b, err)
				tokenLen += len(token)
			}
			b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
		})
	}

	for _, alg := range algs {
		token, err := pascaldekloe_jwt_benchClaims.HMACSign(alg, keysHMAC)
		failIfErr(b, err)

		b.Run("check-"+alg, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := pascaldekloe_jwt.HMACCheck(token, keysHMAC)
				failIfErr(b, err)
			}
		})
	}
}

func Benchmark_pascaldekloe_jwt_RSA(b *testing.B) {
	keys := []*rsa.PrivateKey{testKeyRSA1024, testKeyRSA2048, testKeyRSA4096}

	for _, key := range keys {
		b.Run(fmt.Sprintf("sign-%d-bit", key.Size()*8), func(b *testing.B) {
			var tokenLen int
			for i := 0; i < b.N; i++ {
				token, err := pascaldekloe_jwt_benchClaims.RSASign(pascaldekloe_jwt.RS384, key)
				failIfErr(b, err)
				tokenLen += len(token)
			}
			b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
		})
	}

	for _, key := range keys {
		token, err := pascaldekloe_jwt_benchClaims.RSASign(pascaldekloe_jwt.RS384, key)
		failIfErr(b, err)

		b.Run(fmt.Sprintf("check-%d-bit", key.Size()*8), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := pascaldekloe_jwt.RSACheck(token, &key.PublicKey)
				failIfErr(b, err)
			}
		})
	}
}

func Benchmark_pascaldekloe_jwt_ECDSA(b *testing.B) {
	tests := []struct {
		key *ecdsa.PrivateKey
		alg string
	}{
		{testKeyEC256, pascaldekloe_jwt.ES256},
		{testKeyEC384, pascaldekloe_jwt.ES384},
		{testKeyEC521, pascaldekloe_jwt.ES512},
	}

	for _, test := range tests {
		b.Run("sign-"+test.alg, func(b *testing.B) {
			var tokenLen int
			for i := 0; i < b.N; i++ {
				token, err := pascaldekloe_jwt_benchClaims.ECDSASign(test.alg, test.key)
				failIfErr(b, err)
				tokenLen += len(token)
			}
			b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
		})
	}

	for _, test := range tests {
		token, err := pascaldekloe_jwt_benchClaims.ECDSASign(test.alg, test.key)
		failIfErr(b, err)

		b.Run("check-"+test.alg, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := pascaldekloe_jwt.ECDSACheck(token, &test.key.PublicKey)
				failIfErr(b, err)
			}
		})
	}
}

var golang_jwt_jwt_benchClaims = &golang_jwt_jwt.StandardClaims{
	Issuer:   "benchmark",
	IssuedAt: time.Now().Unix(),
}

func Benchmark_golang_jwt_jwt_EdDSA(b *testing.B) {}

func Benchmark_golang_jwt_jwt_HMAC(b *testing.B) {
	algs := []golang_jwt_jwt.SigningMethod{
		golang_jwt_jwt.SigningMethodHS256,
		golang_jwt_jwt.SigningMethodHS384,
		golang_jwt_jwt.SigningMethodHS512,
	}

	for _, alg := range algs {
		token := golang_jwt_jwt.NewWithClaims(alg, golang_jwt_jwt_benchClaims)

		b.Run("sign-"+alg.Alg(), func(b *testing.B) {
			var tokenLen int
			for i := 0; i < b.N; i++ {
				tokenStr, err := token.SignedString(keysHMAC)
				failIfErr(b, err)
				tokenLen += len(tokenStr)
			}
			b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
		})
	}

	for _, alg := range algs {
		token := golang_jwt_jwt.NewWithClaims(alg, golang_jwt_jwt_benchClaims)
		tokenStr, err := token.SignedString(keysHMAC)
		failIfErr(b, err)

		b.Run("check-"+alg.Alg(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				token, err := golang_jwt_jwt.Parse(tokenStr, func(token *golang_jwt_jwt.Token) (interface{}, error) {
					return keysHMAC, nil
				})
				if err != nil || !token.Valid {
					b.Fatal(err)
				}
			}
		})
	}
}

func Benchmark_golang_jwt_jwt_RSA(b *testing.B) {
	for _, key := range keysRSA {
		token := golang_jwt_jwt.NewWithClaims(golang_jwt_jwt.SigningMethodRS384, golang_jwt_jwt_benchClaims)

		b.Run(fmt.Sprintf("sign-%d-bit", key.Size()*8), func(b *testing.B) {
			var tokenLen int
			for i := 0; i < b.N; i++ {
				tokenStr, err := token.SignedString(key)
				failIfErr(b, err)
				tokenLen += len(tokenStr)
			}
			b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
		})
	}

	for _, key := range keysRSA {
		token := golang_jwt_jwt.NewWithClaims(golang_jwt_jwt.SigningMethodRS384, golang_jwt_jwt_benchClaims)
		tokenStr, err := token.SignedString(key)
		failIfErr(b, err)

		b.Run(fmt.Sprintf("check-%d-bit", key.Size()*8), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				token, err := golang_jwt_jwt.Parse(tokenStr, func(token *golang_jwt_jwt.Token) (interface{}, error) {
					return &key.PublicKey, nil
				})
				failIfErr(b, err)
				if !token.Valid {
					b.Fatal("not valiad")
				}
			}
		})
	}
}

func Benchmark_golang_jwt_jwt_ECDSA(b *testing.B) {
	tests := []struct {
		key *ecdsa.PrivateKey
		alg *golang_jwt_jwt.SigningMethodECDSA
	}{
		{keysECDSA[0], golang_jwt_jwt.SigningMethodES256},
		{keysECDSA[1], golang_jwt_jwt.SigningMethodES384},
		{keysECDSA[2], golang_jwt_jwt.SigningMethodES512},
	}

	for _, test := range tests {
		token := golang_jwt_jwt.NewWithClaims(test.alg, golang_jwt_jwt_benchClaims)

		b.Run(fmt.Sprintf("sign-%d-bit", test.key.Params().BitSize), func(b *testing.B) {
			var tokenLen int
			for i := 0; i < b.N; i++ {
				tokenStr, err := token.SignedString(test.key)
				failIfErr(b, err)
				tokenLen += len(tokenStr)
			}
			b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
		})
	}

	for _, test := range tests {
		token := golang_jwt_jwt.NewWithClaims(test.alg, golang_jwt_jwt_benchClaims)
		tokenStr, err := token.SignedString(test.key)
		failIfErr(b, err)

		b.Run(fmt.Sprintf("check-%d-bit", test.key.Params().BitSize), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				token, err := golang_jwt_jwt.Parse(tokenStr, func(token *golang_jwt_jwt.Token) (interface{}, error) {
					return &test.key.PublicKey, nil
				})
				failIfErr(b, err)
				if !token.Valid {
					b.Fatal("not valiad")
				}
			}
		})
	}
}

var lestrrat_go_jwx_benchClaims = lestrrat_go_jwx.DecoderSettings
