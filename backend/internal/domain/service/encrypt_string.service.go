package domain_service

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

type EncryptStringServiceInterface interface {
	EncryptString(str string, cost int) ([]byte, error)
}
