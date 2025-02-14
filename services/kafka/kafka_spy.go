package kafka

import "context"

type KafkaSpy struct {
	ProduceError error
}

func (k KafkaSpy) Produce(ctx context.Context, topic string, key string, body any) error {
	return k.ProduceError
}
