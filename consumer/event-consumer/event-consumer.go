package event_consumer

import (
	"github.com/just1yaroslav/go_telegrambot/events"
	"log"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (consumer Consumer) Start() error {
	for {
		gotEvents, err := consumer.fetcher.Fetch(consumer.batchSize)
		if err != nil {
			log.Printf("[ERROR] consumer error: %s", err.Error())
			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		if err := consumer.handleFunction(gotEvents); err != nil {
			log.Print(err)
			continue
		}
	}
}

func (consumer Consumer) handleFunction(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		if err := consumer.processor.Precess(event); err != nil {
			log.Printf("cants handle event: %s", err.Error())
			continue
		}
	}
	return nil
}
