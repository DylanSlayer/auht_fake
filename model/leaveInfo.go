package model

type LeaveInfo struct {
	StudentName string `json:"name"`
	StudentId   string `json:"id"`
	Academy     string `json:"academy"`
	Telephone   string `json:"telephone"`
	OutTime     string `json:"outTime"`
	ReturnTime  string `json:"returnTime"`
	OutAddress  string `json:"outAddress"`
	Reason      string `json:"outReason"`
	ApplyTime   string `json:"applyTime"`
	Approver    string `json:"approver"`
	ApproveTime string `json:"approveTime"`
}
