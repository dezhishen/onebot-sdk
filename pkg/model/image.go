package model

type Image struct {
	Size     int32  `json:"size"`
	Filename string `json:"filename"`
	Url      string `json:"url"`
}

func (a *Image) ToGRPC() *ImageGRPC {
	return &ImageGRPC{
		Size:     a.Size,
		Filename: a.Filename,
		Url:      a.Url,
	}
}

func (a *ImageGRPC) ToStruct() *Image {
	return &Image{
		Size:     a.Size,
		Filename: a.Filename,
		Url:      a.Url,
	}
}

type ImageResult struct {
	Data    *Image `json:"data"`
	Retcode int64  `json:"retcode"`
	Status  string `json:"status"`
	Msg     string `json:"msg"`
	Wording string `json:"wording"`
}

func (a *ImageResult) ToGRPC() *ImageResultGRPC {
	result := &ImageResultGRPC{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToGRPC()
	}
	return result
}

func (a *ImageResultGRPC) ToStruct() *ImageResult {
	result := &ImageResult{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToStruct()
	}
	return result
}

type Coordinate struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

func (a *Coordinate) ToGRPC() *CoordinateGRPC {
	return &CoordinateGRPC{
		X: a.X,
		Y: a.Y,
	}
}

func (a *CoordinateGRPC) ToStruct() *Coordinate {
	return &Coordinate{
		X: a.X,
		Y: a.Y,
	}
}

type TextDetection struct {
	Text        string        `json:"text"`
	Confidence  int32         `json:"confidence"`
	Coordinates []*Coordinate `json:"coordinates"`
}

func CoordinateArray2CoordinateGRPCArray(a []*Coordinate) []*CoordinateGRPC {
	var result []*CoordinateGRPC
	for _, v := range a {
		result = append(result, v.ToGRPC())
	}
	return result
}

func (a *TextDetection) ToGRPC() *TextDetectionGRPC {
	return &TextDetectionGRPC{
		Text:        a.Text,
		Confidence:  a.Confidence,
		Coordinates: CoordinateArray2CoordinateGRPCArray(a.Coordinates),
	}
}

func CoordinateGRPCArray2CoordinateArray(a []*CoordinateGRPC) []*Coordinate {
	var result []*Coordinate
	for _, v := range a {
		result = append(result, v.ToStruct())
	}
	return result
}

func (a *TextDetectionGRPC) ToStruct() *TextDetection {
	return &TextDetection{
		Text:        a.Text,
		Confidence:  a.Confidence,
		Coordinates: CoordinateGRPCArray2CoordinateArray(a.Coordinates),
	}
}

type OcrImage struct {
	Texts    []*TextDetection `json:"texts"`
	Language string           `json:"language"`
}

func (a *OcrImage) ToGRPC() *OcrImageGRPC {
	var result []*TextDetectionGRPC
	for _, v := range a.Texts {
		result = append(result, v.ToGRPC())
	}
	return &OcrImageGRPC{
		Texts:    result,
		Language: a.Language,
	}
}

func (a *OcrImageGRPC) ToStruct() *OcrImage {
	var result []*TextDetection
	for _, v := range a.Texts {
		result = append(result, v.ToStruct())
	}
	return &OcrImage{
		Texts:    result,
		Language: a.Language,
	}
}

type OcrImageResult struct {
	Data    *OcrImage `json:"data"`
	Retcode int64     `json:"retcode"`
	Status  string    `json:"status"`
	Msg     string    `json:"msg"`
	Wording string    `json:"wording"`
}

func (a *OcrImageResult) ToGRPC() *OcrImageResultGRPC {
	result := &OcrImageResultGRPC{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToGRPC()
	}
	return result
}

func (a *OcrImageResultGRPC) ToStruct() *OcrImageResult {
	result := &OcrImageResult{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToStruct()
	}
	return result
}
