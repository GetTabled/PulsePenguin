/*
  PulsePenguin - Lightweight and easy-to-use open-source audio beat detection application
  Copyright (C) 2023 GetTabled
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package wave

import (
	"encoding/binary"
	"fmt"
	"os"
)

func ReadWaveHeader(filename string) {

	file, err := os.Open(filename) // Open given .wav file for reading
	if err != nil {
		fmt.Println("Error opening WAVE file:", err)
		return
	}
	defer file.Close()

	// Parse the WAVE file header
	header := waveHeader{}
	err = binary.Read(file, binary.LittleEndian, &header)
	if err != nil {
		fmt.Println("Error reading WAVE file header:", err)
		return
	}

	// Ensure that the file is a valid WAVE file
	if string(header.ChunkID[:]) != "RIFF" || string(header.Format[:]) != "WAVE" {
		fmt.Println("Error: not a valid WAVE file")
		return
	}

	// Create output file
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}

	fmt.Fprintf(outputFile, "File: %s\n\n", file.Name())
	fmt.Fprintf(outputFile, "ChunkID: %s\n", header.ChunkID)
	fmt.Fprintf(outputFile, "ChunkSize: %d\n", header.ChunkSize)
	fmt.Fprintf(outputFile, "Format: %s\n", header.Format)
	fmt.Fprintf(outputFile, "Subchunk1ID: %s\n", header.SubChunk1ID)
	fmt.Fprintf(outputFile, "Subchunk1Size: %d\n", header.SubChunk1Size)
	fmt.Fprintf(outputFile, "AudioFormat: %d\n", header.AudioFormat)
	fmt.Fprintf(outputFile, "NumChannels: %d\n", header.NumChannels)
	fmt.Fprintf(outputFile, "SampleRate: %d\n", header.SampleRate)
	fmt.Fprintf(outputFile, "ByteRate: %d\n", header.ByteRate)
	fmt.Fprintf(outputFile, "BlockAlign: %d\n", header.BlockAlign)
	fmt.Fprintf(outputFile, "BitsPerSample: %d\n", header.BitsPerSample)
	fmt.Fprintf(outputFile, "Subchunk2ID: %s\n", header.SubChunk2ID)
	fmt.Fprintf(outputFile, "Subchunk2Size: %d\n", header.SubChunk2Size)

	fmt.Println("Header written to output.txt")
}
