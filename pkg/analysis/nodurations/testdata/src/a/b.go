package a

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DurationAliasB time.Duration // want "type DurationAliasB should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

type DurationAliasPtrB *time.Duration // want "type DurationAliasPtrB pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

type DurationAliasBWithMetaV1 metav1.Duration // want "type DurationAliasBWithMetaV1 should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."

type DurationAliasPtrBWithMetaV1 *metav1.Duration // want "type DurationAliasPtrBWithMetaV1 pointer should not use a Duration. Use an integer type with units in the name to avoid the need for clients to implement Go style duration parsing."
