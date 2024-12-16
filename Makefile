BUILD_BINARY=.\bin\app.exe
GO_BINARY=app.exe

build:
	@echo start build
	@go build -o ${BUILD_BINARY} .\cmd\app
	@echo - done

build-prod:
	@echo start build
	@go build -o ${BUILD_BINARY} -ldflags '-s' .\cmd\app
	@echo - done

run: build
	@echo start run
	@start /B ${BUILD_BINARY} &
	@echo - done

stop:
	@echo killing proces
	@taskkill /IM ${GO_BINARY} /F
	@echo - done