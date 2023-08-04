package manager


type InfraManager interface {
	GetFile() string
	GetCustomerFilePath() string
}

type infraManagerImpl struct {
	filepath    string
	customerFile string
}

func (im *infraManagerImpl) initDB() {
	im.filepath = "your_db_connection_string_here"
}

func (im *infraManagerImpl) initCustomerFile() {
	im.customerFile = "customer.json"
}

func (im *infraManagerImpl) GetFile() string {
	return im.filepath
}

func (im *infraManagerImpl) GetCustomerFilePath() string {
	return im.customerFile
}

func NewInfraManager() InfraManager {
	infra := infraManagerImpl{}
	infra.initDB()
	infra.initCustomerFile()
	return &infra
}
