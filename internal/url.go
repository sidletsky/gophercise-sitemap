package internal

import (
	"fmt"
)

type Url struct {
	Loc string
}

func (u Url) String() string {
	return fmt.Sprintf(`        <url>
			<loc>%s</loc>
		</url>`, u.Loc)
}
