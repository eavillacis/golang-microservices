PUBSUB_EMULATOR_HOST := localhost:8580
PUBSUB_PROJECT_ID := digital-platforms-302300

start-pubsub-emulator:
	gcloud beta emulators pubsub start --project=$(PUBSUB_PROJECT_ID) --host-port 0.0.0.0:8580

migrate-pubsub:
	PUBSUB_EMULATOR_HOST=$(PUBSUB_EMULATOR_HOST) go run ./scripts/migrate-pubsub.go
