package transaction

type service struct {
	repository Repository
}

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTrnsactionsInput) ([]Transaction, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTrnsactionsInput) ([]Transaction, error) {
	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}