package ai_insight

import (
	"encoding/json"
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/db/model"
	"github.com/1337Bart/improve-yourself/internal/service"
	"gorm.io/gorm"
)

type AiInsight struct {
	SqlDb *gorm.DB
}

func NewAiInsightService(sqlDbConn *gorm.DB) *AiInsight {
	return &AiInsight{
		SqlDb: sqlDbConn,
	}
}

func (a *AiInsight) AddAiInsight(userID string, insight service.AiRecommendation) error {
	insightsJSON, err := json.Marshal(insight.Insights)
	if err != nil {
		return fmt.Errorf("error serializing insights: %v", err)
	}

	activitiesJSON, err := json.Marshal(insight.ProposedActivities)
	if err != nil {
		return fmt.Errorf("error serializing proposed activities: %v", err)
	}

	upsertRec := model.AiRecommendation{
		UserUUID:           userID,
		Insights:           string(insightsJSON),
		ProposedActivities: string(activitiesJSON),
	}

	tx := a.SqlDb.Save(&upsertRec)

	if tx.Error != nil {
		return fmt.Errorf("error upserting record: %v", tx.Error)
	}

	return nil
}

func (a *AiInsight) GetAiRecommendation(userID string) (service.AiRecommendation, error) {
	var aiRec model.AiRecommendation
	serviceAiRec := service.AiRecommendation{}

	tx := a.SqlDb.Where("uuid = ?", userID).First(&aiRec)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return serviceAiRec, fmt.Errorf("record not found")
		}
		return serviceAiRec, fmt.Errorf("error retrieving record: %v", tx.Error)
	}

	var insights []service.Insight
	if err := json.Unmarshal([]byte(aiRec.Insights), &insights); err != nil {
		return serviceAiRec, fmt.Errorf("error unmarshalling insights: %v", err)
	}

	var proposedActivities []service.ProposedActivity
	if err := json.Unmarshal([]byte(aiRec.ProposedActivities), &proposedActivities); err != nil {
		return serviceAiRec, fmt.Errorf("error unmarshalling proposed activities: %v", err)
	}

	serviceAiRec.Insights = insights
	serviceAiRec.ProposedActivities = proposedActivities

	return serviceAiRec, nil
}
