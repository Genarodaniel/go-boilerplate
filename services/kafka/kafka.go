package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaInterface interface {
	Produce(ctx context.Context, topic string, key string, body any) error
}

type Kafka struct {
	Client *kgo.Client
}

func NewKafka(seeds []string, topics []string) (KafkaInterface, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ConsumeTopics(topics...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka client: %s", err.Error())
	}

	return &Kafka{
		Client: client,
	}, nil
}

func (k *Kafka) Produce(ctx context.Context, topic string, key string, body any) error {
	payload, err := k.SerializePayload(body)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(1)
	record := &kgo.Record{Topic: topic, Value: payload, Key: []byte(key)}

	var errReturn error
	k.Client.Produce(ctx, record, func(_ *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			mutex.Lock()
			errReturn = fmt.Errorf("record had a produce error: %v\n", err)
			mutex.Unlock()
		}

	})
	wg.Wait()
	return errReturn
}

func (k *Kafka) SerializePayload(body any) ([]byte, error) {
	response, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("Error while serializing payload %s", err.Error())
	}
	return response, nil
}
