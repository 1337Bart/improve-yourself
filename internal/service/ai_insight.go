package service

type AiRecommendation struct {
	Insights           []Insight          `json:"insights"`
	ProposedActivities []ProposedActivity `json:"proposed_activities"`
}

type Insight struct {
	Description string `json:"description"`
	Details     string `json:"details"`
}

type ProposedActivity struct {
	Description string `json:"description"`
	Details     string `json:"details"`
}
