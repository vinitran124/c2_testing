package main

import "testing"

func TestIsEven(t *testing.T) {
	// Test cho số chẵn
	if IsEven(4) {
		t.Error("4 phải là số chẵn")
	}

	// Test cho số lẻ
	if IsEven(7) {
		t.Error("7 không phải là số chẵn")
	}

	// Test cho số 0
	if !IsEven(0) {
		t.Error("0 phải là số chẵn")
	}

	// Test cho số âm chẵn
	if !IsEven(-10) {
		t.Error("-10 phải là số chẵn")
	}

	// Test cho số âm lẻ
	if IsEven(-3) {
		t.Error("-3 không phải là số chẵn")
	}
}
