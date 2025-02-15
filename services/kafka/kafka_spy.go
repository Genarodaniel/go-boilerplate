package kafka

import "context"

type KafkaMock struct {
	ProduceError error
}

func (k KafkaMock) Produce(ctx context.Context, topic string, key string, body any) error {
	return k.ProduceError
}
