package main

import ejbca "github.com/Keyfactor/ejbca-go-client-sdk/api/ejbca"

func main() {
    configuration := ejbca.NewConfiguration()
    apiClient := ejbca.NewAPIClient(configuration)

}
