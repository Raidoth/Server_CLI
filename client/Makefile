APP_NAME 	= client.exe
BUILD_DIR 	= bin
GO_BUILD 	= go build
CREATE_DIR 	= mkdir
DELETE_DIR 	= rmdir
FLAG_DELETE = /S /Q
makedir:
	@IF NOT EXIST $(BUILD_DIR) ($(CREATE_DIR) $(BUILD_DIR))
build:makedir
	@$(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME) $(GO_MAIN)
clean:
	@IF EXIST $(BUILD_DIR) $(DELETE_DIR)  $(FLAG_DELETE) $(BUILD_DIR)