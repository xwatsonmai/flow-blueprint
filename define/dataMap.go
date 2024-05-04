package define

type IDataMap interface {
	Get(key string) any
	Set(key string, value any)
}
