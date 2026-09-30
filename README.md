##

swag init


---打包liunx
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

protoc --go_out=. BypSendResponse.proto
protoc --go_out=. FinderGetTopic.proto


{
		1: "47204784bbe5abe5f38fe8857569e868"
		2: "HUAWEI-MRX-W09"
		3: "pad-android-29"
		4: 1615386584
	}

--部署
上传conf文件夹
上传swagger文件夹
上传doc文件夹
上传main文件
执行脚本，完成


