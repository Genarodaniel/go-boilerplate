package kafka

import "context"

type KafkaMock struct {
	ProduceError             error
	SerializePayloadError    error
	SerializePayloadResponse []byte
}

func (k KafkaMock) Produce(ctx context.Context, topic string, key string, body any) error {
	return k.ProduceError
}

func (k KafkaMock) SerializePayload(body any) ([]byte, error) {
	return k.SerializePayloadResponse, k.SerializePayloadError
}
