/* Copyright (C) 2019 Monomax Software Pty Ltd
 *
 * This file is part of Dnote CLI.
 *
 * Dnote CLI is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Dnote CLI is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Dnote CLI.  If not, see <https://www.gnu.org/licenses/>.
 */

package core

import (
	"fmt"
	"time"

	"github.com/dnote/dnote/cli/log"
)

// PrintNoteInfo prints a note information
func PrintNoteInfo(info NoteInfo) {
	log.Infof("book name: %s\n", info.BookLabel)
	log.Infof("created at: %s\n", time.Unix(0, info.AddedOn).Format("Jan 2, 2006 3:04pm (MST)"))
	if info.EditedOn != 0 {
		log.Infof("updated at: %s\n", time.Unix(0, info.EditedOn).Format("Jan 2, 2006 3:04pm (MST)"))
	}
	log.Infof("note id: %d\n", info.RowID)
	log.Infof("note uuid: %s\n", info.UUID)

	fmt.Printf("\n------------------------content------------------------\n")
	fmt.Printf("%s", info.Content)
	fmt.Printf("\n-------------------------------------------------------\n")
}