package models

type AuthToken struct {
	Token string `json:"token"`
}

type User struct {
	UserName string `json:"username"`
	IsStaff  bool   `json:"is_staff"`
	IsActive bool   `json:"is_active"`
	Id       int    `json:"id"`
}

type FeaturePageFeatureId struct {
	FeatureId int `json:"feature_id"`
}

type AllFeatureRequests struct {
	FeatureRequests []FeatureRequest `json:"feature_requests"`
}

type FeatureRequest struct {
	Id                 int     `json:"id"`
	Name               string  `json:"name"`
	ProblemDescription string  `json:"problem_description"`
	Sample             string  `json:"sample"`
	PrdLink            string  `json:"prd_link"`
	EnggDocLink        string  `json:"engg_doc_link"`
	EpicLink           string  `json:"epic_link"`
	DesignLink         string  `json:"design_link"`
	OnboardingDoc      string  `json:"onboarding_doc"`
	CreatedAt          string  `json:"created_at"`
	ApprovedDate       string  `json:"approved_date"`
	CompletedDate      string  `json:"completed_date"`
	InternalReleseDate string  `json:"internal_release_date"`
	WillingToPay       float64 `json:"willing_to_pay"`
	Status             string  `json:"status"`
	ApprovedUserId     int     `json:"apporved_user_id"`
	BackendDevId       int     `json:"backend_dev_id"`
	FrontendDevId      int     `json:"frontend_dev_id"`
	PrdUserId          int     `json:"prd_user_id"`
	RequestedUserId    int     `json:"requested_user_id"`
	Type               string  `json:"type"`
	ProductId          int     `json:"product_id"`
	UpdatedAt          string  `json:"updated_at"`
	Urgency            string  `json:"urgency"`
}
type CreateFeature struct {
	Name               string  `json:"name"`
	ProblemDescription string  `json:"problem_description"`
	ProductId          int     `json:"product_id"`
	Type               string  `json:"type"`
	Urgency            string  `json:"urgency"`
	WillingToPay       float64 `json:"willing_to_pay"`
}
