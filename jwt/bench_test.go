package jwt_test

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"fmt"
	"testing"
	"time"

	cristalhq "github.com/cristalhq/jwt/v4"
	golang_jwt "github.com/golang-jwt/jwt"
	lestrrat_go "github.com/lestrrat-go/jwx"
	pascaldekloe "github.com/pascaldekloe/jwt"
)

var cristalhq_benchClaims = &struct {
	cristalhq.RegisteredClaims
}{
	RegisteredClaims: cristalhq.RegisteredClaims{
		Issuer:   "benchmark",
		IssuedAt: cristalhq.NewNumericDate(time.Now()),
	},
}

func Benchmark_cristalhq_EdDSA(b *testing.B) {
	signer, err := cristalhq.NewSignerEdDSA(testKeyEd25519Private)
	failIfErr(b, err)

	bui := cristalhq.NewBuilder(signer)
	b.Run("sign-"+cristalhq.EdDSA.String(), func(b *testing.B) {
		var tokenLen int
		for i := 0; i < b.N; i++ {
			token, err := bui.Build(cristalhq_benchClaims)
			failIfErr(b, err)
			tokenLen += len(token.Bytes())
		}
		b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
	})

	token, err := cristalhq.NewBuilder(signer).Build(cristalhq_benchClaims)
	failIfErr(b, err)

	verifier, err := cristalhq.NewVerifierEdDSA(testKeyEd25519Public)
	failIfErr(b, err)

	b.Run("check-"+cristalhq.EdDSA.String(), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := verifier.Verify(token)
			failIfErr(b, err)
		}
	})
}

func Benchmark_cristalhq_HMAC(b *testing.B) {
	algs := []cristalhq.Algorithm{cristalhq.HS256, cristalhq.HS384, cristalhq.HS512}

	for _, alg := range algs {
		signer, err := cristalhq.NewSignerHS(alg, keysHMAC)
		failIfErr(b, err)
		bui := cristalhq.NewBuilder(signer)
		b.Run("sign-"+alg.String(), func(b *testing.B) {
			var tokenLen int
			for i := 0; i < b.N; i++ {
				token, err := bui.Build(cristalhq_benchClaims)
				failIfErr(b, err)
				tokenLen += len(token.Bytes())
			}
			b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
		})
	}

	for _, alg := range algs {
		signer, err := cristalhq.NewSignerHS(alg, keysHMAC)
		failIfErr(b, err)
		token, err := cristalhq.NewBuilder(signer).Build(cristalhq_benchClaims)
		failIfErr(b, err)

		verifier, err := cristalhq.NewVerifierHS(alg, keysHMAC)
		failIfErr(b, err)
		b.Run("check-"+alg.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := verifier.Verify(token)
				failIfErr(b, err)
			}
		})
	}
}

func Benchmark_cristalhq_RSA(b *testing.B) {
	keys := []*rsa.PrivateKey{testKeyRSA1024, testKeyRSA2048, testKeyRSA4096}

	for _, key := range keys {
		signer, err := cristalhq.NewSignerRS(cristalhq.RS384, key)
		failIfErr(b, err)
		bui := cristalhq.NewBuilder(signer)
		b.Run(fmt.Sprintf("sign-%d-bit", key.Size()*8), func(b *testing.B) {
			var tokenLen int
			for i := 0; i < b.N; i++ {
				token, err := bui.Build(cristalhq_benchClaims)
				failIfErr(b, err)
				tokenLen += len(token.Bytes())
			}
			b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
		})
	}

	for _, key := range keys {
		signer, err := cristalhq.NewSignerRS(cristalhq.RS384, key)
		failIfErr(b, err)
		token, err := cristalhq.NewBuilder(signer).Build(cristalhq_benchClaims)
		failIfErr(b, err)

		verifier, err := cristalhq.NewVerifierRS(cristalhq.RS384, &key.PublicKey)
		failIfErr(b, err)
		b.Run(fmt.Sprintf("check-%d-bit", key.Size()*8), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := verifier.Verify(token)
				failIfErr(b, err)
			}
		})
	}
}

func Benchmark_cristalhq_ECDSA(b *testing.B) {
	tests := []struct {
		key *ecdsa.PrivateKey
		alg cristalhq.Algorithm
	}{
		{keysECDSA[0], cristalhq.ES256},
		{keysECDSA[1], cristalhq.ES384},
		{keysECDSA[2], cristalhq.ES512},
	}

	for _, test := range tests {
		signer, err := cristalhq.NewSignerES(test.alg, test.key)
		failIfErr(b, err)
		bui := cristalhq.NewBuilder(signer)
		b.Run("sign-"+test.alg.String(), func(b *testing.B) {
			var tokenLen int
			for i := 0; i < b.N; i++ {
				token, err := bui.Build(cristalhq_benchClaims)
				failIfErr(b, err)
				tokenLen += len(token.Bytes())
			}
			b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
		})
	}

	for _, test := range tests {
		signer, err := cristalhq.NewSignerES(test.alg, test.key)
		failIfErr(b, err)
		bui := cristalhq.NewBuilder(signer)
		token, err := bui.Build(cristalhq_benchClaims)
		failIfErr(b, err)

		verifier, err := cristalhq.NewVerifierES(test.alg, &test.key.PublicKey)
		failIfErr(b, err)
		b.Run("check-"+test.alg.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := verifier.Verify(token)
				failIfErr(b, err)
			}
		})
	}
}

var pascaldekloe_benchClaims = &pascaldekloe.Claims{
	Registered: pascaldekloe.Registered{
		Issuer: "benchmark",
		Issued: pascaldekloe.NewNumericTime(time.Now()),
	},
}

func Benchmark_pascaldekloe_EdDSA(b *testing.B) {
	b.Run("sign-"+pascaldekloe.EdDSA, func(b *testing.B) {
		var tokenLen int
		for i := 0; i < b.N; i++ {
			token, err := pascaldekloe_benchClaims.EdDSASign(testKeyEd25519Private)
			failIfErr(b, err)
			tokenLen += len(token)
		}
		b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
	})

	b.Run("check-"+pascaldekloe.EdDSA, func(b *testing.B) {
		token, err := pascaldekloe_benchClaims.EdDSASign(testKeyEd25519Private)
		failIfErr(b, err)

		for i := 0; i < b.N; i++ {
			_, err := pascaldekloe.EdDSACheck(token, testKeyEd25519Public)
			failIfErr(b, err)
		}
	})
}

func Benchmark_pascaldekloe_HMAC(b *testing.B) {
	algs := []string{pascaldekloe.HS256, pascaldekloe.HS384, pascaldekloe.HS512}

	for _, alg := range algs {
		b.Run("sign-"+alg, func(b *testing.B) {
			var tokenLen int
			for i := 0; i < b.N; i++ {
				token, err := pascaldekloe_benchClaims.HMACSign(alg, keysHMAC)
				failIfErr(b, err)
				tokenLen += len(token)
			}
			b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
		})
	}

	for _, alg := range algs {
		token, err := pascaldekloe_benchClaims.HMACSign(alg, keysHMAC)
		failIfErr(b, err)

		b.Run("check-"+alg, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := pascaldekloe.HMACCheck(token, keysHMAC)
				failIfErr(b, err)
			}
		})
	}
}

func Benchmark_pascaldekloe_RSA(b *testing.B) {
	keys := []*rsa.PrivateKey{testKeyRSA1024, testKeyRSA2048, testKeyRSA4096}

	for _, key := range keys {
		b.Run(fmt.Sprintf("sign-%d-bit", key.Size()*8), func(b *testing.B) {
			var tokenLen int
			for i := 0; i < b.N; i++ {
				token, err := pascaldekloe_benchClaims.RSASign(pascaldekloe.RS384, key)
				failIfErr(b, err)
				tokenLen += len(token)
			}
			b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
		})
	}

	for _, key := range keys {
		token, err := pascaldekloe_benchClaims.RSASign(pascaldekloe.RS384, key)
		failIfErr(b, err)

		b.Run(fmt.Sprintf("check-%d-bit", key.Size()*8), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := pascaldekloe.RSACheck(token, &key.PublicKey)
				failIfErr(b, err)
			}
		})
	}
}

func Benchmark_pascaldekloe_ECDSA(b *testing.B) {
	tests := []struct {
		key *ecdsa.PrivateKey
		alg string
	}{
		{testKeyEC256, pascaldekloe.ES256},
		{testKeyEC384, pascaldekloe.ES384},
		{testKeyEC521, pascaldekloe.ES512},
	}

	for _, test := range tests {
		b.Run("sign-"+test.alg, func(b *testing.B) {
			var tokenLen int
			for i := 0; i < b.N; i++ {
				token, err := pascaldekloe_benchClaims.ECDSASign(test.alg, test.key)
				failIfErr(b, err)
				tokenLen += len(token)
			}
			b.ReportMetric(float64(tokenLen)/float64(b.N), "B/token")
		})
	}

	for _, test := range tests {
		token, err := pascaldekloe_benchClaims.ECDSASign(test.alg, test.key)
		failIfErr(b, err)

		b.Run("check-"+test.alg, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := pascaldekloe.ECDSACheck(token, &test.key.PublicKey)
				failIfErr(b, err)
			}
		})
	}
}

var golang_jwt_benchClaims = &golang_jwt.StandardClaims{
	Issuer:   "benchmark",
	IssuedAt: time.Now().Unix(),
}

func Benchmark_golang_jwt_EdDSA(b *testing.B) {}

func Benchmark_golang_jwt_HMAC(b *testing.B) {
	algs := []golang_jwt.SigningMethod{
		golang_jwt.SigningMethodHS256,
		golang_jwt.SigningMethodHS384,
		golang_jwt.SigningMethodHS512,
	}

	for _, alg := range algs {
		token := golang_jwt.NewWithClaims(alg, golang_jwt_benchClaims)

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
		token := golang_jwt.NewWithClaims(alg, golang_jwt_benchClaims)
		tokenStr, err := token.SignedString(keysHMAC)
		failIfErr(b, err)

		b.Run("check-"+alg.Alg(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				token, err := golang_jwt.Parse(tokenStr, func(token *golang_jwt.Token) (interface{}, error) {
					return keysHMAC, nil
				})
				if err != nil || !token.Valid {
					b.Fatal(err)
				}
			}
		})
	}
}

func Benchmark_golang_jwt_RSA(b *testing.B) {
	for _, key := range keysRSA {
		token := golang_jwt.NewWithClaims(golang_jwt.SigningMethodRS384, golang_jwt_benchClaims)

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
		token := golang_jwt.NewWithClaims(golang_jwt.SigningMethodRS384, golang_jwt_benchClaims)
		tokenStr, err := token.SignedString(key)
		failIfErr(b, err)

		b.Run(fmt.Sprintf("check-%d-bit", key.Size()*8), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				token, err := golang_jwt.Parse(tokenStr, func(token *golang_jwt.Token) (interface{}, error) {
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

func Benchmark_golang_jwt_ECDSA(b *testing.B) {
	tests := []struct {
		key *ecdsa.PrivateKey
		alg *golang_jwt.SigningMethodECDSA
	}{
		{keysECDSA[0], golang_jwt.SigningMethodES256},
		{keysECDSA[1], golang_jwt.SigningMethodES384},
		{keysECDSA[2], golang_jwt.SigningMethodES512},
	}

	for _, test := range tests {
		token := golang_jwt.NewWithClaims(test.alg, golang_jwt_benchClaims)

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
		token := golang_jwt.NewWithClaims(test.alg, golang_jwt_benchClaims)
		tokenStr, err := token.SignedString(test.key)
		failIfErr(b, err)

		b.Run(fmt.Sprintf("check-%d-bit", test.key.Params().BitSize), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				token, err := golang_jwt.Parse(tokenStr, func(token *golang_jwt.Token) (interface{}, error) {
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

var lestrrat_go_benchClaims = lestrrat_go.DecoderSettings
