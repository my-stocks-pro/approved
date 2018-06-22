package client

import (
	"time"
	"strconv"
	"fmt"
)

const (
	token  = "v2/NjYxMWYtZjk4NjItNWI0ZDEtYjc3ODktNmIyYWEtZWU3NDUvMTEyNDM5ODQ1L2N1c3RvbWVyLzMveEZSaF85amREWHFaOFZuNjZfcjZVaUhWYzlYZmhVVl9kZURtU1EybHBSbDJ0eDFxUXh5czEzUXdRdkV6Wjl6dFBYcXpGbHk4RWhWZUpxU1U4TlFoWll3RjF1YkZKMFhsS0FndDQtSTY5Y0k0TV9nYy0yVFEzLXdzeC02TXdXMlloVFQyQ2kwZzBjZmtsVmNNVE5OUjZCSHdNY1kzSUQ4SW1CZHlwT1dYNFQ5enNIUkFPdUh3VElPRmxaZ214a003dDc1UHE0Tmpta0tIN3ZYM3g4V2xlUQ"
	baseUrl = "https://submit.shutterstock.com/api/catalog_manager/media_types/all/items?filter_type=date_uploaded&filter_value=%s-%s-01%s%s-%s-%s&page_number=%s&per_page=100&sort=popular"
	apiURL = "https://api.shutterstock.com/v2/images?"
)

type ImageFormatType struct {
	Display_name  string
	DPI           int
	File_size     int
	Format        string
	Height        int
	Is_licensable bool
	Width         int
}

type ImageLinksType struct {
	Height int
	URL    string
	Width  int
}

type DataImageType struct {
	Data []struct {
		ID         string
		Added_date string
		Aspect     float64
		Assets struct {
			Small_jpg      ImageFormatType
			Medium_jpg     ImageFormatType
			Huge_jpg       ImageFormatType
			Supersize_jpg  ImageFormatType
			Huge_tiff      ImageFormatType
			Supersize_tiff ImageFormatType
			Preview        ImageLinksType
			Small_thumb    ImageLinksType
			Large_thumb    ImageLinksType
			Huge_thumb     ImageLinksType
		}
		Categories []struct {
			ID   string
			Name string
		}
		Contributor struct {
			ID string
		}
		Description          string
		Image_type           string
		Is_adult             bool
		Is_illustration      bool
		Has_property_release bool
		Keywords             []string
		media_type           string
	}
}

type BaseRespType struct {
	ResponseHeader struct {
		SearchServiceQTime int
		Params struct {
			Q          []string
			Media_type []string
			Namespace  []string
			Start      []string
			Fq         []string
			Source     []string
			Rows       []string
			Sort_order []string
		}
	}
	Total int
	Data []struct {
		Media_id string
	}
}

type TypeChanResp struct {
	Res BaseRespType
	Done bool
}


type TypeRedis struct {
	RedisResp []string
}

type ApprovedType struct {
	TimeStamp time.Time
	CurrDate  string
	CurrYear  string
	CurrMonth string
	CurrDay   string
	BaseURL   string
	ApiURL    string
	NewURL    string
	Token     string
	Session   string
	Redis     TypeRedis
	ChanBaseResp chan TypeChanResp
}

func NewClient() *ApprovedType {
	currDate := time.Now()

	a := &ApprovedType{
		TimeStamp: currDate,
		CurrDate:  currDate.Format("2006-01-02"),
		CurrYear:  strconv.Itoa(currDate.Year()),
		CurrMonth: CheckDate(fmt.Sprintf("%d", currDate.Month())),
		CurrDay:   CheckDate(strconv.Itoa(currDate.Day())),
		BaseURL:   baseUrl,
		ApiURL:    apiURL,
		Session:   "s%3AFLsDQ0KkRmbbHJSFijJz_5VxQPCQI7Ol.t5LQWhFeOPA9qV2S0fqa6JBsFB0Rq%2BrxMDPc1URXyHE",
		Token:     token,
	}

	go a.workerRedis()

	return a
}
