package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"roar-dashboard-backend/database"
	"roar-dashboard-backend/models"
	"time"
)

type Repo struct {
	DB *sql.DB
}

// func (repo *Repo) GetUserIdFromToken(token string) (int, error) {
// 	// query := "SELECT id from auth_utoken where key=(?)"
// 	// ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
// 	// defer cancelfunc()
// 	var userId int
// 	if err := repo.DB.QueryRow("SELECT user_id FROM auth_token WHERE `key`=?", token).Scan(&userId); err != nil {
// 		if err == sql.ErrNoRows {
// 			return 0, err
// 		}
// 		return 0, err
// 	}
// 	return userId, nil
// }

// func GetUserDetailsFromId(userId int) (models.User, error) {
// 	var user models.User
// 	if err := NewRepo().DB.QueryRow("SELECT id, username, is_staff, is_active from auth_user where `id` = ?",
// 		userId).Scan(&user.Id, &user.UserName, &user.IsStaff, &user.IsActive); err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println(err.Error())
// 			return models.User{}, err
// 		}
// 		fmt.Println(err.Error())
// 		return models.User{}, err

// 	}
// 	return user, nil
// }

func (repo *Repo) GetFeatureDetails(feature_id int) (feature_details models.FeatureRequest, err error) {
	// var userId int
	fmt.Println(feature_id)
	if err := repo.DB.QueryRow("SELECT * FROM feedback_management_system_feature WHERE `id`=?", feature_id).Scan(&feature_details.Id, &feature_details.Name, &feature_details.ProblemDescription, &feature_details.Sample, &feature_details.PrdLink, &feature_details.EnggDocLink, &feature_details.EpicLink, &feature_details.DesignLink, &feature_details.OnboardingDoc, &feature_details.CreatedAt, &feature_details.ApprovedDate, &feature_details.CompletedDate, &feature_details.InternalReleseDate, &feature_details.WillingToPay, &feature_details.Status, &feature_details.ApprovedUserId, &feature_details.BackendDevId, &feature_details.FrontendDevId, &feature_details.PrdUserId, &feature_details.RequestedUserId, &feature_details.Type, &feature_details.ProductId, &feature_details.UpdatedAt, &feature_details.Urgency); err != nil {
		if err == sql.ErrNoRows {
			return feature_details, err
		}
		return feature_details, err
	}
	return feature_details, nil
}

func (repo *Repo) GetAllFeatureDetails() ([]models.FeatureRequest, error) {
	var allFeatures []models.FeatureRequest
	var singleFeature models.FeatureRequest
	rows, err := repo.DB.Query("SELECT * FROM feedback_management_system_feature")
	if err != nil {
		log.Fatalf("Database query failed because %s", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&singleFeature.Id, &singleFeature.Name, &singleFeature.ProblemDescription, &singleFeature.Sample, &singleFeature.PrdLink, &singleFeature.EnggDocLink, &singleFeature.EpicLink, &singleFeature.DesignLink, &singleFeature.OnboardingDoc, &singleFeature.CreatedAt, &singleFeature.ApprovedDate, &singleFeature.CompletedDate, &singleFeature.InternalReleseDate, &singleFeature.WillingToPay, &singleFeature.Status, &singleFeature.ApprovedUserId, &singleFeature.BackendDevId, &singleFeature.FrontendDevId, &singleFeature.PrdUserId, &singleFeature.RequestedUserId, &singleFeature.Type, &singleFeature.ProductId, &singleFeature.UpdatedAt, &singleFeature.Urgency)
		if err != nil {
			log.Fatalf("Failed to retrieve row because %s", err)
			return nil, err
		}
		allFeatures = append(allFeatures, singleFeature)
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("Error encountered while iterating over rows: %s", err)
		return nil, err
	}
	return allFeatures, nil
}

func (repo *Repo) CreateFeature(singleFeature models.FeatureRequest) error {
	now := time.Now()
	cmdStr := "INSERT INTO feedback_management_system_feature (id, name, problem_description, sample, prd_link,engg_doc_link, epic_link, design_link, onboarding_doc, created_at, approved_date, completed_date, internal_release_date, willing_to_pay, status, approved_user_id, backend_dev_id, frontend_dev_id, prd_user_id, requested_user_id, type, product_id, updated_at, urgency) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err := repo.DB.ExecContext(context.TODO(), cmdStr, singleFeature.Id, singleFeature.Name, singleFeature.ProblemDescription, singleFeature.Sample, singleFeature.PrdLink, singleFeature.EnggDocLink, singleFeature.EpicLink, singleFeature.DesignLink, singleFeature.OnboardingDoc, now.String(), singleFeature.ApprovedDate,
		singleFeature.CompletedDate, singleFeature.InternalReleseDate, singleFeature.WillingToPay, singleFeature.Status, singleFeature.ApprovedUserId, singleFeature.BackendDevId, singleFeature.FrontendDevId, singleFeature.PrdUserId, singleFeature.RequestedUserId, singleFeature.Type, singleFeature.ProductId, now.String(), singleFeature.Urgency)
	if err != nil {
		log.Fatalf("impossible insert: %s", err)
		return err
	}
	return nil
}

func NewRepo() Repo {
	return Repo{
		DB: database.InitDb(),
	}
}
