package vertex

import (
	"math"
	"testing"
)

func TestVertexCreation(t *testing.T) {
	var v = NewVertex(1, 2, 22)
	if v.x != 1 || v.y != 2 || v.z != 22 {
		t.Errorf("Expected Vertex(1, 2, 22), got Vertex(%f, %f, %f)", v.x, v.y, v.z)
	}
}

func TestCopyVertex(t *testing.T) {
	v1 := NewVertex(1, 2, 3)
	v2 := CopyVertex(v1)
	if v1.x != v2.x || v1.y != v2.y || v1.z != v2.z {
		t.Errorf("Expected copied Vertex to be equal, got Vertex(%f, %f, %f)", v2.x, v2.y, v2.z)
	}

	if &v1 == &v2 {
		t.Errorf("Expected copied Vertex to be a different instance, got same instance")
	}
}

func TestGetOrigo(t *testing.T) {
	v := Origo()
	if v.x != 0.0 || v.y != 0.0 || v.z != 0.0 {
		t.Errorf("Expected Origo(0, 0, 0), got Vertex(%f, %f, %f)", v.x, v.y, v.z)
	}
}

func TestVertexScaling(t *testing.T) {
	tests := []struct {
		v, expected Vertex
		s           float64
	}{
		{Vertex{1, 2, 3}, Vertex{1, 2, 3}, 1.0},
		{*Origo(), *Origo(), 304.908},
		{Vertex{-34.07, -204.9909, 203.076}, Vertex{-102.210000, -614.972700, 609.228000}, 3.0},
		{Vertex{3, -9, 27}, Vertex{1, -3, 9}, (1.0 / 3.0)},
	}

	for _, testData := range tests {
		testData.v.Scale(testData.s)
		if !floatEquals(testData.v.x, testData.expected.x, 1e-6) || !floatEquals(testData.v.y, testData.expected.y, 1e-6) || !floatEquals(testData.v.z, testData.expected.z, 1e-6) {
			t.Errorf("Expected Vertex(%f, %f, %f), got Vertex(%f, %f, %f)", testData.expected.x, testData.expected.y, testData.expected.z, testData.v.x, testData.v.y, testData.v.z)
		}
	}
}

func TestVertexAddition(t *testing.T) {
	tests := []struct {
		v, w     Vertex
		expected Vertex
	}{
		{Vertex{1, 1, 1}, Vertex{2, 2, 2}, Vertex{3, 3, 3}},
		{*Origo(), Vertex{2, 0, 2}, Vertex{2, 0, 2}},
		{Vertex{201, 34.001, 230}, Vertex{-324.99, -99.12, 76.304}, Vertex{-123.99, -65.119, 306.304}},
		{Vertex{-1, -10, -98}, Vertex{0, -4, -123}, Vertex{-1, -14, -221}},
		{*Origo(), *Origo(), *Origo()},
	}

	for _, testData := range tests {
		result := testData.v.Add(&testData.w)
		if !floatEquals(result.x, testData.expected.x, 1e-6) || !floatEquals(result.y, testData.expected.y, 1e-6) || !floatEquals(result.z, testData.expected.z, 1e-6) {
			t.Errorf("Expected Vertex(%f, %f, %f), got Vertex(%f, %f, %f)", testData.expected.x, testData.expected.y, testData.expected.z, result.x, result.y, result.z)
		}
	}
}

func TestVertexSubtraction(t *testing.T) {
	tests := []struct {
		v, w     Vertex
		expected Vertex
	}{
		{Vertex{3, 3, 3}, Vertex{2, 2, 2}, Vertex{1, 1, 1}},
		{Vertex{2, 0, 2}, *Origo(), Vertex{2, 0, 2}},
		{Vertex{201, 34.001, 230}, Vertex{-324.99, -99.12, 76.304}, Vertex{525.99, 133.121, 153.696}},
		{Vertex{-1, -10, -98}, Vertex{0, -4, -123}, Vertex{-1, -6, 25}},
		{*Origo(), *Origo(), *Origo()},
	}

	for _, testData := range tests {
		result := testData.v.Sub(&testData.w)
		if !floatEquals(result.x, testData.expected.x, 1e-6) || !floatEquals(result.y, testData.expected.y, 1e-6) || !floatEquals(result.z, testData.expected.z, 1e-6) {
			t.Errorf("Expected Vertex(%f, %f, %f), got Vertex(%f, %f, %f)", testData.expected.x, testData.expected.y, testData.expected.z, result.x, result.y, result.z)
		}
	}
}

func TestCopyConstructor(t *testing.T) {
	v1 := Vertex{23.23, 807.405, -907.409}
	v2 := CopyVertex(&v1)
	if v1.x != v2.x || v1.y != v2.y || v1.z != v2.z {
		t.Errorf("Expected copied Vertex to be equal, got Vertex(%f, %f, %f)", v2.x, v2.y, v2.z)
	}

	if &v1 == v2 {
		t.Errorf("Expected copied Vertex to be a different instance, got same instance")
	}
}

func TestVertexDistance(t *testing.T) {
	tests := []struct {
		v, w     Vertex
		expected float64
	}{
		{Vertex{1, 1, 1}, Vertex{2, 2, 2}, math.Sqrt(3)},
		{*Origo(), Vertex{2, 0, 2}, math.Sqrt(8)},
		{Vertex{201, 34.001, 230}, Vertex{-324.99, -99.12, 76.304}, 566.346277},
		{Vertex{100, 120, 130}, Vertex{100, 120, 130}, 0.0},
		{Vertex{-1, -10, -98}, Vertex{0, -4, -123}, 21.817424},
		{*Origo(), *Origo(), 0.0},
	}

	for _, testData := range tests {
		result := testData.v.Distance(&testData.w)
		if !floatEquals(result, testData.expected, 1e-6) {
			t.Errorf("Expected distance %f, got %f", testData.expected, result)
		}
	}
}

func TestVertexAbs(t *testing.T) {
	tests := []struct {
		v        Vertex
		expected float64
	}{
		{Vertex{1, 1, 1}, math.Sqrt(3)},
		{*Origo(), 0.0},
		{Vertex{201, 34.001, 230}, math.Sqrt(math.Pow(201, 2) + math.Pow(34.001, 2) + math.Pow(230, 2))},
		{Vertex{-1, -10, -98}, math.Sqrt(math.Pow(-1, 2) + math.Pow(-10, 2) + math.Pow(-98, 2))},
	}

	for _, testData := range tests {
		result := testData.v.Abs()
		if !floatEquals(result, testData.expected, 1e-6) {
			t.Errorf("Expected absolute value %f, got %f", testData.expected, result)
		}
	}
}

func TestToString(t *testing.T) {
	var v = NewVertex(20, 30, 40)
	var expected = "(20.000, 30.000, 40.000)"
	var stringVertex = v.String()
	if stringVertex != expected {
		t.Errorf("Expected string %s, got %s", expected, stringVertex)
	}
}

func TestVertexScaleConstructor(t *testing.T) {
	var v = NewScaledVertex(1, 2, 3, 2)
	var expected = Vertex{2, 4, 6}
	if !floatEquals(v.x, expected.x, 1e-6) || !floatEquals(v.y, expected.y, 1e-6) || !floatEquals(v.z, expected.z, 1e-6) {
		t.Errorf("Expected Vertex(%f, %f, %f), got Vertex(%f, %f, %f)", expected.x, expected.y, expected.z, v.x, v.y, v.z)
	}
}

func floatEquals(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}
