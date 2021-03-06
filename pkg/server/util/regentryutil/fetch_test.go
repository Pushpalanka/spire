package regentryutil

import (
	"context"
	"testing"

	"github.com/spiffe/spire/proto/common"
	"github.com/spiffe/spire/proto/server/datastore"
	"github.com/spiffe/spire/test/fakes/fakedatastore"
	"github.com/stretchr/testify/assert"
)

var (
	ctx = context.Background()
)

func TestFetchRegistrationEntries(t *testing.T) {
	assert := assert.New(t)
	dataStore := fakedatastore.New()

	createRegistrationEntry := func(entry *datastore.RegistrationEntry) *datastore.RegistrationEntry {
		resp, err := dataStore.CreateRegistrationEntry(ctx, &datastore.CreateRegistrationEntryRequest{
			Entry: entry,
		})
		assert.NoError(err)
		return resp.Entry
	}

	setNodeSelectors := func(spiffeID string, selectors ...*common.Selector) {
		_, err := dataStore.SetNodeSelectors(ctx, &datastore.SetNodeSelectorsRequest{
			Selectors: &datastore.NodeSelectors{
				SpiffeId:  spiffeID,
				Selectors: selectors,
			},
		})
		assert.NoError(err)
	}

	rootID := "spiffe://example.org/root"
	oneID := "spiffe://example.org/1"
	twoID := "spiffe://example.org/2"
	threeID := "spiffe://example.org/3"
	fourID := "spiffe://example.org/4"
	fiveID := "spiffe://example.org/5"

	a1 := &common.Selector{Type: "a", Value: "1"}
	b2 := &common.Selector{Type: "b", Value: "2"}

	//
	//        root             4(a1,b2)
	//        /   \           /
	//       1     2         5
	//            /
	//           3
	//
	// node resolvers map from 2 to 4

	oneEntry := createRegistrationEntry(&datastore.RegistrationEntry{
		ParentId: rootID,
		SpiffeId: oneID,
	})

	twoEntry := createRegistrationEntry(&datastore.RegistrationEntry{
		ParentId: rootID,
		SpiffeId: twoID,
	})

	threeEntry := createRegistrationEntry(&datastore.RegistrationEntry{
		ParentId: twoID,
		SpiffeId: threeID,
	})

	fourEntry := createRegistrationEntry(&datastore.RegistrationEntry{
		SpiffeId:  fourID,
		Selectors: []*common.Selector{a1, b2},
	})

	fiveEntry := createRegistrationEntry(&datastore.RegistrationEntry{
		ParentId: fourID,
		SpiffeId: fiveID,
	})

	setNodeSelectors(twoID, a1, b2)

	actual, err := FetchRegistrationEntries(ctx, dataStore, rootID)
	assert.NoError(err)

	expected := []*common.RegistrationEntry{
		oneEntry,
		twoEntry,
		threeEntry,
		fourEntry,
		fiveEntry,
	}
	assert.Equal(expected, actual)
}
