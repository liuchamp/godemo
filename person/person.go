package person

//go:generate mockgen -destination=../mock/male_mock.go -package=mock github.com/liuchamp/godemo/person Male
type Male interface {
	Get(id int64) error
}
