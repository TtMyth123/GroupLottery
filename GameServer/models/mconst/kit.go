package mconst

var mpArticleTypeName map[int]string
var mpBetStatusName map[int]string

func init() {
	mpArticleTypeName = make(map[int]string)
	mpArticleTypeName[ArticleType_1_GG] = ArticleType_N_1_GG
	mpArticleTypeName[ArticleType_2_XW] = ArticleType_N_2_XW
	mpArticleTypeName[ArticleType_3_WZ] = ArticleType_N_3_WZ

	mpBetStatusName = make(map[int]string)
	mpBetStatusName[Bet_Status_1] = Bet_Status_N_1
	mpBetStatusName[Bet_Status_2] = Bet_Status_N_2
	mpBetStatusName[Bet_Status_3] = Bet_Status_N_3

}
func GetArticleTypeName(v int) string {
	return mpArticleTypeName[v]
}

func GetBetStatusName(v int) string {
	return mpBetStatusName[v]
}
