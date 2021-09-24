package transaction

import (
	"time"
)

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatterTransaction := CampaignTransactionFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}

	return formatterTransaction
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var formatterTransactions []CampaignTransactionFormatter

	for _, transaction := range transactions {
		formatterTransaction := CampaignTransactionFormatter{
			ID:        transaction.ID,
			Name:      transaction.User.Name,
			Amount:    transaction.Amount,
			CreatedAt: transaction.CreatedAt,
		}

		formatterTransactions = append(formatterTransactions, formatterTransaction)
	}

	return formatterTransactions
}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Status    string            `json:"status"`
	Amount    int               `json:"amount"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatterCampaign := CampaignFormatter{
		Name:     transaction.Campaign.Name,
		ImageURL: "",
	}

	if len(transaction.Campaign.CampaignImages) > 0 {
		formatterCampaign.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	formatterTransaction := UserTransactionFormatter{
		ID:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
		Campaign:  formatterCampaign,
	}

	return formatterTransaction
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}
	var formatterTransactions []UserTransactionFormatter
	for _, transaction := range transactions {
		formatterTransaction := FormatUserTransaction(transaction)
		formatterTransactions = append(formatterTransactions, formatterTransaction)
	}

	return formatterTransactions
}
