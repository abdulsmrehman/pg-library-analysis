package main

import (
	"testing"
)

func BenchmarkInsertPgx(b *testing.B) {
	conn = getPGXConnection()
	dep := &DeploymentDetails{

		BaseModelDeployment:       BaseModelDeployment{},
		Id:                        201,
		Uuid:                      "sample-analysi-d-pgx",
		ProductID:                 "1afe6cfb-e788-4122-a872-32fdb419a7a9",
		ConsumerID:                "",
		OrgID:                     "",
		SourceOrgID:               "",
		SourceOrgDisplayName:      "",
		URLHashValue:              "",
		SddcID:                    "",
		ResourcePoolID:            "",
		FolderID:                  "",
		VMName:                    "",
		DataStoreID:               "",
		Status:                    "",
		StatusText:                "",
		PowerStatus:               "",
		StartedOn:                 0,
		DeployedOn:                0,
		PoweredOn:                 0,
		PowerON:                   false,
		PublisherName:             "",
		PublisherOrgDisplayName:   "",
		PublisherID:               "",
		ProductName:               "",
		ProductLogo:               "",
		CloudAssemblyOrgID:        "",
		CloudAssemblyProjectID:    "",
		SddcDetails:               nil,
		PublisherOrgID:            "",
		ConsumerOrgID:             "",
		DeploymentPlatformOrgID:   "",
		PublisherOrgName:          "",
		ConsumerOrgName:           "",
		DeploymentPlatformOrgName: "",
		ProductVersion:            "",
		FileID:                    "",
		EulaAccepted:              false,
		DeploymentPlatform:        "",
		SubscriptionURL:           "",
		PlatformRepo:              "",
		ContentCatalogID:          "",
		UpdatesAvailable:          false,
		IsAlreadySubscribed:       false,
		IsAutoSubscribed:          false,
		SubscriptionID:            "",
		ArtifactURL:               "",
		ServiceCatalogID:          "",
		ServiceCatalogSolutionID:  "",
		Eula:                      "",
	}
	for i := 0; i < b.N; i++ {
		//primeNumbers(num)
		addDeployment(dep)
	}
}

func BenchmarkGetPgx(b *testing.B) {
	db = postgresSetUp()
	for i := 0; i < b.N; i++ {
		//primeNumbers(num)
		getSpecificData()
	}
}
