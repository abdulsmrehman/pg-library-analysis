package main

type BaseModelDeployment struct {
	CreatedOn int32 `sql:"created_on"`
	UpdatedOn int32 `sql:"updated_on"`
	DeletedOn int32 `sql:"deleted_on"`
}

type DeploymentStatus string

type PoweredStatus string

type SddcDetails struct {
	BaseModelDeployment
	Id        int32 `sql:",pk"`
	SddcID    string
	Latitude  float64
	Longitude float64
}

type ContainerSubscription struct {
	BaseModelDeployment
	ID                  int32 `sql:",pk"`
	DeploymentType      string
	AppVersion          string
	ChartVersion        string
	DeploymentDetailsID int32
}

type BaseModel struct {
	ID        int32
	CreatedOn int32 `sql:"createdon"`
	UpdatedOn int32 `sql:"updatedon"`
	DeletedOn int32 `sql:"deletedon"`
}

type AutoSubscription struct {
	BaseModel
	Id                        int32 `sql:",pk"`
	ProductId                 string
	CatalogContentLibId       string
	Platform                  string
	DeploymentPlatformOrgID   string
	MaxVersions               int32
	UserID                    string
	OrgID                     string
	DeploymentPlatformOrgName string
	Versions                  []string `sql:"-"`
}

type SubscriptionLog struct {
	AuthData         string
	SubscriptionID   string
	SubscriptionStep string
	TimeStamp        int64
	Platform         string
	SubscriptionLog  string
}

type DeploymentDetails struct {
	BaseModelDeployment
	Id                   int32  `bun:",pk"`
	Uuid                 string `bun:"uuid"`
	ProductID            string `bun:"productid"`
	ConsumerID           string `bun:"consumerid"`
	OrgID                string `bun:"orgid"`
	SourceOrgID          string `bun:"source_org_id"`
	SourceOrgDisplayName string `bun:"-"`
	URLHashValue         string `bun:"url_hash_value"`

	SddcID         string `bun:"sddcid"`
	ResourcePoolID string `bun:"resourcepoolid"`
	FolderID       string `bun:"folderid"`
	VMName         string `bun:"vmname"`
	DataStoreID    string `bun:"datastoreid"`

	Status      DeploymentStatus `bun:"status"`
	StatusText  string           `bun:"statustext"`
	PowerStatus PoweredStatus    `bun:"powerstatus"`
	StartedOn   int32            `bun:"startedon"`
	DeployedOn  int32            `bun:"deployedon"`
	PoweredOn   int32            `bun:"poweredon"`
	PowerON     bool             `bun:"poweron"`

	PublisherName           string `bun:"-"`
	PublisherOrgDisplayName string `bun:"-"`
	PublisherID             string `bun:"publisherid"`
	ProductName             string `bun:"-"`
	ProductLogo             string `bun:"-"`
	CloudAssemblyOrgID      string
	CloudAssemblyProjectID  string
	SddcDetails             *SddcDetails `bun:"-"`

	// additional fields
	PublisherOrgID            string `bun:"publisher_org_id"`
	ConsumerOrgID             string `bun:"consumer_org_id"`
	DeploymentPlatformOrgID   string `bun:"deployment_platform_org_id"`
	PublisherOrgName          string `bun:"publisher_org_name"`
	ConsumerOrgName           string `bun:"consumer_org_name"`
	DeploymentPlatformOrgName string `bun:"deployment_platform_org_name"`
	ProductVersion            string `bun:"product_version"`
	FileID                    string `bun:"fileid"`
	EulaAccepted              bool   `bun:",default:false,notnull"`
	DeploymentPlatform        string `bun:"deployment_platform"`
	SubscriptionURL           string `bun:"subscription_url"`
	PlatformRepo              string `bun:"platform_repo"`
	ContentCatalogID          string            `bun:"platform_repo_id"`
	UpdatesAvailable          bool              `bun:"-"`
	IsAlreadySubscribed       bool              `bun:"-"`
	IsAutoSubscribed          bool              `bun:"is_auto_subscribed,notnull"`
	SubscriptionID            string            `bun:"subscription_id"`
	ArtifactURL               string            `bun:"artifact_url"` // TODO: to be removed
	ServiceCatalogID          string            `bun:"-"`
	ServiceCatalogSolutionID  string            `bun:"-"`
	Eula                      string            `bun:"-"`
}
