package main

type OpenAPI struct {
	Data struct {
		ListTotalCount int `json:"list_total_count"`
		Row            []struct {
			DinrCal    string `json:"dinr_cal"`
			Lunc       string `json:"lunc"`
			SumCal     string `json:"sum_cal"`
			Adspcfd    string `json:"adspcfd"`
			AdspcfdCal string `json:"adspcfd_cal"`
			Dates      string `json:"dates"`
			LuncCal    string `json:"lunc_cal"`
			Brst       string `json:"brst"`
			Dinr       string `json:"dinr"`
			BrstCal    string `json:"brst_cal"`
		} `json:"row"`
	} `json:"DS_TB_MNDT_DATEBYMLSVC_7369"`
}
