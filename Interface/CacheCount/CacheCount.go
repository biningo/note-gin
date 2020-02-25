package CacheCount

type CacheCount interface {
	SetSubFileTotal(FolderID int64, val int)
	GetSubFileTotal(FolderID int64) (int, bool)
	AddSubFileTotal(FolderID int64, val int)
	SubtractSubFileTotal(FolderID int64, val int)
	DelSubFileTotal(FolderID int64)

	SetSubFolderTotal(FolderID int64, val int)
	GetSubFolderTotal(FolderID int64) (int, bool)
	AddSubFolderTotal(FolderID int64, val int)
	SubtractSubFolderTotal(FolderID int64, val int)
	DelSubFolderTotal(FolderID int64)
}
