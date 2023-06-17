APP_NAME	=	server
BUILD_DIR	=	bin
GO_BUILD	=	go build
GO_MAIN		=	main.go
DELETE_DIR 	= 	rm -rf
CREATE_DIR 	= 	mkdir -p
makedir:
	@$(CREATE_DIR) $(BUILD_DIR);


build: makedir
	@$(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME) $(GO_MAIN)

start: build
	@$(BUILD_DIR)/$(APP_NAME)

clean:
	@$(DELETE_DIR) $(BUILD_DIR)