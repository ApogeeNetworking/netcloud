package netcloud

import "time"

// AuthParams ...
type AuthParams struct {
	CpApiID   string
	CpApiKey  string
	EcmApiID  string
	EcmApiKey string
}

type PagingParams struct {
	Offset int
	// A Max of 500 records will be returned for all calls
	// When not specified, by default the Limit is 20
	Limit int
}

// MetaResult used for paging on returned data set for API Calls
type MetaResult struct {
	Limit int `json:"limit"`
	// A URL of the next Dataset to Retrieve From with PagingParams qs
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
}

// Router An object representing a Device
type Router struct {
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	Account           string      `json:"account"`
	ActualFirmware    string      `json:"actual_firmware"`
	AssetID           interface{} `json:"asset_id"`
	ConfigStatus      string      `json:"config_status"`
	ConfigManager     string      `json:"configuration_manager"`
	CreatedAt         time.Time   `json:"created_at"`
	Custom1           string      `json:"custom1"`
	Custom2           interface{} `json:"custom2"`
	Description       string      `json:"description"`
	DeviceType        string      `json:"device_type"`
	FullProductName   string      `json:"full_product_name"`
	Group             string      `json:"group"`
	Ipv4Address       string      `json:"ipv4_address"`
	Lans              string      `json:"lans"`
	LastKnownLocation interface{} `json:"last_known_location"`
	Locality          string      `json:"locality"`
	Mac               string      `json:"mac"`
	OverlayNetBinding string      `json:"overlay_network_binding"`
	Product           string      `json:"product"`
	RebootRequired    bool        `json:"reboot_required"`
	ResourceURL       string      `json:"resource_url"`
	SerialNumber      string      `json:"serial_number"`
	State             string      `json:"state"`
	StateUpdatedAt    time.Time   `json:"state_updated_at"`
	TargetFirmware    string      `json:"target_firmware"`
	UpdatedAt         time.Time   `json:"updated_at"`
	UpgradePending    bool        `json:"upgrade_pending"`
}

// RouterReq ...
type RouterReq struct {
	Data []Router   `json:"data"`
	Meta MetaResult `json:"meta"`
}

// RouterReqParams ...
type RouterReqParams struct {
	ID       string
	Name     string
	Ipv4Addr string
	MacAddr  string
}
