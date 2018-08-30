package keyvault_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"github.com/gobuffalo/envy"
)

func ExampleBaseClient_GetSecretsComplete() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := keyvault.New()
	client.Authorizer, err = auth.NewAuthorizerFromEnvironment()
	if err != nil {
		fmt.Println("Error! ", err)
		return
	}

	vaultURL := os.Getenv("KEYVAULT_BASE_URL")

	res, err := client.GetSecretsComplete(ctx, vaultURL, nil)
	if err != nil {
		fmt.Println("Error! ", err)
		return
	}

	for ; res.NotDone(); res.Next() {
		val := res.Value()
		fmt.Println(*val.ID)
	}

	// Output:
	// https://marstr-kvexamples.vault.azure.net/secrets/anotherthing
	// https://marstr-kvexamples.vault.azure.net/secrets/mything
}

func ExampleBaseClient_GetSecret() {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	vaultURL := os.Getenv("KEYVAULT_BASE_URL")
	secretName, secretVersion := os.Getenv("KEYVAULT_SECRET_NAME"), os.Getenv("KEYVAULT_SECRET_VERSION")

	client := keyvault.New()
	client.Authorizer, err = auth.NewAuthorizerFromEnvironment()
	if err != nil {
		fmt.Println("Error! ", err)
		return
	}

	res, err := client.GetSecret(ctx, vaultURL, secretName, secretVersion)
	if err != nil {
		fmt.Println("Error! ", err)
		return
	}

	fmt.Println(*res.Value)
	// Output:
	// foobar
}

func TestMain(m *testing.M) {
	envy.Load("./.env")
	os.Exit(m.Run())
}
