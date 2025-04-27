package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var MyData string = `{
	"id": 3379243,
	"title": "برنج ایرانی طارم ممتاز معطر گلستان ۴.۵ کیلوگرمی",
	"subtitle": "۴.۵ کیلوگرم",
	"description": "",
	"content": null,
	"max_order_cap": 2,
	"price": 893400,
	"discount_percent": 11,
	"discounted_price": 795126,
	"has_alternative": false,
	"images": [
		{
			"image": "https://api.snapp.market/media/cache/product-variation_image2/uploads/images/vendors/users/app/20211016-189827-1.jpg",
			"thumb": "https://api.snapp.market/media/cache/product_variation_transparent_image/20211016-189827-1.png"
		},
		{
			"image": "https://api.snapp.market/media/cache/product-variation_image2/uploads/images/vendors/users/app/20211016-189827-2.jpg",
			"thumb": "https://api.snapp.market/media/cache/product-variation_image_thumbnail/uploads/images/vendors/users/app/20211016-189827-2.jpg"
		}
	],
	"brand": {
		"id": 1227,
		"title": "‌‌گلستان‌",
		"slug": "golestan",
		"english_title": "Golestan"
	},
	"review_count": null,
	"rating_value": 3,
	"html_description": "",
	"meta_description": "",
	"meta_keywords": "",
	"needs_server_approval": false,
	"tags": [
		{
			"id": 13937,
			"title": "برنج سفید",
			"slug": null
		}
	],
	"coupons": [],
	"badges": [
		{
			"title": "خرید اقساطی",
			"color": "red",
			"icon": null
		}
	],
	"pureTitle": "برنج ایرانی طارم ممتاز معطر گلستان",
	"supplier_id": 0,
	"supplier_title": null
}`

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// body := bodyFrom(os.Args)
	body := MyData
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
