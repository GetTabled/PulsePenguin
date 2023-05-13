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

type waveHeader struct {
	ChunkID       [4]byte // RIFF string
	ChunkSize     uint32  // Overall size of file in bytes
	Format        [4]byte // WAVE string
	SubChunk1ID   [4]byte // fmt string
	SubChunk1Size uint32  // Length of the format data
	AudioFormat   uint16  // Format type | 1 = PCM, 3 = IEEE float, 6 = 8bit A law, 7 = 8bit mu law
	NumChannels   uint16  // Number of channels | 1 = mono , 2 = stereo
	SampleRate    uint32  // Sampling rate (samples per second)
	ByteRate      uint32  // SampleRate * NumChannels * BitsPerSample/8
	BlockAlign    uint16  // NumChannels * BitsPerSample/8
	BitsPerSample uint16  // Bits per sample | 8 = 8bits, 16 = 16 bits, etc.
	SubChunk2ID   [4]byte // data string
	SubChunk2Size uint32  // data size (NumSamples * NumChannels * BitsPerSample/8)
}
