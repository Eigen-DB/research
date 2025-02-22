package product_quantization

import (
	"math"
	"math/rand/v2"
)

var CENTROIDS [][][]float32

func GenerateRandomVector(dim int) []float32 {
	vec := make([]float32, dim)
	for i := 0; i < dim; i++ {
		vec[i] = rand.Float32()*200 - 100 // vectors components range from [-100,100)
	}
	return vec
}

func ComputeL2Distance(u []float32, v []float32) float64 {
	dist := 0.0
	for i := 0; i < len(u); i++ {
		dist += math.Pow(float64(u[i]-v[i]), 2)
	}
	return math.Sqrt(dist)
}

func generateSubVectors(m int, vector []float32) [][]float32 {
	subVectors := make([][]float32, m)
	D_ := len(vector) / m
	for i := 0; i < m; i++ {
		subVectors[i] = vector[i*D_ : (i+1)*D_]
	}
	return subVectors
}

func GenerateCentroids(k int, m int, D_ int) {
	if k%m != 0 {
		panic("k must be divisible by m")
	}

	centroids := make([][][]float32, m)
	k_ := k / m
	for i := range m {
		centroids[i] = make([][]float32, k_)
		for j := range k_ {
			centroids[i][j] = make([]float32, D_)
			centroids[i][j] = GenerateRandomVector(D_)
		}
	}
	CENTROIDS = centroids
}

func compressVector(subVectors [][]float32) []float32 {
	compressed := make([]float32, len(subVectors))
	for i := range len(subVectors) {
		minDist := math.MaxFloat64
		for j := range len(CENTROIDS[i]) {
			dist := ComputeL2Distance(subVectors[i], CENTROIDS[i][j])
			if dist < minDist {
				minDist = dist
				compressed[i] = float32(j)
			}
		}
	}
	return compressed
}

func DecompressVector(compressed []float32) []float32 {
	decompressed := make([]float32, 0)
	for i := range len(compressed) {
		decompressed = append(decompressed, CENTROIDS[i][int(compressed[i])]...)
	}
	return decompressed
}

func CompressManyVectors(num_vectors int, dim int, vectors [][]float32, m int, k int) [][]float32 {
	subVectors := make([][][]float32, num_vectors)
	for i := range num_vectors {
		subVectors[i] = generateSubVectors(m, vectors[i])
		//fmt.Println(subVectors[i])
	}

	compressedVectors := make([][]float32, num_vectors)
	for i := range num_vectors {
		compressedVectors[i] = compressVector(subVectors[i])
		//fmt.Println(compressedVectors[i])
	}

	return compressedVectors
}
