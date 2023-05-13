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

func CreateWaveHeader(numSamples uint32, numChannels uint16, sampleRate uint32, bitsPerSample uint16) waveHeader {
	var header waveHeader
	header.ChunkID = [4]byte{'R', 'I', 'F', 'F'}
	header.ChunkSize = 36 + (numSamples * uint32(numChannels) * uint32(bitsPerSample) / 8)
	header.Format = [4]byte{'W', 'A', 'V', 'E'}
	header.SubChunk1ID = [4]byte{'f', 'm', 't', ' '}
	header.SubChunk1Size = 16
	header.AudioFormat = 1
	header.NumChannels = numChannels
	header.SampleRate = sampleRate
	header.BitsPerSample = bitsPerSample
	header.ByteRate = sampleRate * uint32(numChannels) * uint32(bitsPerSample) / 8
	header.BlockAlign = numChannels * bitsPerSample / 8
	header.SubChunk2ID = [4]byte{'d', 'a', 't', 'a'}
	header.SubChunk2Size = numSamples * uint32(numChannels) * uint32(bitsPerSample) / 8
	return header
}
