package transaction

import "time"

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
