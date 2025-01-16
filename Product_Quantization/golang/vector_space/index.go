package vector_space

import (
	"github.com/Eigen-DB/eigen-db/libs/hnswgo/v3"
)

func CreateIndex(dim int) *hnswgo.Index {
	index, err := hnswgo.New(
		dim,
		32,
		400,
		12345,
		1000,
		"l2",
	)
	if err != nil {
		panic(err)
	}
	return index
}

func InsertVectors(index *hnswgo.Index, vectors [][]float32) {
	for i, v := range vectors {
		if err := index.InsertVector(v, uint64(i)); err != nil {
			panic(err)
		}
	}
}

func Search(index *hnswgo.Index, queryVectorId uint64, k int) ([]uint64, []float32) {
	qVec, err := index.GetVector(queryVectorId)
	if err != nil {
		panic(err)
	}
	nnIds, nnDists, err := index.SearchKNN(qVec, k)
	if err != nil {
		panic(err)
	}
	return nnIds, nnDists
}
