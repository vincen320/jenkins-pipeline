build_mock:
	mockgen -package mock -destination mock/userRepositoryMock.go github.com/vincen320/user-service/repository UserRepository

go_build_user_service:
	env GOOS=linux CGO_ENABLE=0 go build -o user_service 
#WINDOWS:	set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLE=0 && go build -o user_service 


#JALANIN GO_BUILD_USER_SERVICE TERLEBIH DAHULU BARU YANG INI (PAKAI DOCKERFILE YANG AWAL-AWAL)
docker_up: go_build_user_service
#Stopping and Remove Docker Images if running
	docker-compose down
	docker-compose up --build -d



#PAKAI DOCKER FILE YANG BARU (perbedaan di Dockerfile(baru))
#KALAU PAKAI YANG docker_up(file Dockerfile yang awal2) itu build golangnya masi di windows(go_build_user_service), jadi kalau error PATHnya masih menggunakan path windows ( entah pathnya berpengaruh juga dengan PATH di docker atau tidak)
#TAPI JIKA PAKAI Dockerfile yang baru, itu binary golangnya di build di golang sendiri, jadi PATHnya sesuai dengan posisi path yang didefinisikan di Dockerfile, lalu dari binary tsb baru di copas ke OS alpine
docker_up2:
#Stopping and Remove Docker Images if running
	docker-compose down
	docker-compose up --build -d

#docker-up, dockerup2 menggunakan/dipengaruhi oleh file docker-compose, Dockerfile, main.go