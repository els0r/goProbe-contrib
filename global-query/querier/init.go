//go:build goprobe_contrib

package querier

import (
	// register plugins below this line
	_ "github.com/els0r/goProbe-contrib/global-query/querier/dummyquerier"
)
