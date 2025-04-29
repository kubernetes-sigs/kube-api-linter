package a

type Float32AliasB float32 // want "type Float32AliasB should not use a float value because they cannot be reliably round-tripped."

type Float32AliasPtrB *float32 // want "type Float32AliasPtrB pointer should not use a float value because they cannot be reliably round-tripped."
