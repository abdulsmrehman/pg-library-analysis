package main

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"

	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var db *bun.DB

func mai1n() {

	db = postgresSetUp()

	//insertDep()
	getSpecificData()
}

func postgresSetUp() *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr("marketplace-local-db.cffrhwfhhfof.us-east-2.rds.amazonaws.com:5432"),
		pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		pgdriver.WithUser("usermkplc18"),
		pgdriver.WithPassword("QWERasdfmkplc18*&%qwerASDF"),
		pgdriver.WithDatabase("marketplace_deployments"),
	))

	db := bun.NewDB(sqldb, pgdialect.New())

	return db

	//db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

}

func insertDep() {
	ctx := context.Background()
	dep := &DeploymentDetails{
		BaseModelDeployment:       BaseModelDeployment{},
		Id:                        10,
		Uuid:                      "sample-analysi-d",
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

	res, err := db.NewInsert().Model(dep).Exec(ctx)
	if err != nil {
		fmt.Errorf("failed with error:%v", err)
		return
	}

	fmt.Println("Result:", res)
}


func getSpecificData() {
    ctx := context.Background()
    dep := &DeploymentDetails{}
    err := db.NewSelect().Model(dep).Where("uuid = ?","sample-analysi-d").Scan(ctx)
    if err != nil {
        fmt.Errorf("failed with error:%v", err)
        return
    }

    //fmt.Println(dep)
}