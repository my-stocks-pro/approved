package new

import (
	"time"
	"fmt"
	"strconv"
	"github.com/my-stocks-pro/approved/utils"
)

const (
	token  = "v2/NjYxMWYtZjk4NjItNWI0ZDEtYjc3ODktNmIyYWEtZWU3NDUvMTEyNDM5ODQ1L2N1c3RvbWVyLzMveEZSaF85amREWHFaOF" +
		"ZuNjZfcjZVaUhWYzlYZmhVVl9kZURtU1EybHBSbDJ0eDFxUXh5czEzUXdRdkV6Wjl6dFBYcXpGbHk4RWhWZUpxU1U4TlFoWll3RjF1YkZKMFhs" +
		"S0FndDQtSTY5Y0k0TV9nYy0yVFEzLXdzeC02TXdXMlloVFQyQ2kwZzBjZmtsVmNNVE5OUjZCSHdNY1kzSUQ4SW1CZHlwT1dYNFQ5enNIUkFPdUh" +
		"3VElPRmxaZ214a003dDc1UHE0Tmpta0tIN3ZYM3g4V2xlUQ"

	baseUrl = "https://submit.shutterstock.com/api/catalog_manager/media_types/all/items?filter_type=date_uploaded&filter_value=%s-%s-01%s%s-%s-%s&page_number=%s&per_page=100&sort=popular"
)

type ApprovedType struct {
	CurrDate string
	BaseUrl  string
	Token    string
	//BaseResponse BaseResponseType
}

var Approved *ApprovedType

func (a *ApprovedType) New() {
	currDate := time.Now()
	currDateStr := currDate.Format("2006-01-02")

	year := strconv.Itoa(currDate.Year())
	month := utils.CheckDate(fmt.Sprintf("%d", currDate.Month()))
	day := utils.CheckDate(strconv.Itoa(currDate.Day()))

	url := fmt.Sprintf(baseUrl, year, month, "%20", year, month, day)

	a = &ApprovedType{
		CurrDate: currDateStr,
		BaseUrl: url,
	}
}
