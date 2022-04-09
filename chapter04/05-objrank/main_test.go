package main

import "testing"

func TestCase1(t *testing.T) {
	r := FatRateRank{}
	r.inputRecord("王强",0.38)
	r.inputRecord("王强",0.32)
	{
		randOfWQ, fatRateOfWQ := r.getRank("王强")
		if randOfWQ != 1 {
			t.Fatalf("预期 王强 第一，但是得到的是： %d", randOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂是0.32, 但得到的是: %f", fatRateOfWQ)
		}
	}
	r.inputRecord("李静",0.28)
	randOfLJ, fatRateOfLJ := r.getRank("李静")
	{
		if randOfLJ != 1 {
			t.Fatalf("预期 李静 第一，但是得到的是： %d", randOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期王强的体脂是0.28, 但得到的是: %f", fatRateOfLJ)
		}
	}
}

func TestCase2(t *testing.T) {
	r := &FatRateRank{}
	r.inputRecord("王强",0.38)
	r.inputRecord("张伟",0.38)
	r.inputRecord("李静",0.28)
	randOfLJ, fatRateOfLJ := r.getRank("李静")
	{
		if randOfLJ != 1 {
			t.Fatalf("预期 李静 第一，但是得到的是： %d", randOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期王强的体脂是0.28, 但得到的是: %f", fatRateOfLJ)
		}
	}
	{
		randOfWQ, fatRateOfWQ := r.getRank("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期 王强 第二，但是得到的是： %d", randOfWQ)
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("预期王强的体脂是0.32, 但得到的是: %f", fatRateOfWQ)
		}
	}
	{
		randOfWQ, fatRateOfWQ := r.getRank("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期 王强 第二，但是得到的是： %d", randOfWQ)
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("预期王强的体脂是0.32, 但得到的是: %f", fatRateOfWQ)
		}
	}
}

func TestCase3(t *testing.T) {
	r := FatRateRank{}
	r.inputRecord("王强",0.38)
	r.inputRecord("李静",0.23)
	r.inputRecord("张伟")
	randOfLJ, fatRateOfLJ := r.getRank("李静")
	{
		if randOfLJ != 1 {
			t.Fatalf("预期 李静 第一，但是得到的是： %d", randOfLJ)
		}
		if fatRateOfLJ != 0.23 {
			t.Fatalf("预期李静的体脂是0.23, 但得到的是: %f", fatRateOfLJ)
		}
	}
	{
		randOfWQ, fatRateOfWQ := r.getRank("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期 王强 第二，但是得到的是： %d", randOfWQ)
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("预期王强的体脂是0.32, 但得到的是: %f", fatRateOfWQ)
		}
	}
	{
		randOfWQ, _ := r.getRank("张伟")
		if randOfWQ != 3 {
			t.Fatalf("预期 张伟 第三，但是得到的是： %d", randOfWQ)
		}
	}
}

