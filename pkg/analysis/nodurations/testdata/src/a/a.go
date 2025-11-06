package a

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Durations struct {
	ValidString string

	ValidMap map[string]string

	ValidInt32 int32

	ValidInt64 int64

	InvalidDuration time.Duration // want "field Durations.InvalidDuration should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationPtr *time.Duration // want "field Durations.InvalidDurationPtr pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationSlice []time.Duration // want "field Durations.InvalidDurationSlice array element should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationPtrSlice []*time.Duration // want "field Durations.InvalidDurationPtrSlice array element pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationAlias DurationAlias // want "field Durations.InvalidDurationAlias type DurationAlias should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationPtrAlias *DurationAlias // want "field Durations.InvalidDurationPtrAlias pointer type DurationAlias should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationSliceAlias []DurationAlias // want "field Durations.InvalidDurationSliceAlias array element type DurationAlias should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationPtrSliceAlias []*DurationAlias // want "field Durations.InvalidDurationPtrSliceAlias array element pointer type DurationAlias should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidMapStringToDuration map[string]time.Duration // want "field Durations.InvalidMapStringToDuration map value should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidMapStringToDurationPtr map[string]*time.Duration // want "field Durations.InvalidMapStringToDurationPtr map value pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidMapDurationToString map[time.Duration]string // want "field Durations.InvalidMapDurationToString map key should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidMapDurationPtrToString map[*time.Duration]string // want "field Durations.InvalidMapDurationPtrToString map key pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationAliasFromAnotherFile DurationAliasB // want "field Durations.InvalidDurationAliasFromAnotherFile type DurationAliasB should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationPtrAliasFromAnotherFile *DurationAliasB // want "field Durations.InvalidDurationPtrAliasFromAnotherFile pointer type DurationAliasB should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."
}

// DoNothing is used to check that the analyser doesn't report on methods.
func (Durations) DoNothing(a bool) bool {
	return a
}

type DurationAlias time.Duration // want "type DurationAlias should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

type DurationAliasPtr *time.Duration // want "type DurationAliasPtr pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

type DurationAliasSlice []time.Duration // want "type DurationAliasSlice array element should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

type DurationAliasPtrSlice []*time.Duration // want "type DurationAliasPtrSlice array element pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

type MapStringToDurationAlias map[string]time.Duration // want "type MapStringToDurationAlias map value should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

type MapStringToDurationPtrAlias map[string]*time.Duration // want "type MapStringToDurationPtrAlias map value pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

type DurationsWithMetaV1Package struct {
	ValidString string

	ValidMap map[string]string

	ValidInt32 int32

	ValidInt64 int64

	InvalidDuration metav1.Duration // want "field DurationsWithMetaV1Package.InvalidDuration should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationPtr *metav1.Duration // want "field DurationsWithMetaV1Package.InvalidDurationPtr pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationSlice []metav1.Duration // want "field DurationsWithMetaV1Package.InvalidDurationSlice array element should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationPtrSlice []*metav1.Duration // want "field DurationsWithMetaV1Package.InvalidDurationPtrSlice array element pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationAlias DurationAliasWithMetaV1 // want "field DurationsWithMetaV1Package.InvalidDurationAlias type DurationAliasWithMetaV1 should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationPtrAlias *DurationAliasWithMetaV1 // want "field DurationsWithMetaV1Package.InvalidDurationPtrAlias pointer type DurationAliasWithMetaV1 should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationSliceAlias []DurationAliasWithMetaV1 // want "field DurationsWithMetaV1Package.InvalidDurationSliceAlias array element type DurationAliasWithMetaV1 should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationPtrSliceAlias []*DurationAliasWithMetaV1 // want "field DurationsWithMetaV1Package.InvalidDurationPtrSliceAlias array element pointer type DurationAliasWithMetaV1 should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidMapStringToDuration map[string]metav1.Duration // want "field DurationsWithMetaV1Package.InvalidMapStringToDuration map value should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidMapStringToDurationPtr map[string]*metav1.Duration // want "field DurationsWithMetaV1Package.InvalidMapStringToDurationPtr map value pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidMapDurationToString map[metav1.Duration]string // want "field DurationsWithMetaV1Package.InvalidMapDurationToString map key should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidMapDurationPtrToString map[*metav1.Duration]string // want "field DurationsWithMetaV1Package.InvalidMapDurationPtrToString map key pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

	InvalidDurationAliasFromAnotherFile DurationAliasBWithMetaV1 // want "field DurationsWithMetaV1Package.InvalidDurationAliasFromAnotherFile type DurationAliasBWithMetaV1 should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."
}

type DurationAliasWithMetaV1 metav1.Duration // want "type DurationAliasWithMetaV1 should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

type DurationAliasPtrWithMetaV1 *metav1.Duration // want "type DurationAliasPtrWithMetaV1 pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

type DurationAliasSliceWithMetaV1 []metav1.Duration // want "type DurationAliasSliceWithMetaV1 array element should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

type DurationAliasPtrSliceWithMetaV1 []*metav1.Duration // want "type DurationAliasPtrSliceWithMetaV1 array element pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

type MapStringToDurationAliaWithMetaV1 map[string]metav1.Duration // want "type MapStringToDurationAliaWithMetaV1 map value should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

type MapStringToDurationPtrAliasWithMetaV1 map[string]*metav1.Duration // want "type MapStringToDurationPtrAliasWithMetaV1 map value pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."
