package go_shopware_admin_sdk

import "fmt"

type ExtensionList []*ExtensionDetail

func (l ExtensionList) GetByName(name string) *ExtensionDetail {
	for _, detail := range l {
		if detail.Name == name {
			return detail
		}
	}

	return nil
}

func (l ExtensionList) FilterByUpdateable() ExtensionList {
	newList := make(ExtensionList, 0)

	for _, detail := range l {
		if detail.IsUpdateAble() {
			newList = append(newList, detail)
		}
	}

	return newList
}

type ExtensionDetail struct {
	Extensions             []interface{} `json:"extensions"`
	Id                     interface{}   `json:"id"`
	LocalId                string        `json:"localId"`
	Name                   string        `json:"name"`
	Label                  string        `json:"label"`
	Description            string        `json:"description"`
	ShortDescription       interface{}   `json:"shortDescription"`
	ProducerName           string        `json:"producerName"`
	License                string        `json:"license"`
	Version                string        `json:"version"`
	LatestVersion          string        `json:"latestVersion"`
	Languages              []interface{} `json:"languages"`
	Rating                 interface{}   `json:"rating"`
	NumberOfRatings        int           `json:"numberOfRatings"`
	Variants               []interface{} `json:"variants"`
	Faq                    []interface{} `json:"faq"`
	Binaries               []interface{} `json:"binaries"`
	Images                 []interface{} `json:"images"`
	Icon                   interface{}   `json:"icon"`
	IconRaw                *string       `json:"iconRaw"`
	Categories             []interface{} `json:"categories"`
	Permissions            interface{}   `json:"permissions"`
	Active                 bool          `json:"active"`
	Type                   string        `json:"type"`
	IsTheme                bool          `json:"isTheme"`
	Configurable           bool          `json:"configurable"`
	PrivacyPolicyExtension interface{}   `json:"privacyPolicyExtension"`
	StoreLicense           interface{}   `json:"storeLicense"`
	StoreExtension         interface{}   `json:"storeExtension"`
	InstalledAt            *struct {
		Date         string `json:"date"`
		TimezoneType int    `json:"timezone_type"`
		Timezone     string `json:"timezone"`
	} `json:"installedAt"`
	UpdatedAt    interface{}   `json:"updatedAt"`
	Notices      []interface{} `json:"notices"`
	Source       string        `json:"source"`
	UpdateSource string        `json:"updateSource"`
}

func (e ExtensionDetail) Status() string {
	var text string

	switch {
	case e.Source == "store":
		text = "can be downloaded from store"
	case e.Active:
		text = "installed, activated"
	case e.InstalledAt != nil:
		text = "installed, not activated"
	default:
		text = "not installed, not activated"
	}

	if e.IsUpdateAble() {
		text = fmt.Sprintf("%s, update available to %s", text, e.LatestVersion)
	}

	return text
}

func (e ExtensionDetail) IsPlugin() bool {
	return e.Type == "plugin"
}

func (e ExtensionDetail) IsUpdateAble() bool {
	return len(e.LatestVersion) > 0 && e.LatestVersion != e.Version
}
