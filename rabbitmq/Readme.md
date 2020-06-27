## RabitMQ
### 1. Khái niệm
RabbitMQ là một message broker (message-oriented middleware) sử dụng giao thức AMQP — Advanced Message Queue Protocol. RabbitMQ được lập trình bằng ngôn ngữ Erlang. RabbitMQ cung cấp cho lập trình viên một phương tiện trung gian để giao tiếp giữa nhiều thành phần trong một hệ thống lớn. RabbitMQ sẽ nhận message đến từ các thành phần khác nhau trong hệ thống, lưu trữ chúng an toàn trước khi đẩy đến đích.

### 2. Ưu điểm
Một producer không cần phải biết consumer. Nó chỉ việc gửi message đến các queue trong message broker. Consumer chỉ việc đăng ký nhận message từ các queue này
Vì producer giao tiếp với consumer trung gian qua message broker nên dù producer và consumer có khác biệt nhau về ngôn ngữ thì giao tiếp vẫn thành công.(Hiện nay rabbitmq đã hỗ trợ rất nhiều ngôn ngữ khác nhau).
Một đặc tính của rabbitmq là bất đồng bộ(asynchronous). Producer không thể biết khi nào message đến được consumer hay khi nào message được consumer xử lý xong. Đối với producer, đẩy message đến message broker là xong việc. Consumer sẽ lấy message về khi nó muốn. Đặc tính này có thể được tận dụng để xây dựng các hệ thống lưu trữ và xử lý log.

### 3. Flow gửi nhận message:
![](https://miro.medium.com/max/1040/1*gVzRDGT9ZLAvvfeAfkp6dQ.png)

- Producer đẩy message vào exchange. Khi tạo exchange, phải chỉ định nó thuộc loại exchange gì? Các loại exchange sẽ giải thích kĩ hơn ở phần bên dưới.
- Sau khi exchange nhận message, nó chịu trách nhiệm định tuyến message. Exchange sẽ chịu trách về các thuộc tính của message, ví dụ routing key, phụ thuộc loại exchange.
- Việc binding phải được tạo từ exchange đến hàng đợi. Trong trường hợp này, ta sẽ có hai binding đến hai hàng đợi khác nhau từ một exchange. Exchange sẽ định tuyến message vào các hàng đợi dựa trên thuộc tính của của từng message.
- Các message nằm ở hàng đợi đến khi chúng được xử lý bởi một consumer.
- Consumer xử lý message trong queue.

### 4. Các loại Exchange trong RabbitMQ
Có 4 loại Exchange trong RabbitMQ: direct, topic, fanout, headers.

#### Direct Exchange
![](https://miro.medium.com/max/1040/1*3tq1Nj_RAzmtyG2TKb2pQw.png)

Một direct exchange cung cấp các messages tới queues dựa trên một message routing key. Routing key là một thuộc thuộc tính của mesage được thêm vào message header từ producer. Routing key có thể được xem là một địa chỉ mà 

Exchange sử dụng để định tuyến các messages.
Direct exchange được sử dụng trong trường hợp bạn muốn phân biệt các messages published cho cùng một exchange bằng cách sử dụng một chuỗi định danh đơn giản.

Ví dụ:

- Một message với routing key pdf_log được gửi từ exchange pdf_events. Message đó được định tuyến tới pdf_log_queue vì routing key (pdf_log) khớp với binding key (pdf_log).

- Nếu message routing key không khớp với binding key nào thì sẽ bị hủy bỏ.

#### Topic Exchange
![](https://miro.medium.com/max/1040/1*-yj4_wK6UiAe8pjIA1vJGg.png)

Một topic exchange sẽ làm một lá bài (gọi là wildcard) để gắn routing key với một routing pattern khai báo trong binding. Consumer có thể đăng ký những topic mà nó quan tâm. Cú pháp được sử dụng ở đây là * và #.

Ví dụ:

- Một message có routing key agreements.eu.berlin được gửi tới exchange agreements. Message được đính tuyến đến queue berlin_agreements bởi vì routing pattern của “agreements.eu.berlin.#” match với bất kì routing key nào bắt đầu là “agreements.eu.berlin”. Tin nhắn cũng được định tuyến đến queue all_agreements vì routing key (agreements.eu.berlin) match routing pattern (agreements.#)

#### Fanout Exchange
![](https://miro.medium.com/max/1040/1*XIM5AQgGLD-RgGcLG7ed4w.png)

Fanout copy và định tuyến một mesage nhận được tới tất cả queue được ràng buộc với nó.

Exchange này hữu ích với trường hợp ta cần một dữ liệu được gửi tới nhiều thiết bị khác nhau với cùng một message nhưng cách xử lý ở mỗi thiết bị, mỗi nơi là khác nhau.

Ví dụ:

- Một mesage được gửi đến exchange sport news. Message đó được định tuyến tới tất cả các queues (Queue A, Queue B, Queue C).

#### Header Exchange
![](https://miro.medium.com/max/1040/1*GV8aoRn1Yrdo6fj9A5mo9A.png)

Một header exchange sẽ dùng các thuộc tính header của message để định tuyến. Headers Exchange rất giống với Topic Exchange, nhưng nó định tuyến dựa trên các giá trị tiêu đề thay vì các khóa định tuyến.

Một thông điệp được coi là phù hợp nếu giá trị của tiêu đề bằng với giá trị được chỉ định khi ràng buộc.

Ví dụ:

- Message 1 được gửi đến exchange agreements với header là: format: pdf type: report. Message được định tuyến đến Queue A vì giá trị header này khớp với giá trị chỉ định là format: pdf type: report.

### Demo

Start RabbitMQ

` docker run -d -p 5672:5672 -p 15672:15672 -e RABBITMQ_DEFAULT_USER=user -e RABBITMQ_DEFAULT_PASS=password rabbitmq:3-management`

- Vậy là đã hoàn thành, ta sẽ gõ vào trình duyệt địa chỉ: http://127.0.0.1:15672/ để kiểm tra kết quả:

![](https://miro.medium.com/max/1400/1*RDvJhtqEQuTtyWoSLOt3CA.png)

- user và password mặc định sẽ là: guest

![](https://miro.medium.com/max/1400/1*ppAkvs_modQabeR6zSHnUw.png)

#### Sử dụng RabbitMQ trong Go
Get amqp:
`go get github.com/streadway/amqp`

Tạo 1 file sending.go nằm trong folder send và import libary amqp:

```
package main

import (
  "log"

  "github.com/streadway/amqp"
)
```

connect to RabbitMQ server

```
conn, err := amqp.Dial(“amqp://password:user@localhost:5672/”)
failOnError(err, “Failed to connect to RabbitMQ”)
defer conn.Close()
```

create channel

```
ch, err := conn.Channel()
failOnError(err, "Failed to open a channel")
defer ch.Close()
```

khai báo 1 queue, để có thể publish message đến queue

Ở đây ta khai báo với name là "hello" và sẽ publish message là "Hello World!"

```
q, err := ch.QueueDeclare(
  "hello", // name
  false,   // durable
  false,   // delete when unused
  false,   // exclusive
  false,   // no-wait
  nil,     // arguments
)
failOnError(err, "Failed to declare a queue")

body := "Hello World!"
err = ch.Publish(
  "",     // exchange
  q.Name, // routing key
  false,  // mandatory
  false,  // immediate
  amqp.Publishing {
    ContentType: "text/plain",
    Body:        []byte(body),
  })
failOnError(err, "Failed to publish a message")
```

Tạo 1 file receiving.go trong folder receiv

connect và tạo 1 channel. Declare queue, lưu ý name cần khớp với khái báo queue ở file sending.go

Ở đây thì sẽ set name là "hello"

```
conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()

ch, err := conn.Channel()
failOnError(err, "Failed to open a channel")
defer ch.Close()

q, err := ch.QueueDeclare(
  "hello", // name
  false,   // durable
  false,   // delete when unused
  false,   // exclusive
  false,   // no-wait
  nil,     // arguments
)
failOnError(err, "Failed to declare a queue")
```

Đọc các tin nhắn từ channel (amqp: Consume):

```
msgs, err := ch.Consume(
  q.Name, // queue
  "",     // consumer
  true,   // auto-ack
  false,  // exclusive
  false,  // no-local
  false,  // no-wait
  nil,    // args
)
failOnError(err, "Failed to register a consumer")

forever := make(chan bool)

go func() {
  for d := range msgs {
    log.Printf("Received a message: %s", d.Body)
  }
}()

log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
<-forever
```

Bây giờ chúng ta sẽ tiến hành chạy 2 file sending.go và receiving.go

