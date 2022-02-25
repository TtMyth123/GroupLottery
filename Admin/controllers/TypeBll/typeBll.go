package TypeBll

type TypeStruct struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func GetGameTypeAll(sysUserId int) []TypeStruct {
	return nil
}
