package a

// Shouldn't care about Go types
type BasketOfFruit struct {
	RedFruit    []string `json:"redFruit,omitempty"`    // want `naming convention "nofruit": field BasketOfFruit.RedFruit: no fields should contain any variations of the word 'fruit' in their name`
	OrangeFruit []string `json:"orangeFruit,omitempty"` // want `naming convention "nofruit": field BasketOfFruit.OrangeFruit: no fields should contain any variations of the word 'fruit' in their name`
	GreenFruit  []string `json:"greenFruit,omitempty"`  // want `naming convention "nofruit": field BasketOfFruit.GreenFruit: no fields should contain any variations of the word 'fruit' in their name`
	FruitBlue   []string `json:"fruitBlue,omitempty"`   // want `naming convention "nofruit": field BasketOfFruit.FruitBlue: no fields should contain any variations of the word 'fruit' in their name`
	Fruit       []string `json:"fruit,omitempty"`       // want `naming convention "nofruit": field BasketOfFruit.Fruit: no fields should contain any variations of the word 'fruit' in their name`
	AFruit      string   `json:"aFruit,omitempty"`      // want `naming convention "nofruit": field BasketOfFruit.AFruit: no fields should contain any variations of the word 'fruit' in their name`
}

// Shouldn't care about methods
func (b BasketOfFruit) GrabFruit() string {
	return "nothing"
}

// shouldn't care about functions
func IsFruit(in string) {
	return
}

type SpecialBehaviors struct {
	SomethingBehavior string `json:"somethingBehavior,omitempty"` // want `naming convention "preferbehaviour": field SpecialBehaviors.SomethingBehavior: prefer the use of the word 'behaviour' instead of 'behavior'.`
	BehaviorCrazy     bool   `json:"behaviorCrazy,omitempty"`     // want `naming convention "preferbehaviour": field SpecialBehaviors.BehaviorCrazy: prefer the use of the word 'behaviour' instead of 'behavior'.`
}

type Configurations struct {
	SomeSupportedThingy string `json:"someSupportedThingy,omitempty"`
	UnsupportedThingy   string `json:"unsupportedThingy,omitempty"` // want `naming convention "nounsupported": field Configurations.UnsupportedThingy: no fields allowing for unsupported behaviors allowed`
}

type TestSet struct {
	TestName string `json:"testName,omitempty"`  // want `naming convention "notest": field TestSet.TestName: no temporary test fields`
	Other    string `json:"otherTest,omitempty"` // want `naming convention "notest": field TestSet.Other: no temporary test fields`
}
