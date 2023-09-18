package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func encryptAction(cCtx *cli.Context) error {
	msg := cCtx.String("message")
	password := cCtx.String("password")

	enc, err := aesEncrypt(msg, password, Size256bit, 10000, 12)
	if err != nil {
		return err
	}
	
	fmt.Println("ENC", enc)

	return nil
}

func decryptAction(cCtx *cli.Context) error {
	msg := cCtx.String("message")
	password := cCtx.String("password")

	enc, err := aesDecrypt(msg, password, Size256bit, 10000)
	if err != nil {
		return err
	}
	
	fmt.Println("ENC", enc)

	return nil
}


func main() {
	app := &cli.App{
		Name: "crypt",
		Usage: "Encrypt / Decrypt Messages with AES-GCM that uses a derived key from a password with PBKDF2.",
		Commands: []*cli.Command{
			{
				Name: "encrypt",
				Aliases: []string{"e", "enc"},
				Usage: "Encrypt a message",
				Action: encryptAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "message",
						Aliases: []string{"m", "msg"},
						Usage: "The message to encrypt",
						Required: true,
					},
					&cli.StringFlag{
						Name: "password",
						Aliases: []string{"p"},
						Usage: "Password for encryption",
						Required: true,
					},
				},
			},
			{
				Name: "decrypt",
				Aliases: []string{"d", "dec"},
				Usage: "Decrypt a message",
				Action: decryptAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "message",
						Aliases: []string{"m", "msg"},
						Usage: "The message to decrypt",
						Required: true,
					},
					&cli.StringFlag{
						Name: "password",
						Aliases: []string{"p"},
						Usage: "Password for decryption",
						Required: true,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
