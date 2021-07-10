package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

func TestMedianFinder1(t *testing.T) {
	finder := problems.NewMedianFinder()

	finder.AddNum(1)
	finder.AddNum(2)

	m := finder.FindMedian()
	if m != 1.5 {
		t.Errorf("Expected median to be 1.5, got %.1f\n", m)
		t.FailNow()
	}

	finder.AddNum(3)

	m = finder.FindMedian()
	if m != 2.0 {
		t.Errorf("Expected median to be 2.0, got %.1f\n", m)
		t.FailNow()
	}
}

func TestMedianFinder2(t *testing.T) {
	finder := problems.NewMedianFinder()

	m := finder.FindMedian()
	if m != 0.0 {
		t.Errorf("Expected median to be 0.0, got %.1f\n", m)
		t.FailNow()
	}
}

func TestMedianFinder3(t *testing.T) {
	finder := problems.NewMedianFinder()

	finder.AddNum(1)
	finder.AddNum(2)
	finder.AddNum(3)
	finder.AddNum(4)
	finder.AddNum(5)

	m := finder.FindMedian()
	if m != 3.0 {
		t.Errorf("Expected median to be 3.0, got %.1f\n", m)
		t.FailNow()
	}
}

func TestMedianFinder4(t *testing.T) {
	finder := problems.NewMedianFinder()

	finder.AddNum(1)
	finder.AddNum(2)
	finder.AddNum(3)
	finder.AddNum(4)

	m := finder.FindMedian()
	if m != 2.5 {
		t.Errorf("Expected median to be 2.5, got %.1f\n", m)
		t.FailNow()
	}
}

func TestMedianFinder5(t *testing.T) {
	finder := problems.NewMedianFinder()

	finder.AddNum(-1)
	m := finder.FindMedian()
	if m != -1.0 {
		t.Errorf("Expected median to be -1.0, got %.1f\n", m)
		t.FailNow()
	}
	finder.AddNum(-2)
	m = finder.FindMedian()
	if m != -1.5 {
		t.Errorf("Expected median to be -1.5, got %.1f\n", m)
		t.FailNow()
	}
	finder.AddNum(-3)
	m = finder.FindMedian()
	if m != -2.0 {
		t.Errorf("Expected median to be -2.0, got %.1f\n", m)
		t.FailNow()
	}
	finder.AddNum(-4)
	m = finder.FindMedian()
	if m != -2.5 {
		t.Errorf("Expected median to be -2.5, got %.1f\n", m)
		t.FailNow()
	}
	finder.AddNum(-5)
	m = finder.FindMedian()
	if m != -3.0 {
		t.Errorf("Expected median to be -3.0, got %.1f\n", m)
		t.FailNow()
	}
}

func TestMedianFinder6(t *testing.T) {
	finder := problems.NewMedianFinder()

	finder.AddNum(1)
	finder.AddNum(-2)

	m := finder.FindMedian()
	if m != -0.5 {
		t.Errorf("Expected median to be -0.5, got %.1f\n", m)
		t.FailNow()
	}
}
