package entity

import "time"

type Payload struct {
	Account struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"account"`
	Project struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"project"`
	Trigger struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"trigger"`
	Error struct {
		ID             string    `json:"id"`
		ErrorID        string    `json:"errorId"`
		ExceptionClass string    `json:"exceptionClass"`
		Message        string    `json:"message"`
		Context        string    `json:"context"`
		AppVersion     string    `json:"appVersion"`
		ReleaseStage   string    `json:"releaseStage"`
		FirstReceived  time.Time `json:"firstReceived"`
		ReceivedAt     time.Time `json:"receivedAt"`
		UserID         string    `json:"userId"`
		AssignedUserID string    `json:"assignedUserId"`
		URL            string    `json:"url"`
		Severity       string    `json:"severity"`
		Status         string    `json:"status"`
		Unhandled      bool      `json:"unhandled"`
		App            struct {
			ID                   string   `json:"id"`
			Version              string   `json:"version"`
			VersionCode          string   `json:"versionCode"`
			BundleVersion        string   `json:"bundleVersion"`
			CodeBundleID         string   `json:"codeBundleId"`
			BuildUUID            string   `json:"buildUUID"`
			ReleaseStage         string   `json:"releaseStage"`
			Type                 string   `json:"type"`
			DsymUUIDs            []string `json:"dsymUUIDs"`
			Duration             int      `json:"duration"`
			DurationInForeground int      `json:"durationInForeground"`
			InForeground         bool     `json:"inForeground"`
			BinaryArch           string   `json:"binaryArch"`
			RunningOnRosetta     bool     `json:"runningOnRosetta"`
		} `json:"app"`
		Device struct {
			Hostname              string    `json:"hostname"`
			ID                    string    `json:"id"`
			Manufacturer          string    `json:"manufacturer"`
			Model                 string    `json:"model"`
			ModelNumber           string    `json:"modelNumber"`
			OsName                string    `json:"osName"`
			OsVersion             string    `json:"osVersion"`
			FreeMemory            int       `json:"freeMemory"`
			TotalMemory           int       `json:"totalMemory"`
			FreeDisk              int64     `json:"freeDisk"`
			BrowserName           string    `json:"browserName"`
			BrowserVersion        string    `json:"browserVersion"`
			Jailbroken            bool      `json:"jailbroken"`
			Orientation           string    `json:"orientation"`
			Locale                string    `json:"locale"`
			Charging              bool      `json:"charging"`
			BatteryLevel          float64   `json:"batteryLevel"`
			Time                  time.Time `json:"time"`
			Timezone              string    `json:"timezone"`
			MacCatalystiOSVersion string    `json:"macCatalystiOSVersion"`
		} `json:"device"`
		User struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"user"`
		StackTrace []struct {
			File       string `json:"file"`
			LineNumber string `json:"lineNumber"`
			Method     string `json:"method"`
			InProject  bool   `json:"inProject"`
			Code       struct {
				Num121 string `json:"121"`
				Num122 string `json:"122"`
				Num123 string `json:"123"`
				Num124 string `json:"124"`
				Num125 string `json:"125"`
				Num126 string `json:"126"`
				Num127 string `json:"127"`
			} `json:"code,omitempty"`
		} `json:"stackTrace"`
		Breadcrumbs []interface{} `json:"breadcrumbs"`
	} `json:"error"`
}
