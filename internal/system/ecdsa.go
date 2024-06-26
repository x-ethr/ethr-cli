package system

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"io"
	"log/slog"
)

func ECDSA(ctx context.Context, private, public io.Writer) error {
	var encode = func(pkey *ecdsa.PrivateKey) ([]byte, error) {
		pkeyder, err := x509.MarshalECPrivateKey(pkey)
		if err != nil {
			return nil, err
		}

		pkeyblock := &pem.Block{
			Type:  "EC PRIVATE KEY",
			Bytes: pkeyder,
		}

		pkeypem := pem.EncodeToMemory(pkeyblock)
		return pkeypem, nil
	}

	// Generate a private key using the P256 curve
	pkey, e := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if e != nil {
		slog.ErrorContext(ctx, "Error Generating ECDSA Key", slog.String("error", e.Error()))
		return e
	}

	// Encode the private key to PEM format
	pkeypem, e := encode(pkey)
	if e != nil {
		slog.ErrorContext(ctx, "Error Encoding Private Key to PEM", slog.String("error", e.Error()))
		return e
	}

	size, e := private.Write(pkeypem)
	if e != nil {
		slog.ErrorContext(ctx, "Error Writing Private Key Contents to Buffer", slog.String("error", e.Error()), slog.Int("size", size))
		return e
	}

	pubkey := pkey.Public()

	// Marshal the public key into ASN.1 DER encoded form
	x509encodedpublic, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return err
	}

	// Create a PEM block with the encoded public key
	publickeypem := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509encodedpublic,
	}

	if e := pem.Encode(public, &publickeypem); e != nil {
		return e
	}

	slog.DebugContext(ctx, "Successfully Generated and Wrote Private & Public ECDSA Keys to Buffers")

	return nil
}
