package event_consumer

import (
	"log"
	"mybot/events"
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

func (c *Consumer) Start() error {
	for {
		events, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}

		if len(events) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err := c.handleEvents(events); err != nil {
			log.Print(err)

			continue
		}

	}
}

/*
1. Lost events solutions: 1)retry 2)backup to storage 3)fallback
4) confirmation for fetcher
2. Whole batch processing solutions: 1) stop after one or few errors, errors counter
3. Add concurrency (sync.WaitGroup())
*/

func (c *Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		if err := c.processor.Process(event); err != nil {
			log.Printf("can't handle event: %s", err.Error())

			continue
		}
	}

	return nil
}
