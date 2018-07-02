// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package crypto

import (
    "./base58"
)

func AddressFromSecret(secret string, network *Network) (string, error) {
    privateKey, err := PrivateKeyFromSecret(secret, network)

    if err != nil {
        return "", err
    }

    address, err := privateKey.Address()

    if err != nil {
        return "", err
    }

    return address, nil
}

func AddressToBytes(address string) ([]byte, error) {
    bytes, err := base58.Decode(address)

    if err != nil {
        return nil, err
    }

    return bytes[1:], nil
}

func ValidateAddress(address string) (bool, error) {
    return true, nil
}
