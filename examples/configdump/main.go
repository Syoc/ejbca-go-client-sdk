package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/Keyfactor/ejbca-go-client-sdk/api/ejbca"
    "log"
    "strings"
)

func main() {
    configuration := ejbca.NewConfiguration()
    configuration.Debug = false

    apiClient, err := ejbca.NewAPIClient(configuration)
    if err != nil {
        bad, ok := err.(*ejbca.GenericOpenAPIError)
        if ok {
            fmt.Printf("Error: %s", string(bad.Body()))
        }
        log.Fatal(err)
    }

    exclude := []string{
        "ACMECONFIG:*;CRYPTOTOKEN:*",
    }

    include := []string{
        "CERTPROFILE:*",
    }

    _, httpResp, err := apiClient.V1ConfigdumpApi.GetJsonConfigdump(context.Background()).Ignoreerrors(true).Exclude(exclude).Include(include).Execute()
    if !strings.Contains(err.Error(), "json: cannot unmarshal object into Go value of type []string") {
        bad, ok := err.(*ejbca.GenericOpenAPIError)
        if ok {
            fmt.Printf("Error: %s", string(bad.Body()))
        }
        log.Fatal(err)
    }

    var configDump interface{} // Decode JSON body
    err = json.NewDecoder(httpResp.Body).Decode(&configDump)
    if err != nil {
        log.Fatal(err)
    }

    eepMap := configDump.(map[string]interface{})["end-entity-profiles"]
    for key, mapValue := range eepMap.(map[string]interface{}) {
        fmt.Printf("End entity profile %s has value:\n%v\n\n", key, mapValue)
    }
}
