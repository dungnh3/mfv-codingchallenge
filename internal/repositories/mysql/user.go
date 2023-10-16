package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/dungnh3/mfv-codingchallenge/internal/models"
	"gorm.io/gorm/clause"
)

func (q *Queries) UpsertUser(ctx context.Context, user *models.User) error {
	return q.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"name":       user.Name,
			"status":     user.Status,
			"updated_at": time.Now(),
		}),
	}).Create(user.User).Error
}

func (q *Queries) UpsertUserAccount(ctx context.Context, acc *models.UserAccount) error {
	return q.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"name":       acc.Name,
			"status":     acc.Status,
			"updated_at": time.Now(),
		}),
	}).Create(acc.UserAccount).Error
}

const getUserQuery = `
WITH account_ids AS (SELECT id FROM user_accounts WHERE user_id = @id)
SELECT u.*,
       (SELECT JSON_ARRAYAGG(id) FROM account_ids) AS account_ids
FROM users u
WHERE u.id = @id;
`

func (q *Queries) GetUser(ctx context.Context, userId int64) (*models.User, error) {
	var user models.User
	return &user, q.db.WithContext(ctx).Raw(getUserQuery, sql.Named("id", userId)).Take(&user).Error
}

func (q *Queries) ListAccounts(ctx context.Context, userId int64) ([]*models.UserAccount, error) {
	var accounts []*models.UserAccount
	return accounts, q.db.WithContext(ctx).Model(&models.UserAccount{}).
		Where("user_id = ?", userId).Find(&accounts).Error
}

func (q *Queries) GetAccount(ctx context.Context, accountId int64) (*models.UserAccount, error) {
	var account models.UserAccount
	return &account, q.db.WithContext(ctx).Model(&models.UserAccount{}).
		Where("id = ?", accountId).Take(&account).Error
}
