package services

type IngestionService struct {
	Producer KafkaProducer
}

func NewIngestionService(producer KafkaProducer) *IngestionService{
	return &IngestionService{
		Producer:  producer,
	}
}

type KafkaProducer interface{
	SendMessage(key, value string) error
}


func (s IngestionService) Ingest(dataURL string) error{
//TODO:
	return nil
}