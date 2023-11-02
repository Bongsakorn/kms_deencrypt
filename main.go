package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	crypt "gopkg.in/Bongsakorn/cryptototamus.v2"
)

var decryptWithProfile, decryptWithText, encryptWithProfile, encryptWithText, encryptKeyname string

func main() {
	// Subcommands
	decryptCommand := flag.NewFlagSet("decrypt", flag.ExitOnError)
	encryptCommand := flag.NewFlagSet("encrypt", flag.ExitOnError)

	// decrypt command subcommand flag pointers
	// Adding a new choice for --metric of 'substring' and a new --substring flag
	decryptCommand.StringVar(&decryptWithProfile, "profile", "default", "aws credential profile. (optional)")
	decryptCommand.StringVar(&decryptWithText, "text", "", "Decrypted text. (Required)")

	encryptCommand.StringVar(&encryptWithProfile, "profile", "default", "aws credential profile. (optional)")
	encryptCommand.StringVar(&encryptWithText, "text", "", "Text would like to be encrypted. (Required)")
	encryptCommand.StringVar(&encryptKeyname, "keyname", "", "Key name that use for encrypt. (Required)")

	if len(os.Args) < 2 {
		fmt.Println("decrypt or encrypt subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "decrypt":
		decryptCommand.Parse(os.Args[2:])
	case "encrypt":
		encryptCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if decryptCommand.Parsed() {
		if decryptWithText == "" {
			decryptCommand.PrintDefaults()
			os.Exit(1)
		}

		sess := session.Must(session.NewSessionWithOptions(session.Options{
			Profile: decryptWithProfile,
			// Provide SDK Config options, such as Region.
			Config: aws.Config{
				Region: aws.String("ap-southeast-1"),
			},
		}))
		kmsClient := kms.New(sess, aws.NewConfig())
		detailDecrypted, err := crypt.Decrypt(kmsClient, decryptWithText)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("=====================")
		fmt.Println(string(detailDecrypted))
	}

	if encryptCommand.Parsed() {
		if encryptWithText == "" {
			encryptCommand.PrintDefaults()
			os.Exit(1)
		}

		if encryptKeyname == "" {
			encryptCommand.PrintDefaults()
			os.Exit(1)
		}

		sess := session.Must(session.NewSessionWithOptions(session.Options{
			Profile: encryptWithProfile,
			// Provide SDK Config options, such as Region.
			Config: aws.Config{
				Region: aws.String("ap-southeast-1"),
			},
		}))
		kmsClient := kms.New(sess, aws.NewConfig())
		encryptedRequestBody, err := crypt.Encrypt(kmsClient, encryptKeyname, []byte(encryptWithText))
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("=====================")
		fmt.Println(string(encryptedRequestBody))
	}
}
