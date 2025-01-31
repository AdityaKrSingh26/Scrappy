APP_NAME = internship-scraper
BUILD_DIR = build
PLATFORMS = wellfound glassdoor linkedin

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) main.go

clean:
	rm -rf $(BUILD_DIR)

run:
	go run main.go

.PHONY: build clean run