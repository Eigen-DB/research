// Written by Ryan Awad

package main

import (
	"fmt"
	"math"
	"product_quantization/product_quantization"
	"product_quantization/vector_space"
)

func main() {
	dim := 12
	m := 4
	if dim%m != 0 {
		panic("dim must be divisible by m")
	}
	k := int(math.Pow(2, 16))
	uncompressedIndex := vector_space.CreateIndex(dim)
	compressedIndex := vector_space.CreateIndex(m)
	// freeing the indexes from memory when program ends
	defer uncompressedIndex.Free()
	defer compressedIndex.Free()

	num_vectors := 100
	vectors := make([][]float32, num_vectors)
	for i := range num_vectors {
		vectors[i] = product_quantization.GenerateRandomVector(dim)
		fmt.Println(vectors[i])
	}
	vector_space.InsertVectors(uncompressedIndex, vectors)

	product_quantization.GenerateCentroids(k, m, dim/m)

	compressed_vectors := product_quantization.CompressManyVectors(num_vectors, dim, vectors, m, k)
	for _, v := range compressed_vectors {
		fmt.Println(v)
	}
	vector_space.InsertVectors(compressedIndex, compressed_vectors)

	fmt.Println("Comparing similarity search results")
	fmt.Println("Uncompressed:")
	uNnIds, nnDists := vector_space.Search(uncompressedIndex, uint64(2), 4)
	fmt.Printf("Neighbor IDs: %v\n", uNnIds)
	for i := range nnDists {
		nnDists[i] = float32(math.Sqrt(float64(nnDists[i])))
	}
	fmt.Printf("Neighbor distances from query vector: %v\n", nnDists)

	fmt.Println("Compressed:")
	cNnIds, nnDists := vector_space.Search(compressedIndex, uint64(2), 4) // that's probably not how you handle similarity search with compressed vectors..
	fmt.Printf("Neighbor IDs: %v\n", cNnIds)
	for i := range nnDists {
		nnDists[i] = float32(math.Sqrt(float64(nnDists[i])))
	}
	fmt.Printf("Neighbor distances from query vector: %v\n", nnDists)

	maxVectorSpaceDist := 0.0
	for i := 1; i < dim; i++ {
		maxVectorSpaceDist += math.Pow(100, 2) // range of vector componentsi s [-100, 100)
	}

	/*
		CONCLUSION: You cannot mix Product Quanitzation and HNSW together
		This was confirmed in the "Efficient and robust approximate nearest neighbor search using Hierarchical Navigable Small World graphs.pdf" paper
	*/

	//maxVectorSpaceDist = 2 * math.Sqrt(maxVectorSpaceDist)
	//fmt.Printf("Max distance in vector space: %f\n", maxVectorSpaceDist)
	/*
		fmt.Printf("Comparing the compressed vs uncompressed neighbors\nMax distance in vector space: %f\n", maxVectorSpaceDist)
		for i := range uNnIds {
			uVec, err := uncompressedIndex.GetVector(uNnIds[i])
			if err != nil {
				panic(err)
			}
			fmt.Printf("Uncompressed vector %d: %v\n", uNnIds[i], uVec)

			cVec, err := compressedIndex.GetVector(cNnIds[i])
			if err != nil {
				panic(err)
			}
			decompressedVec := product_quantization.DecompressVector(cVec)
			fmt.Printf("DEcompressed vector %d: %v\n", cNnIds[i], decompressedVec)
			fmt.Printf("Distance between the two: %f\n\n", product_quantization.ComputeL2Distance(decompressedVec, uVec))
		}*/
}
