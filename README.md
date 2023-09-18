# Asymmetric Encrpytion with AES-256

## Usage

Install the CLI

```sh
go install
```

Then: 
```sh
crypt -h
```

> **WARNING**: Decryption string contains '$' which need to be escaped in bash '\$' otherwise it will error.

## Build

Build flags for PROD

* Source: https://words.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/

```sh
CGO_ENABLED=0 go build -ldflags="-s -w" -buildvcs=false 
# upx can cause antivirus problems...
# upx --brute crypt
```

## Technical

### Encryption Standard

Using:

* AES-GCM (AEAD tbd)
  * Key created with PBKDF2

### Storage of Encrypted Message

Encrypted Message and Salts for Key and Message are being stored as follows:

```
<salt key>$<salt message>$<ciphered message>
```

