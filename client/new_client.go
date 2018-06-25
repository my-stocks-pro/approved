package client

import (
	"time"
	"github.com/my-stocks-pro/approved/config"
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

//type TypeChanResp struct {
//	Res  *BaseRespType
//	Done bool
//}

type TypeRedis struct {
	ListIDS   []string
	Date      string
	TimeStamp time.Time
	CurrDate  string
	CurrYear  string
	CurrMonth string
	CurrDay   string
}

type ApprovedType struct {
	Config   *config.TypeConfig
	Redis    *TypeRedis
	BaseChan chan *BaseRespType
}

func New() *ApprovedType {

	 a := &ApprovedType{
		config.GetConfig(),
		&TypeRedis{},
		make(chan *BaseRespType),
	}

	go a.BaseWorker()

	return a
}
