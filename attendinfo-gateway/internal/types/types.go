// Code generated by goctl. DO NOT EDIT.
package types

type CourseInfoRequest struct {
	University  string `json:"university" form:"university"`
	Course_id   string `json:"course_id" form:"course_id"`
	Course_main uint32 `json:"course_main" form:"course_main"`
}

type CourseInfoResponse struct {
	Status   uint32      `json:"status"`
	DataInfo interface{} `json:"data"`
	Total    uint32      `json:"total"`
	Error    string      `json:"error"`
}

type CounsellorInfoRequest struct {
	MajorInfoList []MajorInfo `json:"major_info"`
}

type MajorInfo struct {
	Major      string   `json:"major"`
	University string   `json:"university"`
	CourseMain []uint64 `json:"course_main"`
}

type CounsellorInfoResponse struct {
	Status     uint32      `json:"status"`
	MajorData  interface{} `json:"major_data"`
	Total1     uint32      `json:"total1"`
	CourseData interface{} `json:"course_data"`
	Total2     uint32      `json:"total2"`
	Message    string      `json:"message"`
}
