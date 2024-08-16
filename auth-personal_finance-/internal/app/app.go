package app

import (
	"context"

	"auth/api"
	"auth/api/handlers"
	"auth/config"
	kafka "auth/internal/kafka/consumer"
	prd "auth/internal/kafka/producer"
	"auth/internal/logger"
	"auth/internal/storage/postgres"
	"auth/service"

	"github.com/go-redis/redis/v8"
)

// var (
// 	_, b, _, _ = runtime.Caller(0)
// 	basepath   = filepath.Dir(b)
// )

func Run(cfg *config.Config) {
	// basepath = "/home/javohir/Documents/Go_month_5/exam/auth-personal_finance-"
	l := logger.NewLogger(
	// basepath, cfg.LOG_PATH
	)

	// Postgres Connection
	db, err := postgres.NewPostgresStorage(cfg)
	if err != nil {
		l.ERROR.Printf("can't connect to db: %v", err)
	}
	defer db.Db.Close()
	l.INFO.Println("Connected to Postgres")

	// Redis Connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		l.ERROR.Panicf("Failed to connect to Redis: %v", err)
	}
	l.INFO.Println("Connected to Redis")

	authService := service.NewAuthService(db)
	userService := service.NewUserService(db)

	// Kafka
	brokers := []string{"kafka:9092"}
	cm := kafka.NewKafkaConsumerManager()
	pr, err := prd.NewKafkaProducer(brokers)
	if err != nil {
		l.ERROR.Println("Failed to create Kafka producer:", err)
		return
	}

	Reader(brokers, cm, authService, userService, l)

	// HTTP Server
	h := handlers.NewHandler(authService, userService, rdb, &pr, l)

	router := api.Engine(h)
	router.SetTrustedProxies(nil)

	if err := router.Run(cfg.AUTH_PORT); err != nil {
		l.ERROR.Panicf("can't start server: %v", err)
	}

	l.INFO.Printf("REST server started on port %s", cfg.AUTH_PORT)
}
