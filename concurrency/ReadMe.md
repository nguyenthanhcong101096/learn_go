![](https://s3-ap-southeast-1.amazonaws.com/kipalog.com/zss1fjcca6_concurrency.jpg)

- Là khả năng chương trình có thể điều phối tác vụ trong cùng 1 khoảng thời gian và trong quá trình điều phối chỉ cho phép 1 tác vụ chạy trong 1 thời điểm

### Tại sao cần Concurrency
- Tận dùng tối đa CPU
- Tăng tính phản hồi của ứng dụng


![](https://s3-ap-southeast-1.amazonaws.com/kipalog.com/6cpbibh0yq_parallelism.jpg)
- Là khả năng 1 chương trình có thể thực thi 2 hoặc nhiều tasks vụ trong cùng 1 thời điểm. (CPU > 1)


### So sánh Concurrency và Parallelism

![](https://viblo.asia/uploads/410863aa-11cf-4b82-b67d-e9b78b4eeaa1.png)

### Thread và Process
![](https://www.w3.org/People/Frystyk/thesis/MultiStackThread.gif)

**Process**

- 1 **Process** đại diện cho 1 chương trình running trong hệ điều hành
- 1 **Process** bao gồm:
  - Nhiều **thread** con bên trong
  - Chứa **memory** (**static heap code**)

**Thread**

- Bao gồm:
  - **Register** (Thanh ghi)
  - **Stack**

### Cách làm việc của Concurrency
![](https://lh3.googleusercontent.com/AazPtmFbIpTxl2Xzei6DCxkxaLImMnyXfXIwv2tT8zHMrrhg_txHe67jNNZgd-yqwnUTV18wiW6-n02RnhFCm9jyLkM-6Q6XmViWsYN171mGqFDQWllpDr2wvaWZ55kTsJ2nprLmtMXRn4WITbe7qHdKpm4FQGTiAT-SndpW5FQvQpsvGRdlE7OUD27qaCi9SjzCZfCB8TXs-neo0wwM102mPBUz1qxprGM2z-SMXdNBkxCSgMW5c5xqXoqllFMWA14vqHphgUbfkkODkf2QHQ08oB-1gPSnuxb_xNQDCCa8R_6_OMXfpmIVwB1bAXAPXlogyX1jdzyNG3if2GcLfb-cD_S42BN7A_VPH05OEeEe5FVNCBs0d3HqvYB7iikTb5oFuGeX2-G0R1cfKUp5hIe7ujoY7ktXUxmZyQshW6zSl-cIIDCIedQ_-whUuRaseIW8EomQ3Ehwx8aA0enobPtxBFMRpI1ua79UA0kOx5a3gb0Lgbke7Pf3ocvk4gRZZveFsb1EDM9BpmHwcepa2nE4oNcl9v7pXbhUZw7XUrhIU17tHqcYpkRuY4hsIWp1zD1Zwo7dBk7_vDho9cA_QzdTVFDyi0SMBfsnYzzFctIYZKy29qAA5RVwrbau80fE6kbLHJrbGILtRZgFWXtAQdBgExD1-UraCAI9f_9v3XjM6JQKwDz8GJzl7fwd=w1170-h815-no?authuser=1)

**2 Lỗi thường gặp**

- Race Condition

Khi 2 thread cùng đọc cùng đọc 1 biến và cùng 2 thread cùng ghi biến vào vùng nhớ. gây ra sự k đúng đắn dữ liệu

![](https://lh3.googleusercontent.com/KEMI4REAOY883ReKJ7xEYlxuwe5RoegTgM5zszzF8vjPhSN0z5BIEgHIGOQdBabgjc1SAQwBrZUPRPm7xeA2lbHvGhQstFkxorth654fjidkwquZZLNCt4pHOdmglSOtt3vMOnrHtkiZ_586LeOExVQWNurS2v0mkFSOP0PNramWR7tb9i6A55KbcMkum7zyODMSQ-p1NqTgW8fTeV7B2vLkKbRkwmGvVYVohzv5eCUnkKtoIZP2yMw-5fJG912xDLLg112ni0O5x22tT3pTsUcT09C36vlC1zPB29tBjGgY54rf2ry0UqaXxUvfITKzKsAfGRgUV28qU8iwTGSb_7WRqEOOSqOQ2mgQbpW9ums2KfkNpnoFIp13hZj-ryn7o4meoWUj3xk09MWBiVQBgXApc-tcCRB9SgrmbsVc-KlG2IOnrcbNFpp37hr0WYlV2D2HyZJKhlSM-mrqHJArEiLotzoGJV7_w3QoHhCv_67FDaCO1lI1K5-4vdJWIgx9jLytF7kKftj1bA10g57WzLin3HpBR2BpHI8NsFaIziYCD4E1ve89uuUN45qaEySCqaiGCuJQToKachlOTF6Xd1_EdK5a1ekB5RGJrxIuxDVl9YLU5YhUEQYGBieUM7AHEuxnbY9qQM5ro99MPtDOQUNH1Xwb6n4XQXhWApGE_SWFIVzZ8aqLtIJIepG8=w944-h525-no?authuser=1)

- Deadlock

Khi 2 thread cùng chờ resource lẫn nhau


### Concurrency in GO
- `thread` -> `goroutine`

- giao tiếp 2 `goroutine` sử dùng `channel`

#### Cú pháp
- Go routines

`go doSomeWork()`

- Channel

`c <- data` | `data = <- c`

> Giao tiếp giữa Go routine thì k trực tiếp chia sẽ bộ nhớ

### Go Scheduler
- Thành phần của **GoScheduler**
  - **Goroutine**
  - **Machine** (`Worker thread` or `OS thread`) -> Tạo ra và control bởi HĐH
  - **Process** (`context`)
  		-  `local runq` luôn tồn tại (chứa các go routines của chương trình)
  		-  `global` k gắn trục tiếp vào `process`. nó được dùng khi `local runq` k thể chứa hết goroutines -> đẩy `goroutines` vào `global runq`


#### Khởi tạo chương trình
Ví dụ tạo 2 Process (`GOMAXPROCS(2)`)

- go runtime sẽ tạo **Machine(M)** gắn với **Process(P)**
- **Process(P)** không trực tiếp xử lý mà **Machine(M)** sẽ xử lý

Tạo 4 gorouties

```
func g1() { fmt.Println("G1") }
func g2() { fmt.Println("G2") }
func g3() { fmt.Println("G3") }
func g4() { fmt.Println("G4") }
```

Mô hình 

![](https://lh3.googleusercontent.com/c9XzMl-vIFj_26Xa0GFeYTIPtNsstanNmQg2-jdLPEtXzH6IXEjrrIEAYYguq90IlEs-v5bz0yds-q9jF0XuAFzhS-X1eAyHcgrRsFJp7hbx4S3uDVSoj8--TxkPCGRpJKXWP6ctT33JQkYwbowe8FeH6g8ZdKVYXbZ3F4XtZT7HTv2ylcwgIBbH_wZK8XdmS0u5q3GcKyrsTIgurHyohGbe9mjn7avaXGyDXz9ML39qsgrQYVXmMBWykEoXg8Cxp8BTERGRcUvbw1_e57p3Yaa_EWm5zHf5euJswcPKtlaZaWGjsyvJQenXz3AGLhAVn8tr33qD0h3eNuFcM1mNy5XcI2SlxK6SEJOZawq83rx4MIwhg0wsspfOZPGspk9gWdc3oB_9Xfs3wEx7c1DDLTo9T0yn_CBULdWvnw97_9g1jl1L1iatqmSXVz69IL1cp5VIYcctR0mdviGbmTW7mg9F6-RymnvvOzN_kh2gMKWdxkJpRiLk6gv2YKq7wn_OGOuOP5S1UDabyJDYxuQVGSd9pCMUK92SBJNO2XE4h8CJ1PNSeeMFpBYIPZfzY10QVL_emXSvb2peYspp0DtJPOee-4g-qMBk1z7aQG2HUVONdFNprLsDUzQ-GC4Tex0Y59IlCv9PfWGOYL8FPAI6blCyFDSyyboXL2tyXvD42maUbrqsjWdEjLt5R-zD=w878-h479-no?authuser=1)
