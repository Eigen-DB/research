// Written by Ryan Awad

package main

import (
	"fmt"
	"math"
	"product_quantization/product_quantization"
	"product_quantization/vector_space"
)

func main() {
	dim := 128
	m := 32
	if dim%m != 0 {
		panic("dim must be divisible by m")
	}
	k := int(math.Pow(2, 10))
	uncompressedIndex := vector_space.CreateIndex(dim)
	compressedIndex := vector_space.CreateIndex(m)
	// freeing the indexes from memory when program ends
	defer uncompressedIndex.Free()
	defer compressedIndex.Free()

	num_vectors := 10
	vectors := make([][]float32, num_vectors)
	for i := range num_vectors {
		vectors[i] = product_quantization.GenerateRandomVector(dim)
		fmt.Println(vectors[i])
	}
	vector_space.InsertVectors(uncompressedIndex, vectors)

	compressed_vectors := product_quantization.CompressManyVectors(num_vectors, dim, vectors, m, k)
	for _, v := range compressed_vectors {
		fmt.Println(v)
	}
	vector_space.InsertVectors(compressedIndex, compressed_vectors)

	fmt.Println("Comparing similarity search results")

	fmt.Println("Uncompressed:")
	nnIds, nnDists := vector_space.Search(uncompressedIndex, uint64(2), 4)
	fmt.Printf("Neighbor IDs: %v\n", nnIds)
	fmt.Printf("Neighbor distances from query vector: %v\n", nnDists)

	fmt.Println("Compressed:")
	nnIds, nnDists = vector_space.Search(compressedIndex, uint64(2), 4) // that's probably not how you handle similarity search with compressed vectors..
	fmt.Printf("Neighbor IDs: %v\n", nnIds)
	fmt.Printf("Neighbor distances from query vector: %v\n", nnDists)
}
