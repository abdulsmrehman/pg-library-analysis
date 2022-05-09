package main

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v4"
    "os"
)

var conn *pgx.Conn

func main() {
    conn = getPGXConnection()

    dep := &DeploymentDetails{
        BaseModelDeployment:       BaseModelDeployment{},
        Id:                        101,
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

    fmt.Println(dep)
   /* err := addDeployment(dep)
    if err != nil {
        fmt.Errorf("Error:%v",err)
    }*/

    getDeployment()
}

func getPGXConnection() *pgx.Conn {
    userName := "usermkplc18"
    password := "QWERasdfmkplc18*&%qwerASDF"
    host := "marketplace-local-db.cffrhwfhhfof.us-east-2.rds.amazonaws.com"
    port := uint16(5432)
    dbname := "marketplace_deployments"
    var dbURL string = "postgres://postgres:mypassword@localhost:5432/postgres"
	var err error

    pgxConf,_ := pgx.ParseConfig(dbURL)
    pgxConf.Password = password
    pgxConf.User =userName
    pgxConf.Host = host
    pgxConf.Database = dbname
    pgxConf.Port = port

	conn, err = pgx.ConnectConfig(context.Background(),pgxConf)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func addDeployment(dep *DeploymentDetails) error{
    _, err := conn.Exec(context.Background(), "insert into deployment_details(id,uuid,productid,consumerid,orgid,source_org_id,url_hash_value,sddcid,resourcepoolid,vmname) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", dep.Id,dep.Uuid,dep.ProductID,dep.ConsumerID,dep.OrgID,dep.SourceOrgID,dep.SourceOrgDisplayName,dep.URLHashValue,dep.SddcID,dep.VMName)
    return err
}

func getDeployment(){
    input := &DeploymentDetails{
        Uuid: "sample-analysi-d-pgx",
    }

    var output DeploymentDetails

    err := conn.QueryRow(context.Background(), "select $1::json", input).Scan(&output)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Reuslt:",output)

}