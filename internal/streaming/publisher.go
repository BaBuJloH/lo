package streaming

import (
	"encoding/json"
	"log"
	"os"
	"wb-test-task/internal/db"

	stan "github.com/nats-io/stan.go"
)

type Publisher struct {
	sc   *stan.Conn
	name string
}

func NewPublisher(conn *stan.Conn) *Publisher {
	return &Publisher{
		name: "Publisher",
		sc:   conn,
	}
}

// Тестовый скрипт публикации данных Order
func (p *Publisher) Publish() {
	// some date to send
	item1 := db.Items{ChrtID: 9934930, TrackNumber: "WBILMTESTTRACK", Price: 453, Rid: "ab4219087a764ae0btest", Name: "Mascaras", Sale: 30, Size: "0", TotalPrice: 317, NmID: 2389212, Brand: "Vivienne Sabo", Status: 202}
	delivery := db.Delivery{Name: "Test Testov", Phone: "+9720000000", Zip: "2639809", City: "Kiryat Mozkin", Address: "Ploshad Mira 15", Region: "Kraiot", Email: "test@gmail.com"}
	payment := db.Payment{Transaction: "b563feb7b2b84b6test", RequestID: "", Currency: "USD", Provider: "wbpay", Amount: 1817, PaymentDt: 1637907727, Bank: "alpha", DeliveryCost: 1500, GoodsTotal: 317, CustomFee: 0}
	order := db.Order{OrderUID: "b563feb7b2b84b6test", TrackNumber: "WBILMTESTTRACK", Entry: "WBIL", Delivery: delivery, Payment: payment, Items: []db.Items{item1},
		Locale: "en", InternalSignature: "", CustomerID: "test", DeliveryService: "meest", Shardkey: "9", SmID: 99, DateCreated: "2021-11-26T06:22:19Z", OofShard: "1"}
	orderData, err := json.Marshal(order)
	if err != nil {
		log.Printf("%s: json.Marshal error: %v\n", p.name, err)
	}

	// An asynchronous publish API
	ackHandler := func(ackedNuid string, err error) {
		if err != nil {
			log.Printf("%s: error publishing msg id %s: %v\n", p.name, ackedNuid, err.Error())
		} else {
			log.Printf("%s: received ack for msg id: %s\n", p.name, ackedNuid)
		}
	}

	// публикация данных:
	log.Printf("%s: publishing data ...\n", p.name)
	nuid, err := (*p.sc).PublishAsync(os.Getenv("NATS_SUBJECT"), orderData, ackHandler) // returns immediately
	if err != nil {
		log.Printf("%s: error publishing msg %s: %v\n", p.name, nuid, err.Error())
	}
}
