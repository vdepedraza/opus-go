// Copyright © 2015, 2016 Authors (see AUTHORS file)
//
// License for use of this code is detailed in the LICENSE file

package opus

import "testing"

func TestEncoderNew(t *testing.T) {
	enc, err := NewEncoder(48000, 1, APPLICATION_VOIP)
	if err != nil || enc == nil {
		t.Errorf("Error creating new encoder: %v", err)
	}
	enc, err = NewEncoder(12345, 1, APPLICATION_VOIP)
	if err == nil || enc != nil {
		t.Errorf("Expected error for illegal samplerate 12345")
	}
}

func TestEncoderUnitialized(t *testing.T) {
	var enc Encoder
	_, err := enc.Encode(nil, nil)
	if err != errEncUninitialized {
		t.Errorf("Expected \"unitialized encoder\" error: %v", err)
	}
	_, err = enc.EncodeFloat32(nil, nil)
	if err != errEncUninitialized {
		t.Errorf("Expected \"unitialized encoder\" error: %v", err)
	}
}

func TestEncoderDTX(t *testing.T) {
	enc, err := NewEncoder(8000, 1, APPLICATION_VOIP)
	if err != nil || enc == nil {
		t.Errorf("Error creating new encoder: %v", err)
	}
	vals := []bool{true, false}
	for _, dtx := range vals {
		enc.UseDTX(dtx)
		gotv := enc.DTX()
		if gotv != dtx {
			t.Errorf("Error set dtx: expect dtx=%v, got dtx=%v", dtx, gotv)
		}
	}
}

func TestEncoderSampleRate(t *testing.T) {
	sample_rates := []int{8000, 12000, 16000, 24000, 48000}
	for _, f := range sample_rates {
		enc, err := NewEncoder(f, 1, APPLICATION_VOIP)
		if err != nil || enc == nil {
			t.Fatalf("Error creating new encoder with sample_rate %d Hz: %v", f, err)
		}
		f2 := enc.SampleRate()
		if f != f2 {
			t.Errorf("Unexpected sample rate reported by %d Hz encoder: %d", f, f2)
		}
	}
}
