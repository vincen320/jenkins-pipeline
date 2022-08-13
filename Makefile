build_mock:
	mockgen -package mock -destination mock/userRepositoryMock.go github.com/vincen320/user-service/repository UserRepository