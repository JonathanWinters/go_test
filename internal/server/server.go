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
	"github.com/JonathanWinters/go_test/internal/util"
	"github.com/go-redis/redis/v9"
)

const redisServerAddr = "localhost:6380"
const LockKey = "lock_redis"
const LockTimeout = 20 * time.Second

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

	// pubsubs := setRedisListeners()

	// for i := 0; i < len(pubsubs); i++ {

	// 	defer pubsubs[i].Close()
	// }
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

func setRedisListener(ctx context.Context, eventType RedisEventType, cb Callback) (pubsub *redis.PubSub) {

	pubsub = rc.client.PSubscribe(ctx, "__keyevent@0__:"+eventType.String())
	// defer pubsub.Close()

	// Channel to receive messages
	ch := pubsub.Channel()

	fmt.Println("Listening for Redis key expiration events...")

	for msg := range ch {
		fmt.Printf("Received message on channel %s: %s\n", msg.Channel, msg.Payload)
		// The payload will be the name of the expired key
		// You can then check if this key corresponds to an "unlocked" lock
		if isLockKey(msg.Payload) { // Implement your logic to identify lock keys
			fmt.Printf("Lock key '%s' expired (unlocked)\n", msg.Payload)
			// Perform actions related to the unlocked event
			cb(MoveRequestQueue)
		}
	}
	return
}

// func setRedisListeners() (pubsubs []redis.PubSub) {
// 	ctx := context.Background()

// 	setRedisListener(ctx, EventExpired, DequeueMove)
// 	setRedisListener(ctx, EventDel, DequeueMove)
// 	return
// }

// func EnqueueMove(queue []core.MoveRequest, element core.MoveRequest) {
// 	queue = append(queue, element) // Simply append to enqueue.
// }

// func DequeueMove(queue []core.MoveRequest) {
// 	// Slice off the element once it is dequeued.
// 	ReleaseLock(rc.client, LockKey)
// 	slice := queue[1:]
// 	if slice == nil {
// 		return
// 	}
// }

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

func isLockKey(key string) bool {
	// Example: Check if the key starts with a specific prefix for locks
	return len(key) > 4 && key[:4] == "lock"
}

func ProcessMoveRequest(w http.ResponseWriter, moveRequest core.MoveRequest) (rawResult []byte, err error) {
	// Queue Up Requests, then process
	//!NOTED
	// !INFO EFC: for a proper queue and/or multiplayer, we utilize locking in redis (redis???) and/or mutexs
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

func DecodeJson(w http.ResponseWriter, r *http.Request) (mrb MoveRequestBody, err error) {
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
