package initializer

func Initialize(dbURL string) (Infra, Service) {
	db := InitDB(dbURL)
	infra := InitInfra(db)
	service := InitService(infra)

	return infra, service
}
