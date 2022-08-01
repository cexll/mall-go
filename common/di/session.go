package di

func init() {
	//obj := xdi.Object{
	//	Name: "session",
	//	New: func() (i interface{}, e error) {
	//		opts := redis.Options{
	//			Addr:        dotenv.Getenv("REDIS_ADDR").String(),
	//			Password:    dotenv.Getenv("REDIS_PASSWORD").String(),
	//			DB:          int(dotenv.Getenv("REDIS_DATABASE").Int64()),
	//			DialTimeout: time.Duration(dotenv.Getenv("REDIS_DIAL_TIMEOUT").Int64(10)) * time.Second,
	//		}
	//		opt := redis.NewRedisStore(&opts)
	//		return session.NewManager(session.SetStore(opt)), nil
	//	},
	//}
	//if err := xdi.Provide(&obj); err != nil {
	//	panic(err)
	//}
}
