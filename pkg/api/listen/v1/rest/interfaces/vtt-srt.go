// Copyright 2023-2024 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

/*
This package provides the types for the Deepgram PreRecorded API.
*/
package interfacesv1

import (
	"errors"
	"fmt"
	"strings"
)

// ToWebVTT implements output for VTT
//
// Deprecated: This function is deprecated. This will be removed in a future release. For VTT or SRT please
// capture the transcription with utterance enabled, then use the following projects to generate VTT or SRT files:
// - https://github.com/infinitusai/deepgram-python-captions
// - https://github.com/infinitusai/deepgram-js-captions
func (resp *PreRecordedResponse) ToWebVTT() (string, error) {
	if resp.Results.Utterances == nil {
		return "", errors.New("this function requires a transcript that was generated with the utterances feature")
	}

	vtt := "WEBVTT\n\n"
	vtt += "NOTE\nTranscription provided by Deepgram\nRequest ID: " + resp.Metadata.RequestID + "\nCreated: " + resp.Metadata.Created + "\n\n"

	for i, utterance := range resp.Results.Utterances {
		utterance := utterance
		start := SecondsToTimestamp(utterance.Start)
		end := SecondsToTimestamp(utterance.End)
		vtt += fmt.Sprintf("%d\n%s --> %s\n%s\n\n", i+1, start, end, utterance.Transcript)
	}
	return vtt, nil
}

// ToSRT implements output for SRT
//
// Deprecated: This function is deprecated. This will be removed in a future release. For VTT or SRT please
// capture the transcription with utterance enabled, then use the following projects to generate VTT or SRT files:
// - https://github.com/infinitusai/deepgram-python-captions
// - https://github.com/infinitusai/deepgram-js-captions
func (resp *PreRecordedResponse) ToSRT() (string, error) {
	if resp.Results.Utterances == nil {
		return "", errors.New("this function requires a transcript that was generated with the utterances feature")
	}

	srt := ""

	for i, utterance := range resp.Results.Utterances {
		utterance := utterance
		start := SecondsToTimestamp(utterance.Start)
		end := SecondsToTimestamp(utterance.End)
		end = strings.ReplaceAll(end, ".", ",")
		srt += fmt.Sprintf("%d\n%s --> %s\n%s\n\n", i+1, start, end, utterance.Transcript)
	}
	return srt, nil
}

func SecondsToTimestamp(seconds float64) string {
	hours := int(seconds / 3600)
	minutes := int((seconds - float64(hours*3600)) / 60)
	seconds = seconds - float64(hours*3600) - float64(minutes*60)
	return fmt.Sprintf("%02d:%02d:%02.3f", hours, minutes, seconds)
}
