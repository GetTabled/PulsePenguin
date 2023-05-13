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

func WriteWaveHeader(filename string, header waveHeader) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = binary.Write(file, binary.LittleEndian, header)
	if err != nil {
		return err
	}
	fmt.Println("File closed")
	return nil
}
