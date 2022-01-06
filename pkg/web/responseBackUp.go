package web

type ResponseBackUp struct {
	Code        int    `json:"code"`
	Description string `json:"description,omitempty"`
	Error       string `json:"error,omitempty"`
}

func NewResponse(codeStatus int, description string, err string) ResponseBackUp {
	if codeStatus < 300 {
		return ResponseBackUp{codeStatus, description, ""}
	} else {
		return ResponseBackUp{codeStatus, "", err}
	}
}
