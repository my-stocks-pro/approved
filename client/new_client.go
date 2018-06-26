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

type TypeBase struct {
	ListIDS     []string
	Date        string
	TimeStamp   time.Time
	CurrDate    string
	CurrYear    string
	CurrMonth   string
	CurrDay     string
	CountIDS    string
	PerPage     int
	NewURL      string
}

type TypeRedis struct {
	Date    string
	ListIDS []string
}

type ApprovedType struct {
	Config    *config.TypeConfig
	Base      *TypeBase
	ChanBase  chan string
	ChanFull  chan string
	ChanWork  chan string
	ChanSlack chan *DataImageType
	ChanPSQL  chan *DataImageType
	Redis     TypeRedis
	POSTURL   string
}

func New() *ApprovedType {

	a := &ApprovedType{
		config.GetConfig(),
		&TypeBase{},
		make(chan string),
		make(chan string),
		make(chan string),
		make(chan *DataImageType),
		make(chan *DataImageType),
		TypeRedis{},
		"",
	}

	a.Base.PerPage = 1

	return a
}
