run:
	npm run start -- --host 0.0.0.0

#  make sev GOOGLE_APPLICATION_CREDENTIALS=setting_api.json
sev:
	go run server/server.go
