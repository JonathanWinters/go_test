package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/JonathanWinters/go_test/internal/core"
	"github.com/JonathanWinters/go_test/internal/definitions"
	"github.com/JonathanWinters/go_test/internal/util"
	"github.com/go-redis/redis/v9"
)

const redisServerAddr = "localhost:6380"
const LockKey = "lock_redis"
const LockTimeout = 20 * time.Second

type SubmitRequestBody struct {
	UserId string
	Level  [][]int
}

type MoveRequestBody struct {
	PrimaryKey int
	Move       int
	GodMode    bool
}

type RedisEventType int

const (
	EventExpired RedisEventType = iota
	EventDel
)

type RedisClient struct {
	client *redis.Client
}

var redisEventName = map[RedisEventType]string{
	EventExpired: "expired",
	EventDel:     "del",
}

func (ret RedisEventType) String() string {
	return redisEventName[ret]
}

type Callback func(queue []core.MoveRequest)

var rc RedisClient

func StartWebServer() {
	// Start the server and listen on port 8080
	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func StartRedisServer() {
	rc.client = redis.NewClient(&redis.Options{
		Addr:     redisServerAddr,
		Password: "",
	})
}

func StartServers() {

	var wg sync.WaitGroup

	go StartRedisServer()

	// Start web server in a goroutine
	wg.Add(1)
	go StartWebServer()

	// Keep the main goroutine alive until servers are shut down (e.g., via a signal)
	// For a simple example, a select{} block can prevent main from exiting immediately
	select {}

	wg.Wait() // Wait for goroutines to finish (if you implement graceful shutdown)
	fmt.Println("Application shutting down.")
}

func AcquireLock(client *redis.Client, lockKey string, timeout time.Duration) (lockAcquired bool) {
	ctx := context.Background()

	// Try to acquire the lock with SETNX command (SET if Not eXists)
	lockAcquired, err := client.SetNX(ctx, lockKey, "1", timeout).Result()
	if err != nil {
		fmt.Println("Error acquiring lock:", err)
		return
	}

	return
}

func ReleaseLock(client *redis.Client, lockKey string) {
	ctx := context.Background()
	client.Del(ctx, lockKey)
}

func ProcessMoveRequest(w http.ResponseWriter, moveRequest core.MoveRequest) (rawResult []byte, err error) {

	lockAcquired := AcquireLock(rc.client, LockKey, LockTimeout)

	var moveResponse core.MoveResponse
	if lockAcquired {
		moveResponse = core.HandleMove(w, moveRequest)
	} else {
		moveResponse.Error = "RedisThread is Locked"
	}
	ReleaseLock(rc.client, LockKey)
	rawResult, err = json.Marshal(moveResponse)
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}

func ProcessSubmitRequest(w http.ResponseWriter, r *http.Request) (rawResult []byte, err error) {

	srb, err := DecodeSubmitJson(w, r)
	if err != nil {
		log.Fatal(err)
		return
	}

	responseObj := SubmitRequestBody{
		UserId: srb.UserId,
		Level:  srb.Level,
	}

	userid := definitions.UserIDFromString(responseObj.UserId)

	submitRequest := core.SubmitRequest{
		UserID: userid,
		Level:  responseObj.Level,
	}

	submitResponse := core.HandleSubmit(w, submitRequest)

	rawResult, err = json.Marshal(submitResponse)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func DecodeMoveJson(w http.ResponseWriter, r *http.Request) (mrb MoveRequestBody, err error) {
	notJson := util.HandleContentTypeError(w, r)
	if notJson {
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, util.MaxMB)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&mrb)
	if err != nil {
		util.HandleDecodeError(w, err)
		return
	}
	return
}

func DecodeSubmitJson(w http.ResponseWriter, r *http.Request) (srb SubmitRequestBody, err error) {
	notJson := util.HandleContentTypeError(w, r)
	if notJson {
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, util.MaxMB)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&srb)
	if err != nil {
		util.HandleDecodeError(w, err)
		return
	}
	return
}
