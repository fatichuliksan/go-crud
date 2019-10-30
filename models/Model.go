package models

type Model interface {
	GetCollectionName() string
	SetCollectionName(collectionName string)
}
