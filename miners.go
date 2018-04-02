package awesomeminer

import (
	"fmt"
	"time"
)

// MinersResult contains the result of an API call to Awesome Miner's
// "miners" endpoint. e.g. GET http://mypc:17790/api/miners
type MinersResult struct {
	TotalHashrate5s  string         `json:"totalHashrate5s"`
	HasManualAction  bool           `json:"hasManualActions"`
	ManualActionList []manualAction `json:"manualActionList"`
	GroupList        []group        `json:"groupList"`
}

type manualAction struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type group struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	MinerList []Miner `json:"minerList"`
}

// Miner contains the result of individual miners in MinersResult, and also the
// result of an API call to the "miner" endpoint. e.g. GET http://mypc:17790/api/miners/{id}
type Miner struct {
	ID             int          `json:"id"`
	Name           string       `json:"name"`
	Hostname       string       `json:"hostname"`
	GroupID        int          `json:"groupId"`
	Pool           string       `json:"pool"`
	Temperature    string       `json:"temperature"`
	StatusInfo     statusInfo   `json:"statusInfo"`
	ProgressInfo   progressInfo `json:"progressInfo"`
	SpeedInfo      speedInfo    `json:"speedInfo"`
	CoinInfo       coinInfo     `json:"coinInfo"`
	UpdatedUTC     time.Time    `json:"updatedUtc"`
	Updated        string       `json:"updated"`
	PoolList       []pool       `json:"poolList"`
	GPUList        []device     `json:"gpuList"`
	PGAList        []device     `json:"pgaList"`
	AsicList       []device     `json:"asicList"`
	HasPool        bool         `json:"hasPool"`
	HasGPU         bool         `json:"hasGpu"`
	HasPGA         bool         `json:"hasPga"`
	HasAsic        bool         `json:"hasAsic"`
	CanReboot      bool         `json:"canReboot"`
	CanStop        bool         `json:"canStop"`
	CanRestart     bool         `json:"canRestart"`
	CanStart       bool         `json:"canStart"`
	CanPool        bool         `json:"canPool"`
	HasValidStatus bool         `json:"hasValidStatus"`
}

type statusInfo struct {
	StatusDisplay string `json:"statusDisplay"`
	StatusLine3   string `json:"statusLine3"`
}

type progressInfo struct {
	Line1 string `json:"line1"`
	Line2 string `json:"line2"`
	Line3 string `json:"line3"`
}

type speedInfo struct {
	LogInterval int    `json:"logInterval"`
	HashRate    string `json:"hashRate"`
	AvgHashrate string `json:"avgHashrate"`
	WorkUtility string `json:"workUtility"`
}

type coinInfo struct {
	DisplayName        string  `json:"displayName"`
	RevenuePerDay      string  `json:"revenuePerDay"`
	RevenuePerDayValue float64 `json:"revenuePerDayValue"`
	RevenuePerMonth    string  `json:"revenuePerMonth"`
	// The following is undocumented
	ProfitPerDayValue float64 `json:"profitPerDayValue"`
	ProfitPerDay      string  `json:"profitPerDay"`
	ProfitPerMonth    string  `json:"profitPerMonth"`
}

type pool struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	StatusInfo     statusInfo     `json:"statusInfo"`
	AdditionalInfo additionalInfo `json:"additionalInfo"`
	PriorityInfo   priorityInfo   `json:"priorityInfo"`
	ProgressInfo   progressInfo   `json:"progressInfo"`
	CoinName       string         `json:"coinName"`
	MinerID        int            `json:"minerID"`
	MinerName      string         `json:"minerName"`
	CanRemove      bool           `json:"canRemove"`
	CanDisable     bool           `json:"canDisable"`
	CanEnable      bool           `json:"canEnable"`
	CanPrioritize  bool           `json:"canPrioritize"`
}

type additionalInfo struct {
	DisplayURL string `json:"displayUrl"`
	Worker     string `json:"worker"`
}

type priorityInfo struct {
	Priority int `json:"priority"`
	Quote    int `json:"quota"`
}

type device struct {
	Name         string       `json:"name"`
	StatusInfo   statusInfo   `json:"statusInfo"`
	DeviceInfo   deviceInfo   `json:"deviceInfo"`
	ProgressInfo progressInfo `json:"progressInfo"`
	SpeedInfo    speedInfo    `json:"speedInfo"`
}

type deviceInfo struct {
	DeviceType     string `json:"deviceType"`
	GPUActivity    int    `json:"gpuActivity"`
	Intensity      string `json:"intensity"`
	Name           string `json:"name"`
	GPUClock       int    `json:"gpuClock"`
	GPUMemoryClock int    `json:"gpuMemoryClock"`
	GPUVoltage     string `json:"gpuVoltage"`
	GPUPowertune   int    `json:"gpuPowertune"`
	FanSpeed       int    `json:"fanSpeed"`
	FanPercent     int    `json:"fanPercent"`
	Temperature    int    `json:"temperature"`
}

// GetMiners performs a call to the miners endpoint of Awesome Miner and returns
// the result.
func (c *Client) GetMiners() (*MinersResult, error) {
	result := &MinersResult{}
	_, err := c.doGetRequest("miners", result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetMiner performs a call to the miner endpoint of Awesome Miner and returns
// the result.
func (c *Client) GetMiner(id int) (*Miner, error) {
	result := &Miner{}
	_, err := c.doGetRequest(fmt.Sprintf("miners/%d", id), result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
