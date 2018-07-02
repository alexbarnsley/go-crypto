// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package crypto

import (
    "./base58"
    "github.com/btcsuite/btcd/btcec"
    "golang.org/x/crypto/ripemd160"
)

type PublicKey struct {
    *btcec.PublicKey
    isCompressed bool
    network      *Network
}

func PublicKeyFromSecret(secret string, network *Network) (*PublicKey, error) {
    privateKey, err := PrivateKeyFromSecret(secret, network)

    if err != nil {
        return nil, err
    }

    return privateKey.PublicKey, nil
}

func PublicKeyFromBytes(bytes []byte, network *Network) (*PublicKey, error) {
    publicKey, err := btcec.ParsePubKey(bytes, btcec.S256())

    if err != nil {
        return nil, err
    }

    isCompressed := false

    if len(bytes) == btcec.PubKeyBytesLenCompressed {
        isCompressed = true
    }

    return &PublicKey{
        PublicKey:    publicKey,
        isCompressed: isCompressed,
        network:      network,
    }, nil
}

////////////////////////////////////////////////////////////////////////////////
// ADDRESS COMPUTATION /////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

func (publicKey *PublicKey) Address() (string, error) {
    ripeHashedBytes, err := publicKey.AddressBytes()

    if err != nil {
        return "", err
    }

    ripeHashedBytes = append(ripeHashedBytes, 0x0)
    copy(ripeHashedBytes[1:], ripeHashedBytes[:len(ripeHashedBytes)-1])
    ripeHashedBytes[0] = publicKey.network.Version

    return base58.Encode(ripeHashedBytes), nil
}

func (publicKey *PublicKey) Serialise() []byte {
    if publicKey.isCompressed {
        return publicKey.SerializeCompressed()
    }

    return publicKey.SerializeUncompressed()
}

func (publicKey *PublicKey) AddressBytes() ([]byte, error) {
    hash := ripemd160.New()
    _, err := hash.Write(publicKey.Serialise())

    if err != nil {
        return nil, err
    }

    return hash.Sum(nil), nil
}
